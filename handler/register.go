package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
	"github.com/yongjie0203/go-trade-user/model"
	"github.com/yongjie0203/go-universal/crypt"
	"github.com/yongjie0203/go-universal/db"
	"github.com/yongjie0203/go-universal/email"
	"github.com/yongjie0203/go-universal/http"
)

func Register(ctx *gin.Context) {

	userInfo := new(model.UserRegistInfo)
	ctx.ShouldBind(userInfo)
	s, _ := json.Marshal(userInfo)
	fmt.Println(string(s))
	user := new(model.User)
	user.Email = userInfo.Email
	user.Mobile = userInfo.Mobile
	user.AreaCode = userInfo.MobileArea

	user.UID = db.NextId("user")
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Print(err)
	}
	salt := strings.ReplaceAll(uuid.String(), "-", "")
	user.Salt = salt
	user.Time = time.Now().UnixNano()
	user.Pass = crypt.Encode(userInfo.Password)

	db.Coon.Create(user)
	if userInfo.RegistWay == model.REGIST_WAY_EMAIL {
		emails := []string{"360582818@qq.com"}
		go email.SendMail(emails, "新用户激活", "注册成功")
	}

	httpReturn := http.Success("")
	ctx.JSON(200, httpReturn)
}
