package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	//"github.com/rancher/kontainer-engine/drivers/util"
	"github.com/rancher/kontainer-engine/types"
	"github.com/sirupsen/logrus"
	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/rest"
)

const (
	retries          = 5
	pollInterval     = 30
	defaultNamespace = "cattle-system"
)

type XablauDriver struct {
	driverCapabilities types.Capabilities
}

func (d *XablauDriver) ETCDRemoveSnapshot(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) error {
	panic("implement me")
}

// GetDriverCreateOptions returns cli flags that are used in create
func (d *XablauDriver) GetDriverCreateOptions(ctx context.Context) (*types.DriverFlags, error) {
	driverFlag := types.DriverFlags{
		Options: make(map[string]*types.Flag),
	}
    // fillCreateOptions(&driverFlag)

	return &driverFlag, nil
}

// GetDriverUpdateOptions returns cli flags that are used in update
func (d *XablauDriver) GetDriverUpdateOptions(ctx context.Context) (*types.DriverFlags, error) {
	driverFlag := types.DriverFlags{
		Options: make(map[string]*types.Flag),
	}
	driverFlag.Options["description"] = &types.Flag{
		Type:  types.StringType,
		Usage: "An optional description of this cluster",
	}

	return &driverFlag, nil
}

// Create creates the cluster. clusterInfo is only set when we are retrying a failed or interrupted create
func (d *XablauDriver) Create(ctx context.Context, opts *types.DriverOptions, clusterInfo *types.ClusterInfo) (info *types.ClusterInfo, rtnerr error) {
	logrus.Info("creating new cluster")

	//resource cleanup defer
	defer func() {
		// WIP
	}()

	return nil, nil
}

// Update updates the cluster
func (d *XablauDriver) Update(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions) (rtn *types.ClusterInfo, rtnerr error) {
	defer func() {
		if rtnerr != nil {
			logrus.WithError(rtnerr).Info("update return error")
		}
	}()
	logrus.Info("Starting update")
	//
	logrus.Info("update cluster success")
	return clusterInfo, nil
}

// PostCheck does post action after provisioning
func (d *XablauDriver) PostCheck(ctx context.Context, clusterInfo *types.ClusterInfo) (*types.ClusterInfo, error) {
	logrus.Infof("Starting post-check")
	//
	logrus.Info("post-check completed successfully")
	logrus.Debugf("info: %v", *clusterInfo)

	return clusterInfo, nil
}

// Remove removes the cluster
func (d *XablauDriver) Remove(ctx context.Context, clusterInfo *types.ClusterInfo) error {
	return nil
}

func (d *XablauDriver) GetVersion(ctx context.Context, clusterInfo *types.ClusterInfo) (*types.KubernetesVersion, error) {
	capabilities, err := d.GetCapabilities(ctx)
	if err != nil || !capabilities.HasGetVersionCapability() {
		return nil, err
	}
	//
	return nil, nil
}
func (d *XablauDriver) SetVersion(ctx context.Context, clusterInfo *types.ClusterInfo, version *types.KubernetesVersion) error {
	return errors.New("not supported")
}
func (d *XablauDriver) GetClusterSize(ctx context.Context, clusterInfo *types.ClusterInfo) (*types.NodeCount, error) {
	capabilities, err := d.GetCapabilities(ctx)
	if err != nil || !capabilities.HasGetClusterSizeCapability() {
		return nil, err
	}

	count := &types.NodeCount{Count: 0}
	return count, nil
}
func (d *XablauDriver) SetClusterSize(ctx context.Context, clusterInfo *types.ClusterInfo, count *types.NodeCount) error {
	capabilities, err := d.GetCapabilities(ctx)
	if err != nil || !capabilities.HasSetClusterSizeCapability() {
		return err
	}
	// set cluster size

	return nil
}

func (d *XablauDriver) GetCapabilities(ctx context.Context) (*types.Capabilities, error) {
	return &d.driverCapabilities, nil
}

func (d *XablauDriver) GetK8SCapabilities(ctx context.Context, opts *types.DriverOptions) (*types.K8SCapabilities, error) {
	return &types.K8SCapabilities{
		L4LoadBalancer: &types.LoadBalancerCapabilities{
			Enabled: false,
		},
		NodePoolScalingSupported: false,
	}, nil
}

func NewDriver() types.Driver {
	driver := &XablauDriver{
		driverCapabilities: types.Capabilities{
			Capabilities: make(map[int64]bool),
		},
	}

	driver.driverCapabilities.AddCapability(types.GetVersionCapability)
	driver.driverCapabilities.AddCapability(types.GetClusterSizeCapability)
	driver.driverCapabilities.AddCapability(types.SetClusterSizeCapability)

	return driver
}

func (d *XablauDriver) RemoveLegacyServiceAccount(ctx context.Context, clusterInfo *types.ClusterInfo) error {
	return fmt.Errorf("not implemented")
}

func (d *XablauDriver) ETCDSave(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) error {
	return fmt.Errorf("ETCD backup operations are not implemented")
}

func (d *XablauDriver) ETCDRestore(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) (*types.ClusterInfo, error) {
	return nil, fmt.Errorf("ETCD backup operations are not implemented")
}
