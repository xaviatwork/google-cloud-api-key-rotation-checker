package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xaviatwork/api-key-checker/internal/keys"
	"github.com/xaviatwork/api-key-checker/internal/projects"
)

func main() {
	fs := flag.NewFlagSet("apikeycheck", flag.ExitOnError)
	projectId := fs.String("project", "", "[optional] Check API keys in the project identified by ProjectId. By default, check all projects the user has access to.")
	maxDays := fs.Int("max-days", 90, "Max. number of days before API keys should be rotated.")
	format := fs.String("format", "", "Output format for displaying the API keys. Default (empty) is no format. Supported formats: JSON, CSV")
	rotate := fs.Bool("rotate", false, "Display only API keys older than 'max-days'.")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Printf("error parsing flags: %s\n", err.Error())
		os.Exit(1)
	}

	options := keys.Options{
		MaxDays: *maxDays,
		Format:  strings.ToLower(*format),
		Rotate:  *rotate,
	}

	projectList := []string{*projectId}
	if *projectId == "" {
		projectList = projects.Search()
	}

	// get API key list
	var keylist []*keys.Key
	numProjects := len(projectList)
	i := 1
	for _, p := range projectList {
		fmt.Fprintf(os.Stderr, "🕵️‍♂️ [%d/%d] checking API keys on project %s ...", i, numProjects, p)
		kk := keys.List(p)
		if len(kk) == 0 {
			fmt.Fprintf(os.Stderr, " found 0.\n")
		} else {
			fmt.Fprintf(os.Stderr, " found %d 🔑.\n", len(kk))
		}

		keylist = append(keylist, kk...)
		i = i + 1
	}

	// filter API keys
	keylist = keys.Filter(keylist, options)
	// display API keys
	keys.Display(keylist, options)
}
