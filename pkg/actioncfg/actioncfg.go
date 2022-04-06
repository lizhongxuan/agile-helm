package actioncfg

import (
	"agile-helm/pkg/action"
	"encoding/base64"
	"errors"
	"fmt"
	log "github.com/golang/glog"
	"k8s.io/client-go/rest"
	"strings"
)

func InClusterActionCfg(ns string) (*action.Configuration, error) {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	getter,err := NewGetter(cfg)
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

func BasicActionCfg(ns,host,apiAuthBasic string)(*action.Configuration, error)   {
	cfg := &rest.Config{}
	cfg.Insecure = true
	cfg.Host = host
	usernameColonPassword, err := base64.StdEncoding.DecodeString(apiAuthBasic)
	if err != nil {
		return nil, err
	}
	usernamePassword := strings.SplitN(string(usernameColonPassword), ":", 2)
	if len(usernamePassword) >= 2 {
		cfg.Username = usernamePassword[0]
		cfg.Password = usernamePassword[1]
	} else {
		return nil, errors.New("basic auth data incorrect, decode username and password error")
	}
	getter,err := NewGetter(cfg)
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

func TokenActionCfg(ns,host,apiAuthBasic string)(*action.Configuration, error)    {
	cfg := &rest.Config{}
	cfg.BearerToken = apiAuthBasic
	cfg.Insecure = true
	cfg.Host = host
	getter,err := NewGetter(cfg)
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

func TLSActionCfg(ns,host,apiClientCertificateData,apiClientKeyData string,apiClusterAuthData ...string)(*action.Configuration, error)   {
	certData, err := base64.StdEncoding.DecodeString(apiClientCertificateData)
	if err != nil {
		return nil, err
	}
	keyData, err := base64.StdEncoding.DecodeString(apiClientKeyData)
	if err != nil {
		return nil, err
	}
	cfg := &rest.Config{}
	cfg.CertData = certData
	cfg.KeyData = keyData
	cfg.Host = host
	if  len(apiClusterAuthData) != 0 {
		cadata, err := base64.StdEncoding.DecodeString(apiClusterAuthData[0])
		if err != nil {
			return nil, err
		}
		cfg.CAData = cadata
	} else {
		//没有证书则不需要，跳过证书校验
		cfg.Insecure = true
	}
	getter,err := NewGetter(cfg)
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





