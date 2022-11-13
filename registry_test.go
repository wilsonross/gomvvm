package gomvvm

import (
	"reflect"
	"testing"
)

func TestNewViewModelRegistry(t *testing.T) {
	registry := NewViewModelRegistry()
	if registry == nil {
		t.Error("Expected registry to not be nil")
	}
}

func TestViewModelRegistry_Add(t *testing.T) {
	block0 := &testBlock{name: "test0"}
	block1 := &testBlock{name: "test1"}

	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	err := r.Add(block0, block1)
	if err != nil {
		t.Error(err)
	}

	if r.registry[block0.name] != block0 {
		t.Error("Expected block0 to be added to registry")
	}

	if len(r.registry) != 2 {
		t.Error("Expected registry to have 2 blocks")
	}
}

func TestViewModelRegistry_Remove(t *testing.T) {
	block0 := &testBlock{name: "test0"}
	block1 := &testBlock{name: "test1"}

	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	r.registry[block0.name] = block0
	r.registry[block1.name] = block1

	err := r.Remove(block0, block1)
	if err != nil {
		t.Error(err)
	}

	if r.registry[block0.name] != nil {
		t.Error("Expected block to be removed from registry")
	}

	if r.registry[block1.name] != nil {
		t.Error("Expected block to be removed from registry")
	}

	if len(r.registry) != 0 {
		t.Error("Expected registry to be empty")
	}
}

func TestViewModelRegistry_Find(t *testing.T) {
	block := &testBlock{name: "test"}
	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	r.registry[block.name] = block

	b, err := r.Find(block.name)
	if err != nil {
		t.Error(err)
	}

	if b != block {
		t.Error("Expected block to be found in registry")
	}
}

func TestNewViewModelRegistryProxy(t *testing.T) {
	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	proxy := NewViewModelRegistryProxy(r)
	if proxy == nil {
		t.Error("Expected proxy to not be nil")
	}

	if proxy.registry != r {
		t.Error("Expected registry to be the equal")
	}
}

func TestViewModelRegistryProxy_Add(t *testing.T) {
	block0 := &testBlock{name: "test0"}
	block1 := &testBlock{name: "test1"}

	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	proxy := &ViewModelRegistryProxy{
		registry: r,
	}

	err := proxy.Add(block0, block1)
	if err != nil {
		t.Error(err)
	}

	if r.registry[block0.name] != block0 {
		t.Error("Expected block0 to be added to registry")
	}

	if len(r.registry) != 2 {
		t.Error("Expected registry to have 2 blocks")
	}
}

func TestViewModelRegistryProxy_Remove(t *testing.T) {
	block0 := &testBlock{name: "test0"}
	block1 := &testBlock{name: "test1"}

	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	r.registry[block0.name] = block0
	r.registry[block1.name] = block1

	proxy := &ViewModelRegistryProxy{
		registry: r,
	}

	err := proxy.Remove(block0, block1)
	if err != nil {
		t.Error(err)
	}

	if r.registry[block0.name] != nil {
		t.Error("Expected block to be removed from registry")
	}

	if r.registry[block1.name] != nil {
		t.Error("Expected block to be removed from registry")
	}

	if len(r.registry) != 0 {
		t.Error("Expected registry to be empty")
	}
}

func TestViewModelRegistryProxy_Find(t *testing.T) {
	block := &testBlock{name: "test"}
	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	r.registry[block.name] = block

	proxy := &ViewModelRegistryProxy{
		registry: r,
	}

	b, err := proxy.Find(block.name)
	if err != nil {
		t.Error(err)
	}

	if b != block {
		t.Error("Expected block to be found in registry")
	}
}

func TestGetRegistry(t *testing.T) {
	registry := GetRegistry()
	if registry == nil {
		t.Error("Expected registry to not be nil")
	}

	if reflect.DeepEqual(registry, GetRegistry()) == false {
		t.Error("Expected registry to be the equal")
	}
}
