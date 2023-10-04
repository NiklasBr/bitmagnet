// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bitmagnet-io/bitmagnet/internal/model"
)

func newContentCollection(db *gorm.DB, opts ...gen.DOOption) contentCollection {
	_contentCollection := contentCollection{}

	_contentCollection.contentCollectionDo.UseDB(db, opts...)
	_contentCollection.contentCollectionDo.UseModel(&model.ContentCollection{})

	tableName := _contentCollection.contentCollectionDo.TableName()
	_contentCollection.ALL = field.NewAsterisk(tableName)
	_contentCollection.Type = field.NewString(tableName, "type")
	_contentCollection.Source = field.NewString(tableName, "source")
	_contentCollection.ID = field.NewString(tableName, "id")
	_contentCollection.Name = field.NewString(tableName, "name")
	_contentCollection.CreatedAt = field.NewTime(tableName, "created_at")
	_contentCollection.UpdatedAt = field.NewTime(tableName, "updated_at")
	_contentCollection.MetadataSource = contentCollectionBelongsToMetadataSource{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("MetadataSource", "model.MetadataSource"),
	}

	_contentCollection.fillFieldMap()

	return _contentCollection
}

type contentCollection struct {
	contentCollectionDo

	ALL            field.Asterisk
	Type           field.String
	Source         field.String
	ID             field.String
	Name           field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	MetadataSource contentCollectionBelongsToMetadataSource

	fieldMap map[string]field.Expr
}

func (c contentCollection) Table(newTableName string) *contentCollection {
	c.contentCollectionDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c contentCollection) As(alias string) *contentCollection {
	c.contentCollectionDo.DO = *(c.contentCollectionDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *contentCollection) updateTableName(table string) *contentCollection {
	c.ALL = field.NewAsterisk(table)
	c.Type = field.NewString(table, "type")
	c.Source = field.NewString(table, "source")
	c.ID = field.NewString(table, "id")
	c.Name = field.NewString(table, "name")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *contentCollection) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *contentCollection) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["type"] = c.Type
	c.fieldMap["source"] = c.Source
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt

}

func (c contentCollection) clone(db *gorm.DB) contentCollection {
	c.contentCollectionDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c contentCollection) replaceDB(db *gorm.DB) contentCollection {
	c.contentCollectionDo.ReplaceDB(db)
	return c
}

type contentCollectionBelongsToMetadataSource struct {
	db *gorm.DB

	field.RelationField
}

func (a contentCollectionBelongsToMetadataSource) Where(conds ...field.Expr) *contentCollectionBelongsToMetadataSource {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a contentCollectionBelongsToMetadataSource) WithContext(ctx context.Context) *contentCollectionBelongsToMetadataSource {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a contentCollectionBelongsToMetadataSource) Session(session *gorm.Session) *contentCollectionBelongsToMetadataSource {
	a.db = a.db.Session(session)
	return &a
}

func (a contentCollectionBelongsToMetadataSource) Model(m *model.ContentCollection) *contentCollectionBelongsToMetadataSourceTx {
	return &contentCollectionBelongsToMetadataSourceTx{a.db.Model(m).Association(a.Name())}
}

type contentCollectionBelongsToMetadataSourceTx struct{ tx *gorm.Association }

func (a contentCollectionBelongsToMetadataSourceTx) Find() (result *model.MetadataSource, err error) {
	return result, a.tx.Find(&result)
}

func (a contentCollectionBelongsToMetadataSourceTx) Append(values ...*model.MetadataSource) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a contentCollectionBelongsToMetadataSourceTx) Replace(values ...*model.MetadataSource) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a contentCollectionBelongsToMetadataSourceTx) Delete(values ...*model.MetadataSource) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a contentCollectionBelongsToMetadataSourceTx) Clear() error {
	return a.tx.Clear()
}

func (a contentCollectionBelongsToMetadataSourceTx) Count() int64 {
	return a.tx.Count()
}

type contentCollectionDo struct{ gen.DO }

type IContentCollectionDo interface {
	gen.SubQuery
	Debug() IContentCollectionDo
	WithContext(ctx context.Context) IContentCollectionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IContentCollectionDo
	WriteDB() IContentCollectionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IContentCollectionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IContentCollectionDo
	Not(conds ...gen.Condition) IContentCollectionDo
	Or(conds ...gen.Condition) IContentCollectionDo
	Select(conds ...field.Expr) IContentCollectionDo
	Where(conds ...gen.Condition) IContentCollectionDo
	Order(conds ...field.Expr) IContentCollectionDo
	Distinct(cols ...field.Expr) IContentCollectionDo
	Omit(cols ...field.Expr) IContentCollectionDo
	Join(table schema.Tabler, on ...field.Expr) IContentCollectionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IContentCollectionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IContentCollectionDo
	Group(cols ...field.Expr) IContentCollectionDo
	Having(conds ...gen.Condition) IContentCollectionDo
	Limit(limit int) IContentCollectionDo
	Offset(offset int) IContentCollectionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IContentCollectionDo
	Unscoped() IContentCollectionDo
	Create(values ...*model.ContentCollection) error
	CreateInBatches(values []*model.ContentCollection, batchSize int) error
	Save(values ...*model.ContentCollection) error
	First() (*model.ContentCollection, error)
	Take() (*model.ContentCollection, error)
	Last() (*model.ContentCollection, error)
	Find() ([]*model.ContentCollection, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ContentCollection, err error)
	FindInBatches(result *[]*model.ContentCollection, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ContentCollection) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IContentCollectionDo
	Assign(attrs ...field.AssignExpr) IContentCollectionDo
	Joins(fields ...field.RelationField) IContentCollectionDo
	Preload(fields ...field.RelationField) IContentCollectionDo
	FirstOrInit() (*model.ContentCollection, error)
	FirstOrCreate() (*model.ContentCollection, error)
	FindByPage(offset int, limit int) (result []*model.ContentCollection, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IContentCollectionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c contentCollectionDo) Debug() IContentCollectionDo {
	return c.withDO(c.DO.Debug())
}

func (c contentCollectionDo) WithContext(ctx context.Context) IContentCollectionDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c contentCollectionDo) ReadDB() IContentCollectionDo {
	return c.Clauses(dbresolver.Read)
}

func (c contentCollectionDo) WriteDB() IContentCollectionDo {
	return c.Clauses(dbresolver.Write)
}

func (c contentCollectionDo) Session(config *gorm.Session) IContentCollectionDo {
	return c.withDO(c.DO.Session(config))
}

func (c contentCollectionDo) Clauses(conds ...clause.Expression) IContentCollectionDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c contentCollectionDo) Returning(value interface{}, columns ...string) IContentCollectionDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c contentCollectionDo) Not(conds ...gen.Condition) IContentCollectionDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c contentCollectionDo) Or(conds ...gen.Condition) IContentCollectionDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c contentCollectionDo) Select(conds ...field.Expr) IContentCollectionDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c contentCollectionDo) Where(conds ...gen.Condition) IContentCollectionDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c contentCollectionDo) Order(conds ...field.Expr) IContentCollectionDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c contentCollectionDo) Distinct(cols ...field.Expr) IContentCollectionDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c contentCollectionDo) Omit(cols ...field.Expr) IContentCollectionDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c contentCollectionDo) Join(table schema.Tabler, on ...field.Expr) IContentCollectionDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c contentCollectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IContentCollectionDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c contentCollectionDo) RightJoin(table schema.Tabler, on ...field.Expr) IContentCollectionDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c contentCollectionDo) Group(cols ...field.Expr) IContentCollectionDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c contentCollectionDo) Having(conds ...gen.Condition) IContentCollectionDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c contentCollectionDo) Limit(limit int) IContentCollectionDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c contentCollectionDo) Offset(offset int) IContentCollectionDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c contentCollectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IContentCollectionDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c contentCollectionDo) Unscoped() IContentCollectionDo {
	return c.withDO(c.DO.Unscoped())
}

func (c contentCollectionDo) Create(values ...*model.ContentCollection) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c contentCollectionDo) CreateInBatches(values []*model.ContentCollection, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c contentCollectionDo) Save(values ...*model.ContentCollection) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c contentCollectionDo) First() (*model.ContentCollection, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentCollection), nil
	}
}

func (c contentCollectionDo) Take() (*model.ContentCollection, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentCollection), nil
	}
}

func (c contentCollectionDo) Last() (*model.ContentCollection, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentCollection), nil
	}
}

func (c contentCollectionDo) Find() ([]*model.ContentCollection, error) {
	result, err := c.DO.Find()
	return result.([]*model.ContentCollection), err
}

func (c contentCollectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ContentCollection, err error) {
	buf := make([]*model.ContentCollection, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c contentCollectionDo) FindInBatches(result *[]*model.ContentCollection, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c contentCollectionDo) Attrs(attrs ...field.AssignExpr) IContentCollectionDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c contentCollectionDo) Assign(attrs ...field.AssignExpr) IContentCollectionDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c contentCollectionDo) Joins(fields ...field.RelationField) IContentCollectionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c contentCollectionDo) Preload(fields ...field.RelationField) IContentCollectionDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c contentCollectionDo) FirstOrInit() (*model.ContentCollection, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentCollection), nil
	}
}

func (c contentCollectionDo) FirstOrCreate() (*model.ContentCollection, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentCollection), nil
	}
}

func (c contentCollectionDo) FindByPage(offset int, limit int) (result []*model.ContentCollection, count int64, err error) {
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

func (c contentCollectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c contentCollectionDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c contentCollectionDo) Delete(models ...*model.ContentCollection) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *contentCollectionDo) withDO(do gen.Dao) *contentCollectionDo {
	c.DO = *do.(*gen.DO)
	return c
}
