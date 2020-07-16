package model

type AccountPlatform struct {
	Id            uint   `gorm:"column:id;type:int unsigned" json:"id"`                                   // 自增id
	Uid           uint   `gorm:"column:uid;type:int unsigned;default:'0'" json:"uid"`                     // 账号id
	PlatformId    string `gorm:"column:platform_id;type:varchar(60);default:''" json:"platform_id"`       // 平台id
	PlatformToken string `gorm:"column:platform_token;type:varchar(60);default:''" json:"platform_token"` // 平台access_token
	Type          int8   `gorm:"column:type;type:tinyint(1);default:'0'" json:"type"`                     // 平台类型 0:未知,1:facebook,2:google,3:wechat,4:qq,5:weibo,6:twitter
	Nickname      string `gorm:"column:nickname;type:varchar(60);default:''" json:"nickname"`             // 昵称
	Avatar        string `gorm:"column:avatar;type:varchar(255);default:''" json:"avatar"`                // 头像
	CreateAt      int    `gorm:"column:create_at;type:int;default:'0'" json:"create_at"`                  // 创建时间
	UpdateAt      int    `gorm:"column:update_at;type:int;default:'0'" json:"update_at"`                  // 更新时间
}

//get real primary key name
func (accountPlatform *AccountPlatform) GetKey() string {
	return "id"
}

//get primary key in model
func (accountPlatform *AccountPlatform) GetKeyProperty() uint {
	return accountPlatform.Id
}

//set primary key
func (accountPlatform *AccountPlatform) SetKeyProperty(id uint) {
	accountPlatform.Id = id
}

//get real table name
func (accountPlatform *AccountPlatform) TableName() string {
	return "account_platform"
}

func GetAccountPlatformById(id string) (accountPlatform *AccountPlatform, err error) {
	err = DB.Model(accountPlatform).First(accountPlatform, accountPlatform.GetKey()+" = '"+id+"'").Error

	return
}

//func GetAccountPlatformOne(where string, args ...interface{}) (accountPlatform *AccountPlatform) {
//	err := DB,Model(accountPlatform).First(accountPlatform, where, args...).GetErrors()
//	if len(err) > 0 {
//		return nil
//	}
//	return
//}
//
//func GetAccountPlatformList(page, limit int64, where string, condition interface{}) (list []*AccountPlatform) {
//	err := DB,Model(accountPlatform).Limit(limit).Offset((page-1)*limit).Find(list, where, condition).GetErrors()
//	if err != nil {
//		return nil
//	}
//	return
//}

func (accountPlatform *AccountPlatform) Create() error {
	return DB.Model(accountPlatform).Create(accountPlatform).Error
}

func (accountPlatform *AccountPlatform) Update(update AccountPlatform) []error {
	return DB.Model(accountPlatform).UpdateColumns(update).GetErrors()
}

func (accountPlatform *AccountPlatform) Delete() {
	DB.Model(accountPlatform).Delete(accountPlatform)
}
