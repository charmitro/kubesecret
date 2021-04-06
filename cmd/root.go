package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "kubesecret",
		Short:   "Kubesecret.\nPrints the data of k8s secrets and configmaps.",
		Version: "v0.1.0",
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}
