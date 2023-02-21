package manager

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Provider struct {
	clientSet        *kubernetes.Clientset
	kubernetesConfig clientcmd.ClientConfig
	clientConfig     rest.Config
	managedBy        string
	createdBy        string
}

func NewProvider(kubeConfigPath string, contextName string) (*Provider, error) {
	kubernetesConfig := loadKubernetesConfiguration(kubeConfigPath, contextName)
	restClientConfig, err := kubernetesConfig.ClientConfig()
	if err != nil {
		if clientcmd.IsEmptyConfig(err) {
			return nil, fmt.Errorf("couldn't find the kube config file, or file is empty (%s)\n"+
				"you can set alternative kube config file path by adding the kube-config-path field to the %s config file, err:  %w", kubeConfigPath, "todo", err)
		}
		if clientcmd.IsConfigurationInvalid(err) {
			return nil, fmt.Errorf("invalid kube config file (%s)\n"+
				"you can set alternative kube config file path by adding the kube-config-path field to the %s config file, err:  %w", kubeConfigPath, "todo", err)
		}

		return nil, fmt.Errorf("error while using kube config (%s)\n"+
			"you can set alternative kube config file path by adding the kube-config-path field to the %s config file, err:  %w", kubeConfigPath, "todo", err)
	}
	clientSet, err := getClientSet(restClientConfig)
	if err != nil {
		return nil, fmt.Errorf("error while using kube config (%s)\n"+
			"you can set alternative kube config file path by adding the kube-config-path field to the %s config file, err:  %w", kubeConfigPath, "todo", err)
	}
	return &Provider{
		clientSet:        clientSet,
		kubernetesConfig: kubernetesConfig,
		clientConfig:     *restClientConfig,
		managedBy:        "todo",
		createdBy:        "todo",
	}, nil
}

func loadKubernetesConfiguration(kubeConfigPath string, context string) clientcmd.ClientConfig {
	configPathList := filepath.SplitList(kubeConfigPath)
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{}
	if len(configPathList) <= 1 {
		configLoadingRules.ExplicitPath = kubeConfigPath
	} else {
		configLoadingRules.Precedence = configPathList
	}
	contextName := context
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		configLoadingRules,
		&clientcmd.ConfigOverrides{
			CurrentContext: contextName,
		},
	)
}

func getClientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

func (provider *Provider) CurrentNamespace() (string, error) {
	if provider.kubernetesConfig == nil {
		return "", errors.New("kubernetesConfig is nil")
	}
	ns, _, err := provider.kubernetesConfig.Namespace()
	return ns, err
}

func (provider *Provider) ListAllNamespaces(ctx context.Context) ([]core.Namespace, error) {
	namespaces, err := provider.clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces.Items, err
}

func (provider *Provider) GetKubernetesVersion() (string, error) {
	serverVersion, err := provider.clientSet.ServerVersion()
	if err != nil {
		return "", err
	}
	return serverVersion.String(), nil
}
