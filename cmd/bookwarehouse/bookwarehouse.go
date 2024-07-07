// Package main implements the bookwarehouse application
// This create a service which has both inbound as well as outbound service policies
// i.e. bookbuyer makes a GET call to bookstore, bookstore makes a POST call to bookwarehouse
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/draychev/go-toolbox/pkg/logger"
	"github.com/draychev/meshy-bookstore/pkg/common"
	"github.com/draychev/meshy-bookstore/pkg/database"
)

var (
	log      = logger.NewPretty("bookwarehouse")
	identity = flag.String("ident", "unidentified", "the identity of the container where this demo app is running (VM, K8s, etc)")
	port     = flag.Int("port", 14001, "port on which this app is listening for incoming HTTP")
	db       *gorm.DB
)

func getIdentity() string {
	ident := os.Getenv("IDENTITY")
	if ident == "" {
		if identity != nil {
			ident = *identity
		}
	}
	return ident
}

func getBooksStockedRecord() database.Record {
	var record database.Record
	db.Where(&database.Record{Key: database.KeyTotalBooks}).First(&record)
	return record
}

func setHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(common.IdentityHeader, getIdentity())

	if r == nil {
		return
	}

	for _, header := range common.GetTracingHeaderKeys() {
		if v := r.Header.Get(header); v != "" {
			w.Header().Set(header, v)
		}
	}
}

// restockBooks decreases the balance of the given bookwarehouse account.
func restockBooks(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, r)
	var numberOfBooks int
	err := json.NewDecoder(r.Body).Decode(&numberOfBooks)
	if err != nil {
		log.Error().Err(err).Msg("Could not decode request body")
		numberOfBooks = 0
	}

	record := getBooksStockedRecord()
	record.ValueInt += int64(numberOfBooks)
	totalBooks := int(record.ValueInt)
	db.Save(record)

	_, _ = w.Write([]byte(fmt.Sprintf("{\"restocked\":%d}", numberOfBooks)))
	log.Info().Msgf("Restocking bookstore with %d new books; Total so far: %d", numberOfBooks, totalBooks)
	if totalBooks >= 3 {
		fmt.Println(common.Success)
	}
}

func main() {
	flag.Parse()

	db = database.Init()

	//initializing router
	router := mux.NewRouter()

	router.HandleFunc(fmt.Sprintf("/%s", common.RestockWarehouseURL), restockBooks).Methods("POST")
	router.HandleFunc("/", restockBooks).Methods("POST")
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	log.Info().Msgf("Starting BookWarehouse HTTP server on port %d", *port)
	//#nosec G114: Use of net/http serve function that has no support for setting timeouts
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
	log.Fatal().Err(err).Msgf("Failed to start BookWarehouse HTTP server on port %d", *port)
}
