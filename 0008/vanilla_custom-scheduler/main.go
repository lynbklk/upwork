package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Scheduler struct {
	clientset     *kubernetes.Clientset
	schedulerName string
}

func main() {
	fmt.Println("Welcome to the scheduler")
	scheduler := NewScheduler("custom-scheduler")
	scheduler.run()

}

/**
* Create Scheduler structure with k8s client config and scheduler name
* @params schedulerName
* @return Scheduler
**/
func NewScheduler(schedulerName string) Scheduler {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal("error Fatal: ", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return Scheduler{
		clientset:     clientset,
		schedulerName: schedulerName,
	}
}

/**
* Run Scheduler by watching unbound pods and bind them into random node
* @params none
* @return none
**/
func (s *Scheduler) run() {
	watch, err := s.clientset.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{
		FieldSelector: fmt.Sprintf("spec.schedulerName=%s", s.schedulerName),
	})

	if err != nil {
		fmt.Println("lynbklk, watch failed")
		log.Fatal(err)
	}

	for event := range watch.ResultChan() {
		if event.Type != "ADDED" {
			continue
		}
		p := event.Object.(*v1.Pod)
		fmt.Println("found a pod to schedule:", p.Namespace, "/", p.Name)
		n := getRandomNode(s)
		bindPod(s, p, n)
	}

}

/**
* Get random node in the k8s cluster
* @params scheduler
* @return node
**/
func getRandomNode(s *Scheduler) *v1.Node {
	nodes, err := s.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	return &nodes.Items[rand.Intn(len(nodes.Items))]
}

/**
* Bind pod to a given node
* @params Scheduler, a given pod and node
* @return none
**/
func bindPod(s *Scheduler, p *v1.Pod, n *v1.Node) {
	err := s.clientset.CoreV1().Pods(p.Namespace).Bind(context.TODO(), &v1.Binding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Name,
			Namespace: p.Namespace,
		},
		Target: v1.ObjectReference{
			APIVersion: "v1",
			Kind:       "Node",
			Name:       n.Name,
		},
	}, metav1.CreateOptions{})

	if err != nil {
		log.Fatal(err)
	}
}
