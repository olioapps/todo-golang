package db

import (
	log "github.com/Sirupsen/logrus"
	olioDAO "github.com/rachoac/service-skeleton-go/olio/dao"
	olioDb "github.com/rachoac/service-skeleton-go/olio/db"
)

type MigrationsContainer struct {
	connectionManager *olioDAO.ConnectionManager
}

func NewMigrationsContainer(connectionManager *olioDAO.ConnectionManager) MigrationsContainer {
	return MigrationsContainer{connectionManager}
}

func (m MigrationsContainer) GetMigrations() []olioDb.Migration {
	var migrations []olioDb.Migration = []olioDb.Migration{
		m.v1(),
		m.v2(),
		m.v3(),
	}
	return migrations
}

func (m MigrationsContainer) v1() olioDb.Migration {
	return func() error {
		db := m.connectionManager.GetDb()

		log.Info("Create table todo items")
		if err := db.Exec(`
            CREATE TABLE IF NOT EXISTS
            todo_items (
                id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
				name varchar(255) COLLATE utf8_unicode_ci NOT NULL,
                created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
        `).Error; err != nil {
			return err
		}
		return nil
	}
}

func (m MigrationsContainer) v2() olioDb.Migration {
	return func() error {
		db := m.connectionManager.GetDb()

		log.Info("Create table todo list")

		if err := db.Exec(`
		    CREATE TABLE IF NOT EXISTS
			todo_lists (
				id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
				name varchar(255) COLLATE utf8_unicode_ci NOT NULL,
                created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
		`).Error; err != nil {
			return err
		}
		return nil
	}
}

func (m MigrationsContainer) v3() olioDb.Migration {
	return func() error {
		db := m.connectionManager.GetDb()

		log.Info("Add todo list foreign key")

		if err := db.Exec(`
		     	ALTER TABLE todo_items ADD COLUMN todo_list_id bigint(20) unsigned,
				ADD FOREIGN KEY (todo_list_id) REFERENCES todo_lists(id)
		`).Error; err != nil {
			return err
		}
		return nil
	}
}
