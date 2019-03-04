package main

import (
	"database/sql"
	"fmt"
)

type user struct {
	ID        int    `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Age       int    `json:"age,omitempty"`
}

func (u *user) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT firstname, lastname, age FROM users WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Firstname, &u.Lastname, &u.Age)
}

func (u *user) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET firstname='%s', lastname='%s', age=%d WHERE id=%d", u.Firstname, u.Lastname, u.Age, u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}
func (u *user) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users(id, firstname, lastname, age) VALUES('%d', '%s', '%s', '%d')", u.ID, u.Firstname, u.Lastname, u.Age)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := fmt.Sprintf("SELECT id, firstname, lastname, age FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []user{}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
