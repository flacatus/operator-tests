package tests

import (
	"fmt"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"log"
	"time"
)

var _ =  ginkgo.Describe("", func() {
	ginkgo.It("Check CodeReady Pods status", func() {
		var (
			interval = 15 * time.Second
			timeout  = 1 * time.Minute

			requiredRatio float64 = 100
			curRatio      float64
			notReady      []v1.Pod
		)

		err := wait.Poll(interval, timeout, func() (done bool, err error) {
			if curRatio != 0 {
				log.Printf("Checking that all Pods are running or completed (currently %f%%)...", curRatio)
			}
			list, err := clienset.CoreV1().Pods(operatorNamespace).List(metav1.ListOptions{})

			if err != nil {
				return false, err
			}
			Expect(list).NotTo(BeNil())

			notReady = nil
			for _, pod := range list.Items {
				phase := pod.Status.Phase
				if phase != v1.PodRunning && phase != v1.PodSucceeded {
					notReady = append(notReady, pod)
				}
			}

			total := len(list.Items)
			ready := float64(total - len(notReady))
			curRatio = (ready / float64(total)) * 100

			return len(notReady) == 0, nil
		})
		msg := "only %f%% of Pods ready, need %f%%. Not ready: %s"
		Expect(err).NotTo(HaveOccurred(), msg, curRatio, requiredRatio, listPodPhases(notReady))
		Expect(curRatio).Should(Equal(requiredRatio), msg, curRatio, requiredRatio, listPodPhases(notReady))

	})
	ginkgo.It("should not be Failed", func() {
		list, err := clienset.CoreV1().Pods(operatorNamespace).List(metav1.ListOptions{
			FieldSelector: fmt.Sprintf("status.phase=%s", v1.PodFailed),
		})
		Expect(err).NotTo(HaveOccurred(), "couldn't list Pods")
		Expect(list).NotTo(BeNil())
		Expect(list.Items).Should(HaveLen(0), "'%d' Pods are 'Failed'", len(list.Items))
	})
})

func listPodPhases(pods []v1.Pod) (out string) {
	for i, pod := range pods {
		if i != 0 {
			out += ", "
		}
		out += fmt.Sprintf("%s/%s (Phase: %s)", pod.Namespace, pod.Name, pod.Status.Phase)
	}
	return
}
