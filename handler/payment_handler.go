package handler

import (
	"encoding/json"
	"log"
	"net/http"

	data "github.com/form3/data"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	//data base interface
	db data.PaymentProvider
}

func NewApp() *App {
	return &App{}
}

type Response map[string]interface{}

func (a *App) SetMongoProvider(dbConnection *data.MongoDBConn) {
	a.db = &data.PaymentDataBase{dbConnection}
}

// Get list of all payments
func (a *App) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("GetAllPayments  \n")
	payments, err := a.db.ListPayments()
	log.Printf("Payments %+v \n", payments)
	log.Printf("Error %+v \n", err)
	if err != nil {
		SendJson(w, Response{"status": err.Error()})
	} else if len(payments) > 0 {
		SendJson(w, payments)
	} else {
		SendJson(w, Response{"status": "there are not any payments in the collection"})
	}
}

// Get Payment by ID
func (a *App) GetPayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("GetPayment  \n")

	params := mux.Vars(r)
	id := params["id"]

	payment, err := a.db.ListPaymentID(bson.ObjectIdHex(id))
	log.Printf("Payment %+v \n", payment)
	log.Printf("Error %+v \n", err)
	if err != nil {
		SendJson(w, Response{"status": err.Error()})
	} else {
		SendJson(w, payment)
	}
}

// Create payment
func (a *App) CreatePayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("CreatePayment  \n")

	var payment data.Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		SendJsonWithStatus(w, http.StatusBadRequest, Response{"status": err.Error()})
	} else {
		newPayment, err := a.db.CreatePayment(payment)
		log.Printf("Payment %+v \n", newPayment)
		log.Printf("Payment:err %+v \n", err)
		if err != nil {
			SendJson(w, Response{"status": err.Error()})
		} else {
			SendJson(w, newPayment)
		}
	}
}

// Delete payment
func (a *App) DeletePayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("DeletePayment  \n")

	params := mux.Vars(r)
	id := params["id"]
	bsonObjectID := bson.ObjectIdHex(id)
	log.Printf("Params request %+v \n", params)

	err := a.db.RemovePayment(bsonObjectID)
	log.Printf("Error %+v \n", err)
	if err != nil {
		SendJson(w, Response{"status": err.Error()})
	} else {
		SendJson(w, Response{"status": "deleted"})
	}

}

func (a *App) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("UpdatePayment  \n")

	params := mux.Vars(r)
	id := params["id"]
	log.Printf("Params request %+v \n", params)

	var payment data.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		SendJsonWithStatus(w, http.StatusBadRequest, Response{"status": err.Error()})
	} else {
		log.Printf("Payment decoded  %+v \n", payment)
		payment.MongoID = bson.ObjectIdHex(id)
		paymentUpdated, err := a.db.UpdatePayment(payment)
		log.Printf("PaymentUpdated  %+v \n", paymentUpdated)
		log.Printf("Error %+v \n", err)
		if err != nil {
			SendJson(w, Response{"status": err.Error()})
		} else {
			SendJson(w, paymentUpdated)
		}
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
