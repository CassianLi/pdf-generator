package dao

const (
	// QueryCustomsTaxSql 查询报关税金单详细信息
	QueryCustomsTaxSql = `SELECT sca.item_number,
       sca.quantity,
       bd.description,
       bct.tax_type,
       bct.tax_rate,
       bct.tax_fee,
       bct.declared_amount
FROM base_customs_tax bct
         INNER JOIN service_customs_article sca ON bct.customs_id = sca.customs_id AND bct.itemnr = sca.item_number
         INNER JOIN base_description bd ON sca.product_no = bd.product_no AND sca.country = bd.country
WHERE bct.tax_type='A00' and bct.customs_id = ?;`

	// QueryCustomsHasSplitSql 查询指定的customs_id是否是拆分报关
	QueryCustomsHasSplitSql = `
SELECT has_split
FROM stats_customs_info
WHERE  customs_id = ?;`

	// QueryCustomsSplitTaxSql 查询拆分报关的税金信息
	QueryCustomsSplitTaxSql = `
SELECT scsa.item AS item_number,
       sca.quantity,
       bd.description,
       bct.tax_type,
       bct.tax_rate,
       bct.tax_fee,
       bct.declared_amount
FROM base_customs_tax bct
         INNER JOIN service_customs_supply_article scsa ON bct.customs_id = scsa.customs_id AND bct.itemnr = scsa.item
         INNER JOIN service_customs_article sca ON scsa.article_id = sca.id
         INNER JOIN base_description bd ON sca.product_no = bd.product_no AND sca.country = bd.country
WHERE bct.tax_type = 'A00'
  AND bct.customs_id = ?
ORDER BY scsa.item;`

	// QueryCustomsStatusTimeSql 查询税金单时间
	QueryCustomsStatusTimeSql = `SELECT customs_id,
       gmt_create AS status_date
FROM log_clearance_process
WHERE process_code = 'TAX'
  AND customs_id =?;`

	// QueryCustomsImporterAddressSql Query importer address
	QueryCustomsImporterAddressSql = `
	SELECT ba.company_or_name AS company_name,
       ba.country,
       ba.state,
       ba.city,
       ba.postal_code,
       CONCAT_WS(' ', ba.address_line1, ba.address_line2, ba.address_line3) AS address_detail,
       ba.vat_no
FROM base_customs bc
         INNER JOIN base_address ba ON bc.importer = ba.address_code
WHERE bc.customs_id = ?`

	// QueryDeliveryAddressSql Query delivery address
	QueryDeliveryAddressSql = `SELECT ba.company_or_name AS company_name,
       ba.country,
       ba.state,
       ba.city,
       ba.postal_code,
       CONCAT_WS(' ', ba.address_line1, ba.address_line2, ba.address_line3) AS address_detail,
       ba.vat_no
FROM base_customs bc
         INNER JOIN base_address ba ON bc.delivery_address_code = ba.address_code
WHERE bc.customs_id =?`
)
