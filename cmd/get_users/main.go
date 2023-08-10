package main

import (
	"fmt"
	"github.com/AryanshMahato/go-csv-proj/pkg/api"
	"github.com/AryanshMahato/go-csv-proj/pkg/config"
	"github.com/AryanshMahato/go-csv-proj/pkg/constants"
	"github.com/AryanshMahato/go-csv-proj/pkg/processor"
	"github.com/AryanshMahato/go-csv-proj/pkg/saver"
	"log"
	"net/http"
)

func init() {
	fmt.Println("Getting users...")
}

func main() {
	appConfig := config.NewAppConfig()

	appApi := api.NewAppApi(&http.Client{}, appConfig)
	fileSaver := saver.NewFileSaver(constants.CsvFileDir)

	getUsersProcessor := processor.NewGetUsersProcessor(appApi, fileSaver)

	err := getUsersProcessor.Process()
	if err != nil {
		log.Fatal(err)
	}
}
