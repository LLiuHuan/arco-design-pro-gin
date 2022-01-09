package system

type RouterGroup struct {
	JwtRouter
	SysRouter
	BaseRouter
	UserRouter
	MenuRouter
	InitRouter
	CasbinRouter
	AuthorityRouter
}
