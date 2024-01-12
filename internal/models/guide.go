package models

type CourseDetail struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	Desc          string      `json:"desc"`
	TypeActivity  string      `json:"type_activity"`
	ThemeActivity string      `json:"theme_activity"`
	Subtitles     []Subtitles `json:"subtitles"`
}

type Subtitles struct {
	ID        string  `json:"id"`
	Topic     string  `json:"topic"`
	ShortDesc string  `json:"shortdesc"`
	Content   Content `json:"content"`
}

type Content struct {
	Opening string   `json:"opening"`
	Closing string   `json:"closing"`
	Steps   []string `json:"step"`
}

type ListGuide struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	TypeActivity  string `json:"type_activity"`
	ThemeActivity string `json:"theme_activity"`
}

type ListMyGuide struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	TypeActivity  string `json:"type_activity"`
	ThemeActivity string `json:"theme_activity"`
}
