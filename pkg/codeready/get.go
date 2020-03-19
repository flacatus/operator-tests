package codeready

import (
	"errors"
	orgv1 "github.com/eclipse/che-operator/pkg/apis/org/v1"
	"k8s.io/client-go/rest"
	"time"
)

func GetCodeReadyCustomResource(clientset rest.Interface) (*orgv1.CheCluster, error) {
	result := orgv1.CheCluster{}
	err := clientset.
		Get().
		Namespace(operatorNamespace).
		Resource(CheKind).
		Name(crName).
		Do().
		Into(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func VerifyCodeReadyCustomResource(status string, client rest.Interface) (deployed bool, err error) {
	//Sleep to wait to start the pods to deploy
	time.Sleep(1 * time.Minute)
	timeout := time.After(15 * time.Minute)
	tick := time.Tick(10 * time.Second)
	for {
		select {
		case <-timeout:
			return false, errors.New("timed out")
		case <-tick:
			customResource, _ := GetCodeReadyCustomResource(client)
			if customResource.Status.CheClusterRunning != status {
			} else {
				return true, nil
			}
		}
	}
}