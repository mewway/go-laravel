package console

type Stubs struct {
}

func (r Stubs) Model() string {
	return `package DummyPackage

import (
	"time"

	"github.com/mewway/go-laravel/database/orm"
)

DummyConst
type DummyModel struct {
	orm.Model
	DummyField
}
`
}

func (r Stubs) Observer() string {
	return `package DummyPackage

import (
	"github.com/mewway/go-laravel/contracts/database/orm"
)


type DummyObserver struct{}

func (u *DummyObserver) Retrieved(event orm.Event) error {
	return nil
}

func (u *DummyObserver) Creating(event orm.Event) error {
	return nil
}

func (u *DummyObserver) Created(event orm.Event) error {
	return nil
}

func (u *DummyObserver) Updating(event orm.Event) error {
	return nil
}

func (u *DummyObserver) Updated(event orm.Event) error {
	return nil
}

func (u *DummyObserver) Saving(event orm.Event) error {
	return nil
}

func (u *DummyObserver) Saved(event orm.Event) error {
	return nil
}

func (u *DummyObserver) Deleting(event orm.Event) error {
	return nil
}

func (u *DummyObserver) Deleted(event orm.Event) error {
	return nil
}

func (u *DummyObserver) ForceDeleting(event orm.Event) error {
	return nil
}

func (u *DummyObserver) ForceDeleted(event orm.Event) error {
	return nil
}
`
}

func (r Stubs) Seeder() string {
	return `package DummyPackage
	
type DummySeeder struct {
}

// Signature The name and signature of the seeder.
func (s *DummySeeder) Signature() string {
	return "DummySeeder"
}

// Run executes the seeder logic.
func (s *DummySeeder) Run() error {
	return nil
}
`
}
