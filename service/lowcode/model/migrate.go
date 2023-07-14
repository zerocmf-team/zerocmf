package model

import "zerocmf/common/bootstrap/database"

type Lowcode struct {
}

func (l Lowcode) Migrate(db database.MongoDB) (err error) {
	err = new(Theme).Migrate(db)
	if err != nil {
		return err
	}

	err = new(ThemePage).Migrate(db)
	if err != nil {
		return err
	}

	err = new(Settings).Migrate(db)
	if err != nil {
		return err
	}
	return
}
