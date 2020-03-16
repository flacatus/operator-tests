package tests

import (
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ =  ginkgo.Describe("", func() {
	ginkgo.It("Check if config maps are created in cluster", func() {
		_, err := clienset.CoreV1().ConfigMaps(operatorNamespace).Get(CodeReadyConfigMap, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
	})
})
