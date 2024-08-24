package theme

type theme struct {
	Foreground string
	Faint      string
	Prompt     string
	Cyan       string
	Green      string
	Orange     string
	Pink       string
	Purple     string
	Red        string
	Yellow     string
}

var Dracula = theme{
	Foreground: "#f8f8f2",
	Faint:      "#44475a",
	Prompt:     "#8e50e6",
	Cyan:       "#8be9fd",
	Green:      "#50fa7b",
	Orange:     "#ffb86c",
	Pink:       "#ff79c6",
	Purple:     "#bd93f9",
	Red:        "#ff5555",
	Yellow:     "#f1fa8c",
}
