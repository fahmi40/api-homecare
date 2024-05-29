package schedule

import "home-care/app"

// Schedule is the main model of Schedule data. It provides a convenient interface for app.ModelInterface
type Schedule struct {
	app.Model
	ID             app.NullUUID     `json:"id"              db:"m.id"              gorm:"column:id;primaryKey"`
	SchedDate      app.NullInt64    `json:"sched_date"      db:"m.sched_date"      gorm:"column:sched_date"`
	RegistrationId app.NullInt64    `json:"registration_id" db:"m.registration_id" gorm:"column:registration_id"`
	Address        app.NullText     `json:"address"         db:"m.address"         gorm:"column:address"`
	RegencyOrCity  app.NullInt64    `json:"regency_or_city" db:"m.regency_or_city" gorm:"column:regency_or_city"`
	Province       app.NullInt64    `json:"province"        db:"m.province"        gorm:"column:province"`
	PostalCode     app.NullString   `json:"postal_code"     db:"m.postal_code"     gorm:"column:postal_code"`
	CreatedBy      app.NullInt64    `json:"created_by"      db:"m.created_by"      gorm:"column:created_by"`
	ModifiedBy     app.NullString   `json:"modified_by"     db:"m.modified_by"     gorm:"column:modified_by"`
	CreatedAt      app.NullDateTime `json:"created_at"      db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt      app.NullDateTime `json:"updated_at"      db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt      app.NullDateTime `json:"deleted_at"      db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Schedule end point, it used for cache key, etc.
func (Schedule) EndPoint() string {
	return "schedules"
}

// TableVersion returns the versions of the Schedule table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Schedule) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Schedule table in the database.
func (Schedule) TableName() string {
	return "schedules"
}

// TableAliasName returns the table alias name of the Schedule table, used for querying.
func (Schedule) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Schedule data in the database, used for querying.
func (m *Schedule) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Schedule data in the database, used for querying.
func (m *Schedule) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Schedule data in the database, used for querying.
func (m *Schedule) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Schedule data in the database, used for querying.
func (m *Schedule) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Schedule schema, used for querying.
func (m *Schedule) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Schedule schema in the open api documentation.
func (Schedule) OpenAPISchemaName() string {
	return "Schedule"
}

// GetOpenAPISchema returns the Open API Schema of the Schedule in the open api documentation.
func (m *Schedule) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type ScheduleList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the ScheduleList schema in the open api documentation.
func (ScheduleList) OpenAPISchemaName() string {
	return "ScheduleList"
}

// GetOpenAPISchema returns the Open API Schema of the ScheduleList in the open api documentation.
func (p *ScheduleList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Schedule{})
}

// ParamCreate is the expected parameters for create a new Schedule data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Schedule data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Schedule data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Schedule data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
