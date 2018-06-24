// form3 project main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	data "github.com/Form3/data"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

/*
type Handler struct {
	*App
	H func(a *App, w *http.ResponseWriter, r *http.Request) error
}*/

type App struct {
	//data base
	db data.PaymentProvider
}

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

type Response map[string]interface{}

/*
// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h *Handler) ServeHTTP(w *http.ResponseWriter, r *http.Request) {
	log.Printf(" SERVER HTTP!  \n")
	log.Printf("h.App  %+v \n", h.App)
	log.Printf("w  %+v \n", w)
	log.Printf("r  %+v \n", r)
	err := h.H(h.App, w, r)
	fmt.Printf("error %+v \n", err)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}*/

// Get list of all payments
func (a *App) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf(" GetAllPayments  \n")
	payments, err := a.db.ListPayments()
	log.Printf(" payments %+v \n", payments) // TODO test when none
	log.Printf(" err %+v \n", err)
	if err != nil {
		SendJson(w, Response{"status": err.Error()})
	} else if len(payments) > 0 {
		SendJson(w, payments)

	} else {
		SendJson(w, Response{"status": "there are not payments in the collection"})
	}
}

// Get Payment by ID
func (a *App) GetPayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf(" GetPayment  \n")

	params := mux.Vars(r)
	id := params["id"]

	payment, err := a.db.ListPaymentID(bson.ObjectIdHex(id))
	log.Printf(" payment %+v \n", payment)
	log.Printf(" err %+v \n", err)
	if err != nil {
		SendJson(w, Response{"status": err.Error()})
	} else {
		SendJson(w, payment)
	}
}

// Create payment
func (a *App) CreatePayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf(" CreatePayment  \n")

	var payment data.Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), 400)
	}

	newPayment, err := a.db.CreatePayment(payment)
	log.Printf(" payment %+v \n", newPayment)
	log.Printf(" err %+v \n", err)
	//SendJson(w, newPayment) // TODO probably handle the error the same as delete and update
	if err != nil {
		SendJson(w, Response{"status": err.Error()})
	} else {
		SendJson(w, newPayment)
	}

}

// Delete payment
func (a *App) DeletePayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf(" DeletePayment  \n")

	params := mux.Vars(r)
	id := params["id"]
	bsonObjectID := bson.ObjectIdHex(id)
	log.Printf("Params r %+v \n", params)
	log.Printf("Body r %+v \n", r.Form)

	err := a.db.RemovePayment(bsonObjectID)
	log.Printf(" err %+v \n", err)
	if err != nil {
		SendJson(w, Response{"status": err.Error()})
	} else {
		SendJson(w, Response{"status": "deleted"})
	}

}

func (a *App) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf(" UpdatePayment  \n")

	params := mux.Vars(r)
	id := params["id"]
	log.Printf("Params r %+v \n", params)
	log.Printf("Body r %+v \n", r.Form)

	var payment data.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), 400)
	}
	log.Printf("payment decoded  %+v \n", payment)
	payment.MongoID = bson.ObjectIdHex(id)
	paymentUpdated, err := a.db.UpdatePayment(payment)
	log.Printf("paymentUpdated  %+v \n", paymentUpdated)
	log.Printf("err %+v \n", err)
	if err != nil {
		SendJson(w, Response{"status": err.Error()})
	} else {
		SendJson(w, paymentUpdated)
	}
}

// Sets the content type to "application/json" and send the data variable in a JSON format. The output is
// a result of `json.Marshal()` call. If the content cannot be marshaled the result is empty string.
func SendJsonWithStatus(w http.ResponseWriter, statusCode int, data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling response", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(result)
	return
}

// Internally calls `SendJsonWithStatus`
func SendJson(w http.ResponseWriter, data interface{}) {
	SendJsonWithStatus(w, http.StatusOK, data)
}

func main() {
	log.Printf("starting\n")
	r := mux.NewRouter()

	dbConn := data.NewMongoDBConn()
	dbConn.Connect("127.0.0.1:27017", "test")
	log.Printf(" dbConn %+v \n", dbConn)

	errInd := dbConn.SetIndex("id", "test", data.PAYMENT_COLLECTION)
	if errInd != nil {
		log.Fatal(errInd)
	}

	app := &App{&data.PaymentDataBase{*dbConn}}

	//r.Handle("/payments", &Handler{app, GetAllPayments}).Methods("GET")
	r.HandleFunc("/payments", app.GetAllPayments).Methods("GET")
	r.HandleFunc("/payments/{id}", app.GetPayment).Methods("GET")
	r.HandleFunc("/payments", app.CreatePayment).Methods("POST")
	r.HandleFunc("/payments/{id}", app.DeletePayment).Methods("DELETE")
	r.HandleFunc("/payments/{id}", app.UpdatePayment).Methods("PUT")

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}

}

// TODO FIRST tests
// also looking into the handlers error
// other todos think
// comments , do READMe and note there PATCH/PUT and mongo and doc
// deploy
