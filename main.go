package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/moncho/gh-download-count/github"
	"github.com/moncho/gh-download-count/output"
)

var detail = flag.Bool("detail", false, "shows download count per release")
var json = flag.Bool("json", false, "generates a json with the download count")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 || len(args) > 2 {
		fmt.Printf("Got %d args, was expecting 1 or 2.", len(args))
		return
	}

	owner, repo := parseArgs(args)
	p, err := github.NewProject(owner, repo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var w output.ProjectWriter
	if *detail {
		w = output.ASCIITableWriter{}
	} else if *json {
		w = output.JSONWriter{}
	} else {
		w = output.DefaultProjectWriter{}
	}
	w.Write(os.Stdout, p)
}

func parseArgs(args []string) (string, string) {
	argsCount := len(args)
	if argsCount == 1 {
		sep := strings.Index(args[0], "/")
		return args[0][0:sep], args[0][sep+1:]
	}
	return args[0], args[1]
}
