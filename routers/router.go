package routers

import (
	"github.com/astaxie/beego"
	"go-cms/controllers/sys"
)

func init() {
	//ns := beego.NewNamespace("/api",
	//	beego.NSNamespace("/user",
	//		beego.NSInclude(
	//			&sys.UserController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)

	//用户相关
	beego.Router("/api/user/login", &sys.UserController{}, "post:Login")
	beego.Router("/api/user/create", &sys.UserController{}, "post:Create")
	beego.Router("/api/user/info", &sys.UserController{}, "get:UserInfo") // 获取用户消息
	beego.Router("/api/user/list", &sys.UserController{}, "get:UserList") // 获取用户列表
	beego.Router("/api/user/check_token", &sys.UserController{}, "post:CheckToken")
	beego.Router("/api/user/logout", &sys.UserController{}, "get:Logout")
	
	//验证码校验
	beego.Router("/api/captcha/check", &sys.CaptchaController{}, "post:Hander")
	
	//参数设置
	beego.Router("/api/configs/index", &sys.ConfigsController{}, "post:Index")
	beego.Router("/api/configs/create", &sys.ConfigsController{}, "post:Create")
	beego.Router("/api/configs/update", &sys.ConfigsController{}, "post:Update")
	beego.Router("/api/configs/delete", &sys.ConfigsController{}, "post:Delete")
	
	//岗位管理
	beego.Router("/api/post/index", &sys.PostController{}, "post:Index")
	beego.Router("/api/post/create", &sys.PostController{}, "post:Create")
	beego.Router("/api/post/update", &sys.PostController{}, "post:Update")
	beego.Router("/api/post/delete", &sys.PostController{}, "post:Delete")
	
	//菜单管理
	beego.Router("/api/menu/index", &sys.MenuController{}, "post:Index")
	beego.Router("/api/menu/create", &sys.MenuController{}, "post:Create")
	beego.Router("/api/menu/update", &sys.MenuController{}, "post:Update")
	beego.Router("/api/menu/delete", &sys.MenuController{}, "post:Delete")
}
