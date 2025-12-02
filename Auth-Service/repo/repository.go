package repo

import (
	"Auth-Service/interfaces"
	"Auth-Service/service/jwt"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sqlx.DB
	ctx context.Context
}

func NewUserRepository(db *sqlx.DB) interfaces.AuthRepository {
	return &UserRepository{
		db: db,
		ctx: context.Background(),
	}
}

func (r *UserRepository) CreateUser(name string, email string, password string) (interfaces.User, string, error) {
	if r == nil || r.db == nil {
		return interfaces.User{}, "", fmt.Errorf("repository or database is not initialized")
	}
	exists, err := r.GetUserByEmail(email)
	if err != nil {
		return interfaces.User{}, "", err
	}
	if exists {
		return interfaces.User{}, email, fmt.Errorf("user already exists with this email")
	}
	
	hashPassword := HashPassword(password)
	user := interfaces.User{
		Name:     name,
		Email:    email,
		Password: hashPassword,
	}

	query := `
        INSERT INTO users (name, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, password
    `
	ctx := r.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	if err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return interfaces.User{}, "", fmt.Errorf("user already exists with this email")
		}
		return interfaces.User{}, "", err
	}
	
	jwtService := jwt.JwtServices{}
	tokenData, err := jwtService.GenerateToken(user.Email)
	if err != nil {
		return interfaces.User{},"",err
	}

	return user, tokenData.Token, nil
}

func (r *UserRepository) GetUserByEmail(email string) (bool, error) {
	if r == nil || r.db == nil {
		return false, fmt.Errorf("repository or database is not initialized")
	}
	query := `
		SELECT email
		FROM users
		WHERE email = $1;
	`
	var userEmail string
	ctx := r.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	if err := r.db.GetContext(ctx, &userEmail, query, email); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func HashPassword(password string) (string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password
	}
	return string(bytes)
}

func (r *UserRepository) LoginUser(email string, password string) (interfaces.User, error) {
	ctx := r.ctx
    if ctx == nil {
        ctx = context.Background()
    }

    if r.db == nil {
        return interfaces.User{}, fmt.Errorf("db connection is nil")
    }

    query := `
        SELECT id, name, email, password
        FROM users
        WHERE email = $1;
    `
    var user interfaces.User
    if err := r.db.GetContext(ctx, &user, query, email); err != nil {
        if err == sql.ErrNoRows {
            return interfaces.User{}, fmt.Errorf("invalid credentials")
        }
        return interfaces.User{}, err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return interfaces.User{}, fmt.Errorf("invalid password")
    }

    return user, nil
}
