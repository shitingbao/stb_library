package model

type ArgDocx struct {
	HeaderContent  string   `json:"header_content"`           // 页眉内容
	Language       string   `form:"language" json:"language"` // c，c#，java
	HeaderFilters  []string `json:"header_filters"`           //
	ContentFilters []string `json:"content_filters"`          //
	ContentsKey    []string `json:"content_key"`              // 内容关键字集合
	ContentsNum    int      `json:"contents_num"`             // 取多少条内容段
	ContentTitle   string   `json:"content_title"`            // 每个段落前的注释
}

type Code struct {
	Key      string `form:"key" json:"key" gorm:"column:key"`
	Language string `form:"language" json:"language"`
	Content  string `form:"content" json:"content" gorm:"column:content"`
}

type ArgCode struct {
	Codes []Code `form:"codes" json:"codes"`
}

type ArgCodeModel struct {
	Codes string `form:"codes" json:"codes"`
}
