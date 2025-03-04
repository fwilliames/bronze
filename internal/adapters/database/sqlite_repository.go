package database

import (
	"bronze/internal/application/usecases"
	"bronze/internal/domain"
	"database/sql"
	"os"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

// Retorna o caminho do banco no APPDATA
func getDBPath() string {
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		return appData + "\\CadastroUsuarios\\users.db"
	} else {
		return "./temp/users.db"
	}
}

type SQLiteRepository struct {
	db *sql.DB
}

// NewSQLiteRepository cria a conexão e inicializa o banco
func NewSQLiteRepository() (usecases.UserRepository, error) {

	dbPath := getDBPath()

	// Cria o diretório caso não exista
	if runtime.GOOS == "windows" {
		os.MkdirAll(os.Getenv("APPDATA")+"\\CadastroUsuarios", os.ModePerm)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
	if err != nil {
		return nil, err
	}

	return &SQLiteRepository{db: db}, nil
}

// SaveUser insere um usuário no banco
func (r *SQLiteRepository) SaveUser(name string, age int) error {
	_, err := r.db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
	return err
}

// GetAllUsers retorna todos os usuários cadastrados no banco
func (r *SQLiteRepository) GetAllUsers() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
