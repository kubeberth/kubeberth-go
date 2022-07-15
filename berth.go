package kubeberth

import (
	"github.com/kubeberth/kubeberth-apiserver/pkg/berth"
	"github.com/kubeberth/kubeberth-apiserver/pkg/isoimages"
	"github.com/kubeberth/kubeberth-apiserver/pkg/archives"
	"github.com/kubeberth/kubeberth-apiserver/pkg/cloudinits"
	"github.com/kubeberth/kubeberth-apiserver/pkg/disks"
	"github.com/kubeberth/kubeberth-apiserver/pkg/servers"
	"github.com/kubeberth/kubeberth-apiserver/pkg/loadbalancers"
)

type AttachedISOImage     = berth.AttachedISOImage
type AttachedArchive      = berth.AttachedArchive
type AttachedCloudInit    = berth.AttachedCloudInit
type AttachedDisk         = berth.AttachedDisk
type AttachedSource       = berth.AttachedSource
type Destination          = berth.Destination
type Port                 = berth.Port

type ResponseISOImage     = isoimages.ResponseISOImage
type RequestISOImage      = isoimages.RequestISOImage
type Archive              = archives.Archive
type CloudInit            = cloudinits.CloudInit
type ResponseDisk         = disks.ResponseDisk
type RequestDisk          = disks.RequestDisk
type ResponseServer       = servers.ResponseServer
type RequestServer        = servers.RequestServer
type ResponseLoadBalancer = loadbalancers.ResponseLoadBalancer
type RequestLoadBalancer  = loadbalancers.RequestLoadBalancer
