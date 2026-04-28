package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xaviatwork/api-key-checker/internal/keys"
)

func main() {
	fs := flag.NewFlagSet("apikeycheck", flag.ExitOnError)
	projectId := fs.String("project", "", "Check API keys in the project identified by ProjectId")
	maxDays := fs.Int("max-days", 90, "Max. number of days before API keys should be rotated")
	redact := fs.Bool("redact", false, "Obfuscate information when displaying the Key")
	format := fs.String("format", "", "Output format for displaying the API keys")
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Printf("error parsing flags: %s\n", err.Error())
		os.Exit(1)
	}

	options := keys.Options{
		MaxDays: *maxDays,
		Redact:  *redact,
		Format:  strings.ToLower(*format),
	}

	if *projectId == "" {
		fmt.Println("projectId cannot be empty")
		os.Exit(1)
	}

	projectIds := []string{*projectId}

	// get API key list
	var keylist []*keys.Key
	for _, p := range projectIds {
		kk := keys.List(p)
		keylist = append(keylist, kk...)
	}

	// display API keys
	keys.Display(keylist, options)
}
