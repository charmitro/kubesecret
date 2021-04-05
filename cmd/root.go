package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "kubesecret",
		Short:   "Prints the content of k8s secrets",
		Version: "v0.1.0",
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}
