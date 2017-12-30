package demo

import (

    "gitee.com/johng/gf/g/frame/gmvc"
    "gitee.com/johng/gf/g/net/ghttp"
)

// 定义业务相关的控制器对象，
// 建议命名规范中控制器统一使用Controller前缀，后期代码维护时便于区分
type ControllerUser struct {
    gmvc.Controller
}

// 初始化控制器对象，并绑定操作到Web Server
func init() {
    // 绑定控制器到指定URI，所有控制器的公开方法将会映射到指定URI末尾
    // 例如该方法执行后，查看效果可访问：http://127.0.0.1:8199/user/info
    ghttp.GetServer().BindController("/user", &ControllerUser{})
}

// 定义操作逻辑 - 展示模板
func (c *ControllerUser) Info() {
    c.View.Assign("name", "john")
    c.View.Display("user/index")
}



