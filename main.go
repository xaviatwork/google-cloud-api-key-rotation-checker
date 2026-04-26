package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	apikeys "cloud.google.com/go/apikeys/apiv2"
	apikeyspb "cloud.google.com/go/apikeys/apiv2/apikeyspb"
)

func main() {
	fs := flag.NewFlagSet("apikeycheck", flag.ExitOnError)
	projectId := fs.String("project", "", "Check API keys in the project identified by ProjectId")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Printf("error parsing flags: %s\n", err.Error())
		os.Exit(1)
	}

	if *projectId == "" {
		fmt.Println("projectId cannot be empty")
		os.Exit(1)
	}

	ctx := context.Background()
	c, err := apikeys.NewClient(ctx)
	if err != nil {
		fmt.Printf("new client error: %s\n", err.Error())
		os.Exit(1)
	}
	defer c.Close()

	req := &apikeyspb.ListKeysRequest{
		// See https://pkg.go.dev/cloud.google.com/go/apikeys/apiv2/apikeyspb#ListKeysRequest.
		Parent: fmt.Sprintf("projects/%s/locations/global", *projectId),
	}
	for k, err := range c.ListKeys(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
			fmt.Printf("list keys error: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println(k.DisplayName)
	}
}
