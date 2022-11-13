package gomvvm

import (
	"fmt"
	"sync"
)

type Blocker interface {
	GetName() string
}

type Register interface {
	Add(...Blocker) error
	Remove(...Blocker) error
	Find(string) (Blocker, error)
}

type ViewModelRegistry struct {
	registry map[string]Blocker
}

func NewViewModelRegistry() *ViewModelRegistry {
	return &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}
}

func (r *ViewModelRegistry) Add(block ...Blocker) error {
	for _, b := range block {
		_, err := r.Find(b.GetName())
		if err == nil {
			return fmt.Errorf("layout: block with name %s already exists in registry", b.GetName())
		}

		r.registry[b.GetName()] = b
	}

	return nil
}

func (r *ViewModelRegistry) Remove(block ...Blocker) error {
	for _, b := range block {
		_, err := r.Find(b.GetName())
		if err != nil {
			return fmt.Errorf("layout: block with name %s not found in registry", b.GetName())
		}

		delete(r.registry, b.GetName())
	}

	return nil
}

func (r *ViewModelRegistry) Find(name string) (Blocker, error) {
	if r.registry[name] == nil {
		return nil, fmt.Errorf("layout: block with name %s not found in registry", name)
	}

	return r.registry[name], nil
}

type ViewModelRegistryProxy struct {
	registry *ViewModelRegistry
}

func NewViewModelRegistryProxy(registry *ViewModelRegistry) *ViewModelRegistryProxy {
	return &ViewModelRegistryProxy{
		registry: registry,
	}
}

func (r *ViewModelRegistryProxy) Add(block ...Blocker) error {
	return r.registry.Add(block...)
}

func (r *ViewModelRegistryProxy) Remove(block ...Blocker) error {
	return r.registry.Remove(block...)
}

func (r *ViewModelRegistryProxy) Find(name string) (Blocker, error) {
	return r.registry.Find(name)
}

var lock = &sync.Mutex{}
var instance *ViewModelRegistryProxy

func GetRegistry() *ViewModelRegistryProxy {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			registry := NewViewModelRegistry()
			instance = NewViewModelRegistryProxy(registry)
		}
	}

	return instance
}
