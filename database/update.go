package database

import (
	"github.com/yongjie0203/go-trade-user/model"
	"github.com/yongjie0203/go-universal/db"
)

func TableUpdateRegister() {
	db.TableUpdateRegister("user", &model.User{})
}
