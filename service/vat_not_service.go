package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"pdf-generator/dao"
	"pdf-generator/model"
)

// QueryVatNotViewModel 查询VatNote 页面数据
func QueryVatNotViewModel(customsId string) (viewModel model.VatNoteViewModel, err error) {
	customsStatusTime, err := dao.QueryCustomsStatusDateForTax(customsId)
	fmt.Println("customsStatusTime", customsStatusTime)
	if err != nil {
		return viewModel, err
	}
	viewModel.CustomsId = customsStatusTime.CustomsId
	viewModel.StatusDate = customsStatusTime.StatusDate

	importAddress, err := dao.QueryCustomsAddress(customsId, dao.IMPORTER)
	fmt.Println("importAddress", importAddress)

	if err != nil {
		return viewModel, err
	}
	viewModel.ImporterAddress = importAddress

	deliveryAddress, err := dao.QueryCustomsAddress(customsId, dao.DELIVERY)
	fmt.Println("deliveryAddress", deliveryAddress)

	if err != nil {
		return viewModel, err
	}
	viewModel.ConsigneeAddress = deliveryAddress

	customsTaxList, err := dao.QueryCustomsTaxList(customsId)
	fmt.Println("customsTaxList", customsTaxList)

	if err != nil {
		return viewModel, err
	}

	taxFeeTotal, declareAmountTotal := decimal.NewFromFloat(0.0), decimal.NewFromFloat(0.0)
	for _, tax := range customsTaxList {
		taxFeeTotal = taxFeeTotal.Add(decimal.NewFromFloat(tax.TaxFee))
		declareAmountTotal = declareAmountTotal.Add(decimal.NewFromFloat(tax.DeclaredAmount))
	}

	viewModel.TaxFeeTotal, _ = taxFeeTotal.Round(2).Float64()
	viewModel.Articles = customsTaxList
	viewModel.DeclareAmountTotal, _ = declareAmountTotal.Round(2).Float64()

	return viewModel, err
}
