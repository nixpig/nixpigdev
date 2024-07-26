package main

import "fmt"

type page struct {
	title    string
	desc     string
	filepath string
}

func (i page) Title() string {
	return i.title
}

func (i page) Description() string {
	return i.desc
}

func (i page) FilterValue() string {
	return fmt.Sprintf("%s %s", i.title, i.desc)
}
