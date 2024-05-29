package transactiondetail

import "home-care/app"

// TransactionDetail is the main model of TransactionDetail data. It provides a convenient interface for app.ModelInterface
type TransactionDetail struct {
	app.Model
	ID         app.NullUUID     `json:"id"          db:"m.id"              gorm:"column:id;primaryKey"`
	TransId    app.NullInt64    `json:"trans_id"    db:"m.trans_id"        gorm:"column:trans_id"`
	TreatId    app.NullInt64    `json:"treat_id"    db:"m.treat_id"        gorm:"column:treat_id"`
	TreatPrice app.NullFloat64  `json:"treat_price" db:"m.treat_price"     gorm:"column:treat_price"`
	Qty        app.NullInt64    `json:"qty"         db:"m.qty"             gorm:"column:qty"`
	Total      app.NullFloat64  `json:"total"       db:"m.total"           gorm:"column:total"`
	CreatedBy  app.NullInt64    `json:"created_by"  db:"m.created_by"      gorm:"column:created_by"`
	ModifiedBy app.NullInt64    `json:"modified_by" db:"m.modified_by"     gorm:"column:modified_by"`
	CreatedAt  app.NullDateTime `json:"created_at"  db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt  app.NullDateTime `json:"updated_at"  db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt  app.NullDateTime `json:"deleted_at"  db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the TransactionDetail end point, it used for cache key, etc.
func (TransactionDetail) EndPoint() string {
	return "transaction_details"
}

// TableVersion returns the versions of the TransactionDetail table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (TransactionDetail) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the TransactionDetail table in the database.
func (TransactionDetail) TableName() string {
	return "transaction_details"
}

// TableAliasName returns the table alias name of the TransactionDetail table, used for querying.
func (TransactionDetail) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the TransactionDetail data in the database, used for querying.
func (m *TransactionDetail) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the TransactionDetail data in the database, used for querying.
func (m *TransactionDetail) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the TransactionDetail data in the database, used for querying.
func (m *TransactionDetail) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the TransactionDetail data in the database, used for querying.
func (m *TransactionDetail) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the TransactionDetail schema, used for querying.
func (m *TransactionDetail) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the TransactionDetail schema in the open api documentation.
func (TransactionDetail) OpenAPISchemaName() string {
	return "TransactionDetail"
}

// GetOpenAPISchema returns the Open API Schema of the TransactionDetail in the open api documentation.
func (m *TransactionDetail) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TransactionDetailList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the TransactionDetailList schema in the open api documentation.
func (TransactionDetailList) OpenAPISchemaName() string {
	return "TransactionDetailList"
}

// GetOpenAPISchema returns the Open API Schema of the TransactionDetailList in the open api documentation.
func (p *TransactionDetailList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TransactionDetail{})
}

// ParamCreate is the expected parameters for create a new TransactionDetail data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the TransactionDetail data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the TransactionDetail data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the TransactionDetail data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
