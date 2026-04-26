package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xaviatwork/api-key-checker/internal/keys"
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

	projectIds := []string{*projectId}

	// get API key list
	var keylist []*keys.Key
	for _, p := range projectIds {
		kk := keys.List(p)
		keylist = append(keylist, kk...)
	}

	// display API keys
	for _, k := range keylist {
		fmt.Println(k.String())
	}
}
