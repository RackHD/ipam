package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/RackHD/ipam/controllers"
	"github.com/RackHD/ipam/ipam"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
)

var mongo string

func init() {
	flag.StringVar(&mongo, "mongo", "mongodb:27017", "port to connect to mongodb container")
}

func main() {

	// Default to enable mgo debug. Set to false to disable.
	var mgoDebug = true

	if mgoDebug {
		mgo.SetDebug(true)
		var aLogger *log.Logger
		aLogger = log.New(os.Stderr, "", log.LstdFlags)
		mgo.SetLogger(aLogger)
	}

	// Start off with a new mux router.
	router := mux.NewRouter().StrictSlash(true)

	session, err := mgo.Dial(mongo)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer session.Close()

	// Create the IPAM business logic object.
	ipam, err := ipam.NewIpam(session)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Oddly enough don't need to capture the router for it to continue to exist.
	_, err = controllers.NewPoolsController(router, ipam)
	if err != nil {
		log.Fatalf("%s", err)
	}

	_, err = controllers.NewSubnetsController(router, ipam)
	if err != nil {
		log.Fatalf("%s", err)
	}

	_, err = controllers.NewReservationsController(router, ipam)
	if err != nil {
		log.Fatalf("%s", err)
	}

	_, err = controllers.NewLeasesController(router, ipam)
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Show off request logging middleware.
	logged := handlers.LoggingHandler(os.Stdout, router)

	log.Printf("Listening on port 8000...")

	http.ListenAndServe(":8000", logged)
}
