package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/andrescosta/goico/pkg/ioutl"
	"github.com/andrescosta/goico/pkg/yamlutl"
	"github.com/andrescosta/jobico/api/pkg/remote"
	pb "github.com/andrescosta/jobico/api/types"
)

var cmdRollback = &command{
	name:      "rollback",
	usageLine: `cli rollabck < deployment file >.yaml`,
	short:     "rollaback the specified by the deployment file ",
	long:      `Rollback the file`,
}

func initRollback() {
	cmdRollback.flag = *flag.NewFlagSet("rollback", flag.ContinueOnError)
	cmdRollback.run = runRollback
	cmdRollback.flag.Usage = func() {}

}

func runRollback(ctx context.Context, cmd *command, args []string) {
	if len(args) < 1 {
		printHelp(os.Stdout, cmd)
		return
	}
	file := args[0]
	e, err := ioutl.FileExists(file)
	if err != nil {
		printError(os.Stdout, cmd, err)
		return
	}
	if !e {
		fmt.Printf("file %s does not exist.", file)
		return
	}
	f := pb.JobPackage{}
	if err = yamlutl.DecodeFile(file, &f); err != nil {
		printError(os.Stderr, cmd, err)
		return
	}
	client, err := remote.NewControlClient(ctx)
	if err != nil {
		printError(os.Stderr, cmd, err)
		return
	}

	p, err := client.GetPackage(ctx, f.TenantId, &f.ID)
	if err != nil {
		printError(os.Stderr, cmd, err)
		return
	}
	if len(p) < 1 {
		printError(os.Stderr, cmd, fmt.Errorf("package %s does not exist", f.ID))
		return
	}

	err = client.DeletePackage(context.Background(), &f)
	if err != nil {
		printError(os.Stderr, cmd, err)
		return
	}
	fmt.Println("The package was deleted.")
}
