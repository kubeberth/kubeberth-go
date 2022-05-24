package main

import (
	"context"
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

	diskName := "test"
	ok, err := kubeberthClient.DeleteDisk(ctx, diskName)

	if !ok {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
