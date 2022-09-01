package model

type ArgDocx struct {
	HeaderContent  string   `json:"header_content"`
	TitleKey       string   `json:"title_key"`
	TitleFilters   []string `json:"filters"`
	ContentFilters []string `json:"content_filters"`
	ContentsKey    []string `json:"content"`
	ContentsNum    int      `json:"contents_num"`
}
