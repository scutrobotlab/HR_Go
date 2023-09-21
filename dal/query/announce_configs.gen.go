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

func newAnnounceConfig(db *gorm.DB, opts ...gen.DOOption) announceConfig {
	_announceConfig := announceConfig{}

	_announceConfig.announceConfigDo.UseDB(db, opts...)
	_announceConfig.announceConfigDo.UseModel(&model.AnnounceConfig{})

	tableName := _announceConfig.announceConfigDo.TableName()
	_announceConfig.ALL = field.NewAsterisk(tableName)
	_announceConfig.ID = field.NewInt64(tableName, "id")
	_announceConfig.Status = field.NewString(tableName, "status")
	_announceConfig.Body = field.NewString(tableName, "body")

	_announceConfig.fillFieldMap()

	return _announceConfig
}

type announceConfig struct {
	announceConfigDo

	ALL    field.Asterisk
	ID     field.Int64  // ID
	Status field.String // 状态名称
	Body   field.String // 状态内容

	fieldMap map[string]field.Expr
}

func (a announceConfig) Table(newTableName string) *announceConfig {
	a.announceConfigDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a announceConfig) As(alias string) *announceConfig {
	a.announceConfigDo.DO = *(a.announceConfigDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *announceConfig) updateTableName(table string) *announceConfig {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Status = field.NewString(table, "status")
	a.Body = field.NewString(table, "body")

	a.fillFieldMap()

	return a
}

func (a *announceConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *announceConfig) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 3)
	a.fieldMap["id"] = a.ID
	a.fieldMap["status"] = a.Status
	a.fieldMap["body"] = a.Body
}

func (a announceConfig) clone(db *gorm.DB) announceConfig {
	a.announceConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a announceConfig) replaceDB(db *gorm.DB) announceConfig {
	a.announceConfigDo.ReplaceDB(db)
	return a
}

type announceConfigDo struct{ gen.DO }

type IAnnounceConfigDo interface {
	gen.SubQuery
	Debug() IAnnounceConfigDo
	WithContext(ctx context.Context) IAnnounceConfigDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAnnounceConfigDo
	WriteDB() IAnnounceConfigDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAnnounceConfigDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAnnounceConfigDo
	Not(conds ...gen.Condition) IAnnounceConfigDo
	Or(conds ...gen.Condition) IAnnounceConfigDo
	Select(conds ...field.Expr) IAnnounceConfigDo
	Where(conds ...gen.Condition) IAnnounceConfigDo
	Order(conds ...field.Expr) IAnnounceConfigDo
	Distinct(cols ...field.Expr) IAnnounceConfigDo
	Omit(cols ...field.Expr) IAnnounceConfigDo
	Join(table schema.Tabler, on ...field.Expr) IAnnounceConfigDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAnnounceConfigDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAnnounceConfigDo
	Group(cols ...field.Expr) IAnnounceConfigDo
	Having(conds ...gen.Condition) IAnnounceConfigDo
	Limit(limit int) IAnnounceConfigDo
	Offset(offset int) IAnnounceConfigDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAnnounceConfigDo
	Unscoped() IAnnounceConfigDo
	Create(values ...*model.AnnounceConfig) error
	CreateInBatches(values []*model.AnnounceConfig, batchSize int) error
	Save(values ...*model.AnnounceConfig) error
	First() (*model.AnnounceConfig, error)
	Take() (*model.AnnounceConfig, error)
	Last() (*model.AnnounceConfig, error)
	Find() ([]*model.AnnounceConfig, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AnnounceConfig, err error)
	FindInBatches(result *[]*model.AnnounceConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AnnounceConfig) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAnnounceConfigDo
	Assign(attrs ...field.AssignExpr) IAnnounceConfigDo
	Joins(fields ...field.RelationField) IAnnounceConfigDo
	Preload(fields ...field.RelationField) IAnnounceConfigDo
	FirstOrInit() (*model.AnnounceConfig, error)
	FirstOrCreate() (*model.AnnounceConfig, error)
	FindByPage(offset int, limit int) (result []*model.AnnounceConfig, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAnnounceConfigDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a announceConfigDo) Debug() IAnnounceConfigDo {
	return a.withDO(a.DO.Debug())
}

func (a announceConfigDo) WithContext(ctx context.Context) IAnnounceConfigDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a announceConfigDo) ReadDB() IAnnounceConfigDo {
	return a.Clauses(dbresolver.Read)
}

func (a announceConfigDo) WriteDB() IAnnounceConfigDo {
	return a.Clauses(dbresolver.Write)
}

func (a announceConfigDo) Session(config *gorm.Session) IAnnounceConfigDo {
	return a.withDO(a.DO.Session(config))
}

func (a announceConfigDo) Clauses(conds ...clause.Expression) IAnnounceConfigDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a announceConfigDo) Returning(value interface{}, columns ...string) IAnnounceConfigDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a announceConfigDo) Not(conds ...gen.Condition) IAnnounceConfigDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a announceConfigDo) Or(conds ...gen.Condition) IAnnounceConfigDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a announceConfigDo) Select(conds ...field.Expr) IAnnounceConfigDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a announceConfigDo) Where(conds ...gen.Condition) IAnnounceConfigDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a announceConfigDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAnnounceConfigDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a announceConfigDo) Order(conds ...field.Expr) IAnnounceConfigDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a announceConfigDo) Distinct(cols ...field.Expr) IAnnounceConfigDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a announceConfigDo) Omit(cols ...field.Expr) IAnnounceConfigDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a announceConfigDo) Join(table schema.Tabler, on ...field.Expr) IAnnounceConfigDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a announceConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAnnounceConfigDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a announceConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) IAnnounceConfigDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a announceConfigDo) Group(cols ...field.Expr) IAnnounceConfigDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a announceConfigDo) Having(conds ...gen.Condition) IAnnounceConfigDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a announceConfigDo) Limit(limit int) IAnnounceConfigDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a announceConfigDo) Offset(offset int) IAnnounceConfigDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a announceConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAnnounceConfigDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a announceConfigDo) Unscoped() IAnnounceConfigDo {
	return a.withDO(a.DO.Unscoped())
}

func (a announceConfigDo) Create(values ...*model.AnnounceConfig) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a announceConfigDo) CreateInBatches(values []*model.AnnounceConfig, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a announceConfigDo) Save(values ...*model.AnnounceConfig) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a announceConfigDo) First() (*model.AnnounceConfig, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnounceConfig), nil
	}
}

func (a announceConfigDo) Take() (*model.AnnounceConfig, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnounceConfig), nil
	}
}

func (a announceConfigDo) Last() (*model.AnnounceConfig, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnounceConfig), nil
	}
}

func (a announceConfigDo) Find() ([]*model.AnnounceConfig, error) {
	result, err := a.DO.Find()
	return result.([]*model.AnnounceConfig), err
}

func (a announceConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AnnounceConfig, err error) {
	buf := make([]*model.AnnounceConfig, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a announceConfigDo) FindInBatches(result *[]*model.AnnounceConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a announceConfigDo) Attrs(attrs ...field.AssignExpr) IAnnounceConfigDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a announceConfigDo) Assign(attrs ...field.AssignExpr) IAnnounceConfigDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a announceConfigDo) Joins(fields ...field.RelationField) IAnnounceConfigDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a announceConfigDo) Preload(fields ...field.RelationField) IAnnounceConfigDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a announceConfigDo) FirstOrInit() (*model.AnnounceConfig, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnounceConfig), nil
	}
}

func (a announceConfigDo) FirstOrCreate() (*model.AnnounceConfig, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AnnounceConfig), nil
	}
}

func (a announceConfigDo) FindByPage(offset int, limit int) (result []*model.AnnounceConfig, count int64, err error) {
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

func (a announceConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a announceConfigDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a announceConfigDo) Delete(models ...*model.AnnounceConfig) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *announceConfigDo) withDO(do gen.Dao) *announceConfigDo {
	a.DO = *do.(*gen.DO)
	return a
}
