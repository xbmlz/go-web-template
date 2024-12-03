// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"github.com/xbmlz/go-web-template/api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newUserRole(db *gorm.DB, opts ...gen.DOOption) userRole {
	_userRole := userRole{}

	_userRole.userRoleDo.UseDB(db, opts...)
	_userRole.userRoleDo.UseModel(&model.UserRole{})

	tableName := _userRole.userRoleDo.TableName()
	_userRole.ALL = field.NewAsterisk(tableName)
	_userRole.ID = field.NewUint(tableName, "id")
	_userRole.CreatedAt = field.NewTime(tableName, "created_at")
	_userRole.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userRole.UserID = field.NewUint(tableName, "user_id")
	_userRole.RoleID = field.NewUint(tableName, "role_id")

	_userRole.fillFieldMap()

	return _userRole
}

type userRole struct {
	userRoleDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	UserID    field.Uint
	RoleID    field.Uint

	fieldMap map[string]field.Expr
}

func (u userRole) Table(newTableName string) *userRole {
	u.userRoleDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userRole) As(alias string) *userRole {
	u.userRoleDo.DO = *(u.userRoleDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userRole) updateTableName(table string) *userRole {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.UserID = field.NewUint(table, "user_id")
	u.RoleID = field.NewUint(table, "role_id")

	u.fillFieldMap()

	return u
}

func (u *userRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userRole) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 5)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["role_id"] = u.RoleID
}

func (u userRole) clone(db *gorm.DB) userRole {
	u.userRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userRole) replaceDB(db *gorm.DB) userRole {
	u.userRoleDo.ReplaceDB(db)
	return u
}

type userRoleDo struct{ gen.DO }

// SELECT * FROM @@table where id = @id
func (u userRoleDo) FindByID(id uint) (result model.UserRole, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM sys_user_role where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table where id in @ids
func (u userRoleDo) FindByIDs(ids []uint) (result model.UserRole, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("SELECT * FROM sys_user_role where id in ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table SET @field = @value where id = @id
func (u userRoleDo) UpdateByID(id uint, field string, value interface{}) (result model.UserRole, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, field)
	params = append(params, value)
	params = append(params, id)
	generateSQL.WriteString("UPDATE sys_user_role SET ? = ? where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table where id = @id
func (u userRoleDo) DeleteByID(id uint) (result model.UserRole, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("DELETE FROM sys_user_role where id = ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DELETE FROM @@table where id in @ids
func (u userRoleDo) DeleteByIDs(ids []uint) (result model.UserRole, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("DELETE FROM sys_user_role where id in ? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (u userRoleDo) Debug() *userRoleDo {
	return u.withDO(u.DO.Debug())
}

func (u userRoleDo) WithContext(ctx context.Context) *userRoleDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userRoleDo) ReadDB() *userRoleDo {
	return u.Clauses(dbresolver.Read)
}

func (u userRoleDo) WriteDB() *userRoleDo {
	return u.Clauses(dbresolver.Write)
}

func (u userRoleDo) Session(config *gorm.Session) *userRoleDo {
	return u.withDO(u.DO.Session(config))
}

func (u userRoleDo) Clauses(conds ...clause.Expression) *userRoleDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userRoleDo) Returning(value interface{}, columns ...string) *userRoleDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userRoleDo) Not(conds ...gen.Condition) *userRoleDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userRoleDo) Or(conds ...gen.Condition) *userRoleDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userRoleDo) Select(conds ...field.Expr) *userRoleDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userRoleDo) Where(conds ...gen.Condition) *userRoleDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userRoleDo) Order(conds ...field.Expr) *userRoleDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userRoleDo) Distinct(cols ...field.Expr) *userRoleDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userRoleDo) Omit(cols ...field.Expr) *userRoleDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userRoleDo) Join(table schema.Tabler, on ...field.Expr) *userRoleDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userRoleDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) *userRoleDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userRoleDo) Group(cols ...field.Expr) *userRoleDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userRoleDo) Having(conds ...gen.Condition) *userRoleDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userRoleDo) Limit(limit int) *userRoleDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userRoleDo) Offset(offset int) *userRoleDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userRoleDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userRoleDo) Unscoped() *userRoleDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userRoleDo) Create(values ...*model.UserRole) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userRoleDo) CreateInBatches(values []*model.UserRole, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userRoleDo) Save(values ...*model.UserRole) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userRoleDo) First() (*model.UserRole, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRole), nil
	}
}

func (u userRoleDo) Take() (*model.UserRole, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRole), nil
	}
}

func (u userRoleDo) Last() (*model.UserRole, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRole), nil
	}
}

func (u userRoleDo) Find() ([]*model.UserRole, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserRole), err
}

func (u userRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserRole, err error) {
	buf := make([]*model.UserRole, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userRoleDo) FindInBatches(result *[]*model.UserRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userRoleDo) Attrs(attrs ...field.AssignExpr) *userRoleDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userRoleDo) Assign(attrs ...field.AssignExpr) *userRoleDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userRoleDo) Joins(fields ...field.RelationField) *userRoleDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userRoleDo) Preload(fields ...field.RelationField) *userRoleDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userRoleDo) FirstOrInit() (*model.UserRole, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRole), nil
	}
}

func (u userRoleDo) FirstOrCreate() (*model.UserRole, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserRole), nil
	}
}

func (u userRoleDo) FindByPage(offset int, limit int) (result []*model.UserRole, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userRoleDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userRoleDo) Delete(models ...*model.UserRole) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userRoleDo) withDO(do gen.Dao) *userRoleDo {
	u.DO = *do.(*gen.DO)
	return u
}