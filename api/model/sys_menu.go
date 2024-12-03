package model

type Menu struct {
	BaseModel
	ParentID uint   `gorm:"not null;comment:父菜单ID" json:"parent_id"`
	Name     string `gorm:"size:32;not null;comment:菜单名称" json:"name"`
	Type     uint   `gorm:"not null;default:1;comment:菜单类型(1:目录,2:菜单,3:按钮)" json:"type"`
	Path     string `gorm:"size:128;not null;comment:菜单路径" json:"path"`
	Icon     string `gorm:"size:32;not null;comment:菜单图标" json:"icon"`
	Sort     uint   `gorm:"not null;comment:菜单排序" json:"sort"`
	Status   uint   `gorm:"not null;default:1;comment:菜单状态(0:禁用,1:启用)" json:"status"`
	Children []Menu `gorm:"-" json:"children,omitempty"`
}

func (Menu) TableName() string {
	return "sys_menu"
}

type RoleMenu struct {
	BaseModel
	RoleID uint `gorm:"not null;comment:角色ID" json:"role_id"`
	MenuID uint `gorm:"not null;comment:菜单ID" json:"menu_id"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}

func BuildMenuTree(menus []*Menu, pid uint) []Menu {
	tree := make([]Menu, 0)
	for _, menu := range menus {
		if menu.ParentID == pid {
			menu.Children = BuildMenuTree(menus, menu.ID)
			tree = append(tree, *menu)
		}
	}
	return tree
}
