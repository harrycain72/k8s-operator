package visitorsapp

import (
	"context"

	examplev1 "visitors-operator/pkg/apis/example/v1"

	//appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func (r *ReconcileVisitorsApp) ensureService(request reconcile.Request,
	instance *examplev1.VisitorsApp,
	s *corev1.Service,
) (*reconcile.Result, error) {

	found := &corev1.Service{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      s.Name,
		Namespace: instance.Namespace,
	}, found)
	
	// check if we were not able to find the service 
	if err != nil && errors.IsNotFound(err) {

		// Create the service
		log.Info("Creating a new Service", "Service.Namespace", s.Namespace, "Service.Name", s.Name)
		err = r.client.Create(context.TODO(), s)

		if err != nil {
			// Creation failed
			log.Error(err, "Failed to create new Service", "Service.Namespace", s.Namespace, "Service.Name", s.Name)
			return &reconcile.Result{}, err
		} else {
			// Creation was successful
			return nil, nil
		}
	} else if err != nil {
		// Error that isn't due to the service not existing
		log.Error(err, "Failed to get Service")
		return &reconcile.Result{}, err
	}

	return nil, nil
}


func (r *ReconcileVisitorsApp) ensurePod(request reconcile.Request,
	instance *examplev1.VisitorsApp,
	p *corev1.Pod,
) (*reconcile.Result, error) {

	found := &corev1.Pod{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      p.Name,
		Namespace: instance.Namespace,
	}, found)
	
	// check if we were not able to find the pod 
	if err != nil && errors.IsNotFound(err) {

		// Create the pod
		log.Info("Creating a new pod", "Pod.Namespace", p.Namespace, "Pod.Name", p.Name)
		err = r.client.Create(context.TODO(), p)

		if err != nil {
			// Creation failed
			log.Error(err, "Failed to create new Pod", "Pod.Namespace", p.Namespace, "Pod.Name", p.Name)
			return &reconcile.Result{}, err
		} else {
			// Creation was successful
			return nil, nil
		}
	} else if err != nil {
		// Error that isn't due to the pod not existing
		log.Error(err, "Failed to get pod")
		return &reconcile.Result{}, err
	}

	return nil, nil
}