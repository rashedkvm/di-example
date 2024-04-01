package implementation

import (
	"github.com/rashedkvm/di-example/di"
	corev1 "k8s.io/api/core/v1"
)

type ContainerApp struct {
	Env   []string
	Image string
	Probe corev1.Probe
}

func (c *ContainerApp) SetAppConventions() *di.AppConventions {

	return &di.AppConventions{
		Env: c.Env,
	}
}
