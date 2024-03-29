package test

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/andrescosta/goico/pkg/test"
	pb "github.com/andrescosta/jobico/internal/api/types"
	"go.uber.org/goleak"
)

const sendEventURL = "http://listener:1/events/%s/%s"

var (
	//go:embed testdata/schema.json
	schemaV1 []byte

	//go:embed testdata/schema_result_ok.json
	schemaV1Ok []byte

	//go:embed testdata/schema_result_error.json
	schemaV1Error []byte

	//go:embed testdata/echo.wasm
	wasmEcho []byte

	//go:embed testdata/error.wasm
	wasmError []byte

	files = map[string][]byte{
		"sch1":       schemaV1,
		"sch1_ok":    schemaV1Ok,
		"sch1_error": schemaV1Error,
		"run1":       wasmEcho,
		"runerror1":  wasmError,
	}

	//go:embed testdata/schema_updated.json
	schemaV2  []byte
	schemasV2 = map[string][]byte{
		"sch1_v2": schemaV2,
	}
)

func TestStartStop(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.ctl, platform.queue) //, platform.recorder, platform.listener, platform.repo, platform.executor)
	test.Nil(t, err)
}

func TestOk(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.ctl, platform.queue, platform.recorder, platform.listener, platform.repo)
	test.Nil(t, err)
	pkg := newTestPackage()
	addPackageAndFiles(t, cli, pkg)
	ps, err := cli.AllPackages()
	test.Nil(t, err)
	test.NotEmpty(t, ps)
	err = svcGroup.Start(platform.executor)
	test.Nil(t, err)
	_ = sendEvtV1AndValidate(t, pkg, cli)
}

func TestWasmError(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.ctl, platform.queue, platform.recorder, platform.listener, platform.repo)
	test.Nil(t, err)
	pkg := newErrorTestPackage()
	addPackageAndFiles(t, cli, pkg)
	ps, err := cli.AllPackages()
	test.Nil(t, err)
	test.NotEmpty(t, ps)
	err = svcGroup.Start(platform.executor)
	test.Nil(t, err)
	sendEvtV1AndValidateError(t, pkg, cli)
}

func TestStreamingSchemaUpdate(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.ctl, platform.repo, platform.listener, platform.queue, platform.recorder)
	test.Nil(t, err)
	pkg := newTestPackage()
	addPackageAndFiles(t, cli, pkg)
	test.Nil(t, err)
	err = svcGroup.Start(platform.executor)
	test.Nil(t, err)
	url := sendEvtV1AndValidate(t, pkg, cli)
	err = cli.startRecvCacheEvents()
	test.Nil(t, err)
	pkg.Jobs[0].Event.Schema.SchemaRef = "sch1_v2"
	err = cli.uploadSchemas(pkg, schemasV2)
	test.Nil(t, err)
	err = cli.updatePackage(pkg)
	test.Nil(t, err)
	err = cli.waitForCacheEvents()
	test.Nil(t, err)
	_, err = cli.sendEventV1(url)
	test.NotNil(t, err)
	err = cli.sendEventV2(url)
	test.Nil(t, err)
	chkExecOk(t, pkg, nil, cli)
}

func TestStreamingDelete(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.ctl, platform.repo, platform.listener, platform.queue, platform.recorder)
	test.Nil(t, err)
	err = svcGroup.Start(platform.executor)
	test.Nil(t, err)
	pkg := newTestPackage()
	addPackageAndFiles(t, cli, pkg)
	url := sendEvtV1AndValidate(t, pkg, cli)
	err = cli.startRecvCacheEvents()
	test.Nil(t, err)
	err = cli.deletePackage(pkg)
	test.Nil(t, err)
	err = cli.waitForCacheEvents()
	test.Nil(t, err)
	_, err = cli.sendEventV1(url)
	test.NotNil(t, err)
}

func TestEventErrors(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.ctl, platform.repo)
	test.Nil(t, err)
	pkg := newTestPackage()
	addPackageAndFiles(t, cli, pkg)
	test.Nil(t, err)
	err = svcGroup.Start(platform.listener, platform.queue)
	test.Nil(t, err)
	u := fmt.Sprintf("http://listener:1/events/%s/notexist", pkg.Tenant)
	url, err := url.Parse(u)
	test.Nil(t, err)
	_, err = cli.sendEventV1(url)
	test.ErrorIs(t, err, errSend{StatusCode: 500})
	u = "http://listener:1/events/fake/notexist"
	url, err = url.Parse(u)
	test.Nil(t, err)
	_, err = cli.sendEventV1(url)
	test.ErrorIs(t, err, errSend{StatusCode: 500})
	u = fmt.Sprintf(sendEventURL, pkg.Tenant, pkg.Jobs[0].Event.ID)
	url, err = url.Parse(u)
	test.Nil(t, err)
	err = cli.sendEventMalFormed(url)
	test.ErrorIs(t, err, errSend{StatusCode: 400})
}

func TestQueueDown(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.ctl, platform.repo, platform.recorder)
	test.Nil(t, err)
	pkg := newTestPackage()
	addPackageAndFiles(t, cli, pkg)
	err = svcGroup.Start(platform.listener, platform.queue)
	test.Nil(t, err)
	err = svcGroup.Start(platform.executor)
	test.Nil(t, err)
	url := sendEvtV1AndValidate(t, pkg, cli)
	platform.queue.Stop()
	time.Sleep(10 * time.Microsecond)
	_, err = cli.sendEventV1(url)
	test.ErrorIs(t, err, errSend{StatusCode: 500})
}

func TestErroCtl(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	os.Setenv("http.shutdown.timeout", (10 * time.Microsecond).String())
	os.Setenv("http.timeout.write", (5 * time.Microsecond).String())
	os.Setenv("http.timeout.read", (5 * time.Microsecond).String())
	os.Setenv("http.timeout.idle", (5 * time.Microsecond).String())
	os.Setenv("http.timeout.handler", (5 * time.Microsecond).String())
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatformWithTimeout(ctx, 1*time.Second)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	pkg := newPackage(SchemaRefIDs{"sch1", "sch1_ok", "sch1_error"}, "run1")
	err = svcGroup.Start(platform.repo, platform.listener, platform.queue)
	test.Nil(t, err)
	err = sendEvtV1(pkg, cli)
	test.NotNil(t, err)
}

func TestErrorInitQueue(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.repo, platform.ctl)
	test.Nil(t, err)
	pkg := newPackage(SchemaRefIDs{"sch1", "sch1_ok", "sch1_error"}, "run1")
	err = svcGroup.Start(platform.listener)
	test.Nil(t, err)
	u := fmt.Sprintf(sendEventURL, pkg.Tenant, pkg.Jobs[0].Event.ID)
	url, err := url.Parse(u)
	test.Nil(t, err)
	_, err = cli.sendEventV1(url)
	test.NotNil(t, err)
}

func TestErrorRepo(t *testing.T) {
	defer goleak.VerifyNone(t)
	setEnvVars()
	os.Setenv("dial.timeout", (1 * time.Second).String())
	ctx, cancel := context.WithCancel(context.Background())
	platform, err := newPlatform(ctx)
	test.Nil(t, err)
	svcGroup := test.NewServiceGroup()
	cli, err := newTestClient(ctx, platform.conn, platform.conn)
	defer func() {
		cancel()
		cleanUp(t, platform, svcGroup, cli)
	}()
	test.Nil(t, err)
	err = svcGroup.Start(platform.ctl)
	test.Nil(t, err)
	pkg := newPackage(SchemaRefIDs{"sch1", "sch1_ok", "sch1_error"}, "run1")
	addPackage(t, cli, pkg)
	err = svcGroup.Start(platform.queue)
	test.Nil(t, err)
	err = svcGroup.Start(platform.listener)
	test.Nil(t, err)
	err = sendEvtV1(pkg, cli)
	test.NotNil(t, err)
}

func cleanUp(t *testing.T, platform *platform, svcGroup *test.ServiceGroup, cli *testClient) {
	fail := false
	if err := svcGroup.WaitUntilStopped(); err != nil {
		errw := test.ErrService{}
		if errors.As(err, &errw) {
			t.Errorf("error while stopping the service %s: %v", errw.Starter.Addr(), errw.Err)
		} else {
			t.Errorf("error while stopping %v", err)
		}
		fail = true
	}
	if cli != nil {
		if err := cli.close(); err != nil {
			t.Errorf("error stopping service group %v", err)
			fail = true
		}
	}
	err := platform.dispose()
	if err != nil {
		t.Errorf("error disposing services %v", err)
		fail = true
	}
	if fail {
		t.FailNow()
	}
}

func addPackageAndFiles(t *testing.T, cli *testClient, pkg *pb.JobPackage) {
	err := cli.uploadSchemas(pkg, files)
	test.Nil(t, err)
	err = cli.uploadRuntimes(pkg, files)
	test.Nil(t, err)
	addPackage(t, cli, pkg)
}

func sendEvtV1AndValidate(t *testing.T, pkg *pb.JobPackage, cli *testClient) *url.URL {
	u := fmt.Sprintf(sendEventURL, pkg.Tenant, pkg.Jobs[0].Event.ID)
	url, err := url.Parse(u)
	test.Nil(t, err)
	evt, err := cli.sendEventV1(url)
	test.Nil(t, err)
	chkExecOk(t, pkg, &evt, cli)
	return url
}

func sendEvtV1AndValidateError(t *testing.T, pkg *pb.JobPackage, cli *testClient) {
	u := fmt.Sprintf(sendEventURL, pkg.Tenant, pkg.Jobs[0].Event.ID)
	url, err := url.Parse(u)
	test.Nil(t, err)
	evt, err := cli.sendEventV1(url)
	test.Nil(t, err)
	chkExecError(t, pkg, &evt, cli)
}

func sendEvtV1(pkg *pb.JobPackage, cli *testClient) error {
	u := fmt.Sprintf(sendEventURL, pkg.Tenant, pkg.Jobs[0].Event.ID)
	url, err := url.Parse(u)
	if err != nil {
		return err
	}
	_, err = cli.sendEventV1(url)
	return err
}

func addPackage(t *testing.T, cli *testClient, pkg *pb.JobPackage) *pb.JobPackage {
	err := cli.addTenant(pkg.Tenant)
	test.Nil(t, err)
	err = cli.addPackage(pkg)
	test.Nil(t, err)
	return pkg
}

func chkExecOk(t *testing.T, pkg *pb.JobPackage, evt *eventTenantV1, cli *testClient) {
	res, err := cli.dequeue(pkg.Tenant, "queue_id_1_ok")
	test.Nil(t, err)
	test.NotEmpty(t, res)
	results, err := cli.getJobExecutions(pkg, 2)
	test.Nil(t, err)
	test.Len(t, results, 2)
	if evt != nil {
		valResForEvtV1(t, evt, results)
	}
	valLog(t, results)
}

func chkExecError(t *testing.T, pkg *pb.JobPackage, evt *eventTenantV1, cli *testClient) {
	res, err := cli.dequeue(pkg.Tenant, "queue_id_1_error")
	test.Nil(t, err)
	test.NotEmpty(t, res)
	results, err := cli.getJobExecutions(pkg, 2)
	test.Nil(t, err)
	test.Len(t, results, 2)
	if evt != nil {
		valResForEvtV1Error(t, results)
	}
	valLogError(t, results)
}

func valLog(t *testing.T, rs []Result) {
	var result *Result
	for _, r := range rs {
		if r.TypeResult == strings.ToLower(pb.JobResult_Log.String()) {
			rr := r
			result = &rr
			break
		}
	}
	test.NotNil(t, result)
	test.Equals(t, result.Code, 6) // 6- Level info
	test.NotEmpty(t, result.ResultString)
}

func valLogError(t *testing.T, rs []Result) {
	var result *Result
	for _, r := range rs {
		if r.TypeResult == strings.ToLower(pb.JobResult_Log.String()) {
			rr := r
			result = &rr
			break
		}
	}
	test.NotNil(t, result)
	test.Equals(t, result.Code, 3) // 3- Level error
	test.NotEmpty(t, result.ResultString)
}

func valResForEvtV1Error(t *testing.T, rs []Result) {
	var result *Result
	for _, r := range rs {
		if r.TypeResult == strings.ToLower(pb.JobResult_Result.String()) {
			rr := r
			result = &rr
			break
		}
	}
	test.NotNil(t, result)
	test.Equals(t, result.Code, 500)
}

func valResForEvtV1(t *testing.T, evt *eventTenantV1, rs []Result) {
	var result *Result
	for _, r := range rs {
		if r.TypeResult == strings.ToLower(pb.JobResult_Result.String()) {
			rr := r
			result = &rr
			break
		}
	}
	test.NotNil(t, result)
	test.Equals(t, result.Code, 0)
	test.NotNil(t, result.ResultJSON)
	test.Equals(t, result.ResultJSON.Age, evt.Age)
	test.Equals(t, result.ResultJSON.FirstName, evt.FirstName)
	test.Equals(t, result.ResultJSON.LastName, evt.LastName)
}

func setEnvVars() {
	os.Setenv("dial.timeout", (20 * time.Second).String())
	os.Setenv("log.level", "0")
	os.Setenv("log.console.enabled", "true")
	os.Setenv("listener.addr", "listener:1")
	os.Setenv("listener.host", "listener:1")
	os.Setenv("cache_listener.addr", "cache_listener:1")
	os.Setenv("listener.publish.event.cache", "true")

	os.Setenv("ctl.addr", "ctl:1")
	os.Setenv("ctl.host", "ctl:1")

	os.Setenv("repo.addr", "repo:1")
	os.Setenv("repo.host", "repo:1")

	os.Setenv("executor.addr", "exec:1")

	os.Setenv("queue.addr", "queue:1")
	os.Setenv("queue.host", "queue:1")

	os.Setenv("recorder.addr", "recorder:1")
	os.Setenv("recorder.host", "recorder:1")

	os.Setenv("recorder.host", "recorder:1")
	os.Setenv("recorder.addr", "recorder:1")
	os.Setenv("log.console.enabled", "false")
	os.Setenv("log.file.enabled", "false")
	os.Setenv("executor.timeout", (1 * time.Microsecond).String())
}
