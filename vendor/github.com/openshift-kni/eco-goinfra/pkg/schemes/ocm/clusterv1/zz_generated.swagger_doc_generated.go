package clusterv1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_ClientConfig = map[string]string{
	"":         "ClientConfig represents the apiserver address of the managed cluster.",
	"url":      "URL is the URL of apiserver endpoint of the managed cluster.",
	"caBundle": "CABundle is the ca bundle to connect to apiserver of the managed cluster. System certs are used if it is not set.",
}

func (ClientConfig) SwaggerDoc() map[string]string {
	return map_ClientConfig
}

var map_ManagedCluster = map[string]string{
	"":       "ManagedCluster represents the desired state and current status of a managed cluster. ManagedCluster is a cluster-scoped resource. The name is the cluster UID.\n\nThe cluster join process is a double opt-in process. See the following join process steps:\n\n1. The agent on the managed cluster creates a CSR on the hub with the cluster UID and agent name. 2. The agent on the managed cluster creates a ManagedCluster on the hub. 3. The cluster admin on the hub cluster approves the CSR for the UID and agent name of the ManagedCluster. 4. The cluster admin sets the spec.acceptClient of the ManagedCluster to true. 5. The cluster admin on the managed cluster creates a credential of the kubeconfig for the hub cluster.\n\nAfter the hub cluster creates the cluster namespace, the klusterlet agent on the ManagedCluster pushes the credential to the hub cluster to use against the kube-apiserver of the ManagedCluster.",
	"spec":   "Spec represents a desired configuration for the agent on the managed cluster.",
	"status": "Status represents the current status of joined managed cluster",
}

func (ManagedCluster) SwaggerDoc() map[string]string {
	return map_ManagedCluster
}

var map_ManagedClusterClaim = map[string]string{
	"":      "ManagedClusterClaim represents a ClusterClaim collected from a managed cluster.",
	"name":  "Name is the name of a ClusterClaim resource on managed cluster. It's a well known or customized name to identify the claim.",
	"value": "Value is a claim-dependent string",
}

func (ManagedClusterClaim) SwaggerDoc() map[string]string {
	return map_ManagedClusterClaim
}

var map_ManagedClusterList = map[string]string{
	"":         "ManagedClusterList is a collection of managed cluster.",
	"metadata": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
	"items":    "Items is a list of managed clusters.",
}

func (ManagedClusterList) SwaggerDoc() map[string]string {
	return map_ManagedClusterList
}

var map_ManagedClusterSpec = map[string]string{
	"":                            "ManagedClusterSpec provides the information to securely connect to a remote server and verify its identity.",
	"managedClusterClientConfigs": "ManagedClusterClientConfigs represents a list of the apiserver address of the managed cluster. If it is empty, the managed cluster has no accessible address for the hub to connect with it.",
	"hubAcceptsClient":            "hubAcceptsClient represents that hub accepts the joining of Klusterlet agent on the managed cluster with the hub. The default value is false, and can only be set true when the user on hub has an RBAC rule to UPDATE on the virtual subresource of managedclusters/accept. When the value is set true, a namespace whose name is the same as the name of ManagedCluster is created on the hub. This namespace represents the managed cluster, also role/rolebinding is created on the namespace to grant the permision of access from the agent on the managed cluster. When the value is set to false, the namespace representing the managed cluster is deleted.",
	"leaseDurationSeconds":        "LeaseDurationSeconds is used to coordinate the lease update time of Klusterlet agents on the managed cluster. If its value is zero, the Klusterlet agent will update its lease every 60 seconds by default",
	"taints":                      "Taints is a property of managed cluster that allow the cluster to be repelled when scheduling. Taints, including 'ManagedClusterUnavailable' and 'ManagedClusterUnreachable', can not be added/removed by agent running on the managed cluster; while it's fine to add/remove other taints from either hub cluser or managed cluster.",
}

func (ManagedClusterSpec) SwaggerDoc() map[string]string {
	return map_ManagedClusterSpec
}

var map_ManagedClusterStatus = map[string]string{
	"":              "ManagedClusterStatus represents the current status of joined managed cluster.",
	"conditions":    "Conditions contains the different condition statuses for this managed cluster.",
	"capacity":      "Capacity represents the total resource capacity from all nodeStatuses on the managed cluster.",
	"allocatable":   "Allocatable represents the total allocatable resources on the managed cluster.",
	"version":       "Version represents the kubernetes version of the managed cluster.",
	"clusterClaims": "ClusterClaims represents cluster information that a managed cluster claims, for example a unique cluster identifier (id.k8s.io) and kubernetes version (kubeversion.open-cluster-management.io). They are written from the managed cluster. The set of claims is not uniform across a fleet, some claims can be vendor or version specific and may not be included from all managed clusters.",
}

func (ManagedClusterStatus) SwaggerDoc() map[string]string {
	return map_ManagedClusterStatus
}

var map_ManagedClusterVersion = map[string]string{
	"":           "ManagedClusterVersion represents version information about the managed cluster.",
	"kubernetes": "Kubernetes is the kubernetes version of managed cluster.",
}

func (ManagedClusterVersion) SwaggerDoc() map[string]string {
	return map_ManagedClusterVersion
}

var map_Taint = map[string]string{
	"":          "The managed cluster this Taint is attached to has the \"effect\" on any placement that does not tolerate the Taint.",
	"key":       "Key is the taint key applied to a cluster. e.g. bar or foo.example.com/bar. The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)",
	"value":     "Value is the taint value corresponding to the taint key.",
	"effect":    "Effect indicates the effect of the taint on placements that do not tolerate the taint. Valid effects are NoSelect, PreferNoSelect and NoSelectIfNew.",
	"timeAdded": "TimeAdded represents the time at which the taint was added.",
}

func (Taint) SwaggerDoc() map[string]string {
	return map_Taint
}

// AUTO-GENERATED FUNCTIONS END HERE
