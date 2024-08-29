package dao

import (
	"fmt"
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
	// 1. 判断是否是拆分报关
	var hasSplit bool
	err = global.Db.Get(&hasSplit, QueryCustomsHasSplitSql, customsId)
	if err != nil {
		fmt.Printf("Query customs has_split error, continue to make as no-split. %v\n", err)
	}

	// 2. 选择SQL语句
	sql := QueryCustomsTaxSql
	if hasSplit {
		sql = QueryCustomsSplitTaxSql
	}

	err = global.Db.Select(&customsTaxList, sql, customsId)
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
