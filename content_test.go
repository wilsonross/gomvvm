package gomvvm

import (
	"fmt"
	"testing"
	"text/template"
)

func TestNewPageContent(t *testing.T) {
	content := NewPageContent()
	if content == nil {
		t.Error("Expected content to be not nil")
	}
}

func TestUnmarshalPageContent(t *testing.T) {
	data := []byte(`
layout:
  - block: header
    template: header
    children:
      - block: header
        template: header
  - block: footer
    template: footer
`)

	content, err := UnmarshalPageContent(data, &PageContent{})
	if err != nil {
		t.Error(err)
	}

	if content == nil {
		t.Error("Expected content to be not nil")
	}

	if len(content.Layout) != 2 {
		t.Error("Expected content.Layout to have 2 items")
	}

	if len(content.Layout[0].Children) != 1 {
		t.Error("Expected header children to have 1 item")
	}
}

func TestTemplatePageContent(t *testing.T) {
	content := []BlockContent{
		{
			Block:    "header",
			Template: "header",
			Children: []BlockContent{
				{
					Block:    "nav",
					Template: "nav",
				},
			},
		},
		{
			Block:    "footer",
			Template: "footer",
		},
	}

	headerBlock := &testBlock{
		name: "header",
	}

	navBlock := &testBlock{
		name: "nav",
	}

	footerBlock := &testBlock{
		name: "footer",
	}

	r := &ViewModelRegistry{
		registry: make(map[string]Blocker),
	}

	err := r.Add(headerBlock, navBlock, footerBlock)
	if err != nil {
		t.Error(err)
	}

	renderer := &Render{
		parse: func(filenames ...string) (*template.Template, error) {
			var parsed *template.Template
			var err error

			switch filenames[0] {
			case "header":
				parsed, err = template.New("header").Parse("header {{ .GetChildHtml \"nav\" }} ")
			case "nav":
				parsed, err = template.New("nav").Parse("nav")
			case "footer":
				parsed, err = template.New("footer").Parse("footer")
			default:
				err = fmt.Errorf("unexpected template name %s", filenames[0])
			}

			if err != nil {
				return nil, err
			}

			return parsed, nil
		},
	}

	helperFactory := &BlockHelperFactory{
		registry: r,
		renderer: renderer,
	}

	text, err := TemplatePageContent(content, helperFactory)

	if err != nil {
		t.Error(err)
	}

	if text != "header nav footer" {
		t.Errorf("Expected text to be 'header nav footer', got '%s'", text)
	}
}
