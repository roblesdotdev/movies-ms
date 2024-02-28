package consul

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
	"github.com/roblesdotdev/movies-ms/pkg/discovery"
)

type Registry struct {
	client *consul.Client
}

func NewRegistry(addr string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Registry{client: client}, nil
}

func (r *Registry) Register(ctx context.Context, instanceId string, serviceName string, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts) != 2 {
		return errors.New("hostPort must b in a form of <host>:<port>, example: localhost:8080")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: parts[0],
		ID:      instanceId,
		Name:    serviceName,
		Port:    port,
		Check:   &consul.AgentServiceCheck{CheckID: instanceId, TTL: "5s"},
	})
}

func (r *Registry) Deregister(ctx context.Context, instanceId string, _ string) error {
	return r.client.Agent().ServiceDeregister(instanceId)
}

func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	} else if len(entries) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res []string
	for _, e := range entries {
		res = append(res, fmt.Sprintf("%s:%d", e.Service.Address, e.Service.Port))
	}
	return res, nil
}

func (r *Registry) ReportHealthyState(instanceId string, _ string) error {
	return r.client.Agent().PassTTL(instanceId, "")
}
