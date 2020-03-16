package tests

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	configuration, err = kubernetesConfig()
	clienset, _        = getClientSet()
)

func kubernetesConfig() (config *rest.Config, err error) {
	// Check if we run the tests insides of cluster or locally
	host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")

	if len(host) != 0 || len(port) != 0 {
		// creates the in-cluster config
		config, err := rest.InClusterConfig()

		return config, err
	} else {
		var kubeconfig *string

		if home := homeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)

		return config, err
	}
}

func getClientSet() (clientset *kubernetes.Clientset, err2 error) {
	// create the clientset
	clientset, err = kubernetes.NewForConfig(configuration)

	return clientset, err
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
