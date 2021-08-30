package reg

import "github.com/yongjie0203/go-trade-user/wshandler"

func SetupRouter()  {

	group:= New()
	group.On("user_login",wshandler.Login)

}