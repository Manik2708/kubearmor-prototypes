package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"os"
	// "path/filepath"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const (
	appArmorProfile = "kubearmor_kubearmor-prototype-test-2"
)

func main() {
	// Read data from stdin.
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading from stdin:", err)
		os.Exit(1)
	}
    var spec specs.Spec


	err = json.Unmarshal(data, &spec)
	spec.Process.ApparmorProfile = appArmorProfile
	if err != nil {
		fmt.Println("Error unmarshaling data to JSON:", err)
		os.Exit(1)
	}

	if err := json.NewEncoder(os.Stdout).Encode(spec); err != nil {
		fmt.Println("Error encoding data to JSON:", err)
		os.Exit(1)
	}
}
