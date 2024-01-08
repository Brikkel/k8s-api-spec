package openapiv3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// genericOpenAPISpec represents the rough structure of the OpenAPI specification, mostly for debugging purposes
type genericOpenAPISpec struct {
	Components struct {
		Schemas map[string]json.RawMessage `json:"schemas"`
	} `json:"components"`
}

// OpenAPISpec represents the structure of the OpenAPI specification.
type OpenAPISpec struct {
	Components struct {
		Schemas map[string]Schema `json:"schemas"`
	} `json:"components"`
}

// Schema represents a schema in the OpenAPI specification.
type Schema struct {
	// The description of the schema.
	Description string               `json:"description"`
	Type        string               `json:"type"`
	Properties  map[string]Propertie `json:"properties"`
	Required    []string             `json:"required"`
	// XKubernetesGroupVersionKind object only exists for kinds.
	XKubernetesGroupVersionKind []XKubernetesGroupVersionKind `json:"x-kubernetes-group-version-kind"`
}

type XKubernetesGroupVersionKind struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

type Propertie struct {
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Default     interface{} `json:"default"`
	Items       item        `json:"items"`
	AllOf       []AllOf     `json:"allOf"`
}

type item struct {
	Type    string      `json:"type"`
	Default interface{} `json:"default"`
	AllOf   []AllOf     `json:"allOf"`
}

type AllOf struct {
	Ref string `json:"$ref"`
}

// Get the list of resources and their schemas of a provided API version.
func GetSchemas(endpointURL, path string) (map[string]Schema, error) {
	// URL of the OpenAPI specification endpoint
	endpoint := endpointURL + path

	// Make an HTTP GET request
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	// Parse the OpenAPI specification JSON
	var openAPISpec OpenAPISpec
	err = json.Unmarshal(body, &openAPISpec)
	if err != nil {
		return nil, fmt.Errorf("Error parsing OpenAPI JSON: %w", err)
	}

	// Access the components.schemas section
	schemas := openAPISpec.Components.Schemas

	// Return the schemas
	return schemas, nil
}

func GetSchemasJson(endpointURL, path string) (string, error) {
	// URL of the OpenAPI specification endpoint
	endpoint := endpointURL + path

	// Make an HTTP GET request
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", fmt.Errorf("Error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %w", err)
	}

	// Parse the OpenAPI specification JSON
	var openAPISpec genericOpenAPISpec
	err = json.Unmarshal(body, &openAPISpec)
	if err != nil {
		return "", fmt.Errorf("Error parsing OpenAPI JSON: %w", err)
	}

	schemas := openAPISpec.Components.Schemas

	jsonSchemas, err := json.MarshalIndent(schemas, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error marshalling schemas to JSON: %w", err)
	}

	// Return the JSON as a string
	return string(jsonSchemas), nil
}

func GetSchema(endpointURL, path, schemaName string) (*Schema, error) {
	schemas, err := GetSchemas(endpointURL, path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	schema, found := schemas[schemaName]
	if !found {
		return nil, fmt.Errorf("Schema %s not found", schemaName)
	}

	return &schema, nil
}
