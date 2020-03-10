package tests

import (
	"fmt"

	"github.com/flacatus/operator-osd/pkg/metadata"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

const (
	CrName = "checlusters.org.eclipse.che"
)

var _ = ginkgo.Describe("Che Operator Tests", func() {
	defer ginkgo.GinkgoRecover()
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ginkgo.It("Check checlusters.org.eclipse.che CRD exists", func() {
		apiextensions, err := clientset.NewForConfig(cfg)
		Expect(err).NotTo(HaveOccurred())
		// Make sure the CRD exists
		_, err = apiextensions.ApiextensionsV1beta1().CustomResourceDefinitions().Get(CrName, v1.GetOptions{})

		if err != nil {
			metadata.Instance.FoundCRD = false
		} else {
			metadata.Instance.FoundCRD = true
		}

		Expect(err).NotTo(HaveOccurred())
	})
})
