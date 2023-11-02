package main

import (
	"fmt"

	clusterapi "github.com/brikkel/k8s-api-spec/cluster-api"
	openapiv2 "github.com/brikkel/k8s-api-spec/openapi-v2"
)

// Define a struct to store information about the field in Pods.
type FieldInfo struct {
	FieldName        string
	FieldType        string
	FieldDescription string
}

// Define a struct to store resource kinds.
type ResourceKind struct {
	Kind string
}

func main() {
	clientset := clusterapi.AttachMinikube()
	clusterapi.ListResourceKinds(clientset)
	// test()
	fmt.Println("---------------------------")
	schema := openapiv2.GetSchema("http://localhost:12345", "Deployment")
	fmt.Println("---------------------------")
	schemaMap, ok := schema.(map[string]interface{})
	openapiv2.GetLevel1Keys(schemaMap)
	if !ok {
		fmt.Println("Schema is not a map[string]interface{}")
		return
	}

	keyToLookup := "properties"
	value, keyExists := schemaMap[keyToLookup]
	if keyExists {
		fmt.Printf("Value for key '%s': %v\n", keyToLookup, value)
	} else {
		fmt.Printf("Key '%s' not found in the schema\n", keyToLookup)
	}

	propertiesKey := "properties" // The key you want to extract keys from
	fmt.Println("---------------------------")
	// Lookup the properties and check if they exist
	properties, keyExists := schemaMap[propertiesKey].(map[string]interface{})
	if keyExists {
		openapiv2.GetLevel1Keys(properties)
	} else {
		fmt.Printf("Key '%s' not found in the schema\n", propertiesKey)
	}
}
