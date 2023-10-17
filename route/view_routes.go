package route

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"net/http"
	"path/filepath"
	"pdf-generator/service"
	"pdf-generator/utils"
)

func VatNoteForBe(c echo.Context) error {
	customsId := c.Param("customsId")
	if customsId == "" {
		return c.JSON(http.StatusBadRequest, "The customs' ID is required.")
	}
	viewModel, err := service.QueryVatNotViewModel(customsId)
	if err != nil {
		println("errors :", err.Error())
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("The customs' ID %s canot query view model.", customsId))
	}

	return c.Render(http.StatusOK, "vat-note.html", viewModel)
}

func TransferDocBe(c echo.Context) error {
	customsId := c.Param("customsId")
	if customsId == "" {
		return c.JSON(http.StatusBadRequest, "The customs' ID is required.")
	}
	viewModel, err := service.QueryVatNotViewModel(customsId)
	if err != nil {
		println("errors :", err.Error())
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("The customs' ID %s canot query view model.", customsId))
	}

	return c.Render(http.StatusOK, "transfer-doc.html", viewModel)
}

func GeneratePdf(c echo.Context) error {
	customsId := c.Param("customsId")
	viewType := c.Param("viewType")
	if customsId == "" || viewType == "" {
		return c.JSON(http.StatusBadRequest, "The customs' ID and ViewType is required.")
	}

	viewModel, err := service.QueryVatNotViewModel(customsId)
	log.Infof("View model: %v", viewModel)
	if err != nil {
		println("errors :", err.Error())
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("The customs' ID %s canot query view model.", customsId))
	}
	var resourcePath string
	if viewType == "vatNote" {
		resourcePath = "resources/vat-note.html"
	}

	if viewType == "transferDoc" {
		resourcePath = "resources/transfer-doc.html"
	}

	if resourcePath == "" {
		return c.JSON(http.StatusBadRequest, "The view type is not supported.")
	}

	pg, err := utils.ParseTemplate(resourcePath, viewModel)

	if err != nil {
		println("Parse html template, errors :", err.Error())
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Parse html template failed. errors: %v", err))
	}

	savePath := filepath.Join(viper.GetString("tmp-dir"), customsId+".pdf")

	println("save path:", savePath)
	if err = utils.GeneratePageToPDF(pg, savePath); err != nil {
		println("Generate PDF file failed, errors :", err.Error())
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Generate PDF file failed, errors: %v", err))
	}
	if utils.IsExists(savePath) {
		return c.File(savePath)
	}

	return c.Render(http.StatusOK, "vat-note.html", viewModel)
}
