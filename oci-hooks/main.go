package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/opencontainers/runtime-spec/specs-go"
)

type Container struct {
	ContainerId   string
	ContainerName string
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Read input from stdin
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading from stdin:", err)
		os.Exit(1)
	}
	fileN1, err := os.Create(dir+"/state.json")
	if err != nil {
		fmt.Println("Error reading from stdin:", err)
		os.Exit(1)
	}
	fileN1.Write(input)
	// // Unmarshal the JSON payload to get the state
	state := specs.State{}
	err = json.Unmarshal(input, &state)
	if err != nil {
		fmt.Println("Error in unmarshalling:", err)
		os.Exit(1)
	}

	// Read the config.json file from the bundle path
	specBytes, err := os.ReadFile(filepath.Join(state.Bundle, "config.json"))
	if err != nil {
		fmt.Println("Error reading config.json:", err)
		os.Exit(1)
	}
	var spec specs.Spec
	json.Unmarshal(specBytes, &spec)
	// Create the Container struct with ID and Name
	container := &Container{
		ContainerId:   state.ID,
	}
	file, err := os.Create(dir+"/container.txt")
	if err != nil {
		fmt.Println("Error reading from stdin:", err)
		os.Exit(1)
	}
	file.Write([]byte(container.ContainerId))
	fileN, err := os.Create(dir+"/spec.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileN.Write(specBytes)
}