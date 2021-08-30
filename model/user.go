package model

type User struct {
	UID      int64  `json:"uid" gorm:"primary_key"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	AreaCode string `json:"area_code"`
	Mobile   string `json:"mobile"`
	Salt     string `json:"salt"`
	Pass     string `json:"pass"`
	Time     int64  `json:"time"`
	//0:未激活 1：激活正常 2 冻结
	Status     int    `json:"status"`
	Tags       string `json:"tags"`
	UpdateTime int64  `json:"update_time"`
	IsDelete   int    `json:"is_delete"`
	DeleteTime int64  `json:"delete_time"`
}

type UserRegistInfo struct {
	Email            string    `json:"email"`
	Mobile           string    `json:"mobile"`
	MobileArea       string    `json:"mobileArea"`
	Password         string    `json:"password"`
	Inviter          string    `json:"inviter"`
	RegistWay        RegistWay `json:"registWay"`
	ServiceAgreement bool      `json:"serviceAgreement"`
}

type RegistWay int8

const REGIST_WAY_EMAIL RegistWay = 0
const REGIST_WAY_MOBILE RegistWay = 1
