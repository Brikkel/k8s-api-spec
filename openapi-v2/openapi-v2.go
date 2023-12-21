package openapiv2

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getSwaggerData(kubeAPIAddress string) map[string]interface{} {
	// Construct the Swagger UI URL to fetch the Swagger 2.0 JSON documentation
	swaggerURL := fmt.Sprintf("%s/openapi/v2", kubeAPIAddress)

	// Make an HTTP GET request to the Swagger UI endpoint
	resp, err := http.Get(swaggerURL)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// Read the Swagger 2.0 JSON documentation
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}

	// Unmarshal the JSON
	var swaggerData map[string]interface{}
	if err := json.Unmarshal(body, &swaggerData); err != nil {
		fmt.Printf("Error unmarshalling Swagger JSON: %v\n", err)
		return nil
	}

	return swaggerData
}

func GetSchema(kubeAPIAddress string, resource string) interface{} {
	swaggerData := getSwaggerData(kubeAPIAddress)
	// Find and print the "resource" schema as JSON
	definitions := swaggerData["definitions"].(map[string]interface{})
	resourceSchema := findSchemaByKind(definitions, resource)

	if resourceSchema != nil {

		// This is how the schema can be navigated
		//fmt.Println(resourceSchema.(map[string]interface{})["properties"].(map[string]interface{})["spec"].(map[string]interface{})["$ref"])

		return resourceSchema
	} else {
		fmt.Println("Resource schema not found in Swagger JSON documentation")
		return nil
	}
}

func marshalToJson(schema interface{}) []byte {
	// Marshal the schema to JSON
	deploymentSchemaJSON, err := json.Marshal(schema)
	if err != nil {
		fmt.Printf("Error marshalling schema to JSON: %v\n", err)
		return nil
	}
	fmt.Println("Schema for Deployment as JSON:")
	fmt.Println(string(deploymentSchemaJSON))
	return deploymentSchemaJSON
}

// Find the schema definition by kind
func findSchemaByKind(definitions map[string]interface{}, kind string) interface{} {
	for definitionName, schema := range definitions {
		if getKindFromDefinitionName(definitionName) == kind {
			return schema
		}
	}
	return nil
}

// Extracts the kind from a Swagger definition name
func getKindFromDefinitionName(definitionName string) string {
	parts := strings.Split(definitionName, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

// Get the keys of a kind
func GetLevel1Keys(jsonData map[string]interface{}) []string {
	keys := make([]string, 0, len(jsonData))
	for key := range jsonData {
		keys = append(keys, key)
	}

	fmt.Println("Level 1 Keys:")
	for _, key := range keys {
		fmt.Println(key)
	}

	return keys
}
