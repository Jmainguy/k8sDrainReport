package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	//"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	i := 0
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, v := range pods.Items {
		if v.Status.Phase == "Running" {
			i++
		}
	}
	fmt.Printf("Cluster: %s\n", config.Host)
	fmt.Printf("There are %d running pods in the cluster\n", i)
	i = 0
	fmt.Println("========================================")
	fmt.Println("Pods without owners")
	fmt.Println("========================================")
	// Loop through all pods
	for _, v := range pods.Items {
		// List pods without ownership
		if len(v.ObjectMeta.OwnerReferences) == 0 {
			if v.Status.Phase == "Running" {
				// Ignore kube-system, it has non owned pods, this is normal.
				if v.ObjectMeta.Namespace != "kube-system" {
					fmt.Printf("Namespace: %s, Name: %s, NodeName: %s\n", v.ObjectMeta.Namespace, v.ObjectMeta.Name, v.Spec.NodeName)
					i++
				}
			}
		}
	}

	fmt.Printf("There are %d pods without ownership\n", i)

	fmt.Println("")
	fmt.Println("========================================")
	fmt.Println("Pod Distruption Budget, Potential Issues")
	fmt.Println("========================================")
	// Get PodDisruptionBudgets
	pdbs, err := clientset.PolicyV1beta1().PodDisruptionBudgets("").List(metav1.ListOptions{})

	for _, v := range pdbs.Items {
		if v.Status.PodDisruptionsAllowed == 0 {
			fmt.Printf("Namespace: %s, Name: %s, DisruptionsAllowed: %d\n", v.ObjectMeta.Namespace, v.ObjectMeta.Name, v.Status.PodDisruptionsAllowed)
		}
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
