package handler

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/app/service/entity"
)

const swaggerUITemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
	<meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.1.3/swagger-ui-standalone-preset.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.1.3/swagger-ui-bundle.js"></script>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/5.1.3/swagger-ui.css" />
    <title>{{ .Title }}</title>
    <style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }

      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
    </style>
  </head>

  <body>
    <div id="swagger-ui"></div>
    <script>
	  window.onload = function() {
	    SwaggerUIBundle({
          url: "{{ .SpecURL }}",
          dom_id: '#swagger-ui',
          presets: [
            SwaggerUIBundle.presets.apis,
            SwaggerUIStandalonePreset
          ],
          layout: "StandaloneLayout"
        })
      }
    </script>
  </body>
</html>
`

const specURL = "/docs/swagger.yaml"

type tmplParam struct {
	SpecURL string
	Title   string
}

func ServeDoc(c echo.Context) error {
	tmpl := template.Must(template.New("swagger-ui").Parse(swaggerUITemplate))
	buf := bytes.NewBuffer(nil)
	param := tmplParam{
		Title:   "Docs",
		SpecURL: specURL,
	}
	_ = tmpl.Execute(buf, param)
	b := buf.String()

	return c.HTML(http.StatusOK, b)
}

func DocsSpec(c echo.Context) error {
	b, err := os.ReadFile("./docs/swagger.yaml")
	if err != nil {
		return c.JSON(http.StatusNotFound, entity.MessageResponse{Message: "docs spec not found"})
	}

	// Send the contents of the spec file.
	return c.JSONBlob(http.StatusOK, b)
}
