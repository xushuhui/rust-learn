package model

type ProductBrands struct {
	Id       uint   `gorm:"column:id;type:int unsigned" json:"id"`                           // 品牌ID
	Name     string `gorm:"column:name;type:varchar(255);default:''" json:"name"`            // 品牌名称
	Desc     string `gorm:"column:desc;type:varchar(255);default:''" json:"desc"`            // 品牌描述
	LogoUrl  string `gorm:"column:logo_url;type:varchar(255);default:''" json:"logo_url"`    // 品牌logo图片
	CreateAt string `gorm:"column:create_at;type:int unsigned;default:'0'" json:"create_at"` // 创建时间
	CreateBy uint   `gorm:"column:create_by;type:int unsigned;default:'0'" json:"create_by"` // 创建人staff_id
	UpdateAt string `gorm:"column:update_at;type:int unsigned;default:'0'" json:"update_at"` // 更新时间
	UpdateBy uint   `gorm:"column:update_by;type:int unsigned;default:'0'" json:"update_by"` // 修改人staff_id
	Status   uint8  `gorm:"column:status;type:tinyint unsigned;default:'0'" json:"status"`   // 状态 1:enable, 0:disable, -1:deleted
}

//get real primary key name
func (productBrands *ProductBrands) GetKey() string {
	return "id"
}

//get primary key in model
func (productBrands *ProductBrands) GetKeyProperty() uint {
	return productBrands.Id
}

//set primary key
func (productBrands *ProductBrands) SetKeyProperty(id uint) {
	productBrands.Id = id
}

//get real table name
func (productBrands *ProductBrands) TableName() string {
	return "product_brands"
}

func GetProductBrandsById(id string) (productBrands ProductBrands, err []error) {
	err = DB.First(&productBrands, id).GetErrors()
	if len(err) > 0 {
		return
	}
	return
}

func GetProductBrandsOne(where string, condition ...interface{}) (productBrands ProductBrands) {
	errs := DB.First(productBrands, where, condition).GetErrors()
	if len(errs) > 0 {
		return
	}
	return
}

//
func GetProductBrandsList(page, limit int64, where string, condition interface{}) (list []ProductBrands) {

	err := DB.Limit(limit).Offset((page-1)*limit).Find(list, where, condition).GetErrors()
	if err != nil {
		return nil
	}
	return
}
func GetProductBrandsPage(page, limit int64) (list []ProductBrands) {
	err := DB.Limit(limit).Offset((page - 1) * limit).Find(&list).GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func (productBrands *ProductBrands) Create() []error {
	return DB.Create(&productBrands).GetErrors()
}
func UpdateStatus(id string, status uint8) []error {
	brandModel, _ := GetProductBrandsById(id)
	brandModel.Status = status
	return DB.Save(&brandModel).GetErrors()
}
func (productBrands *ProductBrands) Update() []error {
	return DB.Model(&productBrands).Updates(productBrands).GetErrors()
}
func (productBrands *ProductBrands) SoftDelete() {

}
func (productBrands *ProductBrands) Delete() {
	DB.Model(productBrands).Delete(productBrands).GetErrors()
}
