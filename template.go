package gomvvm

import (
	"text/template"
)

type Writeable interface {
	Write(p []byte) (n int, err error)
	ToString() string
}

type TemplateWriter []byte

func NewTemplateWriter() *TemplateWriter {
	return &TemplateWriter{}
}

func (w *TemplateWriter) Write(p []byte) (n int, err error) {
	*w = append(*w, p...)
	return len(*w), nil
}

func (w *TemplateWriter) ToString() string {
	return string(*w)
}

type Renderer interface {
	Template(helper Helper) (string, error)
}

type Render struct {
	parse func(filenames ...string) (*template.Template, error)
}

func NewRender(parse func(filenames ...string) (*template.Template, error)) *Render {
	return &Render{
		parse: parse,
	}
}

func (r Render) Template(helper Helper) (string, error) {
	path := helper.GetTemplatePath()
	t, err := r.GetParseFunction()(path)
	if err != nil {
		return "", err
	}

	writer := NewTemplateWriter()

	err = t.Execute(writer, helper)
	if err != nil {
		return "", err
	}

	return writer.ToString(), nil
}

func (r Render) GetParseFunction() func(filenames ...string) (*template.Template, error) {
	return r.parse
}
