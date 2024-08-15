package main

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main(){
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	ctrList, err := cli.ContainerList(ctx, container.ListOptions{})
	if err!=nil{
		os.Exit(1)
	}
	for _, ctr := range ctrList {
		name, identify := identifyTestContainer(ctr.Names)
		if identify{
			if ctr.State == "running"{
				data, err := os.ReadFile("policy.yaml")
				if err !=nil {
					os.Exit(1)
				}
				str := string(data)
				str = strings.Replace(str, "kubearmor.io/container.name: lb", "kubearmor.io/container.name: "+name, 1)
				file, err := os.Create("new-policy.yaml")
				if err !=nil {
					os.Exit(1)
				}
				_, err = file.WriteString(str)
				if err !=nil {
					os.Exit(1)
				}
				applyPolicy("new-policy.yaml")
			}
		}
	}
}

func identifyTestContainer(names []string) (string, bool) {
	for _, name := range names {
		if strings.Contains(name, "kubearmor-prototype-test"){
			str, _ :=strings.CutPrefix(name, "/")
			return str, true
		}
	}
	return "", false
}

func applyPolicy(policyName string) error {
	cmd := exec.Command("karmor", "vm", "policy", "add", policyName)

	_, err := cmd.Output()
	return err
}