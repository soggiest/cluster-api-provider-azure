/*

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

package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/platform9/azure-provider/pkg/apis"
	"github.com/platform9/azure-provider/pkg/cloud/azure/actuators/cluster"
	"github.com/platform9/azure-provider/pkg/cloud/azure/actuators/machine"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/record"
	clusterapis "sigs.k8s.io/cluster-api/pkg/apis"
	"sigs.k8s.io/cluster-api/pkg/apis/cluster/common"
	"sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
	capicluster "sigs.k8s.io/cluster-api/pkg/controller/cluster"
	capimachine "sigs.k8s.io/cluster-api/pkg/controller/machine"
	//	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

// initLogs is a temporary hack to enable proper logging until upstream dependencies
// are migrated to fully utilize klog instead of glog.
func initLogs() {
	flag.Set("logtostderr", "true")
	flags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(flags)
	flags.Set("alsologtostderr", "true")
	flags.Set("v", "3")
	flag.Parse()
}

func main() {
	initLogs()
	cfg := config.GetConfig()

	// Setup a Manager
	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		klog.Fatalf("Failed to set up overall controller manager: %v", err)
	}

	cs, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Failed to create client from configuration: %v", err)
	}

	if err := prepareEnvironment(); err != nil {
		log.Error(err, "unable to prepare environment for actuators")
		os.Exit(1)
	}

	// Initialize event recorder.
	record.InitFromRecorder(mgr.GetRecorder("aws-controller"))

	// Initialize cluster actuator.
	clusterActuator := cluster.NewActuator(cluster.ActuatorParams{
		Client: cs.ClusterV1alpha1(),
	})

	// Initialize machine actuator.
	machineActuator := machine.NewActuator(machine.ActuatorParams{
		Client: cs.ClusterV1alpha1(),
	})

	// Register our cluster deployer (the interface is in clusterctl and we define the Deployer interface on the actuator)
	common.RegisterClusterProvisioner("aws", clusterActuator)

	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		klog.Fatal(err)
	}

	if err := clusterapis.AddToScheme(mgr.GetScheme()); err != nil {
		klog.Fatal(err)
	}

	capimachine.AddWithActuator(mgr, machineActuator)
	capicluster.AddWithActuator(mgr, clusterActuator)

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		klog.Fatalf("Failed to run manager: %v", err)
	}
}

func prepareEnvironment() error {
	//Parse in environment variables if necessary
	if os.Getenv("AZURE_SUBSCRIPTION_ID") == "" {
		err := godotenv.Load()
		if err == nil && os.Getenv("AZURE_SUBSCRIPTION_ID") == "" {
			return fmt.Errorf("couldn't find environment variable for the Azure subscription: %v", err)
		}
		if err != nil {
			return fmt.Errorf("failed to load environment variables: %v", err)
		}
	}
	return nil
}
