package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Banking-System/domain"
	"github.com/Banking-System/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variale not defined...")
	}

}
func Start() {
	//define routes

	sanityCheck()

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	dbClient := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHander{service.NewAccountService(accountRepositoryDb)}

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb(dbClient))}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet) //http中的约束条件
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// 	//create request multiplexer
	// 	http.HandleFunc("/customers", getAllCustomers)
	// 	//staring server
	// 	//Fatal 系列函数用来写日志消息，然后使用 os.Exit(1)终止程序
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// log.Fatal(http.ListenAndServe("localhost:8000", router))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
	// 取消硬编码，加入函数
	//
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf(("%s:%s@tcp(%s:%s)/%s"), dbUser, dbPasswd, dbAddr, dbPort, dbName)

	// dataSource := fmt.Sprintf("root:pan@tcp(localhost:3306)/banking")

	// client, err := sqlx.Open("mysql", "root:pan@tcp(localhost:3306)/banking")
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
