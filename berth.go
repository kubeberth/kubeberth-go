package kubeberth

import (
	"github.com/kubeberth/berth-apiserver/pkg/berth"
	"github.com/kubeberth/berth-apiserver/pkg/disks"
	"github.com/kubeberth/berth-apiserver/pkg/servers"
)

type AttachedArchive    = berth.AttachedArchive
type AttachedCloudInit  = berth.AttachedCloudInit
type AttachedDisk       = berth.AttachedDisk
type AttachedSource     = berth.AttachedSource
type Disk               = disks.Disk
type Server             = servers.Server
