package wshandler

import (
	"fmt"
	"github.com/yongjie0203/go-universal/websockets"
)

func Register(c *websockets.Client,message *websockets.Message )  {
	fmt.Println("Register request ")
	c.SendMessage([]byte("Register success"))
}