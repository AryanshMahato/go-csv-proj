package main

import (
	"flag"
	"fmt"
	"github.com/AryanshMahato/go-csv-proj/pkg/constants"
	"github.com/AryanshMahato/go-csv-proj/pkg/flagparser"
	"github.com/AryanshMahato/go-csv-proj/pkg/processor"
	"log"
)

var (
	tagFlag string
)

func init() {
	fmt.Println("Searching users...")
	flag.StringVar(&tagFlag, "tag", "", "Search users by tag")
	flag.Parse()
}

func main() {
	flagParser := flagparser.NewAppFlagParser(tagFlag)
	tags := flagParser.GetTags()

	searchUsersProcessor := processor.NewSearchUsersProcessor(tags, constants.CsvFileDir)

	err := searchUsersProcessor.Process()
	if err != nil {
		log.Fatal(err)
	}
}
