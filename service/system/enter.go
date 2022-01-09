package system

type ServiceGroup struct {
	JwtService
	SysService
	UserService
	MenuService
	CasbinService
	InitDBService
	BaseMenuService
	AuthorityService
	OperationRecordService
}
