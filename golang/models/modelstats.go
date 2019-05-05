package models

import (
	"time"
)

// ModelStats holds various stats for each model in the database and can be displayed in the models dashboard
type ModelStats struct {
	ID               int       `schema:"id"`
	CreatedAt        time.Time `schema:"created"`
	UpdatedAt        time.Time `schema:"updated"`
	NameOfModel      string    `schema:"model"`      // Name of model
	Total            uint      `schema:"total"`      // Total is the total number of models in existence. Avaiable in template as {{.models_total}}
	References       uint      `schema:"referenced"` // References is how many times per day the model is referenced.Avaiable in template as {{.model_references}}
	PctActive        uint      `schema:"pct_active"` // Is the rounded, total percent of models that are active. Avaiable in template as {{.model_pct_active}}
	Unused           uint      `schema:"unused"`     // Unused is the total number of models that are not being referenced. Avaiable in template as {{.model_unused}}
	Active           uint      `schema:"active"`     // Active is the total number of models that are being referenced. Avaiable in template as {{.model_active}}
	Archived         uint      `schema:"archived"`   // Archived is the total number of models that are currently archived. Avaiable in template as {{.model_archived}}
	TimeBeforeUnused time.Duration
	ModelsStatus     map[int]Status
}

// Status tracks status of a models data
type Status struct {
	IsActive   bool
	LastAccess time.Time
}

// CRUD is called by each models CRUD handler immediately after a sucessful SQL transaction
// TODO: Find a better way of doing this that is more efficient and does not involve programmer action.
func NewModelStats(model string) *ModelStats {
	return &ModelStats{ModelsStatus: make(map[int]Status), NameOfModel: model}
}

// Refresh updates the model stats
func (m *ModelStats) Refresh() {
	m.Active = 0
	m.Archived = 0
	m.Unused = 0

	for _, v := range m.ModelsStatus {
		if v.IsActive && time.Now().Sub(v.LastAccess) < m.TimeBeforeUnused {
			m.Active++
		}
	}

	m.Unused = m.Total - m.Active
}

// SetTotal sets the total nunber of model records available
func (m *ModelStats) SetTotal(total uint) {
	m.Total = total
}

// AddTotal increments the total number of model records by 1
func (m *ModelStats) AddTotal() {
	m.Total++
}

// MinusTotal decrements the total number of model records by 1
func (m *ModelStats) MinusTotal(id int) {
	m.Total--
	delete(m.ModelsStatus, id)
}

// CRUD is used to count the number of active and inactive model records
func (m *ModelStats) CRUD(id int) {
	m.ModelsStatus[id] = Status{IsActive: true, LastAccess: time.Now()}
}
