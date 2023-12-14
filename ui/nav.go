package ui

// The package implements basic navigation
// through StringValuers.

import (
	"github.com/di4f/pwa/app"
	"strings"
	"log"
)

type uiMap map[string]app.UI
type NavCompo struct {
	app.Compo
	class    string
	parent   *NavCompo
	show     bool
	compos   uiMap
	errCompo app.UI
	value    string
	def      string
	valuer   Valuer[string]
}

func Nav(parent *NavCompo) *NavCompo {
	return &NavCompo{
		show:   true,
		parent: parent,
	}
}

func (n *NavCompo) Parent() *NavCompo {
	return n.parent
}

func (n *NavCompo) Show(
	show bool,
) *NavCompo {
	n.show = show
	return n
}

func (n *NavCompo) OnMount(c app.Context) {
	log.Println("mounting")
	var pth string
	if n.parent == nil {
		pth = app.Window().URL().Path
	} else {
		pth = n.parent.Path()
	}

	values := PathValues(pth)

	i := n.NumParent()
	val := ""
	if i < len(values) {
		val = values[i]
	}

	if val == "" {
		_, ok := n.compos[val]
		if !ok {
			n.SetValue(c, n.def)
		}
	} else {
		n.SetValue(c, val)
	}

	if n.valuer != nil {
		n.valuer.SetValue(c, n.GetValue(c))
	}
}

func (n *NavCompo) Render() app.UI {
	log.Println("rendering")
	// The thing is made to interact
	// with other components so
	// you do not have to use the compo
	// to switch something in your compo.
	if !n.show {
		return app.Text("")
	}

	value := n.GetValue(nil)
	compo, ok := n.compos[value]
	if !ok {
		log.Printf("%q: the path is not registered", value)
		return app.Text("")
	}

	return compo
}

func (n *NavCompo) Path() string {
	if n == nil {
		return "/"
	}

	if n.value == "" {
		return n.parent.Path()
	}

	return n.parent.Path() + n.value + "/"
}

func PathValues(path string) []string {
	if len(path) == 0 {
		return []string{}
	}

	values := strings.Split(path, "/")

	ret := []string{}
	for _, v := range values {
		if v == "" {
			continue
		}
		ret = append(ret, v)
	}

	//ret = append(ret, "")

	return ret
}

func (n *NavCompo) NumParent() int {
	if n == nil {
		return -1
	}

	i := 0
	for {
		n = n.parent
		if n == nil {
			break
		}
		i++
	}

	return i
}

func (n *NavCompo) OnUpdate(c app.Context) {
}

func (n *NavCompo) SetValue(
	c app.Context,
	value string,
) error {
	if n.value == value {
		return nil
	}

	n.value = value
	if n.Mounted() {
		n.Update()
		c.Navigate(n.Path())
	}

	return nil
}

func (n *NavCompo) GetValue(
	c app.Context,
) string {
	return n.value
}

func (n *NavCompo) Valuer(valuer Valuer[string]) *NavCompo {
	n.valuer = valuer
	return n
}

func (n *NavCompo) Define(
	name string,
	ui app.UI,
) *NavCompo {
	if n.compos == nil {
		n.compos = make(uiMap)
	}
	n.compos[name] = ui
	return n
}

func (n *NavCompo) Class(class string) *NavCompo {
	n.class = class
	return n
}

func (n *NavCompo) Default(def string) *NavCompo {
	n.def = def
	return n
}
