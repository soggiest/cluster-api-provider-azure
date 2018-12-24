package main

import (
	"flag"

	"k8s.io/klog"
	//"github.com/golang/glog"

	"github.com/azure-provider/cmd/versioninfo"
	"github.com/platform9/azure-provider/pkg/cloud/azure/actuators/cluster"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/cmd"
	"sigs.k8s.io/cluster-api/pkg/apis/cluster/common"
)

// initLogs is a temporary hack to enable proper logging until upstream dependencies
// are migrated to fully utilize klog instead of glog.
func initLogs() {
	flag.Set("logtostderr", "true")
	flags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(flags)
	flags.Set("alsologtostderr", "true")
	flags.Set("v", "4")
	flag.Parse()
}

func registerCustomCommands() {
	cmd.RootCmd.AddCommand(versioninfo.VersionCmd())
}

func main() {
	initLogs()
	clusterActuator := cluster.NewActuator(cluster.ActuatorParams{})
	common.RegisterClusterProvisioner("azure", clusterActuator)
	registerCustomCommands()
	cmd.Execute()
}

//func main() {
//	var err error
//	machine.Actuator, err = machine.NewMachineActuator(machine.MachineActuatorParams{})
//	if err != nil {
//		glog.Fatalf("Error creating cluster provisioner for azure : %v", err)
//	}
//	common.RegisterClusterProvisioner(machine.ProviderName, machine.Actuator)
//	cmd.Execute()
//}
