package slides

type DotSlide struct {
	Slides []Slide `json:"slides" yaml:"slides"`
}

type Slide struct {
	Cover       string     `json:"cover,omitempty" yaml:"cover"`
	Cap         string     `json:"cap,omitempty" yaml:"cap"`
	Title       string     `json:"title,omitempty" yaml:"title"`
	Description string     `json:"description,omitempty" yaml:"description"`
	Links       []UnitLink `json:"links,omitempty" yaml:"links"`
	Files       []UnitFile `json:"files,omitempty" yaml:"files"`
	Terminal    string     `json:"terminal,omitempty" yaml:"terminal"`
	Image       string     `json:"image,omitempty" yaml:"image"`
	Left        *Unit      `json:"left,omitempty" yaml:"left"`
	Center      *Unit      `json:"center,omitempty" yaml:"center"`
	Right       *Unit      `json:"right,omitempty" yaml:"right"`
	MetaLeft    *Unit      `json:"metaLeft,omitempty" yaml:"metaLeft"`
	MetaCenter  *Unit      `json:"metaCenter,omitempty" yaml:"metaCenter"`
	MetaRight   *Unit      `json:"metaRight,omitempty" yaml:"metaRight"`
	Footer      *Unit      `json:"footer,omitempty" yaml:"footer"`
}

type Unit struct {
	Cap         string     `json:"cap,omitempty" yaml:"cap"`
	Title       string     `json:"title,omitempty" yaml:"title"`
	Description string     `json:"description,omitempty" yaml:"description"`
	Links       []UnitLink `json:"links,omitempty" yaml:"links"`
	Files       []UnitFile `json:"files,omitempty" yaml:"files"`
	Terminal    string     `json:"terminal,omitempty" yaml:"terminal"`
	Image       string     `json:"image,omitempty" yaml:"image"`
}

type UnitLink struct {
	Title string `json:"title,omitempty" yaml:"title"`
	Url   string `json:"url,omitempty" yaml:"url"`
}
type UnitFile struct {
	Filename string `json:"filename,omitempty" yaml:"filename"`
	Code     string `json:"code,omitempty" yaml:"code"`
}
