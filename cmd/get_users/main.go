package main

import (
	"fmt"
	"github.com/AryanshMahato/go-csv-proj/pkg/api"
	"github.com/AryanshMahato/go-csv-proj/pkg/config"
	"net/http"
)

func init() {
	fmt.Println("Getting users...")
}

func main() {
	appConfig := config.NewAppConfig()

	api.NewAppApi(&http.Client{}, appConfig)
}
