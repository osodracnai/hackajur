package database

import (
	"encoding/json"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	p := Propose{

		Debt: Debt{
			Debtor: []Debtor{{
				Id:             uuid.NewString(),
				FiscalDocument: "12345678",
				Name:           "Devedor",
				Email:          "email@email.com",
				TypeOfDebtor:   "individual | company",
				PostalCode:     "38400200",
				City:           "cidade",
				Uf:             "estado",
				Street:         "rua",
				Number:         "0",
				Complement:     "ap 123",
			}},
			Origin:         "",
			DocumentID:     uuid.NewString(),
			ExpirationDate: time.Now().Add(time.Hour * 100000),
			OriginalValue:  300,
			Fee: Fee{
				Name:        "taxa",
				ChargerType: "fixed | percentage",
				Value:       123,
			},
			Interest: Interest{
				Name:        "juros",
				ChargerType: "fixed' | 'percentage",
				Value:       123,
			},
			OtherCharges: nil,
			PresentValue: 1250,
			Collateral: []CollateralGurantee{{
				Name:  "",
				Desc:  "",
				Value: "",
			}},
			Correction: Correction{
				Name:             "correcao",
				Value:            123,
				CorrectionStatus: "correct | no-correct",
			},
		},
		Date: time.Now(),
		Status: Status{
			UpdatedAt: time.Now().Add(time.Hour * 24),
			Situation: "sent\" | \"viewed\" | \"accepted\" | \"completed\" | \"expired\" | \"cancelled\" | \"execution\" | \"error\"",
		},
		ProposeValue:    12345,
		ExpirationDate:  time.Now().Add(time.Hour * 96),
		PaymentDeadline: 3,
		Payments: []Payments{
			{
				Name:   "name",
				Status: false,
			},
		},
		Comunication: []Comunication{{
			Name:   "Email",
			Status: true,
		}, {
			Name:   "SMS",
			Status: true,
		},
			{
				Name:   "WhatsApp",
				Status: true,
			}},
	}
	a, _ := json.Marshal(p)
	println(string(a))
}
