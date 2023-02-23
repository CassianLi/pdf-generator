package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"html/template"
	"io"
	"pdf-generator/route"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Server() {
	// Echo instance
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("resources/*.html")),
	}
	e.Renderer = renderer

	// Middleware
	//e.Use(middleware.Logger())

	e.GET("/vatNoteBe/:customsId", route.VatNoteForBe)

	e.GET("/generatePdf/:viewType/:customsId", route.GeneratePdf)

	port := viper.GetString("port")
	if port == "" {
		port = ":1323"
	} else {
		port = fmt.Sprintf(":%s", port)
	}
	// Start server
	e.Logger.Fatal(e.Start(port))
}
