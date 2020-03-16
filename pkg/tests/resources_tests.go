package tests

import (
	"github.com/flacatus/operator-tests/pkg/metadata"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = ginkgo.Describe("Che Operator Tests", func() {
	ginkgo.It("Check if CRD already exist in Cluster", func() {
		apiextensions, err := clientset.NewForConfig(configuration)
		Expect(err).NotTo(HaveOccurred())
		// Make sure the CRD exist in cluster
		_, err = apiextensions.ApiextensionsV1beta1().CustomResourceDefinitions().Get(CrName, metav1.GetOptions{})

		if err != nil {
			metadata.Instance.FoundCRD = false
		} else {
			metadata.Instance.FoundCRD = true
		}

		Expect(err).NotTo(HaveOccurred())
	})
	ginkgo.It("Check if postgres PVC was created", func() {
		_, err := clienset.CoreV1().PersistentVolumeClaims(operatorNamespace).Get(PostgresPVCName, metav1.GetOptions{})

		if err != nil {
			panic(err)
		}
		Expect(err).NotTo(HaveOccurred())
	})
})
