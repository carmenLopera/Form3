package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	data "github.com/form3/data"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

const DB_ERROR = "database error"

type mockDB struct {
	testCaseEmpty   bool
	testCaseDbError bool
}

func (mdb *mockDB) ListPayments() (payments []data.Payment, err error) {
	if mdb.testCaseDbError != true {
		if mdb.testCaseEmpty != true {
			payments = []data.Payment{
				data.Payment{
					MongoID:        bson.ObjectIdHex("5b290f5b802b0f1479000002"),
					ID:             "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
					Type:           "Payment",
					Version:        0,
					OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
					Attributes: data.Attributes{
						Amount: 0,
						BeneficiaryParty: data.Account{
							AccountName:       "W Owens",
							AccountNumber:     "31926819",
							AccountNumberCode: "BBAN",
							AccountType:       0,
							Address:           "1 The Beneficiary Localtown SE2",
							BankID:            "403000",
							BankIDCode:        "GBDSC Name:Wilfred Jeremiah Owens",
						},
						ChargesInformation: data.ChargesInformation{
							BearerCode: "SHAR",
							SenderCharges: []data.AmountCurrency{
								data.AmountCurrency{Amount: 0, Currency: "GBP"},
								data.AmountCurrency{Amount: 0, Currency: "GBP"},
							},
							ReceiverChargesAmount:   0,
							ReceiverChargesCurrency: "USD",
						},
						DebtorParty: data.Account{
							AccountName:       "EJ Brown Black",
							AccountNumber:     "GB29XABC10161234567801",
							AccountNumberCode: "IBAN",
							AccountType:       0,
							Address:           "10 Debtor Crescent Sourcetown NE1",
							BankID:            "203301",
							BankIDCode:        "GBDSC",
							Name:              "Emelia Jane Brown",
						},
						EndToEndReference: "Wil piano Jan",
						Fx: data.Fx{
							ContractReference: "FX123",
							ExchangeRate:      "2.00000",
							OriginalAmount:    0,
							OriginalCurrency:  "USD",
						},
						NumericReference:     "1002001",
						PaymentID:            "123456789012345678",
						PaymentPurpose:       "Paying for goods/services",
						PaymentScheme:        "FPS",
						PaymentType:          "Credit",
						ProcessingDate:       "2017-01-18",
						Reference:            "Payment for Em's piano lessons",
						SchemePaymentSubType: "InternetBanking",
						SchemePaymentType:    "ImmediatePayment",
						SponsorParty: data.Sponsor{
							AccountNumber: "56781234",
							BankID:        "123123",
							BankIDCode:    "GBDSC",
						},
					},
				},
				data.Payment{
					MongoID:        bson.ObjectIdHex("5b290f5b802b0f1479000003"),
					ID:             "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
					Type:           "Payment",
					Version:        0,
					OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
					Attributes: data.Attributes{
						Amount: 0,
						BeneficiaryParty: data.Account{
							AccountName:       "W Owens",
							AccountNumber:     "31926819",
							AccountNumberCode: "BBAN",
							AccountType:       0,
							Address:           "1 The Beneficiary Localtown SE2",
							BankID:            "403000",
							BankIDCode:        "GBDSC Name:Wilfred Jeremiah Owens",
						},
						ChargesInformation: data.ChargesInformation{
							BearerCode: "SHAR",
							SenderCharges: []data.AmountCurrency{
								data.AmountCurrency{Amount: 0, Currency: "GBP"},
								data.AmountCurrency{Amount: 0, Currency: "GBP"},
							},
							ReceiverChargesAmount:   0,
							ReceiverChargesCurrency: "USD",
						},
						DebtorParty: data.Account{
							AccountName:       "EJ Brown Black",
							AccountNumber:     "GB29XABC10161234567801",
							AccountNumberCode: "IBAN",
							AccountType:       0,
							Address:           "10 Debtor Crescent Sourcetown NE1",
							BankID:            "203301",
							BankIDCode:        "GBDSC",
							Name:              "Emelia Jane Brown",
						},
						EndToEndReference: "Wil piano Jan",
						Fx: data.Fx{
							ContractReference: "FX123",
							ExchangeRate:      "2.00000",
							OriginalAmount:    0,
							OriginalCurrency:  "USD",
						},
						NumericReference:     "1002001",
						PaymentID:            "123456789012345678",
						PaymentPurpose:       "Paying for goods/services",
						PaymentScheme:        "FPS",
						PaymentType:          "Credit",
						ProcessingDate:       "2017-01-18",
						Reference:            "Payment for Em's piano lessons",
						SchemePaymentSubType: "InternetBanking",
						SchemePaymentType:    "ImmediatePayment",
						SponsorParty: data.Sponsor{
							AccountNumber: "56781234",
							BankID:        "123123",
							BankIDCode:    "GBDSC",
						},
					},
				},
			}

		}
		return payments, nil
	} else {
		return nil, errors.New(DB_ERROR)
	}
}

func TestGetAllPayments(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments", &bytes.Buffer{})

	app := &App{db: &mockDB{}}
	http.HandlerFunc(app.GetAllPayments).ServeHTTP(rec, req)

	expected := []data.Payment{
		data.Payment{
			MongoID:        bson.ObjectIdHex("5b290f5b802b0f1479000002"),
			ID:             "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
			Type:           "Payment",
			Version:        0,
			OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
			Attributes: data.Attributes{
				Amount: 0,
				BeneficiaryParty: data.Account{
					AccountName:       "W Owens",
					AccountNumber:     "31926819",
					AccountNumberCode: "BBAN",
					AccountType:       0,
					Address:           "1 The Beneficiary Localtown SE2",
					BankID:            "403000",
					BankIDCode:        "GBDSC Name:Wilfred Jeremiah Owens",
				},
				ChargesInformation: data.ChargesInformation{
					BearerCode: "SHAR",
					SenderCharges: []data.AmountCurrency{
						data.AmountCurrency{Amount: 0, Currency: "GBP"},
						data.AmountCurrency{Amount: 0, Currency: "GBP"},
					},
					ReceiverChargesAmount:   0,
					ReceiverChargesCurrency: "USD",
				},
				DebtorParty: data.Account{
					AccountName:       "EJ Brown Black",
					AccountNumber:     "GB29XABC10161234567801",
					AccountNumberCode: "IBAN",
					AccountType:       0,
					Address:           "10 Debtor Crescent Sourcetown NE1",
					BankID:            "203301",
					BankIDCode:        "GBDSC",
					Name:              "Emelia Jane Brown",
				},
				EndToEndReference: "Wil piano Jan",
				Fx: data.Fx{
					ContractReference: "FX123",
					ExchangeRate:      "2.00000",
					OriginalAmount:    0,
					OriginalCurrency:  "USD",
				},
				NumericReference:     "1002001",
				PaymentID:            "123456789012345678",
				PaymentPurpose:       "Paying for goods/services",
				PaymentScheme:        "FPS",
				PaymentType:          "Credit",
				ProcessingDate:       "2017-01-18",
				Reference:            "Payment for Em's piano lessons",
				SchemePaymentSubType: "InternetBanking",
				SchemePaymentType:    "ImmediatePayment",
				SponsorParty: data.Sponsor{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
			},
		},
		data.Payment{
			MongoID:        bson.ObjectIdHex("5b290f5b802b0f1479000003"),
			ID:             "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
			Type:           "Payment",
			Version:        0,
			OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
			Attributes: data.Attributes{
				Amount: 0,
				BeneficiaryParty: data.Account{
					AccountName:       "W Owens",
					AccountNumber:     "31926819",
					AccountNumberCode: "BBAN",
					AccountType:       0,
					Address:           "1 The Beneficiary Localtown SE2",
					BankID:            "403000",
					BankIDCode:        "GBDSC Name:Wilfred Jeremiah Owens",
				},
				ChargesInformation: data.ChargesInformation{
					BearerCode: "SHAR",
					SenderCharges: []data.AmountCurrency{
						data.AmountCurrency{Amount: 0, Currency: "GBP"},
						data.AmountCurrency{Amount: 0, Currency: "GBP"},
					},
					ReceiverChargesAmount:   0,
					ReceiverChargesCurrency: "USD",
				},
				DebtorParty: data.Account{
					AccountName:       "EJ Brown Black",
					AccountNumber:     "GB29XABC10161234567801",
					AccountNumberCode: "IBAN",
					AccountType:       0,
					Address:           "10 Debtor Crescent Sourcetown NE1",
					BankID:            "203301",
					BankIDCode:        "GBDSC",
					Name:              "Emelia Jane Brown",
				},
				EndToEndReference: "Wil piano Jan",
				Fx: data.Fx{
					ContractReference: "FX123",
					ExchangeRate:      "2.00000",
					OriginalAmount:    0,
					OriginalCurrency:  "USD",
				},
				NumericReference:     "1002001",
				PaymentID:            "123456789012345678",
				PaymentPurpose:       "Paying for goods/services",
				PaymentScheme:        "FPS",
				PaymentType:          "Credit",
				ProcessingDate:       "2017-01-18",
				Reference:            "Payment for Em's piano lessons",
				SchemePaymentSubType: "InternetBanking",
				SchemePaymentType:    "ImmediatePayment",
				SponsorParty: data.Sponsor{
					AccountNumber: "56781234",
					BankID:        "123123",
					BankIDCode:    "GBDSC",
				},
			},
		},
	}

	var result []data.Payment
	err := json.Unmarshal(rec.Body.Bytes(), &result)

	if err != nil {
		t.Errorf("Didn't expect error %v", err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected:\n%+v \nand instead got:\n%+v", expected, result)
	}
}

func TestGetAllPaymentsEmpty(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments", &bytes.Buffer{})

	app := &App{db: &mockDB{testCaseEmpty: true}}
	http.HandlerFunc(app.GetAllPayments).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"there are not any payments in the collection"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestGetAllPaymentsDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments", &bytes.Buffer{})

	app := &App{db: &mockDB{testCaseDbError: true}}
	http.HandlerFunc(app.GetAllPayments).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"database error"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func (mdb *mockDB) ListPaymentID(id bson.ObjectId) (payment *data.Payment, err error) {
	// in database we have
	// MongoID: bson.ObjectIdHex("5b290f5b802b0f1479000002"),
	fmt.Printf("Mock ListPaymentID id %v", id)
	if mdb.testCaseDbError != true {

		if id == bson.ObjectIdHex("5b290f5b802b0f1479000002") {
			fmt.Printf("ids equal")
			payment = &data.Payment{
				MongoID:        bson.ObjectIdHex("5b290f5b802b0f1479000002"),
				ID:             "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
				Type:           "Payment",
				Version:        0,
				OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
				Attributes: data.Attributes{
					Amount: 0,
					BeneficiaryParty: data.Account{
						AccountName:       "W Owens",
						AccountNumber:     "31926819",
						AccountNumberCode: "BBAN",
						AccountType:       0,
						Address:           "1 The Beneficiary Localtown SE2",
						BankID:            "403000",
						BankIDCode:        "GBDSC Name:Wilfred Jeremiah Owens",
					},
					ChargesInformation: data.ChargesInformation{
						BearerCode: "SHAR",
						SenderCharges: []data.AmountCurrency{
							data.AmountCurrency{Amount: 0, Currency: "GBP"},
							data.AmountCurrency{Amount: 0, Currency: "GBP"},
						},
						ReceiverChargesAmount:   0,
						ReceiverChargesCurrency: "USD",
					},
					DebtorParty: data.Account{
						AccountName:       "EJ Brown Black",
						AccountNumber:     "GB29XABC10161234567801",
						AccountNumberCode: "IBAN",
						AccountType:       0,
						Address:           "10 Debtor Crescent Sourcetown NE1",
						BankID:            "203301",
						BankIDCode:        "GBDSC",
						Name:              "Emelia Jane Brown",
					},
					EndToEndReference: "Wil piano Jan",
					Fx: data.Fx{
						ContractReference: "FX123",
						ExchangeRate:      "2.00000",
						OriginalAmount:    0,
						OriginalCurrency:  "USD",
					},
					NumericReference:     "1002001",
					PaymentID:            "123456789012345678",
					PaymentPurpose:       "Paying for goods/services",
					PaymentScheme:        "FPS",
					PaymentType:          "Credit",
					ProcessingDate:       "2017-01-18",
					Reference:            "Payment for Em's piano lessons",
					SchemePaymentSubType: "InternetBanking",
					SchemePaymentType:    "ImmediatePayment",
					SponsorParty: data.Sponsor{
						AccountNumber: "56781234",
						BankID:        "123123",
						BankIDCode:    "GBDSC",
					},
				},
			}

		} else {
			fmt.Printf("not equal")
			return nil, errors.New("not found")
		}
		return payment, nil
	} else {
		return nil, errors.New(DB_ERROR)
	}

}

func TestGetPaymentIDExists(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000002", &bytes.Buffer{})

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{}}
	router.HandleFunc("/payments/{id}", app.GetPayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := data.Payment{
		MongoID:        bson.ObjectIdHex("5b290f5b802b0f1479000002"),
		ID:             "4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
		Type:           "Payment",
		Version:        0,
		OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Attributes: data.Attributes{
			Amount: 0,
			BeneficiaryParty: data.Account{
				AccountName:       "W Owens",
				AccountNumber:     "31926819",
				AccountNumberCode: "BBAN",
				AccountType:       0,
				Address:           "1 The Beneficiary Localtown SE2",
				BankID:            "403000",
				BankIDCode:        "GBDSC Name:Wilfred Jeremiah Owens",
			},
			ChargesInformation: data.ChargesInformation{
				BearerCode: "SHAR",
				SenderCharges: []data.AmountCurrency{
					data.AmountCurrency{Amount: 0, Currency: "GBP"},
					data.AmountCurrency{Amount: 0, Currency: "GBP"},
				},
				ReceiverChargesAmount:   0,
				ReceiverChargesCurrency: "USD",
			},
			DebtorParty: data.Account{
				AccountName:       "EJ Brown Black",
				AccountNumber:     "GB29XABC10161234567801",
				AccountNumberCode: "IBAN",
				AccountType:       0,
				Address:           "10 Debtor Crescent Sourcetown NE1",
				BankID:            "203301",
				BankIDCode:        "GBDSC",
				Name:              "Emelia Jane Brown",
			},
			EndToEndReference: "Wil piano Jan",
			Fx: data.Fx{
				ContractReference: "FX123",
				ExchangeRate:      "2.00000",
				OriginalAmount:    0,
				OriginalCurrency:  "USD",
			},
			NumericReference:     "1002001",
			PaymentID:            "123456789012345678",
			PaymentPurpose:       "Paying for goods/services",
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			ProcessingDate:       "2017-01-18",
			Reference:            "Payment for Em's piano lessons",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "ImmediatePayment",
			SponsorParty: data.Sponsor{
				AccountNumber: "56781234",
				BankID:        "123123",
				BankIDCode:    "GBDSC",
			},
		},
	}
	var result data.Payment
	err := json.Unmarshal(rec.Body.Bytes(), &result)

	if err != nil {
		t.Errorf("Didn't expect error %v", err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected:\n%+v \nand instead got:\n%+v", expected, result)
	}
}

func TestGetPaymentIDDoesNotExist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000003", &bytes.Buffer{})

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{}}
	router.HandleFunc("/payments/{id}", app.GetPayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"not found"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestGetPaymentIDDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000003", &bytes.Buffer{})

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{testCaseDbError: true}}
	router.HandleFunc("/payments/{id}", app.GetPayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"database error"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func (mdb *mockDB) CreatePayment(payment data.Payment) (*data.Payment, error) {
	if mdb.testCaseDbError != true {
		payment.MongoID = bson.ObjectIdHex("5b2ce1c5c089711b0e3bc2fa")
		return &payment, nil
	} else {
		return nil, errors.New(DB_ERROR)
	}
}

func TestCreatePayment(t *testing.T) {
	rec := httptest.NewRecorder()
	newPayment := []byte(`
		 {
        "id": "new_payment_test",
        "type": "Payment",
        "version": 1,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "currency": "GBP"
                    },
                    {
                        "currency": "USD"
                    }
                ],
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    }`)
	req, _ := http.NewRequest("GET", "/payments", bytes.NewReader(newPayment))

	app := &App{db: &mockDB{}}
	http.HandlerFunc(app.CreatePayment).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := data.Payment{
		MongoID:        bson.ObjectIdHex("5b2ce1c5c089711b0e3bc2fa"),
		ID:             "new_payment_test",
		Type:           "Payment",
		Version:        1,
		OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Attributes: data.Attributes{
			Amount: 0.0,
			BeneficiaryParty: data.Account{
				AccountName:       "W Owens",
				AccountNumber:     "31926819",
				AccountNumberCode: "BBAN",
				AccountType:       0,
				Address:           "1 The Beneficiary Localtown SE2",
				BankID:            "403000",
				BankIDCode:        "GBDSC",
				Name:              "Wilfred Jeremiah Owens",
			},
			ChargesInformation: data.ChargesInformation{
				BearerCode: "SHAR",
				SenderCharges: []data.AmountCurrency{
					data.AmountCurrency{Amount: 0.0, Currency: "GBP"},
					data.AmountCurrency{Amount: 0.0, Currency: "USD"},
				},
				ReceiverChargesAmount:   0.0,
				ReceiverChargesCurrency: "USD",
			},
			Currency: "GBP",
			DebtorParty: data.Account{
				AccountName:       "EJ Brown Black",
				AccountNumber:     "GB29XABC10161234567801",
				AccountNumberCode: "IBAN",
				AccountType:       0,
				Address:           "10 Debtor Crescent Sourcetown NE1",
				BankID:            "203301",
				BankIDCode:        "GBDSC",
				Name:              "Emelia Jane Brown",
			},
			EndToEndReference: "Wil piano Jan",
			Fx: data.Fx{
				ContractReference: "FX123",
				ExchangeRate:      "2.00000",
				OriginalAmount:    0.0,
				OriginalCurrency:  "USD",
			},
			NumericReference:     "1002001",
			PaymentID:            "123456789012345678",
			PaymentPurpose:       "Paying for goods/services",
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			ProcessingDate:       "2017-01-18",
			Reference:            "Payment for Em's piano lessons",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "ImmediatePayment",
			SponsorParty: data.Sponsor{
				AccountNumber: "56781234",
				BankID:        "123123",
				BankIDCode:    "GBDSC",
			},
		},
	}

	var result data.Payment
	err := json.Unmarshal(rec.Body.Bytes(), &result)

	if err != nil {
		t.Errorf("Didn't expect error %v", err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected:\n%+v \nand instead got:\n%+v", expected, result)
	}
}

func TestCreatePaymentDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	newPayment := []byte(`
		 {
        "id": "new_payment_test",
        "type": "Payment",
        "version": 1,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "currency": "GBP"
                    },
                    {
                        "currency": "USD"
                    }
                ],
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    }`)
	req, _ := http.NewRequest("GET", "/payments", bytes.NewReader(newPayment))

	app := &App{db: &mockDB{testCaseDbError: true}}
	http.HandlerFunc(app.CreatePayment).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"database error"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestCreatePaymentDecodeError(t *testing.T) {
	rec := httptest.NewRecorder()
	newPayment := []byte(`
		 {
        "_id": "new_payment_test",
        "type": "Payment",
        "version": 1,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "currency": "GBP"
                    },
                    {
                        "currency": "USD"
                    }
                ],
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    }`)
	req, _ := http.NewRequest("GET", "/payments", bytes.NewReader(newPayment))

	app := &App{db: &mockDB{testCaseDbError: true}}
	http.HandlerFunc(app.CreatePayment).ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"invalid ObjectId in JSON: \"new_payment_test\""}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func (mdb *mockDB) RemovePayment(id bson.ObjectId) error {
	// in database we have
	// MongoID: bson.ObjectIdHex("5b290f5b802b0f1479000002"),
	fmt.Printf("Mock ListPaymentID id %v", id)
	if mdb.testCaseDbError != true {

		if id == bson.ObjectIdHex("5b290f5b802b0f1479000002") {
			return nil
		} else {
			return errors.New("not found")
		}
	} else {
		return errors.New(DB_ERROR)
	}
}

func TestRemovePaymentDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000002", &bytes.Buffer{})

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{testCaseDbError: true}}
	router.HandleFunc("/payments/{id}", app.DeletePayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"database error"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestRemovePaymentIDDoesNotExist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000003", &bytes.Buffer{})

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{}}
	router.HandleFunc("/payments/{id}", app.DeletePayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"not found"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestRemovePaymentIDExists(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000002", &bytes.Buffer{})

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{}}
	router.HandleFunc("/payments/{id}", app.DeletePayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"deleted"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func (mdb *mockDB) UpdatePayment(payment data.Payment) (*data.Payment, error) {
	// in database we have
	// MongoID:bson.ObjectIdHex("5b290f5b802b0f1479000002"),
	fmt.Printf("Mock ListPaymentID id %v", payment.MongoID)
	if mdb.testCaseDbError != true {

		if payment.MongoID == bson.ObjectIdHex("5b290f5b802b0f1479000002") {
			return &payment, nil
		} else {
			return nil, errors.New("not found")
		}
	} else {
		return nil, errors.New(DB_ERROR)
	}
}

func TestUpdatePaymentIDExists(t *testing.T) {
	rec := httptest.NewRecorder()
	updatePayment := []byte(`
		 {
		"_id": "5b290f5b802b0f1479000002",
        "id": "update_payment_test",
        "type": "Payment",
        "version": 1,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "currency": "GBP"
                    },
                    {
                        "currency": "USD"
                    }
                ],
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    }`)
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000002", bytes.NewReader(updatePayment))

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{}}
	router.HandleFunc("/payments/{id}", app.UpdatePayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := data.Payment{
		MongoID:        bson.ObjectIdHex("5b290f5b802b0f1479000002"),
		ID:             "update_payment_test",
		Type:           "Payment",
		Version:        1,
		OrganisationID: "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
		Attributes: data.Attributes{
			Amount: 0.0,
			BeneficiaryParty: data.Account{
				AccountName:       "W Owens",
				AccountNumber:     "31926819",
				AccountNumberCode: "BBAN",
				AccountType:       0,
				Address:           "1 The Beneficiary Localtown SE2",
				BankID:            "403000",
				BankIDCode:        "GBDSC",
				Name:              "Wilfred Jeremiah Owens",
			},
			ChargesInformation: data.ChargesInformation{
				BearerCode: "SHAR",
				SenderCharges: []data.AmountCurrency{
					data.AmountCurrency{Amount: 0.0, Currency: "GBP"},
					data.AmountCurrency{Amount: 0.0, Currency: "USD"},
				},
				ReceiverChargesAmount:   0.0,
				ReceiverChargesCurrency: "USD",
			},
			Currency: "GBP",
			DebtorParty: data.Account{
				AccountName:       "EJ Brown Black",
				AccountNumber:     "GB29XABC10161234567801",
				AccountNumberCode: "IBAN",
				AccountType:       0,
				Address:           "10 Debtor Crescent Sourcetown NE1",
				BankID:            "203301",
				BankIDCode:        "GBDSC",
				Name:              "Emelia Jane Brown",
			},
			EndToEndReference: "Wil piano Jan",
			Fx: data.Fx{
				ContractReference: "FX123",
				ExchangeRate:      "2.00000",
				OriginalAmount:    0.0,
				OriginalCurrency:  "USD",
			},
			NumericReference:     "1002001",
			PaymentID:            "123456789012345678",
			PaymentPurpose:       "Paying for goods/services",
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			ProcessingDate:       "2017-01-18",
			Reference:            "Payment for Em's piano lessons",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "ImmediatePayment",
			SponsorParty: data.Sponsor{
				AccountNumber: "56781234",
				BankID:        "123123",
				BankIDCode:    "GBDSC",
			},
		},
	}

	var result data.Payment
	err := json.Unmarshal(rec.Body.Bytes(), &result)

	if err != nil {
		t.Errorf("Didn't expect error %v", err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected:\n%+v \nand instead got:\n%+v", expected, result)
	}
}

func TestUpdatePaymentIDDoesNotExist(t *testing.T) {
	rec := httptest.NewRecorder()
	updatePayment := []byte(`
		 {
		"_id": "5b290f5b802b0f1479000003",
        "id": "update_payment_test",
        "type": "Payment",
        "version": 1,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "currency": "GBP"
                    },
                    {
                        "currency": "USD"
                    }
                ],
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    }`)
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000003", bytes.NewReader(updatePayment))

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{}}
	router.HandleFunc("/payments/{id}", app.UpdatePayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"not found"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestUpdatePaymentDBError(t *testing.T) {
	rec := httptest.NewRecorder()
	updatePayment := []byte(`
		 {
		"_id": "5b290f5b802b0f1479000002",
        "id": "update_payment_test",
        "type": "Payment",
        "version": 1,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "currency": "GBP"
                    },
                    {
                        "currency": "USD"
                    }
                ],
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    }`)
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000003", bytes.NewReader(updatePayment))

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{testCaseDbError: true}}
	router.HandleFunc("/payments/{id}", app.UpdatePayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"database error"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestUpdatePaymentDecodeError(t *testing.T) {
	rec := httptest.NewRecorder()
	updatePayment := []byte(`
		 {
		"_id": 5b290f5b802b0f1479000002,
        "id": "update_payment_test",
        "type": "Payment",
        "version": 1,
        "organisation_id": "743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
        "attributes": {
            "beneficiary_party": {
                "account_name": "W Owens",
                "account_number": "31926819",
                "account_number_code": "BBAN",
                "address": "1 The Beneficiary Localtown SE2",
                "bank_id": "403000",
                "bank_id_code": "GBDSC",
                "name": "Wilfred Jeremiah Owens"
            },
            "charges_information": {
                "bearer_code": "SHAR",
                "sender_charges": [
                    {
                        "currency": "GBP"
                    },
                    {
                        "currency": "USD"
                    }
                ],
                "receiver_charges_currency": "USD"
            },
            "currency": "GBP",
            "debtor_party": {
                "account_name": "EJ Brown Black",
                "account_number": "GB29XABC10161234567801",
                "account_number_code": "IBAN",
                "address": "10 Debtor Crescent Sourcetown NE1",
                "bank_id": "203301",
                "bank_id_code": "GBDSC",
                "name": "Emelia Jane Brown"
            },
            "end_to_end_reference": "Wil piano Jan",
            "fx": {
                "contract_reference": "FX123",
                "exchange_rate": "2.00000",
                "original_currency": "USD"
            },
            "numeric_reference": "1002001",
            "payment_id": "123456789012345678",
            "payment_purpose": "Paying for goods/services",
            "payment_scheme": "FPS",
            "payment_type": "Credit",
            "processing_date": "2017-01-18",
            "reference": "Payment for Em's piano lessons",
            "scheme_payment_sub_type": "InternetBanking",
            "scheme_payment_type": "ImmediatePayment",
            "sponsor_party": {
                "account_number": "56781234",
                "bank_id": "123123",
                "bank_id_code": "GBDSC"
            }
        }
    }`)
	req, _ := http.NewRequest("GET", "/payments/5b290f5b802b0f1479000003", bytes.NewReader(updatePayment))

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	app := &App{db: &mockDB{testCaseDbError: true}}
	router.HandleFunc("/payments/{id}", app.UpdatePayment).Methods("GET")

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("%+v != %+v", rec.Code, http.StatusOK)
	}

	expected := `{"status":"invalid character 'b' after object key:value pair"}`

	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
