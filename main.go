// form3 project main.go
package main

import (
	"log"
	"net/http"
	"os"

	data "github.com/form3/data"
	handler "github.com/form3/handler"
	"github.com/gorilla/mux"
)

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func main() {
	log.Printf("starting\n")
	r := mux.NewRouter()

	dbConn := data.NewMongoDBConn()
	host := getEnv("MONGO_URI", "localhost:27017")
	log.Printf("Host %+v \n", host)
	dbConn.Connect(host, "form3_db")
	log.Printf("DB Connection %+v \n", dbConn)

	errInd := dbConn.SetIndex("id", "form3_db", data.PAYMENT_COLLECTION)
	if errInd != nil {
		log.Fatal(errInd)
	}

	app := handler.NewApp()
	app.SetMongoProvider(dbConn)

	r.HandleFunc("/payments", app.GetAllPayments).Methods("GET")
	r.HandleFunc("/payments/{id}", app.GetPayment).Methods("GET")
	r.HandleFunc("/payments", app.CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", app.DeletePayment).Methods("DELETE")
	r.HandleFunc("/payments/{id}", app.UpdatePayment).Methods("PUT")

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}

}
