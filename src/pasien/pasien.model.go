package pasien

import "home-care-api/app"

// Pasien is the main model of Pasien data. It provides a convenient interface for app.ModelInterface
type Pasien struct {
	app.Model
	ID        app.NullUUID     `json:"id"         db:"m.id"              gorm:"column:id;primaryKey"`
	Nik       app.NullString   `json:"nik"        db:"m.nik"             gorm:"column:nik"`
	Name      app.NullString   `json:"name"       db:"m.name"            gorm:"column:name"`
	Phone     app.NullString   `json:"phone"      db:"m.phone"           gorm:"column:phone"`
	Gender    app.NullString   `json:"gender"     db:"m.gender"          gorm:"column:gender"`
	Address   app.NullString   `json:"address"    db:"m.address"         gorm:"column:address"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Pasien end point, it used for cache key, etc.
func (Pasien) EndPoint() string {
	return "pasien"
}

// TableVersion returns the versions of the Pasien table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Pasien) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Pasien table in the database.
func (Pasien) TableName() string {
	return "pasien"
}

// TableAliasName returns the table alias name of the Pasien table, used for querying.
func (Pasien) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Pasien data in the database, used for querying.
func (m *Pasien) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Pasien data in the database, used for querying.
func (m *Pasien) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Pasien data in the database, used for querying.
func (m *Pasien) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Pasien data in the database, used for querying.
func (m *Pasien) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Pasien schema, used for querying.
func (m *Pasien) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Pasien schema in the open api documentation.
func (Pasien) OpenAPISchemaName() string {
	return "Pasien"
}

// GetOpenAPISchema returns the Open API Schema of the Pasien in the open api documentation.
func (m *Pasien) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type PasienList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the PasienList schema in the open api documentation.
func (PasienList) OpenAPISchemaName() string {
	return "PasienList"
}

// GetOpenAPISchema returns the Open API Schema of the PasienList in the open api documentation.
func (p *PasienList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Pasien{})
}

// ParamCreate is the expected parameters for create a new Pasien data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Pasien data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Pasien data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Pasien data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
