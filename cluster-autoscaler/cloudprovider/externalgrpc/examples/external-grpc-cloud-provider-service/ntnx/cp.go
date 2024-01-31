package ntnx

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	k8scp "k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
)

type ntnxCP struct {
}

func (cp *ntnxCP) Name() string {
	return "ntnx"
}

func (cp *ntnxCP) NodeGroups() []k8scp.NodeGroup {
	// TO DO
	// /v2/nodepool/list for cmsp

	// convert msp nodepools to NodeGroup
	return nil
}

func (cp *ntnxCP) NodeGroupForNode(*apiv1.Node) (k8scp.NodeGroup, error) {
	// TO DO
	return nil, nil
}

func (cp *ntnxCP) HasInstance(*apiv1.Node) (bool, error) {
	// 	// TO DO
	// always true since we will disable scaledown or talk to msp???
	return true, nil
}

// GetResourceLimiter returns struct containing limits (max, min) for resources (cores, memory etc.).
func (cp *ntnxCP) GetResourceLimiter() (*k8scp.ResourceLimiter, error) {
	resourceLimiter := k8scp.NewResourceLimiter(
		map[string]int64{k8scp.ResourceNameCores: 1, k8scp.ResourceNameMemory: 10000000},
		map[string]int64{k8scp.ResourceNameCores: 10, k8scp.ResourceNameMemory: 100000000})
	return resourceLimiter, nil
	// took from aws_cloud_provider_test.go
}

func (cp *ntnxCP) GetNodeGpuConfig(*apiv1.Node) *k8scp.GpuConfig {
	return nil
}

func (cp *ntnxCP) GetAvailableGPUTypes() map[string]struct{} {
	return nil
}

// GPULabel returns the label added to nodes with GPU resource.
func (cp *ntnxCP) GPULabel() string {
	return "GPU"
	// not sure about this.
}

func (cp *ntnxCP) Cleanup() error {
	return nil
}

func (cp *ntnxCP) Refresh() error {
	return nil
}

// -----------------
func (cp *ntnxCP) NewNodeGroup(machineType string, labels map[string]string, systemLabels map[string]string,
	taints []apiv1.Taint, extraResources map[string]resource.Quantity) (k8scp.NodeGroup, error) {
	// Implementation optional.
	return nil, k8scp.ErrNotImplemented
}

func (cp *ntnxCP) Pricing() (k8scp.PricingModel, errors.AutoscalerError) {
	// optional
	return nil, k8scp.ErrNotImplemented
}

func (cp *ntnxCP) GetAvailableMachineTypes() ([]string, error) {
	//optional
	return nil, k8scp.ErrNotImplemented
}
