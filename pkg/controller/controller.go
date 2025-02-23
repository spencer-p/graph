package controller

import (
	cloudevents "github.com/cloudevents/sdk-go"
	"k8s.io/client-go/dynamic"
	"net/http"
	"sync"
)

type Controller struct {
	root string
	mux  *http.ServeMux
	once sync.Once

	namespace string
	client    dynamic.Interface
	CE        cloudevents.Client
}

func New(root, namespace string, client dynamic.Interface) *Controller {
	return &Controller{root: root, namespace: namespace, client: client}
}

func (c *Controller) Mux() *http.ServeMux {
	c.once.Do(func() {
		m := http.NewServeMux()
		m.HandleFunc("/ui", c.RootHandler)
		m.HandleFunc("/list", c.ListHandler)
		m.HandleFunc("/list/delete", c.ListDeleteHandler)
		m.HandleFunc("/tasks", c.TasksHandler)
		m.HandleFunc("/tasks/", c.TasksHandler)
		c.mux = m
	})

	return c.mux
}
