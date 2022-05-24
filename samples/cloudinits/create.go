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
disable_root: false
ssh_authorized_keys:
- ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCzlOwyoT8qOMpkb9TafGPSM8lXxjZgvIAwHyhNLF1OUOBe8w55KMQ0IR6Q5w1lkKTmMsx7294Fd+xe5ak1BfuwwtF8eOcvWWibDyOr/aPmCFT/N6sZVe2BmN756U1PNDzhufNBH0Yq/AWpZsYn4EQL68hKZuUlA8awOBZS/EfZyPLLCNN5sGSo9nGTBT8DWnC6cEzWJ7ZrBuC69sInYF3haItnYVlafbus07H7waca6WXqZJUpeW0A8Acvsp2EUhNl8Kng/nlnnW4TuuccIGgTNn0Hx1QF6dnLMibD3uqkfAz2QBkJES4K3WWKApGJQxP6h4tw6llDrX7l6m7vHZpn`,
	}

	cloudinit, err := kubeberthClient.CreateCloudInit(ctx, cloudinit)

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
