package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	vmName     string
	cpuCount   int
	memorySize string
	enableSSH  bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "firecracker-vm1",
		Short: "CLI tool for launching Firecracker VM using Ignite",
		Run:   runCommand,
	}

	// Define flags
	rootCmd.Flags().StringVarP(&vmName, "name", "n", "my-vm", "Name for the Firecracker VM")
	rootCmd.Flags().IntVarP(&cpuCount, "cpus", "c", 2, "Number of CPUs for the VM")
	rootCmd.Flags().StringVarP(&memorySize, "memory", "m", "1GB", "Amount of RAM for the VM")
	rootCmd.Flags().BoolVarP(&enableSSH, "ssh", "s", false, "Enable SSH for the VM")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runCommand(cmd *cobra.Command, args []string) {
	// Build the Ignite command to launch the Firecracker VM
	igniteCmd := exec.Command("ignite", "run", "weaveworks/ignite-ubuntu",
		"--cpus", fmt.Sprintf("%d", cpuCount),
		"--memory", memorySize,
	)

	if enableSSH {
		igniteCmd.Args = append(igniteCmd.Args, "--ssh")
	}

	igniteCmd.Args = append(igniteCmd.Args, "--name", vmName)

	// Redirect standard input, output, and error streams
	igniteCmd.Stdin = os.Stdin
	igniteCmd.Stdout = os.Stdout
	igniteCmd.Stderr = os.Stderr

	// Run the Ignite command
	err := igniteCmd.Run()
	if err != nil {
		fmt.Printf("Error launching Firecracker VM: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Firecracker VM %s launched successfully\n", vmName)
}
