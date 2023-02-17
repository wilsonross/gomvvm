package gomvvm

import (
	"text/template"
)

func AddBlocks(blocks ...Blocker) error {
	return GetRegistry().Add(blocks...)
}

func TemplateLayout(layout []byte) (string, error) {
	pageContent, err := UnmarshalPageContent(layout, NewPageContent())
	if err != nil {
		return "", err
	}

	render := NewRender(template.ParseFiles)
	helperFactory := NewBlockHelperFactory(GetRegistry(), render)

	root, err := TemplatePageContent(pageContent.Layout, helperFactory)
	if err != nil {
		return "", err
	}

	return root, nil
}
