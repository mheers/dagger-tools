// A generated module for Ci functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/ci/internal/dagger"
	"fmt"
)

type Ci struct{}

const (
	alpineImage      = "alpine:3.20" // keep this until https://github.com/alpinelinux/docker-alpine/issues/93 is fixed
	daggerImage      = "registry.dagger.io/engine"
	daggerTag        = "v0.16.3"
	daggerToolsImage = "mheers/dagger-tools"
	username         = "mheers"
)

var tools = []string{
	"bash",
	"bind-tools",
	"busybox-extras",
	"curl",
	"ethtool",
	"file",
	"git",
	"iproute2",
	"iputils",
	"jq",
	"nano",
	"ncurses",
	"netcat-openbsd",
	"nmap",
	"numactl",
	"openldap-clients",
	"openssh-client",
	"openssl",
	"procps",
	"sudo",
	"sysstat",
	"tcpdump",
	"tree",
	"util-linux",
}

func (m *Ci) BuildAndPushImage(registryToken *dagger.Secret) (string, error) {
	daggerToolsImageName := fmt.Sprintf("%s:%s", daggerToolsImage, daggerTag)

	alpineContainer := dag.Container().From(alpineImage)

	installCmd := []string{"apk", "add", "--no-cache"}
	installCmd = append(installCmd, tools...)

	return dag.Container().From(fmt.Sprintf("%s:%s", daggerImage, daggerTag)).
		WithDirectory("/etc/apk", alpineContainer.Directory("/etc/apk")).
		WithExec(installCmd).
		WithRegistryAuth(daggerToolsImageName, username, registryToken).
		Publish(context.Background(), daggerToolsImageName)
}
