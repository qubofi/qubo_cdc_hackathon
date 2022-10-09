package dal

import (
	"qubo/qubo-backend/domain/entity"
	"qubo/qubo-backend/infrastructure/database"
)

func InsertIntoUsers(user *entity.User) (*entity.User, error) {
	println(user.Password)
	err := database.Db.QueryRow(`INSERT INTO "users"(first_name, last_name, email, password_hash) VALUES($1, $2, $3, $4) RETURNING id`,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func SelectFromUsersByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := database.Db.QueryRow(`SELECT id, first_name, last_name, email, password_hash FROM "users" WHERE email = $1`, email).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func SelectFromUsersById(id uint64) (*entity.User, error) {
	user := &entity.User{}
	err := database.Db.QueryRow(`SELECT id, first_name, last_name, email, password_hash FROM "users" WHERE id = $1`, id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
