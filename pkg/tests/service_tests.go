package tests

import (
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ =  ginkgo.Describe("", func() {
	ginkgo.It("Check if operator create all services in cluster", func() {
		// TODO: Try to improve this checks with new err messages...
		services, err := clienset.CoreV1().Services(operatorNamespace).List(metav1.ListOptions{})
		Expect(services).NotTo(BeNil())

		confmap := map[string]string{}
		for _ ,v:= range services.Items {
			confmap[v.Name]= v.Name
		}

		Expect(confmap["che-host"]).NotTo(BeEmpty())
		Expect(confmap["plugin-registry"]).NotTo(BeEmpty())
		Expect(confmap["postgres"]).NotTo(BeEmpty())
		Expect(confmap["devfile-registry"]).NotTo(BeEmpty())

		Expect(err).NotTo(HaveOccurred())
	})
})
