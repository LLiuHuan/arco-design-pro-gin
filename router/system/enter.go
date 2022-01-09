package system

type RouterGroup struct {
	ApiRouter
	JwtRouter
	SysRouter
	BaseRouter
	UserRouter
	MenuRouter
	InitRouter
	CasbinRouter
	AuthorityRouter
}
