package tests

import (
	"github.com/flacatus/operator-tests/pkg/client"
	"github.com/flacatus/operator-tests/pkg/metadata"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

var _ =  ginkgo.Describe("Che-operator e2e tests dedicated", func() {
	client := client.New()

	ginkgo.It("Check CodeReady Pods status", func() {
		pods, err := client.Kube().CoreV1().Pods(operatorNamespace).List(metav1.ListOptions{})

		for _, pod := range pods.Items {
			if pod.Status.Phase == kubev1.PodPending {
				panic("Pod is not running" +pod.Name)
			}

			if strings.Contains(pod.Name, "codeready-operator") {
				metadata.Instance.CodereadyOperatorStatus = string(pod.Status.Phase)
			}

			if strings.Contains(pod.Name, "postgres") {
				metadata.Instance.PostgreSQLStatus = string(pod.Status.Phase)
			}
			if strings.Contains(pod.Name, "keycloak") {
				metadata.Instance.KeycloackStatus = string(pod.Status.Phase)
			}
			if strings.Contains(pod.Name, "plugin-registry") {
				metadata.Instance.PluginRegistryStatus = string(pod.Status.Phase)
			}
			if strings.Contains(pod.Name, "devfile-registry") {
				metadata.Instance.DevFileStatus = string(pod.Status.Phase)
			}
			if strings.Contains(pod.Name, "codeready") && ! strings.Contains(pod.Name, "codeready-operator") {
				metadata.Instance.CodereadyStatus = string(pod.Status.Phase)
			}
		}

		if err != nil {
			panic(err)
		}

		Expect(err).NotTo(HaveOccurred())
	})

})

