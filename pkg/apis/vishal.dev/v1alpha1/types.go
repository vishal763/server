package v1alpha1
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type server struct{
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec serverSpec
}