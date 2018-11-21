package http

// View テンプレート用のビューオブジェクト
type View struct {
	Head ViewHead
	Body ViewBody
}

type ViewHead struct {
	Title         string
	Scripts       []ViewScript
	WidgetScripts []ViewWidgetScripts
}

type ViewBody struct {
	Nav  []ViewNav
	Main interface{}
}

type ViewScript struct {
	Src  string
	Attr string
}

type ViewWidgetScripts struct {
	Src     string
	TagName string
}

type ViewNav struct {
	Title   string
	Submenu []ViewNavLink
}

type ViewNavLink struct {
	Link string
	Text string
}
