package projects

import (
	"context"
	"fmt"
	"os"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
)

func Search() []string {
	ctx := context.Background()
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		fmt.Println("resourcemanager client", err.Error())
		os.Exit(1)
	}
	defer c.Close()

	req := &resourcemanagerpb.SearchProjectsRequest{
		// See https://pkg.go.dev/cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb#SearchProjectsRequest.
		// Leaving empty for now
	}

	var projectLIst []string
	for project, err := range c.SearchProjects(ctx, req).All() {
		if err != nil {
			fmt.Println("resource manager search project:", project, err.Error())
			os.Exit(1)
		}
		projectLIst = append(projectLIst, project.ProjectId)
	}
	return projectLIst
}
