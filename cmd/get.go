package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	clientset *kubernetes.Clientset

	// Flags
	kubeconfig string
	namespace  string
	secret     string

	// getCmd represents the get command
	getCmd = &cobra.Command{
		Use: "get",
		Short: `'get' returns secret filenames from provided namespace
or the secret values from provided secret.`,
		Example: "kubesecret get --namespace default --secret=app-secrets",
		Run: func(cmd *cobra.Command, args []string) {
			getSecrets()
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "")
	getCmd.Flags().StringVarP(&secret, "secret", "s", "", "")
	if home := homedir.HomeDir(); home != "" {
		getCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	}
}

func getSecrets() {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		cobra.CheckErr(err)
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		cobra.CheckErr(err)
	}

	if secret != "" {
		// Get a specific secret.
		secrets, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), secret, v1.GetOptions{})
		if err != nil {
			cobra.CheckErr(err)
		}

		for i, s := range secrets.Data {
			cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v: %v\n", i, string(s)))
		}
	} else {
		// Else list all available secrets ref names from the namespace
		secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			cobra.CheckErr(err)
		}

		for _, s := range secrets.Items {
			cobra.WriteStringAndCheck(os.Stdout, fmt.Sprintf("%v\n", s.Name))
		}
	}
}
