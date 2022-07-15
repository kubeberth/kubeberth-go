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

	requestISOImage := &kubeberth.RequestISOImage{
		Name:       "test",
		Size:       "4Gi"
		Repository: "http://minio.home.arpa:9000/kubeberth/images/ubuntu-20.04-server-cloudimg-arm64.img",
	}

	responseISOImage, err := kubeberthClient.UpdateISOImage(ctx, "test", requestISOImage)

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.Marshal(responseISOImage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
