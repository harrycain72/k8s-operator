package visitorsapp


import (
	examplev1 "visitors-operator/pkg/apis/example/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *ReconcileVisitorsApp) newPodForCR(v *examplev1.VisitorsApp) *corev1.Pod {
	

	labels := labels(v, "backend")

	p := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      v.Name + "-pod",
			Namespace: v.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}

	log.Info("Pod Spec", "Pod.Name", p.ObjectMeta.Name)

	controllerutil.SetControllerReference(v, p, r.scheme)

	return p
}