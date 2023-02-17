## About

<div>
  <br />
  <p>Provides a foundation for creating a model-view-viewmodel style architecture in Go. Define page layouts in YAML files that look up associated templates and view models—designed to allow for more orthogonal templating in Go and reusable components. Future releases aim to include layout overrides, caching, improved error handling and escapement.</p>
  <br />
  <p>It is inspired by Magento 2 layouts and layout instructions. Where ‘blocks’ are created that contain methods the templates can invoke. In my opinion, this makes for a more manageable templating experience for large, unwieldy pages. Instead of bloated Magento layout files, this library uses layouts written in a more straightforward schema.</p>
  <br />
</div>

## Lifecycle

<div>
  <br />
  <p>Blocks that match an interface can be added to a registry. The interface only requires a single receiver function that fetches the block’s name. A layout file will then look up templates and their associated block. We can create a simple system of reusable nested components containing their own model from these basic building blocks. Each child component can include its unique data and receiver functions.</p>
  <br />
</div>

## Installation

<br />

```bash
go get github.com/wilsonross/gomvvm
```

<br />

## Getting Started

<div>
  <br />
  <p>The following code snippet shows how to create a basic block and add it to a registry. We can also define additional data on the struct and receiver functions. This next section looks verbose but would be suitable when grouping by type. I have intentionally added two components to demonstrate how nested children work.</p>
  <br />
</div>

```go
package main

import (
	layout "github.com/wilsonross/gomvvm"
)

func main() {
	header := Header{
		name: "header",
	}

	menu := Menu{
		name:      "menu",
		menuLinks: []string{"Home", "About", "Contact"},
	}

	// Add the blocks to the registry
	layout.AddBlocks(menu, header)
}

type Header struct {
	name string
}

// Required for our interface
func (h Header) GetName() string {
	return h.name
}

type Menu struct {
	name      string
	menuLinks []string
}

// Required for our interface
func (m Menu) GetName() string {
	return m.name
}

// Data we would like to pass to our template
// We could also export the field directly from our struct
func (m Menu) GetMenuLinks() []string {
	return m.menuLinks
}
```

<div>
  <br />
  <p>Let’s now focus on our layout. We have a root <code>layout</code> key, followed by three simple directives. The <code>block</code> directive is required and is the identifier of the block we have previously set. In our example, it will be <code>menu</code>. The ‘template’ directive is a required path to our HTML file. The <code>children</code> directive is an optional array where we can set child components using the same schema.</p>
  <br />
</div>

```yaml
layout:
  - block: "header"
    template: "header.html"
    children:
      - block: "menu"
        template: "menu.html"
```

<div>
  <br />
  <p>Here are the two HTML templates that we just defined. We are invoking <code>getChildHtml</code> to produce our nested HTML. Additionally, we are retrieving the data we described earlier. Because it is necessary to embed the HTML content in the parent component, we are losing the escapement of the Go templater. I am currently investigating ways around this caveat. However, I strongly recommend you escape the final string for the moment to prevent XSS.</p>
  <br />
</div>

```html
<div class="header">
    {{ .GetChildHtml "menu" }}
</div>
```

<br />

```html
<ul class="menu">
  {{range $name := .GetBlock.GetMenuLinks }}
  <li>{{$name}}</li>
  {{end}}
</ul>
```

<div>
  <br />
  <p>Currently, the library requires the user to read the layouts. This will likely change at a later date. We can finally call the TemplateLayout function to produce a string of HTML. By adding more components, more complex layouts can be created.</p>
  <br />
</div>

```go
func main() {
	header := Header{
		name: "header",
	}

	menu := Menu{
		name:      "menu",
		menuLinks: []string{"Home", "About", "Contact"},
	}

	// Add the blocks to the registry
	layout.AddBlocks(menu, header)

	data, err := os.ReadFile("layout.yml")
	if err != nil {
		panic(err)
	}

	// Render the template
	output, err := layout.TemplateLayout(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
	/*<div class="header">
	  <ul class="menu">
	    <li>Home</li>
	    <li>About</li>
	    <li>Contact</li>
	  </ul>
	</div>*/
}
```

<br />

## Roadmap

<br />

- [x] First release
- [ ] Escapement
- [ ] Improved error handling
- [ ] Caching and improved file reading
- [ ] Layout overrides

<br />

## Contact

<br />

Ross Wilson - [ross.wilson.190298@gmail.com](mailto:ross.wilson.190298@gmail.com)

<br />
