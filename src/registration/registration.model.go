package registration

import "home-care/app"

// Registration is the main model of Registration data. It provides a convenient interface for app.ModelInterface
type Registration struct {
	app.Model
	ID          app.NullUUID     `json:"id"           db:"m.id"              gorm:"column:id;primaryKey"`
	RegDate     app.NullInt64    `json:"reg_date"     db:"m.reg_date"        gorm:"column:reg_date"`
	RegNumber   app.NullString   `json:"reg_number"   db:"m.reg_number"      gorm:"column:reg_number"`
	PatientId   app.NullInt64    `json:"patient_id"   db:"m.patient_id"      gorm:"column:patient_id"`
	PhoneNumber app.NullString   `json:"phone_number" db:"m.phone_number"    gorm:"column:phone_number"`
	Complaints  app.NullText     `json:"complaints"   db:"m.complaints"      gorm:"column:complaints"`
	VisitPlan   app.NullInt64    `json:"visit_plan"   db:"m.visit_plan"      gorm:"column:visit_plan"`
	CreatedBy   app.NullInt64    `json:"created_by"   db:"m.created_by"      gorm:"column:created_by"`
	ModifiedBy  app.NullInt64    `json:"modified_by"  db:"m.modified_by"     gorm:"column:modified_by"`
	CreatedAt   app.NullDateTime `json:"created_at"   db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt   app.NullDateTime `json:"updated_at"   db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt   app.NullDateTime `json:"deleted_at"   db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Registration end point, it used for cache key, etc.
func (Registration) EndPoint() string {
	return "registrations"
}

// TableVersion returns the versions of the Registration table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Registration) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Registration table in the database.
func (Registration) TableName() string {
	return "registrations"
}

// TableAliasName returns the table alias name of the Registration table, used for querying.
func (Registration) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Registration data in the database, used for querying.
func (m *Registration) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Registration data in the database, used for querying.
func (m *Registration) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Registration data in the database, used for querying.
func (m *Registration) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Registration data in the database, used for querying.
func (m *Registration) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Registration schema, used for querying.
func (m *Registration) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Registration schema in the open api documentation.
func (Registration) OpenAPISchemaName() string {
	return "Registration"
}

// GetOpenAPISchema returns the Open API Schema of the Registration in the open api documentation.
func (m *Registration) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type RegistrationList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the RegistrationList schema in the open api documentation.
func (RegistrationList) OpenAPISchemaName() string {
	return "RegistrationList"
}

// GetOpenAPISchema returns the Open API Schema of the RegistrationList in the open api documentation.
func (p *RegistrationList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Registration{})
}

// ParamCreate is the expected parameters for create a new Registration data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Registration data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Registration data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Registration data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
