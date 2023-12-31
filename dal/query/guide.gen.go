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

func newGuide(db *gorm.DB, opts ...gen.DOOption) guide {
	_guide := guide{}

	_guide.guideDo.UseDB(db, opts...)
	_guide.guideDo.UseModel(&model.Guide{})

	tableName := _guide.guideDo.TableName()
	_guide.ALL = field.NewAsterisk(tableName)
	_guide.ID = field.NewInt64(tableName, "id")
	_guide.Group_ = field.NewString(tableName, "group")
	_guide.Guide = field.NewString(tableName, "guide")
	_guide.DeletedAt = field.NewField(tableName, "deleted_at")
	_guide.CreatedAt = field.NewTime(tableName, "created_at")
	_guide.UpdatedAt = field.NewTime(tableName, "updated_at")

	_guide.fillFieldMap()

	return _guide
}

type guide struct {
	guideDo

	ALL       field.Asterisk
	ID        field.Int64
	Group_    field.String // 组别
	Guide     field.String // 指南
	DeletedAt field.Field
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (g guide) Table(newTableName string) *guide {
	g.guideDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g guide) As(alias string) *guide {
	g.guideDo.DO = *(g.guideDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *guide) updateTableName(table string) *guide {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.Group_ = field.NewString(table, "group")
	g.Guide = field.NewString(table, "guide")
	g.DeletedAt = field.NewField(table, "deleted_at")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")

	g.fillFieldMap()

	return g
}

func (g *guide) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *guide) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 6)
	g.fieldMap["id"] = g.ID
	g.fieldMap["group"] = g.Group_
	g.fieldMap["guide"] = g.Guide
	g.fieldMap["deleted_at"] = g.DeletedAt
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
}

func (g guide) clone(db *gorm.DB) guide {
	g.guideDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g guide) replaceDB(db *gorm.DB) guide {
	g.guideDo.ReplaceDB(db)
	return g
}

type guideDo struct{ gen.DO }

type IGuideDo interface {
	gen.SubQuery
	Debug() IGuideDo
	WithContext(ctx context.Context) IGuideDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGuideDo
	WriteDB() IGuideDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGuideDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGuideDo
	Not(conds ...gen.Condition) IGuideDo
	Or(conds ...gen.Condition) IGuideDo
	Select(conds ...field.Expr) IGuideDo
	Where(conds ...gen.Condition) IGuideDo
	Order(conds ...field.Expr) IGuideDo
	Distinct(cols ...field.Expr) IGuideDo
	Omit(cols ...field.Expr) IGuideDo
	Join(table schema.Tabler, on ...field.Expr) IGuideDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGuideDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGuideDo
	Group(cols ...field.Expr) IGuideDo
	Having(conds ...gen.Condition) IGuideDo
	Limit(limit int) IGuideDo
	Offset(offset int) IGuideDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGuideDo
	Unscoped() IGuideDo
	Create(values ...*model.Guide) error
	CreateInBatches(values []*model.Guide, batchSize int) error
	Save(values ...*model.Guide) error
	First() (*model.Guide, error)
	Take() (*model.Guide, error)
	Last() (*model.Guide, error)
	Find() ([]*model.Guide, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Guide, err error)
	FindInBatches(result *[]*model.Guide, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Guide) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGuideDo
	Assign(attrs ...field.AssignExpr) IGuideDo
	Joins(fields ...field.RelationField) IGuideDo
	Preload(fields ...field.RelationField) IGuideDo
	FirstOrInit() (*model.Guide, error)
	FirstOrCreate() (*model.Guide, error)
	FindByPage(offset int, limit int) (result []*model.Guide, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGuideDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g guideDo) Debug() IGuideDo {
	return g.withDO(g.DO.Debug())
}

func (g guideDo) WithContext(ctx context.Context) IGuideDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g guideDo) ReadDB() IGuideDo {
	return g.Clauses(dbresolver.Read)
}

func (g guideDo) WriteDB() IGuideDo {
	return g.Clauses(dbresolver.Write)
}

func (g guideDo) Session(config *gorm.Session) IGuideDo {
	return g.withDO(g.DO.Session(config))
}

func (g guideDo) Clauses(conds ...clause.Expression) IGuideDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g guideDo) Returning(value interface{}, columns ...string) IGuideDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g guideDo) Not(conds ...gen.Condition) IGuideDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g guideDo) Or(conds ...gen.Condition) IGuideDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g guideDo) Select(conds ...field.Expr) IGuideDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g guideDo) Where(conds ...gen.Condition) IGuideDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g guideDo) Order(conds ...field.Expr) IGuideDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g guideDo) Distinct(cols ...field.Expr) IGuideDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g guideDo) Omit(cols ...field.Expr) IGuideDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g guideDo) Join(table schema.Tabler, on ...field.Expr) IGuideDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g guideDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGuideDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g guideDo) RightJoin(table schema.Tabler, on ...field.Expr) IGuideDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g guideDo) Group(cols ...field.Expr) IGuideDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g guideDo) Having(conds ...gen.Condition) IGuideDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g guideDo) Limit(limit int) IGuideDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g guideDo) Offset(offset int) IGuideDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g guideDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGuideDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g guideDo) Unscoped() IGuideDo {
	return g.withDO(g.DO.Unscoped())
}

func (g guideDo) Create(values ...*model.Guide) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g guideDo) CreateInBatches(values []*model.Guide, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g guideDo) Save(values ...*model.Guide) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g guideDo) First() (*model.Guide, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guide), nil
	}
}

func (g guideDo) Take() (*model.Guide, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guide), nil
	}
}

func (g guideDo) Last() (*model.Guide, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guide), nil
	}
}

func (g guideDo) Find() ([]*model.Guide, error) {
	result, err := g.DO.Find()
	return result.([]*model.Guide), err
}

func (g guideDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Guide, err error) {
	buf := make([]*model.Guide, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g guideDo) FindInBatches(result *[]*model.Guide, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g guideDo) Attrs(attrs ...field.AssignExpr) IGuideDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g guideDo) Assign(attrs ...field.AssignExpr) IGuideDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g guideDo) Joins(fields ...field.RelationField) IGuideDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g guideDo) Preload(fields ...field.RelationField) IGuideDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g guideDo) FirstOrInit() (*model.Guide, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guide), nil
	}
}

func (g guideDo) FirstOrCreate() (*model.Guide, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guide), nil
	}
}

func (g guideDo) FindByPage(offset int, limit int) (result []*model.Guide, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g guideDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g guideDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g guideDo) Delete(models ...*model.Guide) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *guideDo) withDO(do gen.Dao) *guideDo {
	g.DO = *do.(*gen.DO)
	return g
}
