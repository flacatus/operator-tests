package codeready

import (
	orgv1 "github.com/eclipse/che-operator/pkg/apis/org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func createCodeReadyCluster() *orgv1.CheCluster {
	return &orgv1.CheCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crName,
			Namespace: operatorNamespace,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       CodeReadyKind,
			APIVersion: CodeReadyAPIVersion,
		},
		Spec: orgv1.CheClusterSpec{
			Server: orgv1.CheClusterSpecServer{
				SelfSignedCert: true,
				TlsSupport:true,
				CheFlavor:CodeReadyCFlavor,
			},
		},
	}
}

func CreateCodeReadyCustomResource(clientset rest.Interface) (err error) {
	result := orgv1.CheCluster{}
	cheCluster := createCodeReadyCluster()
	err = clientset.
		Post().
		Namespace(operatorNamespace).
		Resource(CheKind).
		Name(crName).
		Body(cheCluster).
		Do().
		Into(&result)

	if err != nil {
		return err
	}

	return nil
}