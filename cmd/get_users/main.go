package main

import (
	"fmt"
	"github.com/AryanshMahato/go-csv-proj/pkg/config"
	"github.com/AryanshMahato/go-csv-proj/pkg/constants"
	"github.com/AryanshMahato/go-csv-proj/pkg/getter"
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

	apiGetter := getter.NewApiGetter(&http.Client{}, appConfig)
	fileSaver := saver.NewFileSaver(constants.CsvFileDir)

	saveUsersProcessor := processor.NewSaveUsersProcessor(apiGetter, fileSaver)

	err := saveUsersProcessor.Process()
	if err != nil {
		log.Fatal(err)
	}
}
