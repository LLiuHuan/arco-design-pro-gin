// Package system
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 10:42
package system

import (
	"github.com/lliuhuan/arco-design-pro-gin/global"
	"github.com/lliuhuan/arco-design-pro-gin/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

func (m *menu) TableName() string {
	return "sys_base_menus"
}

func (m *menu) Initialize() error {
	entities := []system.SysBaseMenu{
		{MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Component: "view/dashboard/index.vue", Sort: 1, Meta: system.Meta{Title: "仪表盘", Icon: "setting", Hidden: false}},
		{MenuLevel: 0, ParentId: "0", Path: "about", Name: "about", Component: "view/about/index.vue", Sort: 7, Meta: system.Meta{Title: "关于我们", Icon: "info", Hidden: false}},
		{MenuLevel: 0, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: system.Meta{Title: "超级管理员", Icon: "user-solid", Hidden: false}},
		{MenuLevel: 0, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: system.Meta{Title: "角色管理", Icon: "s-custom", Hidden: false}},
		{MenuLevel: 0, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: system.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true, Hidden: false}},
		{MenuLevel: 0, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: system.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true, Hidden: false}},
		{MenuLevel: 0, ParentId: "3", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 4, Meta: system.Meta{Title: "用户管理", Icon: "coordinate", Hidden: false}},
		{MenuLevel: 0, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 4, Meta: system.Meta{Title: "个人信息", Icon: "message-solid", Hidden: false}},
		{MenuLevel: 0, ParentId: "0", Path: "example", Name: "example", Component: "view/example/index.vue", Sort: 6, Meta: system.Meta{Title: "示例文件", Icon: "s-management", Hidden: false}},
		{MenuLevel: 0, ParentId: "9", Path: "excel", Name: "excel", Component: "view/example/excel/excel.vue", Sort: 4, Meta: system.Meta{Title: "excel导入导出", Icon: "s-marketing", Hidden: false}},
		{MenuLevel: 0, ParentId: "9", Path: "upload", Name: "upload", Component: "view/example/upload/upload.vue", Sort: 5, Meta: system.Meta{Title: "媒体库（上传下载）", Icon: "upload", Hidden: false}},
		{MenuLevel: 0, ParentId: "9", Path: "breakpoint", Name: "breakpoint", Component: "view/example/breakpoint/breakpoint.vue", Sort: 6, Meta: system.Meta{Title: "断点续传", Icon: "upload", Hidden: false}},
		{MenuLevel: 0, ParentId: "9", Path: "customer", Name: "customer", Component: "view/example/customer/customer.vue", Sort: 7, Meta: system.Meta{Title: "客户列表（资源示例）", Icon: "s-custom", Hidden: false}},
		{MenuLevel: 0, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: system.Meta{Title: "系统工具", Icon: "s-cooperation", Hidden: false}},
		{MenuLevel: 0, ParentId: "14", Path: "autoCode", Name: "autoCode", Component: "view/systemTools/autoCode/index.vue", Sort: 1, Meta: system.Meta{Title: "代码生成器", Icon: "cpu", KeepAlive: true, Hidden: false}},
		{MenuLevel: 0, ParentId: "14", Path: "formCreate", Name: "formCreate", Component: "view/systemTools/formCreate/index.vue", Sort: 2, Meta: system.Meta{Title: "表单生成器", Icon: "magic-stick", KeepAlive: true, Hidden: false}},
		{MenuLevel: 0, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: system.Meta{Title: "系统配置", Icon: "s-operation", Hidden: false}},
		{MenuLevel: 0, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: system.Meta{Title: "字典管理", Icon: "notebook-2", Hidden: false}},
		{MenuLevel: 0, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: system.Meta{Title: "字典详情", Icon: "s-order", Hidden: false}},
		{MenuLevel: 0, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: system.Meta{Title: "操作历史", Icon: "time", Hidden: false}},
		{MenuLevel: 0, ParentId: "9", Path: "simpleUploader", Name: "simpleUploader", Component: "view/example/simpleUploader/simpleUploader", Sort: 6, Meta: system.Meta{Title: "断点续传（插件版）", Icon: "upload", Hidden: false}},
		{MenuLevel: 0, ParentId: "0", Path: "https://www.gin-vue-admin.com", Name: "https://www.gin-vue-admin.com", Component: "/", Sort: 0, Meta: system.Meta{Title: "官方网站", Icon: "s-home", Hidden: false}},
		{MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Component: "view/system/state.vue", Sort: 6, Meta: system.Meta{Title: "服务器状态", Icon: "cloudy", Hidden: false}},
		{MenuLevel: 0, ParentId: "14", Path: "autoCodeAdmin", Name: "autoCodeAdmin", Component: "view/systemTools/autoCodeAdmin/index.vue", Sort: 1, Meta: system.Meta{Title: "自动化代码管理", Icon: "s-finance", Hidden: false}},
		{MenuLevel: 0, ParentId: "14", Path: "autoCodeEdit/:id", Name: "autoCodeEdit", Component: "view/systemTools/autoCode/index.vue", Sort: 0, Meta: system.Meta{Title: "自动化代码（复用）", Icon: "s-finance", Hidden: false}},
	}
	if err := global.AdpDb.Create(&entities).Error; err != nil { // 创建 model.User 初始化数据
		return errors.Wrap(err, m.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (m *menu) CheckDataExist() bool {
	if errors.Is(global.AdpDb.Where("path = ?", "autoCodeEdit/:id").First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
