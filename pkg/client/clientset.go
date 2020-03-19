package client

import (
	orgv1 "github.com/eclipse/che-operator/pkg/apis/org/v1"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	groupName        = "org.eclipse.che"
)

type C struct {
	rest *rest.Config
}

type CRClient struct {
	rest rest.Interface
}


// New creates H, a client used to expose common testing functions.
func New() *C {
	helper := &C{}

	return helper
}

// Kube returns the clientset for Kubernetes upstream.
func (c *C) Kube() kubernetes.Interface {
	cfg, err := config.GetConfig()
	client, err := kubernetes.NewForConfig(cfg)
	Expect(err).ShouldNot(HaveOccurred(), "failed to configure Kubernetes clientset")
	return client
}

func (c *C) KubeRest() rest.Interface {
	cfg, _ := config.GetConfig()
	cfg.ContentConfig.GroupVersion = &schema.GroupVersion{Group: groupName, Version: orgv1.SchemeGroupVersion.Version}
	cfg.APIPath = "/apis"
	cfg.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	cfg.UserAgent = rest.DefaultKubernetesUserAgent()

	//client, _ := kubernetes.NewForConfig(cfg)
	client, _ := rest.RESTClientFor(cfg)
	return client
}
