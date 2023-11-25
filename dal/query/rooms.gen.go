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

func newRoom(db *gorm.DB, opts ...gen.DOOption) room {
	_room := room{}

	_room.roomDo.UseDB(db, opts...)
	_room.roomDo.UseModel(&model.Room{})

	tableName := _room.roomDo.TableName()
	_room.ALL = field.NewAsterisk(tableName)
	_room.ID = field.NewInt64(tableName, "id")
	_room.Name = field.NewString(tableName, "name")
	_room.Location = field.NewString(tableName, "location")
	_room.Status = field.NewInt32(tableName, "status")
	_room.StatusUpdatedAt = field.NewTime(tableName, "status_updated_at")
	_room.ApplicantTimeID = field.NewInt64(tableName, "applicant_time_id")
	_room.InterviewerComment = field.NewString(tableName, "interviewer_comment")
	_room.ReceiverComment = field.NewString(tableName, "receiver_comment")
	_room.GroupLabel = field.NewString(tableName, "group_label")
	_room.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_room.DeletedAt = field.NewField(tableName, "deleted_at")
	_room.CreatedAt = field.NewTime(tableName, "created_at")
	_room.UpdatedAt = field.NewTime(tableName, "updated_at")

	_room.fillFieldMap()

	return _room
}

type room struct {
	roomDo

	ALL                field.Asterisk
	ID                 field.Int64
	Name               field.String // 房间名称
	Location           field.String // 房间位置
	Status             field.Int32  // 状态 0-已停用 1-休息中 2-等待中 3-已占用
	StatusUpdatedAt    field.Time   // 状态更新时间
	ApplicantTimeID    field.Int64  // 申请时间ID
	InterviewerComment field.String // 面试官留言
	ReceiverComment    field.String // 接待者留言
	GroupLabel         field.String // 组别标签
	UpdatedBy          field.Int64  // 更新者
	DeletedAt          field.Field
	CreatedAt          field.Time
	UpdatedAt          field.Time

	fieldMap map[string]field.Expr
}

func (r room) Table(newTableName string) *room {
	r.roomDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r room) As(alias string) *room {
	r.roomDo.DO = *(r.roomDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *room) updateTableName(table string) *room {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.Name = field.NewString(table, "name")
	r.Location = field.NewString(table, "location")
	r.Status = field.NewInt32(table, "status")
	r.StatusUpdatedAt = field.NewTime(table, "status_updated_at")
	r.ApplicantTimeID = field.NewInt64(table, "applicant_time_id")
	r.InterviewerComment = field.NewString(table, "interviewer_comment")
	r.ReceiverComment = field.NewString(table, "receiver_comment")
	r.GroupLabel = field.NewString(table, "group_label")
	r.UpdatedBy = field.NewInt64(table, "updated_by")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")

	r.fillFieldMap()

	return r
}

func (r *room) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *room) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 13)
	r.fieldMap["id"] = r.ID
	r.fieldMap["name"] = r.Name
	r.fieldMap["location"] = r.Location
	r.fieldMap["status"] = r.Status
	r.fieldMap["status_updated_at"] = r.StatusUpdatedAt
	r.fieldMap["applicant_time_id"] = r.ApplicantTimeID
	r.fieldMap["interviewer_comment"] = r.InterviewerComment
	r.fieldMap["receiver_comment"] = r.ReceiverComment
	r.fieldMap["group_label"] = r.GroupLabel
	r.fieldMap["updated_by"] = r.UpdatedBy
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
}

func (r room) clone(db *gorm.DB) room {
	r.roomDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r room) replaceDB(db *gorm.DB) room {
	r.roomDo.ReplaceDB(db)
	return r
}

type roomDo struct{ gen.DO }

type IRoomDo interface {
	gen.SubQuery
	Debug() IRoomDo
	WithContext(ctx context.Context) IRoomDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRoomDo
	WriteDB() IRoomDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRoomDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRoomDo
	Not(conds ...gen.Condition) IRoomDo
	Or(conds ...gen.Condition) IRoomDo
	Select(conds ...field.Expr) IRoomDo
	Where(conds ...gen.Condition) IRoomDo
	Order(conds ...field.Expr) IRoomDo
	Distinct(cols ...field.Expr) IRoomDo
	Omit(cols ...field.Expr) IRoomDo
	Join(table schema.Tabler, on ...field.Expr) IRoomDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRoomDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRoomDo
	Group(cols ...field.Expr) IRoomDo
	Having(conds ...gen.Condition) IRoomDo
	Limit(limit int) IRoomDo
	Offset(offset int) IRoomDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRoomDo
	Unscoped() IRoomDo
	Create(values ...*model.Room) error
	CreateInBatches(values []*model.Room, batchSize int) error
	Save(values ...*model.Room) error
	First() (*model.Room, error)
	Take() (*model.Room, error)
	Last() (*model.Room, error)
	Find() ([]*model.Room, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Room, err error)
	FindInBatches(result *[]*model.Room, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Room) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRoomDo
	Assign(attrs ...field.AssignExpr) IRoomDo
	Joins(fields ...field.RelationField) IRoomDo
	Preload(fields ...field.RelationField) IRoomDo
	FirstOrInit() (*model.Room, error)
	FirstOrCreate() (*model.Room, error)
	FindByPage(offset int, limit int) (result []*model.Room, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRoomDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r roomDo) Debug() IRoomDo {
	return r.withDO(r.DO.Debug())
}

func (r roomDo) WithContext(ctx context.Context) IRoomDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roomDo) ReadDB() IRoomDo {
	return r.Clauses(dbresolver.Read)
}

func (r roomDo) WriteDB() IRoomDo {
	return r.Clauses(dbresolver.Write)
}

func (r roomDo) Session(config *gorm.Session) IRoomDo {
	return r.withDO(r.DO.Session(config))
}

func (r roomDo) Clauses(conds ...clause.Expression) IRoomDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roomDo) Returning(value interface{}, columns ...string) IRoomDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roomDo) Not(conds ...gen.Condition) IRoomDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roomDo) Or(conds ...gen.Condition) IRoomDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roomDo) Select(conds ...field.Expr) IRoomDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roomDo) Where(conds ...gen.Condition) IRoomDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roomDo) Order(conds ...field.Expr) IRoomDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roomDo) Distinct(cols ...field.Expr) IRoomDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roomDo) Omit(cols ...field.Expr) IRoomDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roomDo) Join(table schema.Tabler, on ...field.Expr) IRoomDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roomDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRoomDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roomDo) RightJoin(table schema.Tabler, on ...field.Expr) IRoomDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roomDo) Group(cols ...field.Expr) IRoomDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roomDo) Having(conds ...gen.Condition) IRoomDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roomDo) Limit(limit int) IRoomDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roomDo) Offset(offset int) IRoomDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roomDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRoomDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roomDo) Unscoped() IRoomDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roomDo) Create(values ...*model.Room) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roomDo) CreateInBatches(values []*model.Room, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roomDo) Save(values ...*model.Room) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roomDo) First() (*model.Room, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Room), nil
	}
}

func (r roomDo) Take() (*model.Room, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Room), nil
	}
}

func (r roomDo) Last() (*model.Room, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Room), nil
	}
}

func (r roomDo) Find() ([]*model.Room, error) {
	result, err := r.DO.Find()
	return result.([]*model.Room), err
}

func (r roomDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Room, err error) {
	buf := make([]*model.Room, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roomDo) FindInBatches(result *[]*model.Room, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roomDo) Attrs(attrs ...field.AssignExpr) IRoomDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roomDo) Assign(attrs ...field.AssignExpr) IRoomDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roomDo) Joins(fields ...field.RelationField) IRoomDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roomDo) Preload(fields ...field.RelationField) IRoomDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roomDo) FirstOrInit() (*model.Room, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Room), nil
	}
}

func (r roomDo) FirstOrCreate() (*model.Room, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Room), nil
	}
}

func (r roomDo) FindByPage(offset int, limit int) (result []*model.Room, count int64, err error) {
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

func (r roomDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roomDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roomDo) Delete(models ...*model.Room) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roomDo) withDO(do gen.Dao) *roomDo {
	r.DO = *do.(*gen.DO)
	return r
}
