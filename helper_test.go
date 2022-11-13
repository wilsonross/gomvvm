package gomvvm

import (
	"testing"
)

type testBlock struct {
	name string
}

func (t *testBlock) GetName() string {
	return t.name
}

type testRenderer struct{}

func (t *testRenderer) Template(_ Helper) (string, error) {
	return "test", nil
}

func TestBlockHelper_AddChild(t *testing.T) {
	helper := &BlockHelper{templatePath: "test"}
	child := &BlockHelper{templatePath: "test"}

	helper.AddChild(child)

	if len(helper.children) != 1 {
		t.Error("Expected 1 child")
	}
}

func TestBlockHelper_GetBlock(t *testing.T) {
	block := &testBlock{name: "test"}
	helper := &BlockHelper{block: block}

	if helper.GetBlock() == nil {
		t.Error("Expected block")
	}
}

func TestBlockHelper_GetTemplatePath(t *testing.T) {
	helper := &BlockHelper{templatePath: "test"}

	if helper.GetTemplatePath() != "test" {
		t.Error("Expected template path to be 'test'")
	}
}

func TestBlockHelperFactory_GetChildHtml(t *testing.T) {
	block := &testBlock{name: "test"}
	renderer := &testRenderer{}

	helper := &BlockHelper{block: block, templatePath: "test", renderer: renderer}
	helper.children = append(helper.children, helper)

	childHtml := helper.GetChildHtml("test")
	if childHtml != "test" {
		t.Error("Expected child html to be test")
	}
}

func TestBlockHelper_GetRender(t *testing.T) {
	renderer := &Render{}
	helper := &BlockHelper{renderer: renderer}

	if helper.GetRender() == nil {
		t.Error("Expected renderer")
	}
}

func TestNewBlockHelperFactory(t *testing.T) {
	registry := &ViewModelRegistry{}
	renderer := &Render{}
	factory := NewBlockHelperFactory(registry, renderer)

	if factory == nil {
		t.Error("Expected factory")
	}

	if factory.registry == nil {
		t.Error("Expected registry")
	}

	if factory.renderer == nil {
		t.Error("Expected renderer")
	}
}

func TestBlockHelperFactory_New(t *testing.T) {
	block := &testBlock{name: "test"}
	factory := &BlockHelperFactory{}
	helper := factory.New(block, BlockContent{})

	if helper == nil {
		t.Error("Expected helper")
	}
}

func TestBlockHelperFactory_NewComposite(t *testing.T) {
	block := &testBlock{name: "test"}
	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	r.registry[block.name] = block
	content := BlockContent{
		Block:    "test",
		Template: "test",
		Children: []BlockContent{
			{
				Block:    "test",
				Template: "test",
			},
		},
	}

	factory := &BlockHelperFactory{registry: r}
	helper := factory.NewComposite(content)

	if helper == nil {
		t.Error("Expected helper")
	}

	if helper.(*BlockHelper).block == nil {
		t.Error("Expected block")
	}

	if len(helper.(*BlockHelper).children) != 1 {
		t.Error("Expected 1 child")
	}
}

func TestBlockHelperFactory_GetRegistry(t *testing.T) {
	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	factory := &BlockHelperFactory{registry: r}

	if factory.GetRegistry() == nil {
		t.Error("Expected registry")
	}
}

func TestBlockHelperFactory_GetRenderer(t *testing.T) {
	r := &Render{}
	factory := &BlockHelperFactory{renderer: r}

	if factory.GetRenderer() == nil {
		t.Error("Expected renderer")
	}
}
