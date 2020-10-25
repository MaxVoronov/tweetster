package consul

import (
	"net"
	"strconv"

	"github.com/hashicorp/consul/api"

	"github.com/maxvoronov/tweetster/pkg/utils"
)

type Registrar struct {
	apiClient *api.Client
}

func NewServiceRegistrar(consulHost string, consulPort int) (*Registrar, error) {
	config := api.DefaultConfig()
	config.Address = net.JoinHostPort(consulHost, strconv.Itoa(consulPort))
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Registrar{apiClient: client}, nil
}

func (r *Registrar) Register(name string, host string, port int) (string, error) {
	serviceID := name + "-" + utils.RandomString(5)
	asr := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    name,
		Tags:    nil,
		Address: host,
		Port:    port,
	}
	if err := r.apiClient.Agent().ServiceRegister(asr); err != nil {
		return "", err
	}

	return serviceID, nil
}

func (r *Registrar) Deregister(id string) error {
	return r.apiClient.Agent().ServiceDeregister(id)
}
