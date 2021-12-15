package system

type RouterGroup struct {
	JwtRouter
	BaseRouter
	UserRouter
	MenuRouter
	InitRouter
	CasbinRouter
}
