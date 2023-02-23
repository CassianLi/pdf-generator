package dao

import (
	global "pdf-generator/config"
	"pdf-generator/model"
)

type AddressType int

const (
	DELIVERY AddressType = iota
	IMPORTER
)

// QueryCustomsStatusDateForTax Query customs tax status time
func QueryCustomsStatusDateForTax(customsId string) (customsStatusDate model.CustomsStatusDate, err error) {
	err = global.Db.Get(&customsStatusDate, QueryCustomsStatusTimeSql, customsId)
	return customsStatusDate, err
}

// QueryCustomsTaxList Query tax list of customs
func QueryCustomsTaxList(customsId string) (customsTaxList []model.CustomsTax, err error) {
	err = global.Db.Select(&customsTaxList, QueryCustomsTaxSql, customsId)
	return customsTaxList, err
}

// QueryCustomsAddress Query address of customs
func QueryCustomsAddress(customsId string, addressType AddressType) (customsAddress model.Address, err error) {
	var sql string
	switch addressType {
	case DELIVERY:
		sql = QueryDeliveryAddressSql
	case IMPORTER:
		sql = QueryCustomsImporterAddressSql
	default:
		sql = ""
	}
	err = global.Db.Get(&customsAddress, sql, customsId)
	return customsAddress, err
}
