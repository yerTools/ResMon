// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: migrationscript.sql

package model

import (
	"context"
)

const insertMigrationScript = `-- name: InsertMigrationScript :exec
INSERT OR IGNORE INTO
    "migration_script" ("version", "identifier", "up", "down")
VALUES
    (?1, ?2, ?3, ?4)
`

type InsertMigrationScriptParams struct {
	Version    int64  `db:"version" json:"version"`
	Identifier string `db:"identifier" json:"identifier"`
	Up         string `db:"up" json:"up"`
	Down       string `db:"down" json:"down"`
}

// InsertMigrationScript
//
//	INSERT OR IGNORE INTO
//	    "migration_script" ("version", "identifier", "up", "down")
//	VALUES
//	    (?1, ?2, ?3, ?4)
func (q *Queries) InsertMigrationScript(ctx context.Context, arg InsertMigrationScriptParams) error {
	_, err := q.exec(ctx, q.insertMigrationScriptStmt, insertMigrationScript,
		arg.Version,
		arg.Identifier,
		arg.Up,
		arg.Down,
	)
	return err
}

const migrationScripts = `-- name: MigrationScripts :many
SELECT
    "version"   ,
    "identifier",
    "up"        ,
    "down"
FROM
    "migration_script"
`

// MigrationScripts
//
//	SELECT
//	    "version"   ,
//	    "identifier",
//	    "up"        ,
//	    "down"
//	FROM
//	    "migration_script"
func (q *Queries) MigrationScripts(ctx context.Context) ([]MigrationScript, error) {
	rows, err := q.query(ctx, q.migrationScriptsStmt, migrationScripts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MigrationScript
	for rows.Next() {
		var i MigrationScript
		if err := rows.Scan(
			&i.Version,
			&i.Identifier,
			&i.Up,
			&i.Down,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
