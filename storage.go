package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) CreateTable() {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS games(
		Channel TEXT,
		Game TEXT
	);
	`

	_, err := r.db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func (r *Repository) MarkPublished(id string, game Game) {
	add := `
INSERT OR REPLACE INTO games(
Channel,
Game) VALUES (?,?)`

	stmt, err := r.db.Prepare(add)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, game.Title)
	if err != nil {
		panic(err)
	}
}

func (r *Repository) IsPublished(id string, game Game) bool {
	query := `SELECT * FROM games WHERE Channel = "` + id + `" AND Game = "` + game.Title + `"`
	rows, err := r.db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	return rows.Next()
}

func CreateRepository(storage string) (*Repository, error) {
	db, err := sql.Open("sqlite3", storage+"/games.sqlite")
	if err != nil {
		return nil, err
	}

	repo := Repository{
		db: db,
	}

	repo.CreateTable()
	return &repo, nil
}
