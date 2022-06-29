package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/kubeberth/kubeberth-go"
)

func main() {
	url := "http://api.kubeberth.k8s.arpa/api/v1alpha1/"
	config := kubeberth.NewConfig(url)
	kubeberthClient := kubeberth.NewClient(config)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cpu    := resource.MustParse("2")
	memory := resource.MustParse("2Gi")
	server := &kubeberth.Server{
		Name: "test",
		Running: false,
		CPU: &cpu,
		Memory: &memory,
		MACAddress: "52:42:00:4f:8a:2b",
		Hostname: "test",
		Hosting: "node-1.k8s.home.arpa",
		Disk: &kubeberth.AttachedDisk{
			Name: "test",
		},
		CloudInit: &kubeberth.AttachedCloudInit{
			Name: "test",
		},
	}

	server, err := kubeberthClient.UpdateServer(ctx, "test", server)

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
