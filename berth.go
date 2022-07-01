package kubeberth

import (
	"github.com/kubeberth/kubeberth-apiserver/pkg/berth"
	"github.com/kubeberth/kubeberth-apiserver/pkg/archives"
	"github.com/kubeberth/kubeberth-apiserver/pkg/cloudinits"
	"github.com/kubeberth/kubeberth-apiserver/pkg/disks"
	"github.com/kubeberth/kubeberth-apiserver/pkg/servers"
)

type AttachedArchive   = berth.AttachedArchive
type AttachedCloudInit = berth.AttachedCloudInit
type AttachedDisk      = berth.AttachedDisk
type AttachedSource    = berth.AttachedSource

type Archive           = archives.Archive
type CloudInit         = cloudinits.CloudInit
type ResponseDisk      = disks.ResponseDisk
type RequestDisk       = disks.RequestDisk
type ResponseServer    = servers.ResponseServer
type RequestServer     = servers.RequestServer
