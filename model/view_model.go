package model

import (
	"database/sql"
)

type (
	Address struct {
		CompanyName   string         `db:"company_name"`
		Country       string         `db:"country"`
		State         sql.NullString `db:"state"`
		City          string         `db:"city"`
		PostalCode    string         `db:"postal_code"`
		AddressDetail string         `db:"address_detail"`
		VatNumber     sql.NullString `db:"vat_no"`
	}

	CustomsTax struct {
		ItemNumber     int            `db:"item_number"`
		Quantity       int            `db:"quantity"`
		Description    string         `db:"description"`
		TaxType        sql.NullString `db:"tax_type"`
		TaxRate        float64        `db:"tax_rate"`
		TaxFee         float64        `db:"tax_fee"`
		DeclaredAmount float64        `db:"declared_amount"`
	}

	CustomsStatusDate struct {
		CustomsId  string `db:"customs_id"`
		StatusDate string `db:"status_date"`
	}

	// VatNoteViewModel The model of vat not view
	VatNoteViewModel struct {
		CustomsId  string
		StatusDate string
		// A00 tax fee total
		TaxFeeTotal        float64
		DeclareAmountTotal float64
		ImporterAddress    Address
		ConsigneeAddress   Address
		Articles           []CustomsTax
	}
)
