package openapiv3

import (
	"fmt"
	"net/url"
	"strings"
)

type Resource struct {
	FullResourceName string
	Path             string
	Name             string
	Description      string
	Type             string
	Parameters       []Parameter
}

type Parameter struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	Type              string `json:"type"`
	Required          bool   `json:"required"`
	Default           string `json:"default,omitempty"`
	ResourceReference string `json:"resourceReference,omitempty"`
	Items             struct {
		Default           string `json:"default,omitempty"`
		ResourceReference string `json:"resourceReference,omitempty"`
		Type              string `json:"type,omitempty"`
	} `json:"items,omitempty"`
}

// GetResource returns a resource from a map of schemas
func GetResource(endpointURL, path, FullResourceName string) (*Resource, error) {
	// get the path from the raw path
	rawPath, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// get the schemas for the path
	schemas, err := GetSchemas(endpointURL, rawPath)
	if err != nil {
		return nil, err
	}

	// create a new resource
	var newResource Resource

	// get the schema for the resource
	schema, ok := schemas[FullResourceName]
	if !ok {
		return nil, fmt.Errorf("resource %s not found", FullResourceName)
	}

	// set the properties of the resource
	newResource.FullResourceName = FullResourceName
	newResource.Name = getSimpleName(FullResourceName)
	newResource.Path = rawPath
	newResource.Description = schema.Description
	newResource.Type = schema.Type

	newResource.Parameters = interpretProperties(schema.Properties, schema.Required)

	return &newResource, nil
}

// GetResourceParameters returns a list of parameters for a resource
func GetResourceParameters(endpointURL, path, resourceName string) ([]Parameter, error) {
	// get the schemas for the path
	schemas, err := GetSchemas(endpointURL, path)
	if err != nil {
		return nil, err
	}

	// get the schema for the resource
	schema, ok := schemas[resourceName]
	if !ok {
		return nil, fmt.Errorf("resource %s not found", resourceName)
	}

	return interpretProperties(schema.Properties, schema.Required), nil
}

// interpretProperties interprets the properties of a schema and returns a list of parameters
func interpretProperties(properties map[string]Propertie, required []string) []Parameter {
	var parameters []Parameter

	for key, value := range properties {
		var newParameter Parameter
		newParameter.Name = key
		newParameter.Description = value.Description

		// switch case for different types
		switch value.Type {
		case "":
			// Type is only empty for objects
			newParameter.Type = "object"
			newParameter.ResourceReference = stripReference(value.AllOf[0].Ref)
		case "array":
			newParameter.Type = value.Type

			// check if the items type is empty
			if value.Items.Type == "" {
				// Type is only empty for objects
				newParameter.Items.Type = "object"
				newParameter.Items.ResourceReference = stripReference(value.Items.AllOf[0].Ref)
			} else {
				newParameter.Items.Type = value.Items.Type
			}

			newParameter.Items.Default = fmt.Sprintf("%v", value.Items.Default)
		default:
			newParameter.Type = value.Type
		}

		// check if the parameter is required
		if contains(required, key) {
			newParameter.Required = true
		}

		// convert the interface{} default to a string, or "" if there is no default
		if value.Default == nil {
			newParameter.Default = ""
		} else {
			newParameter.Default = fmt.Sprintf("%v", value.Default)
		}

		parameters = append(parameters, newParameter)
	}

	return parameters
}

// contains checks if a string is in a slice of strings
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// getSimpleName returns the name of a resource from its resource name
func getSimpleName(resourceName string) string {
	// split the resource name by .
	split := strings.Split(resourceName, ".")
	// get the last element of the split
	name := split[len(split)-1]
	// return the name
	return name
}

// Get the resource name from the reference
func stripReference(refrence string) string {
	// Split on /
	split := strings.Split(refrence, "/")

	// Return the last element of the split
	return split[len(split)-1]
}
