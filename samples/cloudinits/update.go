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

	cloudinit := &kubeberth.CloudInit{
		Name: "test",
		UserData: `#cloud-config
timezone: Asia/Tokyo
ssh_pwauth: True
password: ubuntu
chpasswd: { expire: False }
disable_root: false`,
	}

	cloudinit, err := kubeberthClient.UpdateCloudInit(ctx, "test", cloudinit)

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.Marshal(cloudinit)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
