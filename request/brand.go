package request

type Brand struct {
	Id       int64  `json:"id,omitempty"`
	Name     string ` json:"name" binding:"required" comment:"品牌名称"` // 品牌名称
	Desc     string `json:"desc" binding:"required" comment:"品牌描述"`  // 品牌描述
	LogoUrl  string ` json:"logo_url" binding:"required"`
	CreateBy uint
}

type Staff struct {
	Id       int    `json:"id,omitempty"`
	Uid      int    `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Avatar   string `json:"avatar,omitempty"`
	Gender   int    `json:"gender"`
}
type Category struct {
	Id       uint   ` json:"id,omitempty"`           // 分类ID
	Pid      uint   `json:"pid" binding:"required"`  // 父ID
	Name     string `json:"name" binding:"required"` // 分类名称
	Desc     string ` json:"desc"`                   // 分类描述
	PicUrl   string ` json:"pic_url"`                // 分类图片
	Path     string `json:"path"`                    // 分类地址{pid}-{child_id}-...
	CreateBy uint   ` json:"create_by"`              // 创建人staff_id
	UpdateBy uint   ` json:"update_by"`              // 修改人staff_id
	Status   uint8  ` json:"status"`                 // 状态 1:enable, 0:disable, -1:deleted
}
type Sku struct {
}
type Spu struct {
}
