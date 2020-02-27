package visitorsapp


import (

	examplev1 "visitors-operator/pkg/apis/example/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

)

const frontendPort = 3000
const frontendServicePort = 30686


//frontendServiceName creates the name for the service
func frontendServiceName(v *examplev1.VisitorsApp) string {
	return v.Name + "-frontend-service"
}



//frontendService creates the service structure
func (r *ReconcileVisitorsApp) frontendService(v *examplev1.VisitorsApp) *corev1.Service {
	labels := labels(v, "frontend")

	s := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:		frontendServiceName(v),
			Namespace: 	v.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{{
				Protocol: corev1.ProtocolTCP,
				Port: frontendPort,
				TargetPort: intstr.FromInt(frontendPort),
				NodePort: frontendServicePort,
			}},
			Type: corev1.ServiceTypeNodePort,
		},
	}

	log.Info("Service Spec", "Service.Name", s.ObjectMeta.Name)

	controllerutil.SetControllerReference(v, s, r.scheme)
	return s
}