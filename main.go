package main

import (
	"context"
	"fmt"
	"os"

	apikeys "cloud.google.com/go/apikeys/apiv2"
	apikeyspb "cloud.google.com/go/apikeys/apiv2/apikeyspb"
)

func main() {
	ctx := context.Background()
	c, err := apikeys.NewClient(ctx)
	if err != nil {
		fmt.Printf("new client error: %s\n", err.Error())
		os.Exit(1)
	}
	defer c.Close()

	req := &apikeyspb.ListKeysRequest{
		// See https://pkg.go.dev/cloud.google.com/go/apikeys/apiv2/apikeyspb#ListKeysRequest.
		Parent: fmt.Sprintf("projects/%s/locations/global", os.Getenv("PROJECTID")),
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
