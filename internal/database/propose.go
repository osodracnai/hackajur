package database

import "time"

type Propose struct {
	Id              string         `json:"id"`
	Debt            Debt           `json:"debt"`
	Date            time.Time      `json:"date"`
	Status          Status         `json:"status"`
	ProposeValue    int            `json:"proposeValue"`
	ExpirationDate  interface{}    `json:"expirationDate"`
	PaymentDeadline int            `json:"paymentDeadline"`
	Payments        []Payments     `json:"payments"`
	Comunication    []Comunication `json:"comunication"`
}
type Status struct {
	UpdatedAt time.Time `json:"updatedAt"`
	Situation string    `json:"situation"`
}

type Payments struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
type Comunication struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
type Debt struct {
	Id             string               `json:"id"`
	Debtor         []Debtor             `json:"debtor"`
	Origin         string               `json:"origin"`
	DocumentID     string               `json:"documentId"`
	ExpirationDate time.Time            `json:"expirationDate"`
	OriginalValue  int                  `json:"originalValue"`
	Fee            Fee                  `json:"fee"`
	Interest       Interest             `json:"interest"`
	OtherCharges   []OtherCharges       `json:"otherCharges"`
	PresentValue   int                  `json:"presentValue"`
	Collateral     []CollateralGurantee `json:"collateral"`
	Correction     Correction           `json:"correction"`
}
type CollateralGurantee struct {
	Name  string
	Desc  string
	Value string
}
type Correction struct {
	Name             string `json:"name"`
	Value            int    `json:"value"`
	CorrectionStatus string `json:"correctionStatus"`
}
type OtherCharges struct {
	Name        string `json:"name"`
	ChargerType string `json:"chargerType"`
	Value       int    `json:"value"`
}
type Interest struct {
	Name        string `json:"name"`
	ChargerType string `json:"chargerType"`
	Value       int    `json:"value"`
}
type Fee struct {
	Name        string `json:"name"`
	ChargerType string `json:"chargerType"`
	Value       int    `json:"value"`
}
