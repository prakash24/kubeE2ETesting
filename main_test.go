package e2e

import (
    "os"
	"testing"
    "sigs.k8s.io/e2e-framework/pkg/env"
	"k8s.io/client-go/util/homedir"

)

var testenv env.Environment


func TestMain(m *testing.M) {
    home := homedir.HomeDir()
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if home != "" && kubeconfigPath == "" {
		kubeconfigPath = home + "/.kube/config"
	}
    testenv = env.NewWithKubeConfig(kubeconfigPath)
    os.Exit(testenv.Run(m))
}

