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

	archives, err := kubeberthClient.GetAllArchives(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	for num, archive := range archives {
		fmt.Printf("[#%d] ", num+1)
		b, err := json.Marshal(archive)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(b))
	}
}