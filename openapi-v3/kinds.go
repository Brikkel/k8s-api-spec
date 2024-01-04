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
	Description string `json:"description"`
	Type        string `json:"type"`
	// XKubernetesGroupVersionKind object only exists for kinds.
	XKubernetesGroupVersionKind []struct {
		Group   string `json:"group"`
		Version string `json:"version"`
		Kind    string `json:"kind"`
	} `json:"x-kubernetes-group-version-kind"`
}

// Struct that represents apiversions with kinds
type ApiVersion struct {
	Name           string
	ApiVersionPath string
	Kinds          []Kind
}

// Struct that represents kinds
type Kind struct {
	Name        string
	Description string
	Type        string
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

func SchemasToJson(schemas map[string]Schema) (string, error) {
	// Marshal the schemas to JSON
	jsonSchemas, err := json.MarshalIndent(schemas, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error marshalling schemas to JSON: %w", err)
	}

	// Return the JSON
	return string(jsonSchemas), nil
}

// Loop through the available API versions and return all kinds.
func GetAllKinds(endpointURL string) ([]ApiVersion, error) {
	// Get the paths to the available API versions
	paths, err := GetOpenAPIV3Paths(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var ApiVersionKinds []ApiVersion
	// For each API version, get the kinds and combine it into one list
	for _, path := range paths {
		kindsPerVersion, err := GetKindsPerVersion(endpointURL, "/openapi/v3/"+path)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		var newApiVersion ApiVersion
		newApiVersion.Name = path
		newApiVersion.ApiVersionPath = path
		newApiVersion.Kinds = kindsPerVersion
		ApiVersionKinds = append(ApiVersionKinds, newApiVersion)
	}

	// Return all the kinds
	return ApiVersionKinds, nil
}

// Get the kinds for a specified API version.
func GetKindsPerVersion(endpointURL, path string) ([]Kind, error) {
	// Filter out resources with "x-kubernetes-group-version-kind" parameter
	//filteredSchemas := make(map[string]Schema)
	var kinds []Kind

	schemas, err := GetSchemas(endpointURL, path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	for _, schema := range schemas {
		if len(schema.XKubernetesGroupVersionKind) > 0 {
			// Resource has "x-kubernetes-group-version-kind" parameter
			//filteredSchemas[resourceName] = schema
			var newKind Kind
			newKind.Name = schema.XKubernetesGroupVersionKind[0].Kind
			newKind.Description = schema.Description
			newKind.Type = schema.Type
			kinds = append(kinds, newKind)
		}
	}

	// Print the list of kinds
	return kinds, nil
}

func GetKindDescription(endpointURL, path, kind string) (string, error) {
	// Filter out resources with "x-kubernetes-group-version-kind" parameter
	var description string

	schemas, err := GetSchemas(endpointURL, path)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	for _, schema := range schemas {
		if len(schema.XKubernetesGroupVersionKind) > 0 {
			// Resource has "x-kubernetes-group-version-kind" parameter
			if schema.XKubernetesGroupVersionKind[0].Kind == kind {
				description = schema.Description
			}
		}
	}

	// Print the list of kinds
	return description, nil
}
