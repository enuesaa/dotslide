package main

type SlideDetail struct {
	Title       string `yaml:"title" json:"title"`
	Description string `yaml:"description" json:"description"`
}

type Slide struct {
	Title       string       `yaml:"title" json:"title"`
	TitleStyle  string       `yaml:"titleStyle,omitempty" json:"titleStyle,omitempty"`
	Description string       `yaml:"description,omitempty" json:"description,omitempty"`
	Image       string       `yaml:"image,omitempty" json:"image,omitempty"`
	Left        *SlideDetail `yaml:"left,omitempty" json:"left,omitempty"`
	Right       *SlideDetail `yaml:"right,omitempty" json:"right,omitempty"`
}
