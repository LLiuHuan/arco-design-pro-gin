package system

type ServiceGroup struct {
	JwtService
	UserService
	MenuService
	CasbinService
	InitDBService
	OperationRecordService
}
