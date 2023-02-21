package manager

import (
	"context"
	"reflect"
	"testing"

	core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func TestNewProvider(t *testing.T) {
	type args struct {
		kubeConfigPath string
		contextName    string
	}
	tests := []struct {
		name    string
		args    args
		want    *Provider
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProvider(tt.args.kubeConfigPath, tt.args.contextName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_CurrentNamespace(t *testing.T) {
	type fields struct {
		clientSet        *kubernetes.Clientset
		kubernetesConfig clientcmd.ClientConfig
		clientConfig     rest.Config
		managedBy        string
		createdBy        string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := &Provider{
				clientSet:        tt.fields.clientSet,
				kubernetesConfig: tt.fields.kubernetesConfig,
				clientConfig:     tt.fields.clientConfig,
				managedBy:        tt.fields.managedBy,
				createdBy:        tt.fields.createdBy,
			}
			got, err := provider.CurrentNamespace()
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.CurrentNamespace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Provider.CurrentNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_ListAllNamespaces(t *testing.T) {
	type fields struct {
		clientSet        *kubernetes.Clientset
		kubernetesConfig clientcmd.ClientConfig
		clientConfig     rest.Config
		managedBy        string
		createdBy        string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []core.Namespace
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := &Provider{
				clientSet:        tt.fields.clientSet,
				kubernetesConfig: tt.fields.kubernetesConfig,
				clientConfig:     tt.fields.clientConfig,
				managedBy:        tt.fields.managedBy,
				createdBy:        tt.fields.createdBy,
			}
			got, err := provider.ListAllNamespaces(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.ListAllNamespaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.ListAllNamespaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
