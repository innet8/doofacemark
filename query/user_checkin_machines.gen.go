// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"checkin/query/model"
)

func newUserCheckinMachine(db *gorm.DB, opts ...gen.DOOption) userCheckinMachine {
	_userCheckinMachine := userCheckinMachine{}

	_userCheckinMachine.userCheckinMachineDo.UseDB(db, opts...)
	_userCheckinMachine.userCheckinMachineDo.UseModel(&model.UserCheckinMachine{})

	tableName := _userCheckinMachine.userCheckinMachineDo.TableName()
	_userCheckinMachine.ALL = field.NewAsterisk(tableName)
	_userCheckinMachine.ID = field.NewInt(tableName, "id")
	_userCheckinMachine.Sn = field.NewString(tableName, "sn")
	_userCheckinMachine.Devinfo = field.NewString(tableName, "devinfo")
	_userCheckinMachine.CreatedAt = field.NewTime(tableName, "created_at")
	_userCheckinMachine.UpdatedAt = field.NewTime(tableName, "updated_at")

	_userCheckinMachine.fillFieldMap()

	return _userCheckinMachine
}

type userCheckinMachine struct {
	userCheckinMachineDo userCheckinMachineDo

	ALL       field.Asterisk
	ID        field.Int
	Sn        field.String
	Devinfo   field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (u userCheckinMachine) Table(newTableName string) *userCheckinMachine {
	u.userCheckinMachineDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userCheckinMachine) As(alias string) *userCheckinMachine {
	u.userCheckinMachineDo.DO = *(u.userCheckinMachineDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userCheckinMachine) updateTableName(table string) *userCheckinMachine {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt(table, "id")
	u.Sn = field.NewString(table, "sn")
	u.Devinfo = field.NewString(table, "devinfo")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")

	u.fillFieldMap()

	return u
}

func (u *userCheckinMachine) WithContext(ctx context.Context) IUserCheckinMachineDo {
	return u.userCheckinMachineDo.WithContext(ctx)
}

func (u userCheckinMachine) TableName() string { return u.userCheckinMachineDo.TableName() }

func (u userCheckinMachine) Alias() string { return u.userCheckinMachineDo.Alias() }

func (u userCheckinMachine) Columns(cols ...field.Expr) gen.Columns {
	return u.userCheckinMachineDo.Columns(cols...)
}

func (u *userCheckinMachine) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userCheckinMachine) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 5)
	u.fieldMap["id"] = u.ID
	u.fieldMap["sn"] = u.Sn
	u.fieldMap["devinfo"] = u.Devinfo
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
}

func (u userCheckinMachine) clone(db *gorm.DB) userCheckinMachine {
	u.userCheckinMachineDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userCheckinMachine) replaceDB(db *gorm.DB) userCheckinMachine {
	u.userCheckinMachineDo.ReplaceDB(db)
	return u
}

type userCheckinMachineDo struct{ gen.DO }

type IUserCheckinMachineDo interface {
	gen.SubQuery
	Debug() IUserCheckinMachineDo
	WithContext(ctx context.Context) IUserCheckinMachineDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserCheckinMachineDo
	WriteDB() IUserCheckinMachineDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserCheckinMachineDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserCheckinMachineDo
	Not(conds ...gen.Condition) IUserCheckinMachineDo
	Or(conds ...gen.Condition) IUserCheckinMachineDo
	Select(conds ...field.Expr) IUserCheckinMachineDo
	Where(conds ...gen.Condition) IUserCheckinMachineDo
	Order(conds ...field.Expr) IUserCheckinMachineDo
	Distinct(cols ...field.Expr) IUserCheckinMachineDo
	Omit(cols ...field.Expr) IUserCheckinMachineDo
	Join(table schema.Tabler, on ...field.Expr) IUserCheckinMachineDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserCheckinMachineDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserCheckinMachineDo
	Group(cols ...field.Expr) IUserCheckinMachineDo
	Having(conds ...gen.Condition) IUserCheckinMachineDo
	Limit(limit int) IUserCheckinMachineDo
	Offset(offset int) IUserCheckinMachineDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserCheckinMachineDo
	Unscoped() IUserCheckinMachineDo
	Create(values ...*model.UserCheckinMachine) error
	CreateInBatches(values []*model.UserCheckinMachine, batchSize int) error
	Save(values ...*model.UserCheckinMachine) error
	First() (*model.UserCheckinMachine, error)
	Take() (*model.UserCheckinMachine, error)
	Last() (*model.UserCheckinMachine, error)
	Find() ([]*model.UserCheckinMachine, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserCheckinMachine, err error)
	FindInBatches(result *[]*model.UserCheckinMachine, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserCheckinMachine) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserCheckinMachineDo
	Assign(attrs ...field.AssignExpr) IUserCheckinMachineDo
	Joins(fields ...field.RelationField) IUserCheckinMachineDo
	Preload(fields ...field.RelationField) IUserCheckinMachineDo
	FirstOrInit() (*model.UserCheckinMachine, error)
	FirstOrCreate() (*model.UserCheckinMachine, error)
	FindByPage(offset int, limit int) (result []*model.UserCheckinMachine, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserCheckinMachineDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userCheckinMachineDo) Debug() IUserCheckinMachineDo {
	return u.withDO(u.DO.Debug())
}

func (u userCheckinMachineDo) WithContext(ctx context.Context) IUserCheckinMachineDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userCheckinMachineDo) ReadDB() IUserCheckinMachineDo {
	return u.Clauses(dbresolver.Read)
}

func (u userCheckinMachineDo) WriteDB() IUserCheckinMachineDo {
	return u.Clauses(dbresolver.Write)
}

func (u userCheckinMachineDo) Session(config *gorm.Session) IUserCheckinMachineDo {
	return u.withDO(u.DO.Session(config))
}

func (u userCheckinMachineDo) Clauses(conds ...clause.Expression) IUserCheckinMachineDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userCheckinMachineDo) Returning(value interface{}, columns ...string) IUserCheckinMachineDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userCheckinMachineDo) Not(conds ...gen.Condition) IUserCheckinMachineDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userCheckinMachineDo) Or(conds ...gen.Condition) IUserCheckinMachineDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userCheckinMachineDo) Select(conds ...field.Expr) IUserCheckinMachineDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userCheckinMachineDo) Where(conds ...gen.Condition) IUserCheckinMachineDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userCheckinMachineDo) Order(conds ...field.Expr) IUserCheckinMachineDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userCheckinMachineDo) Distinct(cols ...field.Expr) IUserCheckinMachineDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userCheckinMachineDo) Omit(cols ...field.Expr) IUserCheckinMachineDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userCheckinMachineDo) Join(table schema.Tabler, on ...field.Expr) IUserCheckinMachineDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userCheckinMachineDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserCheckinMachineDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userCheckinMachineDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserCheckinMachineDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userCheckinMachineDo) Group(cols ...field.Expr) IUserCheckinMachineDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userCheckinMachineDo) Having(conds ...gen.Condition) IUserCheckinMachineDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userCheckinMachineDo) Limit(limit int) IUserCheckinMachineDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userCheckinMachineDo) Offset(offset int) IUserCheckinMachineDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userCheckinMachineDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserCheckinMachineDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userCheckinMachineDo) Unscoped() IUserCheckinMachineDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userCheckinMachineDo) Create(values ...*model.UserCheckinMachine) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userCheckinMachineDo) CreateInBatches(values []*model.UserCheckinMachine, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userCheckinMachineDo) Save(values ...*model.UserCheckinMachine) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userCheckinMachineDo) First() (*model.UserCheckinMachine, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCheckinMachine), nil
	}
}

func (u userCheckinMachineDo) Take() (*model.UserCheckinMachine, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCheckinMachine), nil
	}
}

func (u userCheckinMachineDo) Last() (*model.UserCheckinMachine, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCheckinMachine), nil
	}
}

func (u userCheckinMachineDo) Find() ([]*model.UserCheckinMachine, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserCheckinMachine), err
}

func (u userCheckinMachineDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserCheckinMachine, err error) {
	buf := make([]*model.UserCheckinMachine, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userCheckinMachineDo) FindInBatches(result *[]*model.UserCheckinMachine, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userCheckinMachineDo) Attrs(attrs ...field.AssignExpr) IUserCheckinMachineDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userCheckinMachineDo) Assign(attrs ...field.AssignExpr) IUserCheckinMachineDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userCheckinMachineDo) Joins(fields ...field.RelationField) IUserCheckinMachineDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userCheckinMachineDo) Preload(fields ...field.RelationField) IUserCheckinMachineDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userCheckinMachineDo) FirstOrInit() (*model.UserCheckinMachine, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCheckinMachine), nil
	}
}

func (u userCheckinMachineDo) FirstOrCreate() (*model.UserCheckinMachine, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCheckinMachine), nil
	}
}

func (u userCheckinMachineDo) FindByPage(offset int, limit int) (result []*model.UserCheckinMachine, count int64, err error) {
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

func (u userCheckinMachineDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userCheckinMachineDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userCheckinMachineDo) Delete(models ...*model.UserCheckinMachine) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userCheckinMachineDo) withDO(do gen.Dao) *userCheckinMachineDo {
	u.DO = *do.(*gen.DO)
	return u
}