package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kubeberth/kubeberth-go"
)

func main() {
	url := "http://api.kubeberth.k8s.arpa/api/v1alpha1/"
	config := kubeberth.NewConfig(url)
	kubeberthClient := kubeberth.NewClient(config)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	requestDisk := &kubeberth.RequestDisk{
		Name: "test",
		Size: "16Gi",
		Source: &kubeberth.AttachedSource{
			Archive: &kubeberth.AttachedArchive{
				Name: "test",
			},
		},
	}

	responseDisk, err := kubeberthClient.CreateDisk(ctx, requestDisk)

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.Marshal(responseDisk)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
