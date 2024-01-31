package ntnx

import (
	apiv1 "k8s.io/api/core/v1"
	k8scp "k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

type nodepool struct {
}

func (np *nodepool) MaxSize() int {
	return 5
}

func (np *nodepool) MinSize() int {
	return 0
}

// // TargetSize returns the current target size of the node group. It is possible that the
// // number of nodes in Kubernetes is different at the moment but should be equal
// // to Size() once everything stabilizes (new nodes finish startup and registration or
// // removed nodes are deleted completely). Implementation required.
func (np *nodepool) TargetSize() (int, error) {
	// TO DO: msp api
	return 2, nil
}

// // IncreaseSize increases the size of the node group. To delete a node you need
// // to explicitly name it and use DeleteNode. This function should wait until
// // node group size is updated. Implementation required.
func (np *nodepool) IncreaseSize(delta int) error {
	// To Do: node extension api
	return nil
}

// // DeleteNodes deletes nodes from this node group. Error is returned either on
// // failure or if the given node doesn't belong to this node group. This function
// // should wait until node group size is updated. Implementation required.
func (np *nodepool) DeleteNodes([]*apiv1.Node) error {
	// No need to implement for hackathon
	return nil
}

// // DecreaseTargetSize decreases the target size of the node group. This function
// // doesn't permit to delete any existing node and can be used only to reduce the
// // request for new nodes that have not been yet fulfilled. Delta should be negative.
// // It is assumed that cloud provider will not delete the existing nodes when there
// // is an option to just decrease the target. Implementation required.
func (np *nodepool) DecreaseTargetSize(delta int) error {
	// No need to implement for hackathon ?????
	return nil
}

// // Id returns an unique identifier of the node group.
func (np *nodepool) Id() string {
	// TO DO: get nodepool uuid
	return "211"
}

func (np *nodepool) Debug() string {
	return "NOT IM"
}

// // Nodes returns a list of all nodes that belong to this node group.
// // It is required that Instance objects returned by this method have Id field set.
// // Other fields are optional.
// // This list should include also instances that might have not become a kubernetes node yet.
func (np *nodepool) Nodes() ([]k8scp.Instance, error) {
	// Get nodepool while nodepool expansion is happening?
	return nil, nil
}

// // TemplateNodeInfo returns a schedulerframework.NodeInfo structure of an empty
// // (as if just started) node. This will be used in scale-up simulations to
// // predict what would a new node look like if a node group was expanded. The returned
// // NodeInfo is expected to have a fully populated Node object, with all of the labels,
// // capacity and allocatable information as well as all pods that are started on
// // the node by default, using manifest (most likely only kube-proxy). Implementation optional.
func (np *nodepool) TemplateNodeInfo() (*schedulerframework.NodeInfo, error) {
	return nil, k8scp.ErrNotImplemented
}

// // Exist checks if the node group really exists on the cloud provider side. Allows to tell the
// // theoretical node group from the real one. Implementation required.
func (np *nodepool) Exist() bool {
	return true
}

// 	// Create creates the node group on the cloud provider side. Implementation optional.
// 	Create() (NodeGroup, error)

// 	// Delete deletes the node group on the cloud provider side.
// 	// This will be executed only for autoprovisioned node groups, once their size drops to 0.
// 	// Implementation optional.
// 	Delete() error

// 	// Autoprovisioned returns true if the node group is autoprovisioned. An autoprovisioned group
// 	// was created by CA and can be deleted when scaled to 0.
// 	Autoprovisioned() bool

// 	// GetOptions returns NodeGroupAutoscalingOptions that should be used for this particular
// 	// NodeGroup. Returning a nil will result in using default options.
// 	// Implementation optional.
// 	GetOptions(defaults config.NodeGroupAutoscalingOptions) (*config.NodeGroupAutoscalingOptions, error)
// }
