package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/NlaakStudiosLLC/GoWAF-Framework/logger"

	"github.com/casbin/gorm-adapter"

	"github.com/NlaakStudiosLLC/GoWAF-Framework/database"

	"github.com/NlaakStudiosLLC/GoWAF-Framework/config"
	"github.com/jinzhu/gorm"

	// support none, mysql, sqlite3 and postgresql

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

// Model facilitate database interactions, supports postgres, mysql and foundation
type Model struct {
	models map[string]reflect.Value
	isOpen bool
	database.Database
	*gorm.DB
	*gormadapter.Adapter
}

// NewModel returns a new Model without opening database connection
func NewModel() *Model {
	return &Model{
		models: make(map[string]reflect.Value),
	}
}

// Count returns the number of registered models
func (m *Model) Count() int {
	return len(m.models)
}

// IsOpen returns true if the Model has already established connection
// to the database
func (m *Model) IsOpen() bool {
	return m.isOpen
}

// OpenWithConfig opens database connection with the settings found in cfg
func (m *Model) OpenWithConfig(cfg *config.Config) error {

	//See if a database was defined in the config
	if len(cfg.Database) < 5 {
		// Not using a database
		return nil
	}

	//try and open a connection to the database defined in the config
	db, err := gorm.Open(cfg.Database, cfg.DatabaseConn)
	if err != nil {
		return err
	}

	//Success we have a database connection
	m.DB = db
	m.Database.DB = db
	m.isOpen = true

	return nil
}

// Register adds the values to the models registry
func (m *Model) Register(values ...interface{}) error {

	// do not work on them.models first, this is like an insurance policy
	// whenever we encounter any error in the values nothing goes into the registry
	models := make(map[string]reflect.Value)
	if len(values) > 0 {
		for _, val := range values {
			rVal := reflect.ValueOf(val)
			if rVal.Kind() == reflect.Ptr {
				rVal = rVal.Elem()
			}
			switch rVal.Kind() {
			case reflect.Struct:
				models[getTypName(rVal.Type())] = reflect.New(rVal.Type())
			default:
				return errors.New("gowaf: models must be structs")
			}
		}
	}

	for k, v := range models {
		m.models[k] = v
	}

	return nil
}

// DropTables Drops All Model Database Tables
func (m *Model) DropTables(drop bool, verbose bool) {
	if drop {
		cnt := len(m.models)
		dberr := ""

		logger.LogThis.Warn(fmt.Sprintf("Dropping All [%d] Tables...", cnt))

		for k, v := range m.models {

			m.DropTableIfExists(v.Interface())
			if verbose {
				if m.DB.Error != nil {
					dberr = "Failed."
				} else {
					dberr = "Success."
				}

				logger.LogThis.Warn(fmt.Sprintf("- Dropping %s...%s", k, dberr))
			}
		}

		logger.LogThis.Info("Done.")
	}
}

// AutoMigrateAll runs migrations for all the registered models
func (m *Model) AutoMigrateAll(migrate bool) {
	if migrate {
		for _, v := range m.models {
			m.DB.AutoMigrate(v.Interface())
		}
	}
}

func getTypName(typ reflect.Type) string {
	if typ.Name() != "" {
		return typ.Name()
	}
	split := strings.Split(typ.String(), ".")
	return split[len(split)-1]
}
