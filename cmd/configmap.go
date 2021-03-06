package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	configMap string

	// configmapCmd represents the configmapCmd command
	configmapCmd = &cobra.Command{
		Use:     "configmap [configmap]",
		Aliases: []string{"cm"},
		Example: `kubesecret get configmap [configmaps] --namespace default.`,
		Run: func(cmd *cobra.Command, args []string) {
			getConfigMaps(args)
		},
	}
)

func init() {
	getCmd.AddCommand(configmapCmd)
}

func determineConfigMapDataType(cfg string) (string, error) {
	t := make(map[string]interface{})

	if err := yaml.Unmarshal([]byte(cfg), &t); err != nil &&
		strings.ContainsAny("cannot unmarshal !!", err.Error()) {
		return cfg, nil
	}

	m, err := yaml.Marshal(&t)
	if err != nil {
		return "", err
	}

	return "\n" + string(m), nil
}

func getConfigMaps(args []string) {
	if len(args) != 0 {
		if len(args) > 1 {
			for _, cm := range args {
				cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("ConfigMap: %s\n\n", cm))

				configMaps, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), cm, v1.GetOptions{})
				if err != nil {
					cobra.CheckErr(err)
				}

				for i, cfg := range configMaps.Data {
					cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v: %v\n", i, cfg))
				}
				cobra.WriteStringAndCheck(os.Stdout, "\n")
			}
		} else {
			configMaps, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), args[0], v1.GetOptions{})
			if err != nil {
				cobra.CheckErr(err)
			}

			for i, cfg := range configMaps.Data {
				data, err := determineConfigMapDataType(cfg)
				if err != nil {
					cobra.CheckErr(err)
				}
				cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v: %v\n", i, data))

			}
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
