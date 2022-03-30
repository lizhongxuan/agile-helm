package actioncfg

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// Getter implements k8s.io/cli-runtime/ahelm/genericclioptions.RESTClientGetter interface.
type Getter struct {
	c clientcmd.ClientConfig
}

// ToRESTMapper is part of k8s.io/cli-runtime/ahelm/genericclioptions.RESTClientGetter interface.
func (c *Getter) ToRESTMapper() (meta.RESTMapper, error) {
	d, err := c.ToDiscoveryClient()
	if err != nil {
		return nil, err
	}

	mapper := restmapper.NewDeferredDiscoveryRESTMapper(d)
	expander := restmapper.NewShortcutExpander(mapper, d)

	return expander, nil
}

// ToDiscoveryClient is part of k8s.io/cli-runtime/ahelm/genericclioptions.RESTClientGetter interface.
func (c *Getter) ToDiscoveryClient() (discovery.CachedDiscoveryInterface, error) {
	cc, err := c.ToRESTConfig()
	if err != nil {
		return nil, fmt.Errorf("getting REST config: %w", err)
	}

	d, err := discovery.NewDiscoveryClientForConfig(cc)
	if err != nil {
		return nil, fmt.Errorf("creating discovery client: %w", err)
	}

	return memory.NewMemCacheClient(d), nil
}

// ToRawKubeConfigLoader is part of k8s.io/cli-runtime/ahelm/genericclioptions.RESTClientGetter interface.
func (c *Getter) ToRawKubeConfigLoader() clientcmd.ClientConfig {
	return c.c
}

// ToRESTConfig is part of k8s.io/cli-runtime/ahelm/genericclioptions.RESTClientGetter interface.
func (c *Getter) ToRESTConfig() (*rest.Config, error) {
	return c.c.ClientConfig()
}

func NewGetter(cfg *rest.Config)(*Getter,error)   {
	name:= "agilehelm"
	apicfg := &clientcmdapi.Config{
		Kind:           "Config",
		APIVersion:     "v1",
		CurrentContext: name,
		Preferences:    *clientcmdapi.NewPreferences(),
		Contexts: map[string]*clientcmdapi.Context{
			name: {
				Cluster:  name,
				AuthInfo: name,
			},
		},
		Clusters: map[string]*clientcmdapi.Cluster{
			name: {
				Server:                cfg.Host,
				InsecureSkipTLSVerify: true,
			},
		},
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			name:{
				Token:     cfg.BearerToken,
				TokenFile: cfg.BearerTokenFile,
				ClientCertificateData: cfg.CertData,
				ClientKeyData:         cfg.KeyData,
			},
		},
	}

	return &Getter{
		c: clientcmd.NewDefaultClientConfig(*apicfg, nil),
	},nil
}