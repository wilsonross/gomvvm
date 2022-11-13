package gomvvm

import (
	"github.com/go-yaml/yaml"
)

type PageContent struct {
	Layout []BlockContent `yaml:"layout"`
}

type BlockContent struct {
	Block    string         `yaml:"block"`
	Template string         `yaml:"template"`
	Children []BlockContent `yaml:"children"`
}

func NewPageContent() *PageContent {
	return &PageContent{}
}

func UnmarshalPageContent(data []byte, content *PageContent) (*PageContent, error) {
	err := yaml.UnmarshalStrict(data, content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func TemplatePageContent(content []BlockContent, helperFactory HelperFactory) (string, error) {
	var html string
	for _, c := range content {
		helper := helperFactory.NewComposite(c)
		renderer := helperFactory.GetRenderer()

		t, err := renderer.Template(helper)
		if err != nil {
			return "", err
		}
		html += t
	}

	return html, nil
}
