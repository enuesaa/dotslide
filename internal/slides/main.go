package slides

type DotSlide struct {
	Slides []Slide `json:"slides"`
}

type Slide struct {
	Cap         string `json:"cap"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Left        *Unit  `json:"left"`
	Center      *Unit  `json:"center"`
	Right       *Unit  `json:"right"`
	MetaLeft    *Unit  `json:"metaLeft"`
	MetaCenter  *Unit  `json:"metaCenter"`
	MetaRight   *Unit  `json:"metaRight"`
}

type Unit struct {
	Cap         string `json:"cap"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
