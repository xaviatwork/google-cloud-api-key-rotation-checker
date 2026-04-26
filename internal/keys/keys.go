package keys

import (
	"context"
	"fmt"
	"os"
	"time"

	apikeys "cloud.google.com/go/apikeys/apiv2"
	"cloud.google.com/go/apikeys/apiv2/apikeyspb"
)

type Key struct {
	CreateTime  time.Time `json:"create_time,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Name        string    `json:"name,omitempty"`
	ProjectId   string    `json:"project_id,omitempty"`
}

func (k Key) String() string {
	const hoursDay = 24
	return fmt.Sprintf("%s on project %s (created: %s, %.0f days ago)", k.DisplayName, k.ProjectId, k.CreateTime.Format(time.RFC1123), time.Since(k.CreateTime).Hours()/hoursDay)
}

func List(projectid string) []*Key {
	ctx := context.Background()
	c, err := apikeys.NewClient(ctx)
	if err != nil {
		fmt.Printf("new client error: %s\n", err.Error())
		os.Exit(1)
	}
	defer c.Close()

	req := &apikeyspb.ListKeysRequest{
		// See https://pkg.go.dev/cloud.google.com/go/apikeys/apiv2/apikeyspb#ListKeysRequest.
		Parent: fmt.Sprintf("projects/%s/locations/global", projectid),
	}

	var keys []*Key
	for k, err := range c.ListKeys(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
			fmt.Printf("list keys error: %s\n", err.Error())
			os.Exit(1)
		}
		key := &Key{
			Name:        k.Name,
			DisplayName: k.DisplayName,
			CreateTime:  k.CreateTime.AsTime(),
			ProjectId:   projectid,
		}
		keys = append(keys, key)
	}

	return keys
}
