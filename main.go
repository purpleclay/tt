package main

import (
	"fmt"
	"log"
	"strings"

	git "github.com/purpleclay/gitz"
)

func main() {
	gitc, err := git.NewClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	tags, err := gitc.Tags(git.WithShellGlob("*.*.*"),
		git.WithSortBy(git.VersionDesc),
		git.WithCount(1))
	if err != nil {
		log.Fatal(err.Error())
	}

	if len(tags) == 0 {
		return
	}

	majorPos := strings.Index(tags[0], ".")
	minorPos := strings.LastIndex(tags[0], ".")

	fmt.Printf("%s,%s,%s", tags[0], tags[0][:majorPos], tags[0][:minorPos])
}
