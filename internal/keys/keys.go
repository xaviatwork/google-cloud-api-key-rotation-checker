package keys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	apikeys "cloud.google.com/go/apikeys/apiv2"
	"cloud.google.com/go/apikeys/apiv2/apikeyspb"
)

type Options struct {
	MaxDays int
	Redact  bool
	Format  string
}

type Key struct {
	CreateTime  time.Time `json:"create_time,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Name        string    `json:"name,omitempty"`
	ProjectId   string    `json:"project_id,omitempty"`
}

func (k Key) String() string {
	return fmt.Sprintf("%s on project %s (created: %s, %d days ago)", k.DisplayName, k.ProjectId, k.CreateTime.Format(time.RFC1123), daysSinceCreated(k.CreateTime))
}

func (k Key) NeedsToBeRotated(options Options) bool {
	if daysSinceCreated(k.CreateTime) <= options.MaxDays {
		return false
	}
	return true
}

func daysSinceCreated(c time.Time) int {
	const hoursDay = 24
	return int(time.Since(c).Hours() / hoursDay)
}

func redact(s string, restr string, mask string) string {
	re, err := regexp.Compile(restr)
	if err != nil {
		log.Println("redact", err.Error())
	}
	return re.ReplaceAllString(s, mask)
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

func Display(keylist []*Key, options Options) {
	const (
		warningIcon string = "⚠️"
		okIcon      string = "✅"
	)
	switch options.Format {
	case "json":
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(&keylist); err != nil {
			log.Println("json encoding", err.Error())
			os.Exit(1)
		}
		fmt.Println(buf.String())
	default:
		icon := okIcon
		for _, k := range keylist {
			if k.NeedsToBeRotated(options) {
				icon = warningIcon
			}
			// Redact as many fields as we want from the API Key
			if options.Redact {
				k.ProjectId = redact(k.ProjectId, "[a-zA-Z]", "░")
			}
			fmt.Println(icon, k.String())
		}
	}
}
