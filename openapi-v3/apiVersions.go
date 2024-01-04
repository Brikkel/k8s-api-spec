// A package that ontains function to interprate the Kubernetes openapi spec v3
//
/*
The k8s-API-Spec-openapiv3 GO package will allow users to effectively find kubernetes
resources and their information. An exmple usage is to navigate all the resources that can be described in a configuration .yaml file.
*/
package openapiv3

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// A struct to represent the base API v3 response in JSON structure.
type OpenAPIV3BaseResponse struct {
	Paths map[string]struct {
		ServerRelativeURL string `json:"serverRelativeURL"`
	} `json:"paths"`
}

// The kubernetes API is split up into several smaller API's, this returns a list of the available ones when provided with a k8s api endpoint.
func GetAPIVersions(endpointURL string) ([]string, error) {
	paths, err := GetOpenAPIV3Paths(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("Can not acquire API paths: %w", err)
	}

	// List of APIVSersions
	var APIVersions []string

	for _, path := range paths {
		parts := strings.Split(path, "/")
		APIversion := strings.Join(parts[1:], "/")

		APIVersions = append(APIVersions, APIversion)
	}

	return APIVersions, nil
}

// This function connects to the k8s api and retreives the base openapi v3 object,
// it then parses through the object to obtain the paths to the different API versions.
func GetOpenAPIV3Paths(endpointURL string) ([]string, error) {
	// Acquire the various openapi v3 paths
	body, err := getJson(endpointURL + "/openapi/v3")
	if err != nil {
		return nil, fmt.Errorf("Error retreiving openapi v3 paths: %w", err)
	}

	// Parse JSON
	var openAPIResp OpenAPIV3BaseResponse
	err = json.Unmarshal(body, &openAPIResp)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %w", err)
	}

	// Extract paths into an array
	var paths []string
	for path := range openAPIResp.Paths {
		// Filter out any non API related paths
		if strings.HasPrefix(path, "api") {
			paths = append(paths, path)
		}
	}

	// Identify unique base paths
	uniquePaths := make(map[string][]string)
	for _, path := range paths {
		basePath := getBasePath(path)
		uniquePaths[basePath] = append(uniquePaths[basePath], path)
	}

	// For each unique base path, find the highest non-beta version
	var filteredPaths []string
	for _, paths := range uniquePaths {
		highestVersionPath := findHighestNonBetaVersion(paths)
		filteredPaths = append(filteredPaths, highestVersionPath)
	}

	// Return the filtered paths
	return filteredPaths, nil
}

// getBasePath extracts the base path (without version) from a given path
func getBasePath(path string) string {
	// Split the path by "/"
	parts := strings.Split(path, "/")

	// Check if there are more than one segment in the path
	if len(parts) > 2 {
		// Join all segments except the last one
		return strings.Join(parts[:len(parts)-1], "/")
	}

	// If there is only one segment, return the original path
	return path
}

// findHighestNonBetaVersion finds the path with the highest non-beta version
func findHighestNonBetaVersion(paths []string) string {
	var nonBetaPaths []string

	// Filter out paths with beta versions
	for _, path := range paths {
		if !strings.Contains(path, "beta") {
			nonBetaPaths = append(nonBetaPaths, path)
		}
	}

	// Sort non-beta paths in descending order
	sort.Sort(sort.Reverse(sort.StringSlice(nonBetaPaths)))

	// Return the highest non-beta version path
	if len(nonBetaPaths) > 0 {
		return nonBetaPaths[0]
	}
	return ""
}

///////////////////////
// Depricated
///////////////////////

func getAPIVersionPaths(endpointURL string) ([]string, error) {
	paths, err := GetOpenAPIV3Paths(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("Can not acquire API paths: %w", err)
	}

	var APIVersionPaths []string

	for _, path := range paths {
		parts := strings.Split(path, "/")
		if len(parts) > 1 {
			APIVersionPaths = append(APIVersionPaths, path)
		}
	}

	return APIVersionPaths, nil
}
