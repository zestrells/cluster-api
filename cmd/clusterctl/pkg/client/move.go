/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

func (c *clusterctlClient) Move(options MoveOptions) error {
	// Get the client for interacting with the source management cluster.
	fromCluster, err := c.clusterClientFactory(options.FromKubeconfig)
	if err != nil {
		return err
	}

	// Ensures the custom resource definitions required by clusterctl are in place.
	if err := fromCluster.ProviderInventory().EnsureCustomResourceDefinitions(); err != nil {
		return err
	}

	// Get the client for interacting with the target management cluster.
	toCluster, err := c.clusterClientFactory(options.ToKubeconfig)
	if err != nil {
		return err
	}

	// Ensures the custom resource definitions required by clusterctl are in place
	if err := toCluster.ProviderInventory().EnsureCustomResourceDefinitions(); err != nil {
		return err
	}

	if err := fromCluster.ObjectMover().Move(options.Namespace, toCluster); err != nil {
		return err
	}

	return nil
}
