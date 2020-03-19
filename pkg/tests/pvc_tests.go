package tests

import (
	"github.com/flacatus/operator-tests/pkg/client"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ =  ginkgo.Describe("Che-operator e2e tests dedicated", func() {
	client := client.New()

	ginkgo.It("Check if postgres PVC was created", func() {
		_, err := client.Kube().CoreV1().PersistentVolumeClaims(operatorNamespace).Get(PostgresPVCName, metav1.GetOptions{})

		if err != nil {
			panic(err)
		}
		Expect(err).NotTo(HaveOccurred())
	})
})

