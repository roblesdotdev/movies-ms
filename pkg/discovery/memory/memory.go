package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/roblesdotdev/movies-ms/pkg/discovery"
)

type Registry struct {
	serviceAddrs map[string]map[string]*serviceInstance
	sync.RWMutex
}

type serviceInstance struct {
	lastActive time.Time
	hostPort   string
}

func NewRegistry() *Registry {
	return &Registry{serviceAddrs: map[string]map[string]*serviceInstance{}}
}

// Register creates a service record in the registry.
func (r *Registry) Register(ctx context.Context, instanceId string, sname string, hostPort string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[sname]; !ok {
		r.serviceAddrs[sname] = map[string]*serviceInstance{}
	}
	r.serviceAddrs[sname][instanceId] = &serviceInstance{hostPort: hostPort, lastActive: time.Now()}
	return nil
}

// Deregister removes a service record from the registry.
func (r *Registry) Deregister(ctx context.Context, instanceId string, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return nil
	}
	delete(r.serviceAddrs[serviceName], instanceId)
	return nil
}

// ReportHealthyState is a puch mechanism for reporting healthy state to th registry.
func (r *Registry) ReportHealthyState(instanceId string, serviceName string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return errors.New("service is not registered yet")
	}
	if _, ok := r.serviceAddrs[serviceName][instanceId]; !ok {
		return errors.New("service instance is not registered yet")
	}
	r.serviceAddrs[serviceName][instanceId].lastActive = time.Now()
	return nil
}

// ServiceAddresses returns the list of active instance of the given service.
func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.serviceAddrs[serviceName]) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res []string
	for _, i := range r.serviceAddrs[serviceName] {
		if i.lastActive.Before(time.Now().Add(-5 * time.Second)) {
			continue
		}
		res = append(res, i.hostPort)
	}
	return res, nil
}
