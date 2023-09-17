package database

import (
	"context"
	"github.com/google/uuid"
)

type Debtor struct {
	Id             string `db:"id"`
	FiscalDocument string `db:"fiscal_document"`
	Name           string `db:"name"`
	Email          string `db:"email"`
	TypeOfDebtor   string `db:"type_of_debtor"`
	PostalCode     string `db:"postal_code"`
	City           string `db:"city"`
	Uf             string `db:"uf"`
	Street         string `db:"street"`
	Number         string `db:"number"`
	Complement     string `db:"complement"`
}

func (db *DB) InsertDebtor(fiscalDocument, name, email, typeOfDebtor, postalCode, city, uf, number, complement string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	query := `
		INSERT INTO debtors (id,fiscal_document, name, email, type_of_debtor, postal_code, city, uf, number, complement)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = db.ExecContext(ctx, query, id, fiscalDocument, name, email, typeOfDebtor, postalCode, city, uf, number, complement)
	if err != nil {
		return "", err
	}

	return id.String(), err
}
