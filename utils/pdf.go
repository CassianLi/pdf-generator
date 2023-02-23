package utils

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/spf13/viper"
	"strings"
)

// GeneratePageToPDF Generates PDF file from html page string
func GeneratePageToPDF(pg string, savePath string) (err error) {
	wkhtmltopdfPath := viper.GetString("wkhtmltopdf-path")
	if wkhtmltopdfPath != "" {
		wkhtmltopdf.SetPath(wkhtmltopdfPath)
	}
	// Client code
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(pg)))
	pdfg.Dpi.Set(600)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	// Create PDF document in internal buffer
	if err = pdfg.Create(); err != nil {
		return err
	}

	// Write buffer contents to file on disk
	if err = pdfg.WriteFile(savePath); err != nil {
		return err
	}
	return nil
}
