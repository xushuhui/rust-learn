package request

type Brand struct {
	Name     string ` json:"name" binding:"required"` // 品牌名称
	Desc     string `json:"desc" binding:"required"`  // 品牌描述
	LogoUrl  string ` json:"logo_url" binding:"required"`
	CreateBy uint
}
