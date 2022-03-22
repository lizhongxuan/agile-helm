package pkg

import (
	"agile-helm/helm/v3/pkg/action"
	"agile-helm/helm/v3/pkg/kube"
	"fmt"
	log "github.com/golang/glog"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func GetActionConfig(ns string) (*action.Configuration, error) {
	getter,err := BuildGetter()
	if err !=nil {
		log.Error(err, ns)
		return nil, err
	}
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(getter, ns, "", logf); err != nil {
		log.Error(err, ns)
		return nil, err
	}
	return actionConfig, nil
}

func logf(format string, v ...interface{}) {
	if len(v) > 0 {
		log.Info(fmt.Sprintf(format, v...))
	} else {
		log.Info(format)
	}
}


func BuildGetter()(*Getter,error)   {
		cfg, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

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
		return NewGetter(clientcmd.NewDefaultClientConfig(*apicfg, nil)),nil
}

func BuildKubeClient()(*kube.Client,error)  {
	getter,err := BuildGetter()
	if err !=nil {
		return nil,err
	}
	return kube.New(getter),nil
}
