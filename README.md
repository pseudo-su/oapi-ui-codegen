OpenAPI Spec and Swagger UI Code Generator
----------------------------------------

When working with services, it can be convienient to embed the openapi spec and the SwaggerUI tool
into your server so that you can eaily explore your API during development.

This tool is designed to function alongside [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen) to embed the
data required in to create routes that expose your OpenAPI spec and the SwaggerUI tool.

## Using `oapi-ui-codegen`

Using the following `go:generate` directives will generate an embedded spec and a swaggerui file 
using the provided OpenAPI spec (it's recommended to use [`gobin`](https://github.com/myitcv/gobin)
for reproducible/consistent builds).

```go
//go:generate gobin -m -run github.com/pseudo-su/oapi-ui-codegen/cmd/oapi-ui-codegen --package=main --generate spec -o ./spec.gen.go ./openapi.yaml
//go:generate gobin -m -run github.com/pseudo-su/oapi-ui-codegen/cmd/oapi-ui-codegen --package=main --generate swaggerui -o ./swagger_ui.gen.go ./openapi.yaml
```

A working example can also be found here for reference [pseudo-su/golang-service-template](https://github.com/pseudo-su/golang-service-template)

### Options

The default options for `oapi-ui-codegen` will generate the swagger spec and swagger-ui
files, but you can generate subsets of those via the `-generate` flag. It defaults to `spec,swaggerui`.

- `spec`: embed the OpenAPI spec into the generated code as a gzipped blob.
- `swaggerui`: embed the SwaggerUI html pages into the generated code as a gzipped blob.
- `skip-fmt`: skip running `go fmt` on the generated code. This is useful for debugging the generated file in case the spec contains weird strings.

So, for example, if you would like to produce only the embedded spec, you could
run `oapi-ui-generate --generate spec`.
