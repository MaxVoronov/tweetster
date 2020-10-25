package consul

import (
	"net"
	"strconv"

	"github.com/go-logr/logr"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
)

const Scheme = "consul"

func RegisterDefaultResolver(logger logr.Logger) {
	resolver.Register(&consulBuilder{
		logger: logger,
	})
}

type consulBuilder struct {
	logger logr.Logger
}

func (cb *consulBuilder) Build(target resolver.Target, conn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	cb.logger.Info("building consul resolver", "params", target)

	config := api.DefaultConfig()
	config.Address = target.Authority
	apiClient, err := api.NewClient(config)
	if err != nil {
		cb.logger.Error(err, "error create consul client")
		return nil, err
	}

	cr := &consulResolver{
		apiClient:   apiClient,
		clientConn:  conn,
		serviceName: target.Endpoint,
		lastIndex:   0,
		logger:      cb.logger,
	}
	go cr.watcher()

	return cr, nil
}

func (cb *consulBuilder) Scheme() string {
	return Scheme
}

type consulResolver struct {
	apiClient   *api.Client
	clientConn  resolver.ClientConn
	serviceName string
	lastIndex   uint64
	logger      logr.Logger
}

func (cr *consulResolver) watcher() {
	for {
		services, meta, err := cr.apiClient.Health().Service(cr.serviceName, "", true, &api.QueryOptions{WaitIndex: cr.lastIndex})
		if err != nil {
			cr.logger.Error(err, "failed retrieving instances from Consul")
			continue
		}

		cr.lastIndex = meta.LastIndex
		addresses := make([]resolver.Address, 0, len(services))
		for _, svc := range services {
			addresses = append(addresses, resolver.Address{
				Addr:       net.JoinHostPort(svc.Service.Address, strconv.Itoa(svc.Service.Port)),
				ServerName: cr.serviceName,
			})
		}

		cr.logger.Info("adding service addresses", "addresses", addresses)
		cr.clientConn.UpdateState(resolver.State{Addresses: addresses})
	}
}

func (cr *consulResolver) ResolveNow(opt resolver.ResolveNowOptions) {
	// No need to refresh regularly like dns_resolver
}

func (cr *consulResolver) Close() {
	// No need do anything
}
