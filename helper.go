package gomvvm

type Helper interface {
	AddChild(child Helper)
	GetBlock() Blocker
	GetTemplatePath() string
	GetChildHtml(name string) string
	GetRender() Renderer
}

type BlockHelper struct {
	block        Blocker
	children     []Helper
	templatePath string
	renderer     Renderer
}

func (b *BlockHelper) AddChild(child Helper) {
	b.children = append(b.children, child)
}

func (b *BlockHelper) GetBlock() Blocker {
	return b.block
}

func (b *BlockHelper) GetTemplatePath() string {
	return b.templatePath
}

func (b *BlockHelper) GetChildHtml(name string) string {
	for _, c := range b.children {
		if c.GetBlock().GetName() != name {
			continue
		}

		t, _ := c.GetRender().Template(c)
		return t
	}

	return ""
}

func (b *BlockHelper) GetRender() Renderer {
	return b.renderer
}

type HelperFactory interface {
	New(Blocker, BlockContent) Helper
	NewComposite(BlockContent) Helper
	GetRenderer() Renderer
}

type BlockHelperFactory struct {
	registry Register
	renderer Renderer
}

func NewBlockHelperFactory(registry Register, renderer Renderer) *BlockHelperFactory {
	return &BlockHelperFactory{
		registry: registry,
		renderer: renderer,
	}
}

func (b BlockHelperFactory) New(block Blocker, content BlockContent) Helper {
	return &BlockHelper{
		block:        block,
		templatePath: content.Template,
		renderer:     b.GetRenderer(),
	}
}

func (b BlockHelperFactory) NewComposite(content BlockContent) Helper {
	block, _ := b.GetRegistry().Find(content.Block)
	helper := b.New(block, content)

	for _, c := range content.Children {
		child := b.NewComposite(c)
		helper.AddChild(child)
	}

	return helper
}

func (b BlockHelperFactory) GetRegistry() Register {
	return b.registry
}

func (b BlockHelperFactory) GetRenderer() Renderer {
	return b.renderer
}
