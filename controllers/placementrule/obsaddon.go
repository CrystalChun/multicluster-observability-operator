// Copyright (c) 2021 Red Hat, Inc.

package placementrule

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	obv1beta1 "github.com/open-cluster-management/multicluster-monitoring-operator/api/v1beta1"
)

const (
	obsAddonName = "observability-addon"
)

func deleteObsAddon(client client.Client, namespace string) error {
	found := &obv1beta1.ObservabilityAddon{}
	err := client.Get(context.TODO(), types.NamespacedName{Name: obsAddonName, Namespace: namespace}, found)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		log.Error(err, "Failed to check endpoint config cr", "namespace", namespace)
		return err
	}
	err = client.Delete(context.TODO(), found)
	if err != nil {
		log.Error(err, "Failed to delete observabilityaddon", "namespace", namespace)
	}

	err = removeObservabilityAddon(client, namespace)
	if err != nil {
		return err
	}

	log.Info("observabilityaddon is deleted", "namespace", namespace)
	return nil
}

func createObsAddon(client client.Client, namespace string) error {
	ec := &obv1beta1.ObservabilityAddon{
		ObjectMeta: metav1.ObjectMeta{
			Name:      obsAddonName,
			Namespace: namespace,
			Labels: map[string]string{
				ownerLabelKey: ownerLabelValue,
			},
		},
	}
	found := &obv1beta1.ObservabilityAddon{}
	err := client.Get(context.TODO(), types.NamespacedName{Name: obsAddonName, Namespace: namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		log.Info("Creating endpoint config cr", "namespace", namespace)
		err = client.Create(context.TODO(), ec)
		if err != nil {
			log.Error(err, "Failed to create endpoint config cr")
			return err
		}
		return nil
	} else if err != nil {
		log.Error(err, "Failed to check endpoint config cr")
		return err
	}

	log.Info("endponitmetrics already existed/unchanged", "namespace", namespace)
	return nil
}