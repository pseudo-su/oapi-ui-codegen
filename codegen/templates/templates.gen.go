package templates

import (
	"strings"
	"text/template"
)

var templates = map[string]string{"imports.tmpl": `// Package {{.PackageName}} provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/pseudo-su/oapi-ui-codegen DO NOT EDIT.
package {{.PackageName}}

{{if .Imports}}
import (
{{range .Imports}} "{{.}}"
{{end}})
{{end}}
`,
	"inline-ui.tmpl": `// data{{.Code}} is a Base64 encoded, gzipped, json marshaled html template
var data{{.Code}} = []string{
{{range .Parts}}
    "{{.}}",{{end}}
}

// Get{{.Code}} returns the {{.Description}}.
func Get{{.Code}}(specURL string, redirectURL string) ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(data{{.Code}}, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding data{{.Code}}: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	s := buf.String()
	// idx := strings.Index(s, "<!DOCTYPE html>")

	tmpl, err := template.
		New("oapi-ui-codegen {{.Code}}").
		Delims("__LEFT_DELIM__", "__RIGHT_DELIM__").
		Parse(s)

	if err != nil {
		return nil, fmt.Errorf("error loading {{.Code}} template: %s", err)
	}
	buf.Reset()
	data := struct {
		SpecURL string
		SwaggerUIRedirectURL string
	} {
		SpecURL: specURL,
		SwaggerUIRedirectURL: redirectURL,
	}
	
	tmpl.Execute(&buf, data)
	return buf.Bytes(), nil
}
`,
	"inline.tmpl": `// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{
{{range .}}
    "{{.}}",{{end}}
}

// GetOpenAPISpec returns the Swagger specification corresponding to the generated code
// in this file.
func GetOpenAPISpec() (*openapi3.T, error) {
    zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
    if err != nil {
        return nil, fmt.Errorf("error base64 decoding spec: %s", err)
    }
    zr, err := gzip.NewReader(bytes.NewReader(zipped))
    if err != nil {
        return nil, fmt.Errorf("error decompressing spec: %s", err)
    }
    var buf bytes.Buffer
    _, err = buf.ReadFrom(zr)
    if err != nil {
        return nil, fmt.Errorf("error decompressing spec: %s", err)
    }

    swagger, err := openapi3.NewLoader().LoadFromData(buf.Bytes())
    if err != nil {
        return nil, fmt.Errorf("error loading Swagger: %s", err)
    }
    return swagger, nil
}
`,
}

func Parse(t *template.Template) (*template.Template, error) {
	for name, s := range templates {
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		ss := strings.ReplaceAll(s, "__TEMP_BACKTICK__", "`")
		if _, err := tmpl.Parse(ss); err != nil {
			return nil, err
		}
	}
	return t, nil
}
