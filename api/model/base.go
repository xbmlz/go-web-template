package model

import (
	"time"

	"github.com/xbmlz/go-web-template/internal/logger"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp;" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;" json:"updatedAt"`
}

type Querier interface {
	// SELECT * FROM @@table where id = @id
	FindByID(id uint) (*gen.T, error)
	// SELECT * FROM @@table where id in @ids
	FindByIDs(ids []uint) (*gen.T, error)

	// UPDATE @@table SET @field = @value where id = @id
	UpdateByID(id uint, field string, value interface{}) (*gen.T, error)

	// DELETE FROM @@table where id = @id
	DeleteByID(id uint) (*gen.T, error)
	// DELETE FROM @@table where id in @ids
	DeleteByIDs(ids []uint) (*gen.T, error)
}

func AllModels() []interface{} {
	return []interface{}{
		&User{},
		&Role{},
		&UserRole{},
		&Menu{},
		&RoleMenu{},
	}
}

func MigrateAndSeed(db *gorm.DB) {
	if err := db.AutoMigrate(AllModels()...); err != nil {
		logger.Errorf("migrate models failed: %v", err)
	}
	// seed data
	Seed(db)
}

func Seed(db *gorm.DB) {
	// seed roles
	db.Where("name = ?", "超级管理员").FirstOrCreate(&Role{Name: "超级管理员", Code: "super_admin", Desc: "Administrator"})
	db.Where("name = ?", "管理员").FirstOrCreate(&Role{Name: "管理员", Code: "admin", Desc: "Administrator"})
	db.Where("name = ?", "普通用户").FirstOrCreate(&Role{Name: "普通用户", Code: "user", Desc: "User"})

	// seed users, check if exists, if not, create
	db.Where("username = ?", "admin").FirstOrCreate(&User{Username: "admin", Password: "admin", Email: "admin@example.com"})
	db.Where("username = ?", "user").FirstOrCreate(&User{Username: "user", Password: "user", Email: "user@example.com"})

	// seed user roles
	db.Where("user_id = ? AND role_id = ?", 1, 2).FirstOrCreate(&UserRole{UserID: 1, RoleID: 2})
	db.Where("user_id = ? AND role_id = ?", 2, 3).FirstOrCreate(&UserRole{UserID: 2, RoleID: 3})

	// seed menus
	db.Where("parent_id = ? AND name = ?", 0, "系统管理").FirstOrCreate(&Menu{BaseModel: BaseModel{ID: 1}, ParentID: 0, Name: "系统管理", Type: 1, Path: "", Icon: "el-icon-setting", Sort: 1, Status: 1})
	db.Where("parent_id = ? AND name = ?", 1, "用户管理").FirstOrCreate(&Menu{BaseModel: BaseModel{ID: 2}, ParentID: 1, Name: "用户管理", Type: 2, Path: "/users", Icon: "el-icon-user", Sort: 1, Status: 1})
	db.Where("parent_id = ? AND name = ?", 1, "角色管理").FirstOrCreate(&Menu{BaseModel: BaseModel{ID: 3}, ParentID: 1, Name: "角色管理", Type: 2, Path: "/roles", Icon: "el-icon-s-custom", Sort: 2, Status: 1})
	db.Where("parent_id = ? AND name = ?", 1, "菜单管理").FirstOrCreate(&Menu{BaseModel: BaseModel{ID: 4}, ParentID: 1, Name: "菜单管理", Type: 2, Path: "/menus", Icon: "el-icon-menu", Sort: 3, Status: 1})

	// seed role menus
	db.Where("role_id = ? AND menu_id = ?", 2, 2).FirstOrCreate(&RoleMenu{RoleID: 2, MenuID: 1})
	db.Where("role_id = ? AND menu_id = ?", 2, 2).FirstOrCreate(&RoleMenu{RoleID: 2, MenuID: 2})
	db.Where("role_id = ? AND menu_id = ?", 2, 3).FirstOrCreate(&RoleMenu{RoleID: 2, MenuID: 3})
	db.Where("role_id = ? AND menu_id = ?", 2, 4).FirstOrCreate(&RoleMenu{RoleID: 2, MenuID: 4})
}
