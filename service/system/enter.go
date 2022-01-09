package system

type ServiceGroup struct {
	ApiService
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
