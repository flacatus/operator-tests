package tests

import (
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/flacatus/operator-tests/pkg/client"
)

var _ =  ginkgo.Describe("", func() {
	ginkgo.It("Check if certs secret was created", func() {
		client := client.New()
		_, err := client.Kube().CoreV1().Secrets(operatorNamespace).Get(secretSelfSignedCrt, metav1.GetOptions{})

		Expect(err).NotTo(HaveOccurred(), "failed to get secretName %v\n", secretSelfSignedCrt)
	})
})