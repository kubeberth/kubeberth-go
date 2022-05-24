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

	disk := &kubeberth.Disk{
		Name: "test",
		Size: "32Gi",
		Source: &kubeberth.AttachedSource{
			Archive: &kubeberth.AttachedArchive{
				Name: "test",
			},
		},
	}

	disk, err := kubeberthClient.UpdateDisk(ctx, "test", disk)

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.Marshal(disk)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
