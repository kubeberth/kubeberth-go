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

	server := &kubeberth.Server{
		Name: "test",
		Running: "false",
		CPU: "1",
		Memory: "1Gi",
		MACAddress: "52:42:00:00:00:00",
		HostName: "test",
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
