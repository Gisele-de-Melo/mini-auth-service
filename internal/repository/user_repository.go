package repository

import (
	"database/sql"
	"errors"

	"mini-auth-service/internal/model"

	_ "modernc.org/sqlite"
)

type UserRepository struct {
	db *sql.DB
}

func NewSQLiteDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created_at DATETIME NOT NULL
    );
    `

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	query := `
        INSERT INTO users (email, password, created_at)
        VALUES (?, ?, ?)
    `

	_, err := r.db.Exec(query, user.Email, user.Password, user.CreatedAt)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	query := `
        SELECT id, email, password, created_at
        FROM users
        WHERE email = ?
    `

	row := r.db.QueryRow(query, email)

	var u model.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindByID(id int) (*model.User, error) {
	query := `
        SELECT id, email, password, created_at
        FROM users
        WHERE id = ?
    `

	row := r.db.QueryRow(query, id)

	var u model.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
