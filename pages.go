package seed

import "github.com/qlova/seed/script"
import "github.com/qlova/seed/style/css"

//Page is a page of an app, or seed.
type Page struct {
	Seed
}

//NewPage returns a new Page.
func NewPage() Page {
	seed := New()
	seed.SetCol()

	seed.page = true
	seed.class = "page"

	seed.SetWillChange(css.Property.Display)
	seed.SetWillChange(css.Property.Transform)

	seed.SetPosition(css.Absolute)
	seed.SetTop(css.Zero)
	seed.SetLeft(css.Zero)
	//seed.Style.Style.SetWidth(css.Number(100).Vw())
	//seed.Style.SetHeight(100)*/
	seed.Style.SetSize(100, 100)

	return Page{seed}
}

//SetTag sets a tag associated with this page.
func (page Page) SetTag(name string) {
	if page.tags == nil {
		page.tags = make(map[string]bool)
	}
	page.tags[name] = true
}

type pages map[string]Page

func (p pages) Get(key string) Page {
	return p[key]
}

//NewPage returns a NewPage attached to a given seed.
func (seed Seed) NewPage() Page {
	return AddPageTo(seed)
}

//AddPageTo adds a page to a parent.
func AddPageTo(parent Interface) Page {
	var page = NewPage()
	parent.Root().Add(page)
	return page
}

//SetBack sets the page that this page should go to when a back button is pressed.
func (page Page) SetBack(back Page) {
	page.SetAttributes(page.Attributes() + ` data-back="` + back.ID() + `"`)
}

//SyncVisibilityWith sets the given seed to be visible when the page is visible and hidden when the page is hidden.
func (page Page) SyncVisibilityWith(seed Interface) {
	var root = seed.Root()
	page.OnPageEnter(func(q Script) {
		root.Script(q).SetVisible()
	})
	page.OnPageExit(func(q Script) {
		root.Script(q).SetHidden()
	})
}

//Script returns a script interface to the page.
func (page Page) Script(q Script) script.Page {
	return script.Page{page.Seed.Script(q)}
}

//OnBack is triggered before the back action is triggered, return q.Bool(true) to prevent default behaviour.
func (page Page) OnBack(f func(q Script)) {
	page.OnReady(func(q Script) {
		q.Javascript("{")
		q.Javascript("let old_onback = " + page.Script(q).Element() + ".onback;")
		q.Javascript(page.Script(q).Element() + ".onback = function() {")
		q.Javascript("if (old_onback) old_onback();")
		f(q)
		q.Javascript("};")
		q.Javascript("}")
	})
}

//OnPageEnter runs a script when this page is entered/ongoto.
func (page Page) OnPageEnter(f func(Script)) {
	page.On("pageenter", func(q Script) {
		f(q)
	})
}

//OnPageExit runs a script when leaving this page (onleave).
func (page Page) OnPageExit(f func(Script)) {
	page.On("pageexit", func(q Script) {
		f(q)
	})
}
