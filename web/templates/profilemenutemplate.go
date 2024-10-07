package templates

import "html/template"

type Menu struct {
	menu template.HTML
}

func LoadProfilMenu() Menu {
	return Menu{}
}
