package codegen

import (
	"go/format"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golangci/lint-1"
	"github.com/stretchr/testify/assert"
)

func TestSpecEmbedding(t *testing.T) {

	// Input vars for code generation:
	packageName := "testswagger"
	opts := Options{
		Spec:      true,
		SwaggerUI: false,
	}

	// Get a spec from the test definition in this file:
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData([]byte(testOpenAPIDefinition))
	assert.NoError(t, err)

	// Run our code generation:
	code, err := Generate(swagger, packageName, opts)
	assert.NoError(t, err)
	assert.NotEmpty(t, code)

	// Check that we have valid (formattable) code:
	_, err = format.Source([]byte(code))
	assert.NoError(t, err)

	// Check that we have a package:
	assert.Contains(t, code, "package testswagger")

	assert.Contains(t, code, "func GetOpenAPISpec() (*openapi3.Swagger, error) {")

	// Write snapshot
	cupaloy.SnapshotT(t, code)

	// Make sure the generated code is valid:
	linter := new(lint.Linter)
	problems, err := linter.Lint("test.gen.go", []byte(code))
	assert.NoError(t, err)
	assert.Len(t, problems, 0)
}
func TestSwaggerUIEmbedding(t *testing.T) {

	// Input vars for code generation:
	packageName := "testswagger"
	opts := Options{
		Spec:      false,
		SwaggerUI: true,
	}

	// Get a spec from the test definition in this file:
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData([]byte(testOpenAPIDefinition))
	assert.NoError(t, err)

	// Run our code generation:
	code, err := Generate(swagger, packageName, opts)
	assert.NoError(t, err)
	assert.NotEmpty(t, code)

	// Check that we have valid (formattable) code:
	_, err = format.Source([]byte(code))
	assert.NoError(t, err)

	// Check that we have a package:
	assert.Contains(t, code, "package testswagger")

	assert.Contains(t, code, "func GetOauth2RedirectPage(specURL string, redirectURL string) ([]byte, error) {")
	assert.Contains(t, code, "func GetSwaggerUIPage(specURL string, redirectURL string) ([]byte, error) {")

	// Write snapshot
	cupaloy.SnapshotT(t, code)

	// Make sure the generated code is valid:
	linter := new(lint.Linter)
	problems, err := linter.Lint("test.gen.go", []byte(code))
	assert.NoError(t, err)
	assert.Len(t, problems, 0)
}

const testOpenAPIDefinition = `
openapi: 3.0.1

info:
  title: OpenAPI-CodeGen Test
  description: 'This is a test OpenAPI Spec'
  version: 1.0.0

servers:
- url: https://test.oapi-codegen.com/v2
- url: http://test.oapi-codegen.com/v2

paths:
  /test/{name}:
    get:
      tags:
      - test
      summary: Get test
      operationId: getTestByName
      parameters:
      - name: name
        in: path
        required: true
        schema:
          type: string
      responses:
        200:
          description: Success
          content:
            application/xml:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Test'
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Test'
        422:
          description: InvalidArray
          content:
            application/xml:
              schema:
                type: array
            application/json:
              schema:
                type: array
        default:
          description: Error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
  /cat:
    get:
      tags:
      - cat
      summary: Get cat status
      operationId: getCatStatus
      responses:
        200:
          description: Success
          content:
            application/json:
              schema:
                oneOf:
                - $ref: '#/components/schemas/CatAlive'
                - $ref: '#/components/schemas/CatDead'
            application/xml:
              schema:
                anyOf:
                - $ref: '#/components/schemas/CatAlive'
                - $ref: '#/components/schemas/CatDead'
            application/yaml:
              schema:
                allOf:
                - $ref: '#/components/schemas/CatAlive'
                - $ref: '#/components/schemas/CatDead'
        default:
          description: Error
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'

components:
  schemas:

    Test:
      properties:
        name:
          type: string
        cases:
          type: array
          items:
            $ref: '#/components/schemas/TestCase'

    TestCase:
      properties:
        name:
          type: string
        command:
          type: string

    Error:
      properties:
        code:
          type: integer
          format: int32
        message:
          type: string

    CatAlive:
      properties:
        name:
          type: string
        alive_since:
          type: string
          format: date-time

    CatDead:
      properties:
        name:
          type: string
        dead_since:
          type: string
          format: date-time
        cause:
          type: string
          enum: [car, dog, oldage]
`
