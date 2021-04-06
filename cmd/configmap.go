package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	configMap string

	// configmapCmd represents the configmapCmd command
	configmapCmd = &cobra.Command{
		Use:     "configmap [configmap]",
		Aliases: []string{"cm"},
		Example: `kubesecret get configmap [configmap] --namespace default.`,
		Run: func(cmd *cobra.Command, args []string) {
			getConfigMaps(args)
		},
	}
)

func init() {
	getCmd.AddCommand(configmapCmd)
	configmapCmd.Flags().StringVarP(&configMap, "configmap", "c", "", "configmap to search for.")
}

func getConfigMaps(args []string) {
	if len(args) != 0 {
		configMaps, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), args[0], v1.GetOptions{})
		if err != nil {
			cobra.CheckErr(err)
		}

		for i, cfg := range configMaps.Data {
			cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v: %v\n", i, string(cfg)))
		}

	} else {
		configMaps, err := clientset.CoreV1().ConfigMaps(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			cobra.CheckErr(err)
		}

		if len(configMaps.Items) != 0 {
			fmt.Printf("Printing all available configmaps in namespace '%s'.\n\n", namespace)
			for _, c := range configMaps.Items {
				cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v\n", c.Name))
			}
		} else {
			fmt.Printf("No available configmaps in namespace '%s'.", namespace)
		}

		fmt.Println()
	}
}
