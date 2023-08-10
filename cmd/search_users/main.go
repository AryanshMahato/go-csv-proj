package main

import (
	"flag"
	"fmt"
	"github.com/AryanshMahato/go-csv-proj/pkg/flagparser"
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
	fmt.Println("tags", flagParser.GetTags())
}
