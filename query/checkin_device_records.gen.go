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

func newCheckinDeviceRecord(db *gorm.DB, opts ...gen.DOOption) checkinDeviceRecord {
	_checkinDeviceRecord := checkinDeviceRecord{}

	_checkinDeviceRecord.checkinDeviceRecordDo.UseDB(db, opts...)
	_checkinDeviceRecord.checkinDeviceRecordDo.UseModel(&model.CheckinDeviceRecord{})

	tableName := _checkinDeviceRecord.checkinDeviceRecordDo.TableName()
	_checkinDeviceRecord.ALL = field.NewAsterisk(tableName)
	_checkinDeviceRecord.ID = field.NewInt(tableName, "id")
	_checkinDeviceRecord.Sn = field.NewString(tableName, "sn")
	_checkinDeviceRecord.Mode = field.NewInt(tableName, "mode")
	_checkinDeviceRecord.Inout = field.NewInt(tableName, "inout")
	_checkinDeviceRecord.Event = field.NewInt(tableName, "event")
	_checkinDeviceRecord.Enrollid = field.NewInt(tableName, "enrollid")
	_checkinDeviceRecord.ReportTime = field.NewTime(tableName, "report_time")
	_checkinDeviceRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_checkinDeviceRecord.UpdatedAt = field.NewTime(tableName, "updated_at")

	_checkinDeviceRecord.fillFieldMap()

	return _checkinDeviceRecord
}

type checkinDeviceRecord struct {
	checkinDeviceRecordDo checkinDeviceRecordDo

	ALL        field.Asterisk
	ID         field.Int
	Sn         field.String
	Mode       field.Int
	Inout      field.Int
	Event      field.Int
	Enrollid   field.Int
	ReportTime field.Time
	CreatedAt  field.Time
	UpdatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (c checkinDeviceRecord) Table(newTableName string) *checkinDeviceRecord {
	c.checkinDeviceRecordDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c checkinDeviceRecord) As(alias string) *checkinDeviceRecord {
	c.checkinDeviceRecordDo.DO = *(c.checkinDeviceRecordDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *checkinDeviceRecord) updateTableName(table string) *checkinDeviceRecord {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt(table, "id")
	c.Sn = field.NewString(table, "sn")
	c.Mode = field.NewInt(table, "mode")
	c.Inout = field.NewInt(table, "inout")
	c.Event = field.NewInt(table, "event")
	c.Enrollid = field.NewInt(table, "enrollid")
	c.ReportTime = field.NewTime(table, "report_time")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *checkinDeviceRecord) WithContext(ctx context.Context) ICheckinDeviceRecordDo {
	return c.checkinDeviceRecordDo.WithContext(ctx)
}

func (c checkinDeviceRecord) TableName() string { return c.checkinDeviceRecordDo.TableName() }

func (c checkinDeviceRecord) Alias() string { return c.checkinDeviceRecordDo.Alias() }

func (c checkinDeviceRecord) Columns(cols ...field.Expr) gen.Columns {
	return c.checkinDeviceRecordDo.Columns(cols...)
}

func (c *checkinDeviceRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *checkinDeviceRecord) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["sn"] = c.Sn
	c.fieldMap["mode"] = c.Mode
	c.fieldMap["inout"] = c.Inout
	c.fieldMap["event"] = c.Event
	c.fieldMap["enrollid"] = c.Enrollid
	c.fieldMap["report_time"] = c.ReportTime
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c checkinDeviceRecord) clone(db *gorm.DB) checkinDeviceRecord {
	c.checkinDeviceRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c checkinDeviceRecord) replaceDB(db *gorm.DB) checkinDeviceRecord {
	c.checkinDeviceRecordDo.ReplaceDB(db)
	return c
}

type checkinDeviceRecordDo struct{ gen.DO }

type ICheckinDeviceRecordDo interface {
	gen.SubQuery
	Debug() ICheckinDeviceRecordDo
	WithContext(ctx context.Context) ICheckinDeviceRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICheckinDeviceRecordDo
	WriteDB() ICheckinDeviceRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICheckinDeviceRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICheckinDeviceRecordDo
	Not(conds ...gen.Condition) ICheckinDeviceRecordDo
	Or(conds ...gen.Condition) ICheckinDeviceRecordDo
	Select(conds ...field.Expr) ICheckinDeviceRecordDo
	Where(conds ...gen.Condition) ICheckinDeviceRecordDo
	Order(conds ...field.Expr) ICheckinDeviceRecordDo
	Distinct(cols ...field.Expr) ICheckinDeviceRecordDo
	Omit(cols ...field.Expr) ICheckinDeviceRecordDo
	Join(table schema.Tabler, on ...field.Expr) ICheckinDeviceRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICheckinDeviceRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICheckinDeviceRecordDo
	Group(cols ...field.Expr) ICheckinDeviceRecordDo
	Having(conds ...gen.Condition) ICheckinDeviceRecordDo
	Limit(limit int) ICheckinDeviceRecordDo
	Offset(offset int) ICheckinDeviceRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICheckinDeviceRecordDo
	Unscoped() ICheckinDeviceRecordDo
	Create(values ...*model.CheckinDeviceRecord) error
	CreateInBatches(values []*model.CheckinDeviceRecord, batchSize int) error
	Save(values ...*model.CheckinDeviceRecord) error
	First() (*model.CheckinDeviceRecord, error)
	Take() (*model.CheckinDeviceRecord, error)
	Last() (*model.CheckinDeviceRecord, error)
	Find() ([]*model.CheckinDeviceRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CheckinDeviceRecord, err error)
	FindInBatches(result *[]*model.CheckinDeviceRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CheckinDeviceRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICheckinDeviceRecordDo
	Assign(attrs ...field.AssignExpr) ICheckinDeviceRecordDo
	Joins(fields ...field.RelationField) ICheckinDeviceRecordDo
	Preload(fields ...field.RelationField) ICheckinDeviceRecordDo
	FirstOrInit() (*model.CheckinDeviceRecord, error)
	FirstOrCreate() (*model.CheckinDeviceRecord, error)
	FindByPage(offset int, limit int) (result []*model.CheckinDeviceRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICheckinDeviceRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c checkinDeviceRecordDo) Debug() ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Debug())
}

func (c checkinDeviceRecordDo) WithContext(ctx context.Context) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c checkinDeviceRecordDo) ReadDB() ICheckinDeviceRecordDo {
	return c.Clauses(dbresolver.Read)
}

func (c checkinDeviceRecordDo) WriteDB() ICheckinDeviceRecordDo {
	return c.Clauses(dbresolver.Write)
}

func (c checkinDeviceRecordDo) Session(config *gorm.Session) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Session(config))
}

func (c checkinDeviceRecordDo) Clauses(conds ...clause.Expression) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c checkinDeviceRecordDo) Returning(value interface{}, columns ...string) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c checkinDeviceRecordDo) Not(conds ...gen.Condition) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c checkinDeviceRecordDo) Or(conds ...gen.Condition) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c checkinDeviceRecordDo) Select(conds ...field.Expr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c checkinDeviceRecordDo) Where(conds ...gen.Condition) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c checkinDeviceRecordDo) Order(conds ...field.Expr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c checkinDeviceRecordDo) Distinct(cols ...field.Expr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c checkinDeviceRecordDo) Omit(cols ...field.Expr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c checkinDeviceRecordDo) Join(table schema.Tabler, on ...field.Expr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c checkinDeviceRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c checkinDeviceRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c checkinDeviceRecordDo) Group(cols ...field.Expr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c checkinDeviceRecordDo) Having(conds ...gen.Condition) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c checkinDeviceRecordDo) Limit(limit int) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c checkinDeviceRecordDo) Offset(offset int) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c checkinDeviceRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c checkinDeviceRecordDo) Unscoped() ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Unscoped())
}

func (c checkinDeviceRecordDo) Create(values ...*model.CheckinDeviceRecord) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c checkinDeviceRecordDo) CreateInBatches(values []*model.CheckinDeviceRecord, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c checkinDeviceRecordDo) Save(values ...*model.CheckinDeviceRecord) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c checkinDeviceRecordDo) First() (*model.CheckinDeviceRecord, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckinDeviceRecord), nil
	}
}

func (c checkinDeviceRecordDo) Take() (*model.CheckinDeviceRecord, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckinDeviceRecord), nil
	}
}

func (c checkinDeviceRecordDo) Last() (*model.CheckinDeviceRecord, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckinDeviceRecord), nil
	}
}

func (c checkinDeviceRecordDo) Find() ([]*model.CheckinDeviceRecord, error) {
	result, err := c.DO.Find()
	return result.([]*model.CheckinDeviceRecord), err
}

func (c checkinDeviceRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CheckinDeviceRecord, err error) {
	buf := make([]*model.CheckinDeviceRecord, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c checkinDeviceRecordDo) FindInBatches(result *[]*model.CheckinDeviceRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c checkinDeviceRecordDo) Attrs(attrs ...field.AssignExpr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c checkinDeviceRecordDo) Assign(attrs ...field.AssignExpr) ICheckinDeviceRecordDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c checkinDeviceRecordDo) Joins(fields ...field.RelationField) ICheckinDeviceRecordDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c checkinDeviceRecordDo) Preload(fields ...field.RelationField) ICheckinDeviceRecordDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c checkinDeviceRecordDo) FirstOrInit() (*model.CheckinDeviceRecord, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckinDeviceRecord), nil
	}
}

func (c checkinDeviceRecordDo) FirstOrCreate() (*model.CheckinDeviceRecord, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckinDeviceRecord), nil
	}
}

func (c checkinDeviceRecordDo) FindByPage(offset int, limit int) (result []*model.CheckinDeviceRecord, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c checkinDeviceRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c checkinDeviceRecordDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c checkinDeviceRecordDo) Delete(models ...*model.CheckinDeviceRecord) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *checkinDeviceRecordDo) withDO(do gen.Dao) *checkinDeviceRecordDo {
	c.DO = *do.(*gen.DO)
	return c
}
