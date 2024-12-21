package slides

type DotSlide struct {
	Slides []Slide `json:"slides"`
}

type Slide struct {
	Cap         string `json:"cap,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Left        *Unit  `json:"left,omitempty"`
	Center      *Unit  `json:"center,omitempty"`
	Right       *Unit  `json:"right,omitempty"`
	MetaLeft    *Unit  `json:"metaLeft,omitempty"`
	MetaCenter  *Unit  `json:"metaCenter,omitempty"`
	MetaRight   *Unit  `json:"metaRight,omitempty"`
}

type Unit struct {
	Cap         string `json:"cap,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Links       []UnitLink `json:"links,omitempty"`
	Image       string `json:"image,omitempty"`
}

type UnitLink struct {
	Title string `json:"title,omitempty"`
	Url   string `json:"url,omitempty"`
}
