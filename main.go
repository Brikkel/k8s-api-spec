package main

import (
	"encoding/json"
	"fmt"

	clusterapi "github.com/brikkel/k8s-api-spec/cluster-api"
	openapiv2 "github.com/brikkel/k8s-api-spec/openapi-v2"
	openapiv3 "github.com/brikkel/k8s-api-spec/openapi-v3"
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
	// paths, err := openapiv3.GetOpenAPIV3Paths("http://localhost:8001")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(paths)

	// fmt.Println(openapiv3.GetAllKinds("http://localhost:8001"))

	// openapiv3.GetKindsPerVersion("http://localhost:8001", "/openapi/v3/api/v1")
	// openapiv3.GetKindsPerVersion("http://localhost:8001", "/openapi/v3/apis/apps/v1")

	// transform schemas into json
	// schemasJson, err := openapiv3.GetSchemasJson("http://localhost:8001", "/openapi/v3/api/v1")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(schemasJson)

	// description, err := openapiv3.GetKindDescription("http://localhost:8001", "/openapi/v3/api/v1", "Pod")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(description)

	kinds, err := openapiv3.GetAllKinds("http://localhost:8001")
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonKinds, err := json.Marshal(kinds)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonKinds))

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
