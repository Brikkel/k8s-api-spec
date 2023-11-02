package clusterapi

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func AttachMinikube() *kubernetes.Clientset {
	kubeconfigPath := "C:\\Users\\brikkel\\.kube\\config" // kubeconfig file

	// Build the client configuration from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Println("Error creating configuration:", err)
		return nil
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error creating clientset:", err)
		return nil
	}

	return clientset
}

func ListResourceKinds(clientset *kubernetes.Clientset) []string {
	// Retrieve a list of resource kinds available in your cluster.
	discoveryClient := clientset.Discovery()
	apiResourceList, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		fmt.Println("Error retrieving resource kinds:", err)
		return nil
	}

	var kinds []string
	for _, apiResource := range apiResourceList {
		for _, resource := range apiResource.APIResources {
			kinds = append(kinds, resource.Kind)
		}
	}

	// Debugging
	for _, kind := range kinds {
		fmt.Println(kind)
	}

	return kinds
}
