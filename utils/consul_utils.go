package utils

import (
	"fmt"
	consulapi "github.com/armon/consul-api"
	"strings"
)

func GetConsul(url string) *consulapi.Client {
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = url
	consul, err := consulapi.NewClient(consulConfig)
	if err != nil {
		fmt.Println("Can`t get consul config:" + err.Error())
	}
	return consul
}

func LoadConfig(consul *consulapi.Client, cfgPrefix string) map[string]string {
	cfgPrefix = cfgPrefix + "/"
	prefixStripper := strings.NewReplacer(cfgPrefix, "")

	qo := &consulapi.QueryOptions{
		WaitIndex: 0,
	}
	kvPairs, _, err := consul.KV().List(cfgPrefix, qo)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	configs := make(map[string]string)

	for _, item := range kvPairs {
		if item.Key == cfgPrefix {
			continue
		}
		key := prefixStripper.Replace(item.Key)
		configs[key] = string(item.Value)
	}

	return configs
}
