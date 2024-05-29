package medicaltreatment

import "home-care/app"

// MedicalTreatment is the main model of MedicalTreatment data. It provides a convenient interface for app.ModelInterface
type MedicalTreatment struct {
	app.Model
	ID                     app.NullUUID     `json:"id"                        db:"m.id"                        gorm:"column:id;primaryKey"`
	TreatName              app.NullString   `json:"treat_name"                db:"m.treat_name"                gorm:"column:treat_name"`
	TreatPrice             app.NullFloat64  `json:"treat_price"               db:"m.treat_price"               gorm:"column:treat_price"`
	ShareOfOfficersValue   app.NullFloat64  `json:"share_of_officers_value"   db:"m.share_of_officers_value"   gorm:"column:share_of_officers_value"`
	ShareOfClinicsValue    app.NullFloat64  `json:"share_of_clinics_value"    db:"m.share_of_clinics_value"    gorm:"column:share_of_clinics_value"`
	ShareOfOfficersPercent app.NullFloat64  `json:"share_of_officers_percent" db:"m.share_of_officers_percent" gorm:"column:share_of_officers_percent"`
	ShareOfClinicsPercent  app.NullFloat64  `json:"share_of_clinics_percent"  db:"m.share_of_clinics_percent"  gorm:"column:share_of_clinics_percent"`
	CreatedBy              app.NullInt64    `json:"created_by"                db:"m.created_by"                gorm:"column:created_by"`
	ModifiedBy             app.NullInt64    `json:"modified_by"               db:"m.modified_by"               gorm:"column:modified_by"`
	CreatedAt              app.NullDateTime `json:"created_at"                db:"m.created_at"                gorm:"column:created_at"`
	UpdatedAt              app.NullDateTime `json:"updated_at"                db:"m.updated_at"                gorm:"column:updated_at"`
	DeletedAt              app.NullDateTime `json:"deleted_at"                db:"m.deleted_at,hide"           gorm:"column:deleted_at"`
}

// EndPoint returns the MedicalTreatment end point, it used for cache key, etc.
func (MedicalTreatment) EndPoint() string {
	return "medical_treatments"
}

// TableVersion returns the versions of the MedicalTreatment table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (MedicalTreatment) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the MedicalTreatment table in the database.
func (MedicalTreatment) TableName() string {
	return "medical_treatments"
}

// TableAliasName returns the table alias name of the MedicalTreatment table, used for querying.
func (MedicalTreatment) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the MedicalTreatment data in the database, used for querying.
func (m *MedicalTreatment) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the MedicalTreatment data in the database, used for querying.
func (m *MedicalTreatment) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the MedicalTreatment data in the database, used for querying.
func (m *MedicalTreatment) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the MedicalTreatment data in the database, used for querying.
func (m *MedicalTreatment) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the MedicalTreatment schema, used for querying.
func (m *MedicalTreatment) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the MedicalTreatment schema in the open api documentation.
func (MedicalTreatment) OpenAPISchemaName() string {
	return "MedicalTreatment"
}

// GetOpenAPISchema returns the Open API Schema of the MedicalTreatment in the open api documentation.
func (m *MedicalTreatment) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type MedicalTreatmentList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the MedicalTreatmentList schema in the open api documentation.
func (MedicalTreatmentList) OpenAPISchemaName() string {
	return "MedicalTreatmentList"
}

// GetOpenAPISchema returns the Open API Schema of the MedicalTreatmentList in the open api documentation.
func (p *MedicalTreatmentList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&MedicalTreatment{})
}

// ParamCreate is the expected parameters for create a new MedicalTreatment data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the MedicalTreatment data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the MedicalTreatment data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the MedicalTreatment data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
