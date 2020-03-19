package tests

import (
	"github.com/eclipse/che-operator/pkg/controller/che"
	"github.com/flacatus/operator-tests/pkg/client"
	codeready "github.com/flacatus/operator-tests/pkg/codeready"
	"github.com/flacatus/operator-tests/pkg/metadata"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var _  = ginkgo.BeforeSuite(func() {
	client := client.New().KubeRest()

	logrus.Infof("Generating Custom Resource in cluster")
	if err := codeready.CreateCodeReadyCustomResource(client); err != nil {
		logrus.Fatalf("Failed to create custom resources in cluster: %s", err)
	}
	logrus.Infof("Successfully created CodeReady Custom Resources")

	logrus.Infof("Starting to check Code Ready Cluster if is available")
	deploy,_ := codeready.VerifyCodeReadyCustomResource(che.AvailableStatus, client)
	if deploy {
		logrus.Infof("CodeReady Cluster status available")
	}
})

var _ = ginkgo.Describe(osdDescription, func() {
	ginkgo.It("Check if CRD already exist in Cluster", func() {
		cfg, err := config.GetConfig()
		apiextensions, err := clientset.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		// Make sure the CRD exist in cluster
		_, err = apiextensions.ApiextensionsV1beta1().CustomResourceDefinitions().Get(CRDName, metav1.GetOptions{})

		if err != nil {
			metadata.Instance.FoundCRD = false
		} else {
			metadata.Instance.FoundCRD = true
		}

		Expect(err).NotTo(HaveOccurred())
	})
})
