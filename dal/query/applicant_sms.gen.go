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

func newApplicantSm(db *gorm.DB, opts ...gen.DOOption) applicantSm {
	_applicantSm := applicantSm{}

	_applicantSm.applicantSmDo.UseDB(db, opts...)
	_applicantSm.applicantSmDo.UseModel(&model.ApplicantSm{})

	tableName := _applicantSm.applicantSmDo.TableName()
	_applicantSm.ALL = field.NewAsterisk(tableName)
	_applicantSm.ID = field.NewInt64(tableName, "id")
	_applicantSm.ApplicantID = field.NewInt64(tableName, "applicant_id")
	_applicantSm.Typ = field.NewString(tableName, "typ")
	_applicantSm.Status = field.NewInt32(tableName, "status")
	_applicantSm.Args = field.NewString(tableName, "args")
	_applicantSm.Content = field.NewString(tableName, "content")
	_applicantSm.DeletedAt = field.NewField(tableName, "deleted_at")
	_applicantSm.CreatedAt = field.NewTime(tableName, "created_at")
	_applicantSm.UpdatedAt = field.NewTime(tableName, "updated_at")
	_applicantSm.CreatedBy = field.NewInt64(tableName, "created_by")

	_applicantSm.fillFieldMap()

	return _applicantSm
}

type applicantSm struct {
	applicantSmDo

	ALL         field.Asterisk
	ID          field.Int64  // ID
	ApplicantID field.Int64  // 申请者ID
	Typ         field.String // 类型
	Status      field.Int32  // 状态
	Args        field.String // 变量
	Content     field.String // 内容
	DeletedAt   field.Field  // 删除时间
	CreatedAt   field.Time   // 创建时间
	UpdatedAt   field.Time   // 更新时间
	CreatedBy   field.Int64  // 创建者

	fieldMap map[string]field.Expr
}

func (a applicantSm) Table(newTableName string) *applicantSm {
	a.applicantSmDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a applicantSm) As(alias string) *applicantSm {
	a.applicantSmDo.DO = *(a.applicantSmDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *applicantSm) updateTableName(table string) *applicantSm {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.ApplicantID = field.NewInt64(table, "applicant_id")
	a.Typ = field.NewString(table, "typ")
	a.Status = field.NewInt32(table, "status")
	a.Args = field.NewString(table, "args")
	a.Content = field.NewString(table, "content")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.CreatedBy = field.NewInt64(table, "created_by")

	a.fillFieldMap()

	return a
}

func (a *applicantSm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *applicantSm) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["applicant_id"] = a.ApplicantID
	a.fieldMap["typ"] = a.Typ
	a.fieldMap["status"] = a.Status
	a.fieldMap["args"] = a.Args
	a.fieldMap["content"] = a.Content
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["created_by"] = a.CreatedBy
}

func (a applicantSm) clone(db *gorm.DB) applicantSm {
	a.applicantSmDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a applicantSm) replaceDB(db *gorm.DB) applicantSm {
	a.applicantSmDo.ReplaceDB(db)
	return a
}

type applicantSmDo struct{ gen.DO }

type IApplicantSmDo interface {
	gen.SubQuery
	Debug() IApplicantSmDo
	WithContext(ctx context.Context) IApplicantSmDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IApplicantSmDo
	WriteDB() IApplicantSmDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IApplicantSmDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IApplicantSmDo
	Not(conds ...gen.Condition) IApplicantSmDo
	Or(conds ...gen.Condition) IApplicantSmDo
	Select(conds ...field.Expr) IApplicantSmDo
	Where(conds ...gen.Condition) IApplicantSmDo
	Order(conds ...field.Expr) IApplicantSmDo
	Distinct(cols ...field.Expr) IApplicantSmDo
	Omit(cols ...field.Expr) IApplicantSmDo
	Join(table schema.Tabler, on ...field.Expr) IApplicantSmDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IApplicantSmDo
	RightJoin(table schema.Tabler, on ...field.Expr) IApplicantSmDo
	Group(cols ...field.Expr) IApplicantSmDo
	Having(conds ...gen.Condition) IApplicantSmDo
	Limit(limit int) IApplicantSmDo
	Offset(offset int) IApplicantSmDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IApplicantSmDo
	Unscoped() IApplicantSmDo
	Create(values ...*model.ApplicantSm) error
	CreateInBatches(values []*model.ApplicantSm, batchSize int) error
	Save(values ...*model.ApplicantSm) error
	First() (*model.ApplicantSm, error)
	Take() (*model.ApplicantSm, error)
	Last() (*model.ApplicantSm, error)
	Find() ([]*model.ApplicantSm, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ApplicantSm, err error)
	FindInBatches(result *[]*model.ApplicantSm, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ApplicantSm) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IApplicantSmDo
	Assign(attrs ...field.AssignExpr) IApplicantSmDo
	Joins(fields ...field.RelationField) IApplicantSmDo
	Preload(fields ...field.RelationField) IApplicantSmDo
	FirstOrInit() (*model.ApplicantSm, error)
	FirstOrCreate() (*model.ApplicantSm, error)
	FindByPage(offset int, limit int) (result []*model.ApplicantSm, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IApplicantSmDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a applicantSmDo) Debug() IApplicantSmDo {
	return a.withDO(a.DO.Debug())
}

func (a applicantSmDo) WithContext(ctx context.Context) IApplicantSmDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a applicantSmDo) ReadDB() IApplicantSmDo {
	return a.Clauses(dbresolver.Read)
}

func (a applicantSmDo) WriteDB() IApplicantSmDo {
	return a.Clauses(dbresolver.Write)
}

func (a applicantSmDo) Session(config *gorm.Session) IApplicantSmDo {
	return a.withDO(a.DO.Session(config))
}

func (a applicantSmDo) Clauses(conds ...clause.Expression) IApplicantSmDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a applicantSmDo) Returning(value interface{}, columns ...string) IApplicantSmDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a applicantSmDo) Not(conds ...gen.Condition) IApplicantSmDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a applicantSmDo) Or(conds ...gen.Condition) IApplicantSmDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a applicantSmDo) Select(conds ...field.Expr) IApplicantSmDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a applicantSmDo) Where(conds ...gen.Condition) IApplicantSmDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a applicantSmDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IApplicantSmDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a applicantSmDo) Order(conds ...field.Expr) IApplicantSmDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a applicantSmDo) Distinct(cols ...field.Expr) IApplicantSmDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a applicantSmDo) Omit(cols ...field.Expr) IApplicantSmDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a applicantSmDo) Join(table schema.Tabler, on ...field.Expr) IApplicantSmDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a applicantSmDo) LeftJoin(table schema.Tabler, on ...field.Expr) IApplicantSmDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a applicantSmDo) RightJoin(table schema.Tabler, on ...field.Expr) IApplicantSmDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a applicantSmDo) Group(cols ...field.Expr) IApplicantSmDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a applicantSmDo) Having(conds ...gen.Condition) IApplicantSmDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a applicantSmDo) Limit(limit int) IApplicantSmDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a applicantSmDo) Offset(offset int) IApplicantSmDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a applicantSmDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IApplicantSmDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a applicantSmDo) Unscoped() IApplicantSmDo {
	return a.withDO(a.DO.Unscoped())
}

func (a applicantSmDo) Create(values ...*model.ApplicantSm) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a applicantSmDo) CreateInBatches(values []*model.ApplicantSm, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a applicantSmDo) Save(values ...*model.ApplicantSm) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a applicantSmDo) First() (*model.ApplicantSm, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplicantSm), nil
	}
}

func (a applicantSmDo) Take() (*model.ApplicantSm, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplicantSm), nil
	}
}

func (a applicantSmDo) Last() (*model.ApplicantSm, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplicantSm), nil
	}
}

func (a applicantSmDo) Find() ([]*model.ApplicantSm, error) {
	result, err := a.DO.Find()
	return result.([]*model.ApplicantSm), err
}

func (a applicantSmDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ApplicantSm, err error) {
	buf := make([]*model.ApplicantSm, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a applicantSmDo) FindInBatches(result *[]*model.ApplicantSm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a applicantSmDo) Attrs(attrs ...field.AssignExpr) IApplicantSmDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a applicantSmDo) Assign(attrs ...field.AssignExpr) IApplicantSmDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a applicantSmDo) Joins(fields ...field.RelationField) IApplicantSmDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a applicantSmDo) Preload(fields ...field.RelationField) IApplicantSmDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a applicantSmDo) FirstOrInit() (*model.ApplicantSm, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplicantSm), nil
	}
}

func (a applicantSmDo) FirstOrCreate() (*model.ApplicantSm, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplicantSm), nil
	}
}

func (a applicantSmDo) FindByPage(offset int, limit int) (result []*model.ApplicantSm, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a applicantSmDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a applicantSmDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a applicantSmDo) Delete(models ...*model.ApplicantSm) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *applicantSmDo) withDO(do gen.Dao) *applicantSmDo {
	a.DO = *do.(*gen.DO)
	return a
}
