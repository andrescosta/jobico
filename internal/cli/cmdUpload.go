package cli

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/andrescosta/goico/pkg/service"
	"github.com/andrescosta/jobico/internal/api/client"
	pb "github.com/andrescosta/jobico/internal/api/types"
)

func newUpload() *command {
	cmdUpload := &command{
		name:      "upload",
		usageLine: "cli upload <wasm|json> <tenant> <file id> <file name>",
		short:     "upload a wasm or json schema file",
		long: `
	The 'upload' command enables the upload of a WebAssembly or JSON schema file to the file Repository. 
	This file will be referenced by the Job definitions.`,
	}
	cmdUpload.flag = *flag.NewFlagSet("upload", flag.ContinueOnError)
	cmdUpload.run = runUpload
	cmdUpload.flag.Usage = func() {}
	return cmdUpload
}

func runUpload(ctx context.Context, cmd *command, d service.GrpcDialer, args []string) {
	if len(args) < 4 {
		printHelp(os.Stdout, cmd)
		return
	}
	fileTypeStr := args[0]
	tenant := args[1]
	fileID := args[2]
	fullFileName := filepath.Clean(args[3])
	f, err := os.Open(fullFileName)
	if err != nil {
		printError(os.Stderr, cmd, err)
		return
	}
	client, err := client.NewRepo(ctx, d)
	if err != nil {
		return
	}
	var fileType pb.File_FileType
	switch fileTypeStr {
	case "wasm":
		fileType = pb.File_Wasm
	case "json":
		fileType = pb.File_JsonSchema
	default:
		printHelp(os.Stdout, cmd)
		return
	}
	if err = client.AddFile(context.Background(), tenant, fileID, fileType, f); err != nil {
		printError(os.Stderr, cmd, err)
		return
	}
	fmt.Println("file uploaded successfully")
}
