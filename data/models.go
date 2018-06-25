package data

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

const PAYMENT_COLLECTION = "payments"

type Payment struct {
	MongoID        bson.ObjectId `bson:"_id" json:"_id"`
	ID             string        `json:"id,omitempty" bson:"id,omitempty"`
	Type           string        `json:"type,omitempty" bson:"type,omitempty"`
	Version        int           `json:"version" bson:"version"`
	OrganisationID string        `json:"organisation_id,omitempty" bson:"organisation_id,omitempty"`
	Attributes     Attributes    `json:"attributes,omitempty" bson:"attributes,omitempty"`
}

type Attributes struct {
	Amount               float64            `json:"amount,omitempty" bson:"amount,omitempty"`
	BeneficiaryParty     Account            `json:"beneficiary_party,omitempty" bson:"beneficiary_party,omitempty"`
	ChargesInformation   ChargesInformation `json:"charges_information,omitempty" bson:"charges_information,omitempty"`
	Currency             string             `json:"currency,omitempty" bson:"currency,omitempty"`
	DebtorParty          Account            `json:"debtor_party,omitempty" bson:"debtor_party,omitempty"`
	EndToEndReference    string             `json:"end_to_end_reference,omitempty" bson:"end_to_end_reference,omitempty"`
	Fx                   Fx                 `json:"fx,omitempty" bson:"fx,omitempty"`
	NumericReference     string             `json:"numeric_reference,omitempty" bson:"numeric_reference,omitempty"`
	PaymentID            string             `json:"payment_id,omitempty" bson:"payment_id,omitempty"`
	PaymentPurpose       string             `json:"payment_purpose,omitempty" bson:"payment_purpose,omitempty"`
	PaymentScheme        string             `json:"payment_scheme,omitempty" bson:"payment_scheme,omitempty"`
	PaymentType          string             `json:"payment_type,omitempty" bson:"payment_type,omitempty"`
	ProcessingDate       string             `json:"processing_date,omitempty" bson:"processing_date,omitempty"`
	Reference            string             `json:"reference,omitempty" bson:"reference,omitempty"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type,omitempty" bson:"scheme_payment_sub_type,omitempty"`
	SchemePaymentType    string             `json:"scheme_payment_type,omitempty" bson:"scheme_payment_type,omitempty"`
	SponsorParty         Sponsor            `json:"sponsor_party,omitempty" bson:"sponsor_party,omitempty"`
}

type Account struct {
	AccountName       string `json:"account_name,omitempty" bson:"account_name,omitempty"`
	AccountNumber     string `json:"account_number,omitempty" bson:"account_number,omitempty"`
	AccountNumberCode string `json:"account_number_code,omitempty" bson:"account_number_code,omitempty"`
	AccountType       int64  `json:"account_type,omitempty" bson:"account_type,omitempty"`
	Address           string `json:"address,omitempty" bson:"address,omitempty"`
	BankID            string `json:"bank_id,omitempty" bson:"bank_id,omitempty"`
	BankIDCode        string `json:"bank_id_code,omitempty" bson:"bank_id_code,omitempty"`
	Name              string `json:"name,omitempty" bson:"name,omitempty"`
}

type AmountCurrency struct {
	Amount   float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency string  `json:"currency,omitempty" bson:"currency,omitempty"`
}

type ChargesInformation struct {
	BearerCode              string           `json:"bearer_code,omitempty" bson:"bearer_code,omitempty"`
	SenderCharges           []AmountCurrency `json:"sender_charges,omitempty" bson:"sender_charges,omitempty"`
	ReceiverChargesAmount   float64          `json:"receiver_charges_amount,omitempty" bson:"receiver_charges_amount,omitempty"`
	ReceiverChargesCurrency string           `json:"receiver_charges_currency,omitempty" bson:"receiver_charges_currency,omitempty"`
}

type Fx struct {
	ContractReference string  `json:"contract_reference,omitempty" bson:"contract_reference,omitempty"`
	ExchangeRate      string  `json:"exchange_rate,omitempty" bson:"exchange_rate,omitempty"`
	OriginalAmount    float64 `json:"original_amount,omitempty" bson:"original_amount,omitempty"`
	OriginalCurrency  string  `json:"original_currency,omitempty" bson:"original_currency,omitempty"`
}

type Sponsor struct {
	AccountNumber string `json:"account_number,omitempty" bson:"account_number,omitempty"`
	BankID        string `json:"bank_id,omitempty" bson:"bank_id,omitempty"`
	BankIDCode    string `json:"bank_id_code,omitempty" bson:"bank_id_code,omitempty"`
}

type PaymentProvider interface {
	ListPayments() ([]Payment, error)
	ListPaymentID(id bson.ObjectId) (*Payment, error)
	CreatePayment(payment Payment) (*Payment, error)
	RemovePayment(id bson.ObjectId) error
	UpdatePayment(payment Payment) (*Payment, error)
}

type PaymentDataBase struct {
	*MongoDBConn
}

func (p *PaymentDataBase) ListPayments() (payments []Payment, err error) {
	log.Printf("DataBase ListPayments  \n")
	conn := p.GetConn()
	defer conn.Close()
	c := conn.DB(p.db).C(PAYMENT_COLLECTION)
	err = c.Find(bson.M{}).All(&payments)
	return
}

func (p *PaymentDataBase) ListPaymentID(id bson.ObjectId) (payment *Payment, err error) {
	log.Printf("DataBase ListPaymentID  \n")
	conn := p.GetConn()
	defer conn.Close()
	c := conn.DB(p.db).C(PAYMENT_COLLECTION)
	err = c.Find(bson.M{"_id": id}).One(&payment)
	return
}

// Create a payment
func (p *PaymentDataBase) CreatePayment(payment Payment) (*Payment, error) {
	log.Printf("DataBase Create Payment  \n")
	conn := p.GetConn()
	defer conn.Close()
	var err error
	c := conn.DB(p.db).C(PAYMENT_COLLECTION)
	payment.MongoID = bson.NewObjectId()
	err = c.Insert(payment)
	return &payment, err
}

// Delete a payment
func (p *PaymentDataBase) RemovePayment(id bson.ObjectId) error {
	log.Printf("DataBase Remove Payment  \n")
	conn := p.GetConn()
	defer conn.Close()
	c := conn.DB(p.db).C(PAYMENT_COLLECTION)
	err := c.Remove(bson.M{"_id": id})
	return err
}

func (p *PaymentDataBase) UpdatePayment(payment Payment) (*Payment, error) {
	log.Printf("DataBase Update Payment  \n")
	conn := p.GetConn()
	defer conn.Close()
	c := conn.DB(p.db).C(PAYMENT_COLLECTION)

	// update existing object:
	mongoID := payment.MongoID
	err := c.Update(bson.M{"_id": mongoID}, payment)
	log.Printf("Find return update error %+v \n", err)
	if err != nil {
		log.Println("Error could not update:", err.Error())
	} else {
		updatedPayment := &Payment{}
		err = c.Find(bson.M{"_id": mongoID}).One(updatedPayment)
		if err == nil {
			log.Printf("Updated payment in models %+v \n", updatedPayment)
			return updatedPayment, err
		} else {
			log.Printf("Find error %+v \n", err)
		}
	}

	return nil, err
}
