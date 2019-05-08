package repository

import (
	"database/sql"
	"log"

	"example.local/go-pilot/internal/models"
	"example.local/go-pilot/internal/user"
)

type postgresUserRepository struct {
	Conn *sql.DB
}

// NewPostgresUserRepository will create new user postgres repository
func NewPostgresUserRepository(Conn *sql.DB) user.Repository {
	return &postgresUserRepository{Conn}
}

// Fetch will extract users from database
func (p *postgresUserRepository) Fetch() ([]*models.User, error) {
	rows, err := p.Conn.Query(`
		SELECT
			uid,
			name,
			age,
			address,
			salary
		FROM users.users
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(
			&user.UID,
			&user.Name,
			&user.Age,
			&user.Address,
			&user.Salary,
		)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, &user)
	}

	return users, nil
}

// Create will create new user
func (p *postgresUserRepository) Create(m *models.User) error {
	statement := `
		INSERT INTO users.users (
			name, age, address, salary
		) VALUES (
			$1, $2, $3, $4
		)
	`
	_, err := p.Conn.Exec(statement, m.Name, m.Age, m.Address, m.Salary)

	if err != nil {
		return err
	}

	return nil
}

// Update will update user fields
func (p *postgresUserRepository) Update(uid int, m *models.User) error {
	statement := `
		UPDATE users.users
		SET
			name = $2,
			age = $3,
			address = $4,
			salary = $5
		WHERE
			uid = $1
	`

	_, err := p.Conn.Exec(statement, uid, m.Name, m.Age, m.Address, m.Salary)

	if err != nil {
		return err
	}

	return nil
}

// Delete will remove user entity from database based on uid
func (p *postgresUserRepository) Delete(uid int) error {
	statement := `
	DELETE from users.users
	WHERE
		uid = $1
	`
	_, err := p.Conn.Exec(statement, uid)
	if err != nil {
		return err
	}

	return nil
}
