package main

import (
	"context"
	"strings"

	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/containers"
)

// I have tried to prepare a small version of what we will be doing in Mentorship
// Therefore I have written smaller versions of KubeArmor and Podman structs
// Real implementation would be similar but not same!

type PodmanHandler struct{
	PodmanClient context.Context

	// More properties can be added as per the need!
}

type KubeArmorContainer struct {
	ContainerId string
	ContainerName string
	NamespaceName string
	EndPointName string
	ContainerImage string
}


func NewPodmanHandler()(*PodmanHandler, error){
	podman := &PodmanHandler{}
	PodmanClient, err := bindings.NewConnection(context.Background(), "unix:///run/podman/podman.sock")
	if err!=nil{
		return nil, err
	}
	podman.PodmanClient = PodmanClient

	return podman, nil
}

func (pd *PodmanHandler) GetContainerInfo(containerId string) (KubeArmorContainer, error) {
	inspect, err := containers.Inspect(pd.PodmanClient, containerId, &containers.InspectOptions{})
	container := KubeArmorContainer{}
	if err!=nil{
		return KubeArmorContainer{}, err
	}
	container.ContainerId = inspect.ID
	container.ContainerName = strings.TrimLeft(inspect.Name, "/")
	container.NamespaceName = "Unknown"
	container.EndPointName = "Unknown"
	container.NamespaceName = "container_namespace"
	container.ContainerImage = inspect.Config.Image

	return container, nil
}