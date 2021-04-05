package cmd

import "github.com/spf13/cobra"


var (
	// configmapCmd represents the get command
	configmapCmd = &cobra.Command{
		Use: "configmap",
		Short: `'get' returns secret filenames from provided namespace
or the secret values from provided secret.`,
		Example: "kubesecret get --namespace default --secret=app-secrets",
		Run: func(cmd *cobra.Command, args []string) {
			getSecrets()
		},
	}
)

func init() {
	getCmd.AddCommand(configmapCmd)
}


func getConfigMaps() {

}
