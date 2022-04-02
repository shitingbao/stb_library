package biz

type Menu struct {
	Id       int    `form:"id" json:"id" gorm:"column:id"`
	Name     string `form:"name" json:"name" gorm:"column:name"`
	UserId   int    `form:"user_id" json:"user_id" gorm:"column:user_id"`
	ParentId int    `form:"parent_id" json:"parent_id" gorm:"column:parent_id"`
}

type XiaojiUseCase struct {
	sLog           *SlogUseCase
	defaultFileDir DefaultFileDir
	xiaoji         XiaojiRepo
}

type XiaojiRepo interface {
	GetMenuList(userId int) ([]Menu, error)
}

func NewXiaojiCase(defaultDir DefaultFileDir, s *SlogUseCase) *XiaojiUseCase {
	return &XiaojiUseCase{defaultFileDir: defaultDir, sLog: s}
}

func (x *XiaojiUseCase) GetMenuList(userId int) ([]Menu, error) {
	return x.xiaoji.GetMenuList(userId)
}
