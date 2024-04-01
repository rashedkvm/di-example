package di

import corev1 "k8s.io/api/core/v1"

type AppConventions struct {
	Env   []string
	Image string
	Probe corev1.Probe
}

type Resource interface {
	SetAppConventions() *AppConventions
}
