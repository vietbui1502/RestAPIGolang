package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vietbui1502/RestAPIGolang/domain"
	"github.com/vietbui1502/RestAPIGolang/service"
)

func Start() {

	//router := http.NewServeMux()
	router := mux.NewRouter()

	//Create customer service
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//Define route
	router.HandleFunc("/customer", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//Starting server
	log.Fatal(http.ListenAndServe(":8088", router))
}
