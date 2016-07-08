package models

import (
	"time"

	"github.com/satori/go.uuid"

	"gopkg.in/gorp.v1"
)

// Organization identifies uniquely one organization
type Organization struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	DeletedAt int64  `db:"deleted_at"`
}

// PreInsert populates fields before inserting a new organization
func (o *Organization) PreInsert(s gorp.SqlExecutor) error {
	o.ID = uuid.NewV4().String()
	o.CreatedAt = time.Now().Unix()
	o.UpdatedAt = o.CreatedAt
	return nil
}

// PreUpdate populates fields before updating an organization
func (o *Organization) PreUpdate(s gorp.SqlExecutor) error {
	o.UpdatedAt = time.Now().Unix()
	return nil
}

// GetOrganizationByID returns an organization by id
func GetOrganizationByID(db DB, id string) (*Organization, error) {
	obj, err := db.Get(Organization{}, id)
	if err != nil || obj == nil {
		return nil, &ModelNotFoundError{"Organization", "id", id}
	}
	return obj.(*Organization), nil
}

// GetOrganizationByName returns an organization by its name
func GetOrganizationByName(db DB, name string) (*Organization, error) {
	var organization Organization
	err := db.SelectOne(&organization, "SELECT * FROM organizations WHERE name=$1", name)
	if err != nil || &organization == nil {
		return nil, &ModelNotFoundError{"Organization", "name", name}
	}
	return &organization, nil
}

// CreateOrganization creates a new Organization
func CreateOrganization(db DB, Name string) (*Organization, error) {
	organization := &Organization{
		Name: Name,
	}
	err := db.Insert(organization)
	if err != nil {
		return nil, err
	}
	return organization, nil
}