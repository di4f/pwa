package ui

// The package implements basic navigation
// through StringValuers.

import (
	"github.com/omnipunk/pwa/v9/app"
	"strings"
	"log"
)

// INav is the interface that describes navigational
// component.
type INav interface {
	app.UI
	Parent(INav) INav
	Show(bool) INav
	Define(string, app.UI) INav
	Default(string) INav
	Valuer(Valuer[string]) INav
	Class(string) INav

	GetParent() INav
	SetValue(app.Context, string) error
	GetValue(app.Context) string
	Path() string
	NumParent() int
}

type uiMap map[string] app.UI
type navCompo struct {
	app.Compo
	class string
	parent *navCompo
	show bool
	compos uiMap
	errCompo app.UI
	value string
	def string
	valuer Valuer[string]
	path string
}

func Nav() INav {
	return &navCompo{
		show: true,
	}
}

func (n *navCompo) Parent(
	parent INav,
) INav {
	n.parent = parent.(*navCompo)
	return n
}

func (n *navCompo) Show(
	show bool,
) INav {
	n.show = show
	return n
}

func (n *navCompo) OnMount(c app.Context) {
	var pth string
	if n.parent == nil {
		pth = app.Window().URL().Path
		n.path = pth
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

func (n *navCompo) Render() app.UI {
	if !n.Mounted() {
		return app.Text("")
	}
	if n.parent == nil {
		n.path = app.Window().URL().Path
	}

	// The thing is made to interact
	// with other components so
	// you do not have to use the compo
	// to switch something in your compo.
	if !n.show {
		return app.Text("")
	}

	compo, ok := n.compos[n.value]
	if !ok {
		log.Printf("%q: the path is not registered", n.value)
		return app.Text("")
	}

	return compo
}

func (n *navCompo) GetParent() INav {
	return n.parent
}

func (n *navCompo) Path() string {
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

	ret = append(ret, "")

	return ret
}

func (n *navCompo) NumParent() int {
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

func (n *navCompo) SetValue(
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

func (n *navCompo) GetValue(
	c app.Context,
) string {
	return n.value
}

func (n *navCompo) Valuer(valuer Valuer[string]) INav {
	n.valuer = valuer
	return n
}

func (n *navCompo) Define(
	name string,
	ui app.UI,
) INav {
	if n.compos == nil {
		n.compos = make(uiMap)
	}
	n.compos[name] = ui
	return n
}

func (n *navCompo) Class(class string) INav {
	n.class = class
	return n
}

func (n *navCompo) Default(def string) INav {
	n.def = def
	return n
}

