# Kubesecret

Kubesecret is a command-line tool that prints secrets and configmaps data of a kubernetes cluster.

`kubesecret -h` for help pages.

#### Install
`go get github.com/charmitro/kubesecret`

#### Flags
    -h, --help        help for [command]
        --kubeconfig  string    (optional) absolute path to the kubeconfig file (default "/Users/charmitro/.kube/config")
    -n, --namespace   string Namespace for kubesecret to look into. (default "default")

#### Secrets
    kubesecret get secret                             : Prints all the secrets that are in the 'default' namespace.
    kubesecret get secret -n <namespace>              : Prints all the secrets that are in the provided namespace.
    kubesecret get secret <secretname> -n <namespace> : Prints all the data of provided secret.

#### ConfigMaps
    kubesecret get configmap                             : Prints all the configmaps that are in the 'default' namespace.
    kubesecret get configmap -n <namespace>              : Prints all the configmaps that are in the provided namespace.
    kubesecret get configmap <secretname> -n <namespace> : Prints all the data of provided configmap.
    

