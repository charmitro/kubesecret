package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	secret string

	// secretCmd represents the secret command
	secretCmd = &cobra.Command{
		Use:     "secret [secret]",
		Example: "kubesecret get secret [secret] --namespace default.",
		Run: func(cmd *cobra.Command, args []string) {
			getSecrets(args)
		},
	}
)

func init() {
	getCmd.AddCommand(secretCmd)
	secretCmd.Flags().StringVarP(&secret, "secret", "s", "", "secret to search for.")
}

func getSecrets(args []string) {
	if len(args) != 0 {
		if len(args) > 1 {
			for _, scrt := range args {
				cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("Secret: %s\n\n", scrt))

				secrets, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), scrt, v1.GetOptions{})
				if err != nil {
					cobra.CheckErr(err)
				}

				for i, s := range secrets.Data {
					cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v: %v\n", i, string(s)))
				}

				cobra.WriteStringAndCheck(os.Stdout, "\n")
			}
		} else {
			secrets, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), args[0], v1.GetOptions{})
			if err != nil {
				cobra.CheckErr(err)
			}

			for i, s := range secrets.Data {
				cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v: %v\n", i, string(s)))
			}
		}
	} else {
		secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			cobra.CheckErr(err)
		}

		if len(secrets.Items) != 0 {
			fmt.Printf("Printing all available secrets in namespace '%s'.\n\n", namespace)
			for _, s := range secrets.Items {
				cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v\n", s.Name))
			}
		} else {
			fmt.Printf("No available secrets in namespace '%s'.", namespace)
		}

		fmt.Println()
	}

}
