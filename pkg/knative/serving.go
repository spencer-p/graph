package knative

import (
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	servingv1alpha1 "knative.dev/serving/pkg/apis/serving/v1alpha1"
)

func (c *Client) KnServices(namespace string, yv *[]YamlView) []servingv1alpha1.Service {
	gvr := schema.GroupVersionResource{
		Group:    "serving.knative.dev",
		Version:  "v1alpha1",
		Resource: "services",
	}
	like := servingv1alpha1.Service{}

	list, err := c.dc.Resource(gvr).Namespace(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Printf("Failed to List Services, %v", err)
		return nil
	}

	all := make([]servingv1alpha1.Service, len(list.Items))

	for i, item := range list.Items {
		obj := like.DeepCopy()
		if err = runtime.DefaultUnstructuredConverter.FromUnstructured(item.Object, obj); err != nil {
			log.Fatalf("Error DefaultUnstructuree.Dynamiconverter. %v", err)
		}
		obj.ResourceVersion = gvr.Version
		obj.APIVersion = gvr.GroupVersion().String()
		all[i] = *obj

		// Yaml View
		AddToYamlView(item, yv)
	}
	return all
}
