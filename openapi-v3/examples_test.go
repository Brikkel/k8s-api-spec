package openapiv3

import (
	"fmt"
)

func ExampleGetKindDescription() {
	descr, err := GetKindDescription("http://localhost:8001", "/openapi/v3/api/v1", "Pod")
	if err != nil {
		panic(err)
	}
	fmt.Println(descr)
	// Output:
	// Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.
}

func ExampleGetKindsPerVersion() {
	kinds, err := GetKindsPerVersion("http://localhost:8001", "/openapi/v3/api/v1")
	if err != nil {
		panic(err)
	}

	fmt.Println(kinds)
	// Output:
	// [APIResourceList Binding ComponentStatus ComponentStatusList ConfigMap ConfigMapList DeleteOptions Endpoints EndpointsList Event EventList Eviction LimitRange LimitRangeList Namespace NamespaceList Node NodeList PersistentVolume PersistentVolumeClaim PersistentVolumeClaimList PersistentVolumeList Pod PodList PodTemplate PodTemplateList ReplicationController ReplicationControllerList ResourceQuota ResourceQuotaList Scale Secret SecretList Service ServiceAccount ServiceAccountList ServiceList Status TokenRequest WatchEvent]
}

func ExampleGetAllKinds() {
	allKinds, err := GetAllKinds("http://localhost:8001")
	if err != nil {
		panic(err)
	}

	fmt.Println(allKinds)
	// Output:
	// [APIGroup APIGroup APIGroupList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIResourceList APIVersions Binding CSIDriver CSIDriverList CSINode CSINodeList CSIStorageCapacity CSIStorageCapacityList CertificateSigningRequest CertificateSigningRequestList ClusterRole ClusterRoleBinding ClusterRoleBindingList ClusterRoleList ComponentStatus ComponentStatusList ConfigMap ConfigMapList ControllerRevision ControllerRevisionList CronJob CronJobList CustomResourceDefinition CustomResourceDefinitionList DaemonSet DaemonSetList DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions DeleteOptions Deployment DeploymentList EndpointSlice EndpointSliceList Endpoints EndpointsList Event Event EventList EventList Eviction HorizontalPodAutoscaler HorizontalPodAutoscalerList Ingress IngressClass IngressClassList IngressList Job JobList Lease LeaseList LimitRange LimitRangeList LocalSubjectAccessReview MutatingWebhookConfiguration MutatingWebhookConfigurationList Namespace NamespaceList NetworkPolicy NetworkPolicyList Node NodeList PersistentVolume PersistentVolumeClaim PersistentVolumeClaim PersistentVolumeClaimList PersistentVolumeList Pod PodDisruptionBudget PodDisruptionBudgetList PodList PodTemplate PodTemplateList PriorityClass PriorityClassList ReplicaSet ReplicaSetList ReplicationController ReplicationControllerList ResourceQuota ResourceQuotaList Role RoleBinding RoleBindingList RoleList RuntimeClass RuntimeClassList Scale Scale Secret SecretList SelfSubjectAccessReview SelfSubjectRulesReview Service ServiceAccount ServiceAccountList ServiceList StatefulSet StatefulSetList Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status Status StorageClass StorageClassList SubjectAccessReview TokenRequest TokenReview ValidatingWebhookConfiguration ValidatingWebhookConfigurationList VolumeAttachment VolumeAttachmentList WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent WatchEvent]
}
