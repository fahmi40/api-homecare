package transaction

import "home-care/app"

// Transaction is the main model of Transaction data. It provides a convenient interface for app.ModelInterface
type Transaction struct {
	app.Model
	ID               app.NullUUID     `json:"id"                db:"m.id"                gorm:"column:id;primaryKey"`
	TransNumber      app.NullString   `json:"trans_number"      db:"m.trans_number"      gorm:"column:trans_number"`
	TransDate        app.NullInt64    `json:"trans_date"        db:"m.trans_date"        gorm:"column:trans_date"`
	SchedId          app.NullInt64    `json:"sched_id"          db:"m.sched_id"          gorm:"column:sched_id"`
	Notes            app.NullText     `json:"notes"             db:"m.notes"             gorm:"column:notes"`
	TotalPrices      app.NullFloat64  `json:"total_prices"      db:"m.total_prices"      gorm:"column:total_prices"`
	ConfirmationId   app.NullInt64    `json:"confirmation_id"   db:"m.confirmation_id"   gorm:"column:confirmation_id"`
	ConfirmationDate app.NullInt64    `json:"confirmation_date" db:"m.confirmation_date" gorm:"column:confirmation_date"`
	CreatedBy        app.NullInt64    `json:"created_by"        db:"m.created_by"        gorm:"column:created_by"`
	ModifiedBy       app.NullInt64    `json:"modified_by"       db:"m.modified_by"       gorm:"column:modified_by"`
	CreatedAt        app.NullDateTime `json:"created_at"        db:"m.created_at"        gorm:"column:created_at"`
	UpdatedAt        app.NullDateTime `json:"updated_at"        db:"m.updated_at"        gorm:"column:updated_at"`
	DeletedAt        app.NullDateTime `json:"deleted_at"        db:"m.deleted_at,hide"   gorm:"column:deleted_at"`
}

// EndPoint returns the Transaction end point, it used for cache key, etc.
func (Transaction) EndPoint() string {
	return "transactions"
}

// TableVersion returns the versions of the Transaction table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Transaction) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Transaction table in the database.
func (Transaction) TableName() string {
	return "transactions"
}

// TableAliasName returns the table alias name of the Transaction table, used for querying.
func (Transaction) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Transaction data in the database, used for querying.
func (m *Transaction) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Transaction data in the database, used for querying.
func (m *Transaction) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Transaction data in the database, used for querying.
func (m *Transaction) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Transaction data in the database, used for querying.
func (m *Transaction) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Transaction schema, used for querying.
func (m *Transaction) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Transaction schema in the open api documentation.
func (Transaction) OpenAPISchemaName() string {
	return "Transaction"
}

// GetOpenAPISchema returns the Open API Schema of the Transaction in the open api documentation.
func (m *Transaction) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TransactionList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the TransactionList schema in the open api documentation.
func (TransactionList) OpenAPISchemaName() string {
	return "TransactionList"
}

// GetOpenAPISchema returns the Open API Schema of the TransactionList in the open api documentation.
func (p *TransactionList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Transaction{})
}

// ParamCreate is the expected parameters for create a new Transaction data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Transaction data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Transaction data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Transaction data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
