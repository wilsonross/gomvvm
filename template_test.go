package gomvvm

import (
	"testing"
	"text/template"
)

func TestNewTemplateWriter(t *testing.T) {
	tw := NewTemplateWriter()
	if tw == nil {
		t.Error("Expected template writer to be not nil")
	}
}

func TestTemplateWriter_Write(t *testing.T) {
	tw := &TemplateWriter{}

	written, err := tw.Write([]byte("test"))
	if err != nil {
		t.Error(err)
	}

	written, err = tw.Write([]byte(" test"))
	if err != nil {
		t.Error(err)
	}

	if written != 9 {
		t.Error("Expected 9 bytes to be written")
	}

	if string(*tw) != "test test" {
		t.Error("Expected template writer to be test")
	}
}

func TestTemplateWriter_ToString(t *testing.T) {
	tw := &TemplateWriter{}

	_, err := tw.Write([]byte("test"))
	if err != nil {
		t.Error(err)
	}

	if tw.ToString() != "test" {
		t.Error("Expected template writer to be test")
	}
}

func TestNewRender(t *testing.T) {
	r := NewRender(func(filenames ...string) (*template.Template, error) {
		return nil, nil
	})

	if r == nil {
		t.Error("Expected render to be not nil")
	}
}

func TestRender_Template(t *testing.T) {
	r := &Render{
		parse: func(filenames ...string) (*template.Template, error) {
			parsed, _ := template.New("test").Parse("template text")
			return parsed, nil
		},
	}

	block := &testBlock{
		name: "test",
	}

	helper := &BlockHelper{
		block:        block,
		templatePath: "test",
		renderer:     r,
	}

	text, err := r.Template(helper)

	if err != nil {
		t.Error(err)
	}

	if text != "template text" {
		t.Error("Expected template writer to be equal to template text")
	}
}

func TestRender_GetParseFunction(t *testing.T) {
	r := &Render{
		parse: func(filenames ...string) (*template.Template, error) {
			return nil, nil
		},
	}

	if r.GetParseFunction() == nil {
		t.Error("Expected parse function to be not nil")
	}
}
