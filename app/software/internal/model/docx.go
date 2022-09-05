package model

type ArgDocx struct {
	HeaderContent  string   `json:"header_content"`  // 页眉内容
	TitleKey       string   `json:"title_key"`       // c，c#，java
	TitleFilters   []string `json:"filters"`         //
	ContentFilters []string `json:"content_filters"` //
	ContentsKey    []string `json:"content_key"`     // 内容关键字集合
	ContentsNum    int      `json:"contents_num"`    // 取多少条内容段
	ContentTitle   string   `json:"content_title"`   // 每个段落前的注释
}

type Code struct {
	Key     string `form:"key" json:"key" gorm:"column:key"`
	Content string `form:"content" json:"content" gorm:"column:content"`
}

func (Code) TableName() string {
	return "code"
}

type ArgCode struct {
	Codes []Code `form:"codes" json:"codes"`
}

type CodeInfo struct {
	Key     string   `form:"key" json:"key" gorm:"column:key"`
	Filters []string `form:"filters" json:"filters"`
	Num     int      `form:"num" json:"num"`
}
