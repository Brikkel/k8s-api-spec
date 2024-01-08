package main

import (
	"encoding/json"
	"fmt"

	clusterapi "github.com/brikkel/k8s-api-spec/cluster-api"
	openapiv2 "github.com/brikkel/k8s-api-spec/openapi-v2"
	oav3 "github.com/brikkel/k8s-api-spec/openapi-v3"
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
	testOpenapiv3()
}

func testOpenapiv3() {
	endpointURL := "http://localhost:8001"

	// KINDS

	// kinds, err := openapiv3.GetAllKinds("http://localhost:8001")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// jsonKinds, err := json.Marshal(kinds)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(jsonKinds))

	// schema, err := oav3.GetSchema(endpointURL, "/openapi/v3/api/v1", "io.k8s.api.core.v1.Pod")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// jsonSchema, err := json.Marshal(schema)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(jsonSchema))

	// newSchema, err := oav3.GetSchema(endpointURL, "/openapi/v3/api/v1", "io.k8s.api.core.v1.Pod")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// jsonSchema, err := json.Marshal(newSchema)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(jsonSchema))

	// PARAMETERS
	podParameters, err := oav3.GetResource(endpointURL, "/openapi/v3/api/v1", "io.k8s.api.core.v1.Pod")
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonPodParameters, err := json.Marshal(podParameters)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonPodParameters))
}

func testOpenapiv2() {
	clientset := clusterapi.AttachMinikube()
	clusterapi.ListResourceKinds(clientset)
	// test()
	fmt.Println("---------------------------")
	schema := openapiv2.GetSchema("http://localhost:12345", "Pod")
	fmt.Println("---------------------------")
	schemaMap, ok := schema.(map[string]interface{})
	openapiv2.GetLevel1Keys(schemaMap)
	if !ok {
		fmt.Println("Schema is not a map[string]interface{}")
		return
	}

	keyToLookup := "description"
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
