package app

import (
	"log"
	"net/http"

	"github.com/Banking-System/domain"
	"github.com/Banking-System/service"
	"github.com/gorilla/mux"
)

func Start() {
	//define routes

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet) //http中的约束条件

	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// 	//create request multiplexer
	// 	http.HandleFunc("/customers", getAllCustomers)
	// 	//staring server
	// 	//Fatal 系列函数用来写日志消息，然后使用 os.Exit(1)终止程序
	log.Fatal(http.ListenAndServe("localhost:8000", router))
	//
}
