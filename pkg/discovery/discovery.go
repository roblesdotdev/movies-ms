package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Registry defines a service registry.
type Registry interface {
	Register(ctx context.Context, intanceId string, serviceName string, hostPort string) error
	Deregister(ctx context.Context, intanceId string, serviceName string) error
	ServiceAddresses(ctx context.Context, serviceId string) ([]string, error)
	ReportHealthyState(instanceId string, serviceName string) error
}

var ErrNotFound = errors.New("no service addresses found")

// GenerateInstanceId generates a pseuto-random service instance identifier.
func GenerateInstanceId(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
