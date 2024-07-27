package app

import "fmt"

type Page struct {
	PageTitle string
	Desc      string
	Filepath  string
	Content   string
}

func (i Page) Title() string {
	return i.PageTitle
}

func (i Page) Description() string {
	return i.Desc
}

func (i Page) FilterValue() string {
	return fmt.Sprintf("%s %s", i.PageTitle, i.Desc)
}
