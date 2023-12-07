package listener

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrescosta/jobico/api/pkg/remote"
	pb "github.com/andrescosta/jobico/api/types"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Controller struct {
	queueClient *remote.QueueClient
	events      map[string]*Event
}

type Event struct {
	event  *pb.EventDef
	schema *jsonschema.Schema
}

func ConfigureRoutes(ctx context.Context, r *mux.Router) error {
	events, err := buildEventsMap(ctx)
	if err != nil {
		return err
	}

	queueClient, err := remote.NewQueueClient(ctx)
	if err != nil {
		return err
	}

	c := Controller{
		events:      events,
		queueClient: queueClient,
	}

	s := r.PathPrefix("/events").Subrouter()
	s.Methods("POST").Path("/{tenant_id}/{event_id}").HandlerFunc(c.Post)
	s.Methods("GET").HandlerFunc(c.Get)
	return nil
}

func getFullEventId(tenantId string, eventId string) string {
	return tenantId + "/" + eventId
}

func (rr Controller) getEventDef(tenantId string, eventId string) (*Event, error) {
	ev, ok := rr.events[getFullEventId(tenantId, eventId)]
	if !ok {
		return nil, fmt.Errorf("event unknown")
	}
	return ev, nil
}

func (rr Controller) Get(w http.ResponseWriter, r *http.Request) {
}

func (c Controller) Post(writer http.ResponseWriter, request *http.Request) {
	logger := zerolog.Ctx(request.Context())
	event := pb.MerchantData{}
	if err := json.NewDecoder(request.Body).Decode(&event); err != nil {
		logger.Error().Msgf("Failed to decode request body: %s", err)
		http.Error(writer, "Request body illegal", http.StatusBadRequest)
		return
	}

	var items []*pb.QueueItem
	tenantId := chi.URLParam(request, "tenant_id")
	eventId := chi.URLParam(request, "event_id")
	ef, err := c.getEventDef(tenantId, eventId)
	if err != nil {
		logger.Error().Msgf("Data unknown: %s", err)
		http.Error(writer, "Event unknown", http.StatusBadRequest)
	}
	if ef.event.DataType == pb.DataType_Json {

		for _, ev := range event.Data {

			if err = ef.schema.Validate(ev); err != nil {
				logger.Error().Msgf("Failed to validate event: %s", err)
				http.Error(writer, "Failed to validate event", http.StatusBadRequest)
				return
			}

			evBin, err := json.Marshal(ev)
			if err != nil {
				logger.Error().Msgf("Failed to encode event: %s", err)
				http.Error(writer, "Failed to process event", http.StatusBadRequest)
				return
			}
			q := pb.QueueItem{
				Data:    evBin,
				EventId: eventId,
			}
			items = append(items, &q)
		}
	}
	queueRequest := pb.QueueRequest{
		QueueId:  ef.event.SupplierQueueId,
		TenantId: tenantId,
		Items:    items,
	}
	err = c.queueClient.Queue(request.Context(), &queueRequest)
	if err != nil {
		logger.Error().Msgf("Failed to connect to queue server: %s", err)
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func buildEventsMap(ctx context.Context) (map[string]*Event, error) {
	controlClient, err := remote.NewControlClient(ctx)
	if err != nil {
		return nil, err
	}
	repoClient, err := remote.NewRepoClient(ctx)
	if err != nil {
		return nil, err
	}
	events := make(map[string]*Event)

	pkgs, err := controlClient.GetAllPackages(ctx)
	if err != nil {
		return nil, err
	}

	for _, ps := range pkgs {
		tenantId := ps.TenantId
		for _, job := range ps.Jobs {
			event := job.Event
			f, err := repoClient.GetFile(ctx, tenantId, event.Schema.SchemaRef)
			if err != nil {
				return nil, err
			}
			comp := jsonschema.NewCompiler()
			if err := comp.AddResource(getFullEventId(tenantId, event.ID), bytes.NewReader(f)); err != nil {
				return nil, err
			}
			compiledSchema, err := comp.Compile(getFullEventId(tenantId, event.ID))
			if err != nil {
				return nil, err
			}

			events[getFullEventId(tenantId, event.ID)] = &Event{
				event:  event,
				schema: compiledSchema,
			}
		}
	}
	return events, nil
}
