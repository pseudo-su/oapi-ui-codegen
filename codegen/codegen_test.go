package codegen

// import (
// 	"bytes"
// 	"go/format"
// 	"io/ioutil"
// 	"net/http"
// 	"testing"

// 	"github.com/getkin/kin-openapi/openapi3"
// 	"github.com/golangci/lint-1"
// 	"github.com/stretchr/testify/assert"
// )

// func TestExamplePetStoreCodeGeneration(t *testing.T) {

// 	// Input vars for code generation:
// 	packageName := "api"
// 	opts := Options{
// 		GenerateClient:     true,
// 		GenerateEchoServer: true,
// 		GenerateTypes:      true,
// 		EmbedSpec:          true,
// 		EmbedSpecUI:        true,
// 	}

// 	// Get a spec from the example PetStore definition:
// 	swagger, err := examplePetstore.GetOpenAPISpec()
// 	assert.NoError(t, err)

// 	// Run our code generation:
// 	code, err := Generate(swagger, packageName, opts)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, code)

// 	// Check that we have valid (formattable) code:
// 	_, err = format.Source([]byte(code))
// 	assert.NoError(t, err)

// 	// Check that we have a package:
// 	assert.Contains(t, code, "package api")

// 	// Check that the client method signatures return response structs:
// 	assert.Contains(t, code, "func (c *Client) FindPetById(ctx context.Context, id int64) (*http.Response, error) {")

// 	// Make sure the generated code is valid:
// 	linter := new(lint.Linter)
// 	problems, err := linter.Lint("test.gen.go", []byte(code))
// 	assert.NoError(t, err)
// 	assert.Len(t, problems, 0)
// }

// func TestExamplePetStoreParseFunction(t *testing.T) {

// 	bodyBytes := []byte(`{"id": 5, "name": "testpet", "tag": "cat"}`)

// 	cannedResponse := &http.Response{
// 		StatusCode: 200,
// 		Body:       ioutil.NopCloser(bytes.NewReader(bodyBytes)),
// 		Header:     http.Header{},
// 	}
// 	cannedResponse.Header.Add("Content-type", "application/json")

// 	findPetByIDResponse, err := examplePetstoreClient.ParsefindPetByIdResponse(cannedResponse)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, findPetByIDResponse.JSON200)
// 	assert.Equal(t, int64(5), findPetByIDResponse.JSON200.Id)
// 	assert.Equal(t, "testpet", findPetByIDResponse.JSON200.Name)
// 	assert.NotNil(t, findPetByIDResponse.JSON200.Tag)
// 	assert.Equal(t, "cat", *findPetByIDResponse.JSON200.Tag)
// }

// func TestExampleOpenAPICodeGeneration(t *testing.T) {

// 	// Input vars for code generation:
// 	packageName := "testswagger"
// 	opts := Options{
// 		GenerateClient:     true,
// 		GenerateEchoServer: true,
// 		GenerateTypes:      true,
// 		EmbedSpec:          true,
// 		EmbedSpecUI:        true,
// 	}

// 	// Get a spec from the test definition in this file:
// 	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData([]byte(testOpenAPIDefinition))
// 	assert.NoError(t, err)

// 	// Run our code generation:
// 	code, err := Generate(swagger, packageName, opts)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, code)

// 	// Check that we have valid (formattable) code:
// 	_, err = format.Source([]byte(code))
// 	assert.NoError(t, err)

// 	// Check that we have a package:
// 	assert.Contains(t, code, "package testswagger")

// 	// Check that response structs are generated correctly:
// 	assert.Contains(t, code, "type getTestByNameResponse struct {")

// 	// Check that response structs contains fallbacks to interface for invalid types:
// 	// Here an invalid array with no items.
// 	assert.Contains(t, code, `
// type getTestByNameResponse struct {
// 	Body         []byte
// 	HTTPResponse *http.Response
// 	JSON200      *[]Test
// 	XML200       *[]Test
// 	JSON422      *[]interface{}
// 	XML422       *[]interface{}
// 	JSONDefault  *Error
// }`)

// 	// Check that the helper methods are generated correctly:
// 	assert.Contains(t, code, "func (r getTestByNameResponse) Status() string {")
// 	assert.Contains(t, code, "func (r getTestByNameResponse) StatusCode() int {")
// 	assert.Contains(t, code, "func ParsegetTestByNameResponse(rsp *http.Response) (*getTestByNameResponse, error) {")

// 	// Check the client method signatures:
// 	assert.Contains(t, code, "func (c *Client) GetTestByName(ctx context.Context, name string) (*http.Response, error) {")
// 	assert.Contains(t, code, "func (c *ClientWithResponses) GetTestByNameWithResponse(ctx context.Context, name string) (*getTestByNameResponse, error) {")

// 	// Make sure the generated code is valid:
// 	linter := new(lint.Linter)
// 	problems, err := linter.Lint("test.gen.go", []byte(code))
// 	assert.NoError(t, err)
// 	assert.Len(t, problems, 0)
// }

// const testOpenAPIDefinition = `
// openapi: 3.0.1

// info:
//   title: OpenAPI-CodeGen Test
//   description: 'This is a test OpenAPI Spec'
//   version: 1.0.0

// servers:
// - url: https://test.oapi-codegen.com/v2
// - url: http://test.oapi-codegen.com/v2

// paths:
//   /test/{name}:
//     get:
//       tags:
//       - test
//       summary: Get test
//       operationId: getTestByName
//       parameters:
//       - name: name
//         in: path
//         required: true
//         schema:
//           type: string
//       responses:
//         200:
//           description: Success
//           content:
//             application/xml:
//               schema:
//                 type: array
//                 items:
//                   $ref: '#/components/schemas/Test'
//             application/json:
//               schema:
//                 type: array
//                 items:
//                   $ref: '#/components/schemas/Test'
//         422:
//           description: InvalidArray
//           content:
//             application/xml:
//               schema:
//                 type: array
//             application/json:
//               schema:
//                 type: array
//         default:
//           description: Error
//           content:
//             application/json:
//               schema:
//                 $ref: '#/components/schemas/Error'
//   /cat:
//     get:
//       tags:
//       - cat
//       summary: Get cat status
//       operationId: getCatStatus
//       responses:
//         200:
//           description: Success
//           content:
//             application/json:
//               schema:
//                 oneOf:
//                 - $ref: '#/components/schemas/CatAlive'
//                 - $ref: '#/components/schemas/CatDead'
//             application/xml:
//               schema:
//                 anyOf:
//                 - $ref: '#/components/schemas/CatAlive'
//                 - $ref: '#/components/schemas/CatDead'
//             application/yaml:
//               schema:
//                 allOf:
//                 - $ref: '#/components/schemas/CatAlive'
//                 - $ref: '#/components/schemas/CatDead'
//         default:
//           description: Error
//           content:
//             application/json:
//               schema:
//                 $ref: '#/components/schemas/Error'

// components:
//   schemas:

//     Test:
//       properties:
//         name:
//           type: string
//         cases:
//           type: array
//           items:
//             $ref: '#/components/schemas/TestCase'

//     TestCase:
//       properties:
//         name:
//           type: string
//         command:
//           type: string

//     Error:
//       properties:
//         code:
//           type: integer
//           format: int32
//         message:
//           type: string

//     CatAlive:
//       properties:
//         name:
//           type: string
//         alive_since:
//           type: string
//           format: date-time

//     CatDead:
//       properties:
//         name:
//           type: string
//         dead_since:
//           type: string
//           format: date-time
//         cause:
//           type: string
//           enum: [car, dog, oldage]
// `
