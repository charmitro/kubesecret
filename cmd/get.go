package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	clientset *kubernetes.Clientset

	// Flags
	kubeconfig string
	namespace  string

	// getCmd represents the get command
	getCmd = &cobra.Command{
		Use: "get",
		Example: `kubesecret get configmap [configmap] --namespace default
kubesecret get secret [secret] --namespace default`,
	}
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "Namespace for kubesecret to look into.")
	if home := homedir.HomeDir(); home != "" {
		getCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		cobra.CheckErr(err)
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		cobra.CheckErr(err)
	}
}
