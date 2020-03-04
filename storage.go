package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) CreateTable() error {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS games(
		Channel TEXT,
		Game TEXT,
		GameID TEXT,
		Store TEXT
	);
	`

	_, err := r.db.Exec(sql_table)
	if err != nil {
		log.WithError(err).Fatal("Could not create database")
	}
	return err
}

func (r *Repository) MarkPublished(id , title, gameID, store string) error {
	add := `
INSERT OR REPLACE INTO games(
Channel,
Game,
GameID,
Store) VALUES (?,?,?,?)`

	stmt, err := r.db.Prepare(add)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, title, gameID, store)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"channel": id, "game": title}).Error("Could not mark game as published")
	}

	return err
}

func (r *Repository) IsPublished(id, title, store string) (bool, error) {
	query := `SELECT * FROM games WHERE Channel = "` + id + `" AND Store = "` + store + `" AND game = "` + title + `"`
	rows, err := r.db.Query(query)
	if err != nil {
		log.WithError(err).Error("Could not get games from DB")
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

func CreateRepository(storage string) (*Repository, error) {
	db, err := sql.Open("sqlite3", storage+"/games.sqlite")
	if err != nil {
		return nil, err
	}

	repo := Repository{
		db: db,
	}

	err = repo.CreateTable()
	return &repo, err
}
