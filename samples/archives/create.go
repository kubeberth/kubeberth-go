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

	requestArchive := &kubeberth.RequestArchive{
		Name:       "test",
		Repository: "https://cloud-images.ubuntu.com/releases/focal/release/ubuntu-20.04-server-cloudimg-arm64.img",
	}

	responseArchive, err := kubeberthClient.CreateArchive(ctx, requestArchive)

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.Marshal(responseArchive)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
