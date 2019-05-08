package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"example.local/users/internal/user/delivery"
	"example.local/users/internal/user/repository"
	"example.local/users/internal/user/usecase"
)

func main() {
	connStr := PgConnStr()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewPostgresUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	delivery.NewUserHandler(usecase)

	http.ListenAndServe(":8080", nil)
}
