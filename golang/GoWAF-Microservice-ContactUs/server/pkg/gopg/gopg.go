package gopg

import (
	goPg "github.com/AndrewDonelson/go-pg-orm"
)

//New return new instance of goPg.Model
func New(role, database, password string, table interface{}, tables []interface{}) (*goPg.Model, error) {
	var err error
	mod := goPg.NewModel(true, true)

	err = mod.OpenWithDefault(role, database, password)
	if err != nil {
		return nil, err
	}

	//register new model
	err = mod.Register(
		table,
	)
	if err != nil {
		return nil, err
	}

	if tables != nil {
		for _, table := range tables {
			err = mod.Register(
				table,
			)
			if err != nil {
				continue
			}
		}

	}

	//migrate model
	err = mod.AutoMigrateAll()
	if err != nil {
		return nil, err
	}

	return mod, nil
}
