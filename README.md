OpenAPI Spec and Swagger UI Code Generator
----------------------------------------

This package contains a set of utilities for generating Go boilerplate code for
services based on
[OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md)
API definitions.
When working with services, it can be convienient to embed the openapi spec and the SwaggerUI tool
into your server so that you can eaily explore your API during development. This tool is designed
to function alongside [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen) to embed the
data required in to create routes that expose your openapi spec and the SwaggerUI tool.

## Overview

We're going to use the OpenAPI example of the
[Expanded Petstore](https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore-expanded.yaml)
in the descriptions below, please have a look at it.

## Using `oapi-codegen`

The default options for `oapi-ui-codegen` will generate the swagger spec and swagger-ui
files, but you can generate subsets of those via the `-generate` flag. It defaults to `spec,swaggerui`.

- `spec`: embed the OpenAPI spec into the generated code as a gzipped blob.
- `swaggerui`: embed the SwaggerUI html pages into the generated code as a gzipped blob.
- `skip-fmt`: skip running `go fmt` on the generated code. This is useful for debugging the generated file in case the spec contains weird strings.

So, for example, if you would like to produce only the embedded spec, you could
run `oapi-ui-generate --generate spec`.
