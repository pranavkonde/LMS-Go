package db

import (
	"context"
	"database/sql"
)

type User struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Gender    string `db:"gender"`
	Address   string `db:"address"`
	Age       int    `db:"age"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	MobileNum string `db:"mob_no"`
	Role      string `db:"role"`
}

const (
	createUserQuery = `INSERT INTO users (id,first_name, last_name, gender,age,address,email,password,mob_no,role)
    VALUES(?, ?,?,?,?,?,?,?,?,?)`
	listUsersQuery      = `SELECT * FROM users`
	findUserByIDQuery   = `SELECT * FROM users WHERE id = ?`
	deleteUserByIDQuery = `DELETE FROM users WHERE id = ?`
	updateUserQuery     = `UPDATE users SET first_name = ?, last_name=?, gender=?,age=?,address=?,password=?,mob_no = ? where id = ?`
)

func (s *store) CreateUser(ctx context.Context, user *User) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			user.ID,
			user.FirstName,
			user.LastName,
			user.Gender,
			user.Age,
			user.Address,
			user.Email,
			user.Password,
			user.MobileNum,
			user.Role,
		)
		return err
	})
}

func (s *store) ListUsers(ctx context.Context) (users []User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &users, listUsersQuery)
	})
	if err == sql.ErrNoRows {
		return users, ErrUserNotExist
	}
	return
}

func (s *store) FindUserByID(ctx context.Context, id string) (user User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByIDQuery, id)
	})
	if err == sql.ErrNoRows {
		return user, ErrUserNotExist
	}
	return
}

func (s *store) DeleteUserByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteUserByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrUserNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateUser(ctx context.Context, user *User) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateUserQuery,
			user.FirstName,
			user.LastName,
			user.Gender,
			user.Age,
			user.Address,
			user.Password,
			user.MobileNum,
			user.ID,
		)
		return err
	})
}
