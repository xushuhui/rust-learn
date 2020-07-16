package model

type AccountUser struct {
	Id            uint   `gorm:"column:id;type:int unsigned" json:"id"`                                       // 账号id
	Email         string `gorm:"column:email;type:varchar(30);default:''" json:"email"`                       // 邮箱
	Phone         string `gorm:"column:phone;type:varchar(15);default:''" json:"phone"`                       // 手机号
	Password      string `gorm:"column:password;type:char(32);default:''" json:"password"`                    // 密码
	CreateAt      int    `gorm:"column:create_at;type:int;default:'0'" json:"create_at"`                      // 创建时间
	CreateIpAt    string `gorm:"column:create_ip_at;type:varchar(12);default:''" json:"create_ip_at"`         // 创建ip
	LastLoginAt   int    `gorm:"column:last_login_at;type:int;default:'0'" json:"last_login_at"`              // 最后一次登录时间
	LastLoginIpAt string `gorm:"column:last_login_ip_at;type:varchar(12);default:''" json:"last_login_ip_at"` // 最后一次登录ip
	LoginTimes    int    `gorm:"column:login_times;type:int;default:'0'" json:"login_times"`                  // 登录次数
	Status        int8   `gorm:"column:status;type:tinyint(1);default:'0'" json:"status"`                     // 状态 1:enable, 0:disable, -1:deleted
}

//get real primary key name
func (accountUser *AccountUser) GetKey() string {
	return "id"
}

//get primary key in model
func (accountUser *AccountUser) GetKeyProperty() uint {
	return accountUser.Id
}

//set primary key
func (accountUser *AccountUser) SetKeyProperty(id uint) {
	accountUser.Id = id
}

//get real table name
func (accountUser *AccountUser) TableName() string {
	return "account_user"
}

func GetAccountUserById(id string) (accountUser *AccountUser, err error) {
	err = DB.Model(accountUser).First(accountUser, accountUser.GetKey()+" = '"+id+"'").Error

	return
}

func GetAccountUserOne(where string, args ...interface{}) (accountUser *AccountUser, err error) {
	err = DB.Model(accountUser).First(accountUser, where, args).Error

	return
}

func GetAccountUserList(page, limit int64, where string, condition interface{}) (list []*AccountUser, err error) {
	err = DB.Limit(limit).Offset((page-1)*limit).Find(list, where, condition).Error

	return
}

func (accountUser *AccountUser) Create() error {
	return DB.Model(accountUser).Create(accountUser).Error
}

func (accountUser *AccountUser) Update(update AccountUser) error {
	return DB.Model(accountUser).UpdateColumns(update).Error
}

func (accountUser *AccountUser) Delete() {
	DB.Model(accountUser).Delete(accountUser)
}
