package wshandler

import (
	"fmt"
	"github.com/yongjie0203/go-universal/websockets"

)

func Login(c *websockets.Client,message *websockets.Message )  {
	fmt.Println("login request ")
	c.SendMessage([]byte("login success"))
}


