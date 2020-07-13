package model

type StaffInfo struct {
	Id       uint   `gorm:"column:id;type:int unsigned" json:"id"`                    // 员工id
	Uid      uint   `gorm:"column:uid;type:int unsigned;default:'0'" json:"uid"`      // 账号id
	Email    string `gorm:"column:email;type:varchar(30);default:''" json:"email"`    // 员工邮箱
	Password string `gorm:"column:password;type:char(32);default:''" json:"password"` // 密码
	Phone    string `gorm:"column:phone;type:varchar(15);default:''" json:"phone"`    // 员工手机号
	Name     string `gorm:"column:name;type:varchar(30);default:''" json:"name"`      // 员工姓名
	Avatar   string `gorm:"column:avatar;type:varchar(255);default:''" json:"avatar"` // 员工头像(相对路径)
	Gender   int8   `gorm:"column:gender;type:tinyint(1);default:'1'" json:"gender"`  // 员工性别 1 unknow 2 male 3 female
	CreateAt string `gorm:"column:create_at;type:int;default:'0'" json:"create_at"`   // 创建时间
	UpdateAt string `gorm:"column:update_at;type:int;default:'0'" json:"update_at"`   // 更新时间
}

//get real primary key name
func (staffInfo *StaffInfo) GetKey() string {
	return "id"
}

//get primary key in model
func (staffInfo *StaffInfo) GetKeyProperty() uint {
	return staffInfo.Id
}

//set primary key
func (staffInfo *StaffInfo) SetKeyProperty(id uint) {
	staffInfo.Id = id
}

//get real table name
func (staffInfo *StaffInfo) TableName() string {
	return "staff_info"
}

func GetStaffInfoById(id string) (staffInfo *StaffInfo, err error) {
	err = DB.Model(staffInfo).First(staffInfo, staffInfo.GetKey()+" = '"+id+"'").Error

	return
}

func GetStaffInfoOne(where string, args ...interface{}) (staffInfo *StaffInfo, err error) {
	err = DB.Model(staffInfo).First(&staffInfo, where, args).Error
	return
}
func GetStaffInfoPage(page, limit int64) (list []StaffInfo) {
	err := DB.Limit(limit).Offset((page - 1) * limit).Find(&list).GetErrors()
	if len(err) > 0 {
		return nil
	}
	return
}

func (staffInfo *StaffInfo) Create() error {
	return DB.Model(&staffInfo).Create(staffInfo).Error
}

func (staffInfo *StaffInfo) Update() []error {
	return DB.Model(&staffInfo).Update(staffInfo).GetErrors()
}

func (staffInfo *StaffInfo) Delete() {
	DB.Model(staffInfo).Delete(staffInfo)
}
