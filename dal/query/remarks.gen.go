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

	"HR_Go/dal/model"
)

func newRemark(db *gorm.DB, opts ...gen.DOOption) remark {
	_remark := remark{}

	_remark.remarkDo.UseDB(db, opts...)
	_remark.remarkDo.UseModel(&model.Remark{})

	tableName := _remark.remarkDo.TableName()
	_remark.ALL = field.NewAsterisk(tableName)
	_remark.ID = field.NewInt64(tableName, "id")
	_remark.AdminID = field.NewInt64(tableName, "admin_id")
	_remark.ApplicantID = field.NewInt64(tableName, "applicant_id")
	_remark.Remark = field.NewString(tableName, "remark")
	_remark.DeletedAt = field.NewField(tableName, "deleted_at")
	_remark.CreatedAt = field.NewTime(tableName, "created_at")
	_remark.UpdatedAt = field.NewTime(tableName, "updated_at")

	_remark.fillFieldMap()

	return _remark
}

type remark struct {
	remarkDo

	ALL         field.Asterisk
	ID          field.Int64
	AdminID     field.Int64  // 管理员ID
	ApplicantID field.Int64  // 申请ID
	Remark      field.String // 评价
	DeletedAt   field.Field
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (r remark) Table(newTableName string) *remark {
	r.remarkDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r remark) As(alias string) *remark {
	r.remarkDo.DO = *(r.remarkDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *remark) updateTableName(table string) *remark {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.AdminID = field.NewInt64(table, "admin_id")
	r.ApplicantID = field.NewInt64(table, "applicant_id")
	r.Remark = field.NewString(table, "remark")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *remark) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *remark) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 7)
	r.fieldMap["id"] = r.ID
	r.fieldMap["admin_id"] = r.AdminID
	r.fieldMap["applicant_id"] = r.ApplicantID
	r.fieldMap["remark"] = r.Remark
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r remark) clone(db *gorm.DB) remark {
	r.remarkDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r remark) replaceDB(db *gorm.DB) remark {
	r.remarkDo.ReplaceDB(db)
	return r
}

type remarkDo struct{ gen.DO }

type IRemarkDo interface {
	gen.SubQuery
	Debug() IRemarkDo
	WithContext(ctx context.Context) IRemarkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRemarkDo
	WriteDB() IRemarkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRemarkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRemarkDo
	Not(conds ...gen.Condition) IRemarkDo
	Or(conds ...gen.Condition) IRemarkDo
	Select(conds ...field.Expr) IRemarkDo
	Where(conds ...gen.Condition) IRemarkDo
	Order(conds ...field.Expr) IRemarkDo
	Distinct(cols ...field.Expr) IRemarkDo
	Omit(cols ...field.Expr) IRemarkDo
	Join(table schema.Tabler, on ...field.Expr) IRemarkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRemarkDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRemarkDo
	Group(cols ...field.Expr) IRemarkDo
	Having(conds ...gen.Condition) IRemarkDo
	Limit(limit int) IRemarkDo
	Offset(offset int) IRemarkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRemarkDo
	Unscoped() IRemarkDo
	Create(values ...*model.Remark) error
	CreateInBatches(values []*model.Remark, batchSize int) error
	Save(values ...*model.Remark) error
	First() (*model.Remark, error)
	Take() (*model.Remark, error)
	Last() (*model.Remark, error)
	Find() ([]*model.Remark, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Remark, err error)
	FindInBatches(result *[]*model.Remark, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Remark) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRemarkDo
	Assign(attrs ...field.AssignExpr) IRemarkDo
	Joins(fields ...field.RelationField) IRemarkDo
	Preload(fields ...field.RelationField) IRemarkDo
	FirstOrInit() (*model.Remark, error)
	FirstOrCreate() (*model.Remark, error)
	FindByPage(offset int, limit int) (result []*model.Remark, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRemarkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r remarkDo) Debug() IRemarkDo {
	return r.withDO(r.DO.Debug())
}

func (r remarkDo) WithContext(ctx context.Context) IRemarkDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r remarkDo) ReadDB() IRemarkDo {
	return r.Clauses(dbresolver.Read)
}

func (r remarkDo) WriteDB() IRemarkDo {
	return r.Clauses(dbresolver.Write)
}

func (r remarkDo) Session(config *gorm.Session) IRemarkDo {
	return r.withDO(r.DO.Session(config))
}

func (r remarkDo) Clauses(conds ...clause.Expression) IRemarkDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r remarkDo) Returning(value interface{}, columns ...string) IRemarkDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r remarkDo) Not(conds ...gen.Condition) IRemarkDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r remarkDo) Or(conds ...gen.Condition) IRemarkDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r remarkDo) Select(conds ...field.Expr) IRemarkDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r remarkDo) Where(conds ...gen.Condition) IRemarkDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r remarkDo) Order(conds ...field.Expr) IRemarkDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r remarkDo) Distinct(cols ...field.Expr) IRemarkDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r remarkDo) Omit(cols ...field.Expr) IRemarkDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r remarkDo) Join(table schema.Tabler, on ...field.Expr) IRemarkDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r remarkDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRemarkDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r remarkDo) RightJoin(table schema.Tabler, on ...field.Expr) IRemarkDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r remarkDo) Group(cols ...field.Expr) IRemarkDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r remarkDo) Having(conds ...gen.Condition) IRemarkDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r remarkDo) Limit(limit int) IRemarkDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r remarkDo) Offset(offset int) IRemarkDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r remarkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRemarkDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r remarkDo) Unscoped() IRemarkDo {
	return r.withDO(r.DO.Unscoped())
}

func (r remarkDo) Create(values ...*model.Remark) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r remarkDo) CreateInBatches(values []*model.Remark, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r remarkDo) Save(values ...*model.Remark) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r remarkDo) First() (*model.Remark, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) Take() (*model.Remark, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) Last() (*model.Remark, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) Find() ([]*model.Remark, error) {
	result, err := r.DO.Find()
	return result.([]*model.Remark), err
}

func (r remarkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Remark, err error) {
	buf := make([]*model.Remark, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r remarkDo) FindInBatches(result *[]*model.Remark, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r remarkDo) Attrs(attrs ...field.AssignExpr) IRemarkDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r remarkDo) Assign(attrs ...field.AssignExpr) IRemarkDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r remarkDo) Joins(fields ...field.RelationField) IRemarkDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r remarkDo) Preload(fields ...field.RelationField) IRemarkDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r remarkDo) FirstOrInit() (*model.Remark, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) FirstOrCreate() (*model.Remark, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) FindByPage(offset int, limit int) (result []*model.Remark, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r remarkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r remarkDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r remarkDo) Delete(models ...*model.Remark) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *remarkDo) withDO(do gen.Dao) *remarkDo {
	r.DO = *do.(*gen.DO)
	return r
}
