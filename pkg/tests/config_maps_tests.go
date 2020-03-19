package tests

import (
	"github.com/flacatus/operator-tests/pkg/client"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ =  ginkgo.Describe(osdDescription, func() {
	client := client.New()

	ginkgo.It("Check if config maps are created in cluster", func() {
		confmap, err := client.Kube().CoreV1().ConfigMaps(operatorNamespace).Get(CodeReadyConfigMap, metav1.GetOptions{})

		if err != nil {
			panic(err)
		}

		Expect(confmap.Name).To(Equal(CodeReadyConfigMap))
		Expect(err).NotTo(HaveOccurred())
	})
})
