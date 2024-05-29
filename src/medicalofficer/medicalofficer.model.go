package medicalofficer

import "home-care/app"

// MedicalOfficer is the main model of MedicalOfficer data. It provides a convenient interface for app.ModelInterface
type MedicalOfficer struct {
	app.Model
	ID            app.NullUUID     `json:"id"              db:"m.id"              gorm:"column:id;primaryKey"`
	FullName      app.NullString   `json:"full_name"       db:"m.full_name"       gorm:"column:full_name"`
	TitleName     app.NullString   `json:"title_name"      db:"m.title_name"      gorm:"column:title_name"`
	PhoneNumber   app.NullString   `json:"phone_number"    db:"m.phone_number"    gorm:"column:phone_number"`
	Email         app.NullString   `json:"email"           db:"m.email"           gorm:"column:email"`
	NoKtp         app.NullString   `json:"no_ktp"          db:"m.no_ktp"          gorm:"column:no_ktp"`
	NoSip         app.NullString   `json:"no_sip"          db:"m.no_sip"          gorm:"column:no_sip"`
	Birthdate     app.NullDate     `json:"birthdate"       db:"m.birthdate"       gorm:"column:birthdate"`
	PlaceOfBirth  app.NullString   `json:"place_of_birth"  db:"m.place_of_birth"  gorm:"column:place_of_birth"`
	Religion      app.NullString   `json:"religion"        db:"m.religion"        gorm:"column:religion"`
	Gender        app.NullString   `json:"gender"          db:"m.gender"          gorm:"column:gender"`
	Address       app.NullText     `json:"address"         db:"m.address"         gorm:"column:address"`
	RegencyOrCity app.NullInt64    `json:"regency_or_city" db:"m.regency_or_city" gorm:"column:regency_or_city"`
	Province      app.NullInt64    `json:"province"        db:"m.province"        gorm:"column:province"`
	PostalCode    app.NullString   `json:"postal_code"     db:"m.postal_code"     gorm:"column:postal_code"`
	Latitude      app.NullString   `json:"latitude"        db:"m.latitude"        gorm:"column:latitude"`
	Longitude     app.NullString   `json:"longitude"       db:"m.longitude"       gorm:"column:longitude"`
	CreatedBy     app.NullInt64    `json:"created_by"      db:"m.created_by"      gorm:"column:created_by"`
	ModifiedBy    app.NullInt64    `json:"modified_by"     db:"m.modified_by"     gorm:"column:modified_by"`
	CreatedAt     app.NullDateTime `json:"created_at"      db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt     app.NullDateTime `json:"updated_at"      db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt     app.NullDateTime `json:"deleted_at"      db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the MedicalOfficer end point, it used for cache key, etc.
func (MedicalOfficer) EndPoint() string {
	return "medical_officers"
}

// TableVersion returns the versions of the MedicalOfficer table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (MedicalOfficer) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the MedicalOfficer table in the database.
func (MedicalOfficer) TableName() string {
	return "medical_officers"
}

// TableAliasName returns the table alias name of the MedicalOfficer table, used for querying.
func (MedicalOfficer) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the MedicalOfficer data in the database, used for querying.
func (m *MedicalOfficer) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the MedicalOfficer data in the database, used for querying.
func (m *MedicalOfficer) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the MedicalOfficer data in the database, used for querying.
func (m *MedicalOfficer) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the MedicalOfficer data in the database, used for querying.
func (m *MedicalOfficer) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the MedicalOfficer schema, used for querying.
func (m *MedicalOfficer) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the MedicalOfficer schema in the open api documentation.
func (MedicalOfficer) OpenAPISchemaName() string {
	return "MedicalOfficer"
}

// GetOpenAPISchema returns the Open API Schema of the MedicalOfficer in the open api documentation.
func (m *MedicalOfficer) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type MedicalOfficerList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the MedicalOfficerList schema in the open api documentation.
func (MedicalOfficerList) OpenAPISchemaName() string {
	return "MedicalOfficerList"
}

// GetOpenAPISchema returns the Open API Schema of the MedicalOfficerList in the open api documentation.
func (p *MedicalOfficerList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&MedicalOfficer{})
}

// ParamCreate is the expected parameters for create a new MedicalOfficer data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the MedicalOfficer data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the MedicalOfficer data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the MedicalOfficer data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
