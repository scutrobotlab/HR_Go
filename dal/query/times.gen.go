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

func newTime(db *gorm.DB, opts ...gen.DOOption) time {
	_time := time{}

	_time.timeDo.UseDB(db, opts...)
	_time.timeDo.UseModel(&model.Time{})

	tableName := _time.timeDo.TableName()
	_time.ALL = field.NewAsterisk(tableName)
	_time.ID = field.NewInt64(tableName, "id")
	_time.Group_ = field.NewString(tableName, "group")
	_time.Time = field.NewTime(tableName, "time")
	_time.Rank = field.NewInt32(tableName, "rank")
	_time.Location = field.NewString(tableName, "location")
	_time.TotalCnt = field.NewInt32(tableName, "total_cnt")
	_time.Campus = field.NewString(tableName, "campus")
	_time.Grade = field.NewString(tableName, "grade")
	_time.MeetingID = field.NewString(tableName, "meeting_id")
	_time.DeletedAt = field.NewField(tableName, "deleted_at")
	_time.CreatedAt = field.NewTime(tableName, "created_at")
	_time.UpdatedAt = field.NewTime(tableName, "updated_at")

	_time.fillFieldMap()

	return _time
}

type time struct {
	timeDo

	ALL       field.Asterisk
	ID        field.Int64  // ID
	Group_    field.String // 组别
	Time      field.Time   // 时间
	Rank      field.Int32  // 顺序
	Location  field.String // 地点
	TotalCnt  field.Int32  // 总数
	Campus    field.String // 校区
	Grade     field.String // 年级
	MeetingID field.String // 会议ID
	DeletedAt field.Field  // 删除时间
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (t time) Table(newTableName string) *time {
	t.timeDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t time) As(alias string) *time {
	t.timeDo.DO = *(t.timeDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *time) updateTableName(table string) *time {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Group_ = field.NewString(table, "group")
	t.Time = field.NewTime(table, "time")
	t.Rank = field.NewInt32(table, "rank")
	t.Location = field.NewString(table, "location")
	t.TotalCnt = field.NewInt32(table, "total_cnt")
	t.Campus = field.NewString(table, "campus")
	t.Grade = field.NewString(table, "grade")
	t.MeetingID = field.NewString(table, "meeting_id")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")

	t.fillFieldMap()

	return t
}

func (t *time) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *time) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["id"] = t.ID
	t.fieldMap["group"] = t.Group_
	t.fieldMap["time"] = t.Time
	t.fieldMap["rank"] = t.Rank
	t.fieldMap["location"] = t.Location
	t.fieldMap["total_cnt"] = t.TotalCnt
	t.fieldMap["campus"] = t.Campus
	t.fieldMap["grade"] = t.Grade
	t.fieldMap["meeting_id"] = t.MeetingID
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
}

func (t time) clone(db *gorm.DB) time {
	t.timeDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t time) replaceDB(db *gorm.DB) time {
	t.timeDo.ReplaceDB(db)
	return t
}

type timeDo struct{ gen.DO }

type ITimeDo interface {
	gen.SubQuery
	Debug() ITimeDo
	WithContext(ctx context.Context) ITimeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITimeDo
	WriteDB() ITimeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITimeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITimeDo
	Not(conds ...gen.Condition) ITimeDo
	Or(conds ...gen.Condition) ITimeDo
	Select(conds ...field.Expr) ITimeDo
	Where(conds ...gen.Condition) ITimeDo
	Order(conds ...field.Expr) ITimeDo
	Distinct(cols ...field.Expr) ITimeDo
	Omit(cols ...field.Expr) ITimeDo
	Join(table schema.Tabler, on ...field.Expr) ITimeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITimeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITimeDo
	Group(cols ...field.Expr) ITimeDo
	Having(conds ...gen.Condition) ITimeDo
	Limit(limit int) ITimeDo
	Offset(offset int) ITimeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITimeDo
	Unscoped() ITimeDo
	Create(values ...*model.Time) error
	CreateInBatches(values []*model.Time, batchSize int) error
	Save(values ...*model.Time) error
	First() (*model.Time, error)
	Take() (*model.Time, error)
	Last() (*model.Time, error)
	Find() ([]*model.Time, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Time, err error)
	FindInBatches(result *[]*model.Time, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Time) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITimeDo
	Assign(attrs ...field.AssignExpr) ITimeDo
	Joins(fields ...field.RelationField) ITimeDo
	Preload(fields ...field.RelationField) ITimeDo
	FirstOrInit() (*model.Time, error)
	FirstOrCreate() (*model.Time, error)
	FindByPage(offset int, limit int) (result []*model.Time, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITimeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t timeDo) Debug() ITimeDo {
	return t.withDO(t.DO.Debug())
}

func (t timeDo) WithContext(ctx context.Context) ITimeDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timeDo) ReadDB() ITimeDo {
	return t.Clauses(dbresolver.Read)
}

func (t timeDo) WriteDB() ITimeDo {
	return t.Clauses(dbresolver.Write)
}

func (t timeDo) Session(config *gorm.Session) ITimeDo {
	return t.withDO(t.DO.Session(config))
}

func (t timeDo) Clauses(conds ...clause.Expression) ITimeDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timeDo) Returning(value interface{}, columns ...string) ITimeDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timeDo) Not(conds ...gen.Condition) ITimeDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timeDo) Or(conds ...gen.Condition) ITimeDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timeDo) Select(conds ...field.Expr) ITimeDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timeDo) Where(conds ...gen.Condition) ITimeDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITimeDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t timeDo) Order(conds ...field.Expr) ITimeDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timeDo) Distinct(cols ...field.Expr) ITimeDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timeDo) Omit(cols ...field.Expr) ITimeDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timeDo) Join(table schema.Tabler, on ...field.Expr) ITimeDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITimeDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timeDo) RightJoin(table schema.Tabler, on ...field.Expr) ITimeDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timeDo) Group(cols ...field.Expr) ITimeDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timeDo) Having(conds ...gen.Condition) ITimeDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timeDo) Limit(limit int) ITimeDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timeDo) Offset(offset int) ITimeDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITimeDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timeDo) Unscoped() ITimeDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timeDo) Create(values ...*model.Time) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timeDo) CreateInBatches(values []*model.Time, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timeDo) Save(values ...*model.Time) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timeDo) First() (*model.Time, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Time), nil
	}
}

func (t timeDo) Take() (*model.Time, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Time), nil
	}
}

func (t timeDo) Last() (*model.Time, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Time), nil
	}
}

func (t timeDo) Find() ([]*model.Time, error) {
	result, err := t.DO.Find()
	return result.([]*model.Time), err
}

func (t timeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Time, err error) {
	buf := make([]*model.Time, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timeDo) FindInBatches(result *[]*model.Time, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timeDo) Attrs(attrs ...field.AssignExpr) ITimeDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timeDo) Assign(attrs ...field.AssignExpr) ITimeDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timeDo) Joins(fields ...field.RelationField) ITimeDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timeDo) Preload(fields ...field.RelationField) ITimeDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timeDo) FirstOrInit() (*model.Time, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Time), nil
	}
}

func (t timeDo) FirstOrCreate() (*model.Time, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Time), nil
	}
}

func (t timeDo) FindByPage(offset int, limit int) (result []*model.Time, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t timeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timeDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timeDo) Delete(models ...*model.Time) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timeDo) withDO(do gen.Dao) *timeDo {
	t.DO = *do.(*gen.DO)
	return t
}
