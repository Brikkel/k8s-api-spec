package openapiv3

import (
	"fmt"
)

// Struct that represents apiversions with kinds
type ApiVersion struct {
	Name           string
	ApiVersionPath string
	Kinds          []Kind
}

// Struct that represents kinds
type Kind struct {
	ResourceName string
	Name         string
	Description  string
	Type         string
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

	for keyName, schema := range schemas {
		if len(schema.XKubernetesGroupVersionKind) > 0 {
			// Resource has "x-kubernetes-group-version-kind" parameter
			//filteredSchemas[resourceName] = schema
			var newKind Kind
			newKind.ResourceName = keyName
			newKind.Name = schema.XKubernetesGroupVersionKind[0].Kind
			newKind.Description = schema.Description
			newKind.Type = schema.Type
			kinds = append(kinds, newKind)
		}
	}

	// Print the list of kinds
	return kinds, nil
}

// Get the description of a kind, by passing the name of the kind
func GetKindDescriptionByName(endpointURL, path, kindName string) (string, error) {
	var description string

	schemas, err := GetSchemas(endpointURL, path)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	for _, schema := range schemas {
		if len(schema.XKubernetesGroupVersionKind) > 0 {
			// Resource has "x-kubernetes-group-version-kind" parameter
			if schema.XKubernetesGroupVersionKind[0].Kind == kindName {
				description = schema.Description
			}
		}
	}

	// Print the list of kinds
	return description, nil
}
