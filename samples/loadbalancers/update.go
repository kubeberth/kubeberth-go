package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/kubeberth/kubeberth-go"
)

func main() {
	url := "http://api.kubeberth.k8s.arpa/api/v1alpha1/"
	config := kubeberth.NewConfig(url)
	kubeberthClient := kubeberth.NewClient(config)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	requestLoadBalancer := &kubeberth.RequestLoadBalancer{
		Name: "test",
		Backends: []kubeberth.Destination{
			{ Server: "web-a" },
			{ Server: "web-b" },
			{ Server: "web-c" },
		},
		Ports: []kubeberth.Port{
			{
				Name: "http",
				Protocol: "TCP",
				Port: 80,
				TargetPort: intstr.FromInt(80),
			},
		},
	}

	responseLoadBalancer, err := kubeberthClient.UpdateLoadBalancer(ctx, "test", requestLoadBalancer)

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.Marshal(responseLoadBalancer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
