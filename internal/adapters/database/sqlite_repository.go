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
		return appData + "\\SuperMarketTracker\\products.db"
	} else {
		return "./temp/products.db"
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
		os.MkdirAll(os.Getenv("APPDATA")+"\\SuperMarketTracker", os.ModePerm)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Criação das tabelas
	statements := []string{
		`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY, 
		name TEXT, 
		data TEXT, 
		value FLOAT,
		market TEXT
	)`,
		`CREATE TABLE IF NOT EXISTS markets (
		id INTEGER PRIMARY KEY, 
		name TEXT UNIQUE
	)`,
	}

	// Executa cada comando de criação de tabela
	for _, stmt := range statements {
		if _, err := db.Exec(stmt); err != nil {
			db.Close() // Fecha a conexão se houver erro
			return nil, err
		}
	}

	return &SQLiteRepository{db: db}, nil
}

// SaveUser insere um usuário no banco
func (r *SQLiteRepository) SaveProduct(name, data, market string, value float64) error {
	_, err := r.db.Exec("INSERT INTO products (name, data, value, market) VALUES (?, ?, ?, ?)", name, data, value, market)
	return err
}

func (r *SQLiteRepository) SaveMarket(name string) error {
	_, err := r.db.Exec("INSERT INTO markets (name) VALUES (?)", name)
	return err
}

// GetAllUsers retorna todos os usuários cadastrados no banco
func (r *SQLiteRepository) GetAllProducts() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, name, data, value FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Data, &product.Value); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// GetAllUsers retorna todos os usuários cadastrados no banco
func (r *SQLiteRepository) GetAllProductsbyFilter(filter string) ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT * FROM products WHERE data = ?", filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Data, &product.Value, &product.Market); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// GetAllUsers retorna todos os usuários cadastrados no banco
func (r *SQLiteRepository) GetAllMarkets() ([]string, error) {
	rows, err := r.db.Query("SELECT DISTINCT name FROM markets ORDER BY name DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []string
	for rows.Next() {
		var market string
		if err := rows.Scan(&market); err != nil {
			return nil, err
		}
		markets = append(markets, market)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return markets, nil
}

func (r *SQLiteRepository) GetUniqueDates() ([]string, error) {
	rows, err := r.db.Query("SELECT DISTINCT data FROM products ORDER BY data DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datas []string
	for rows.Next() {
		var data string
		if err := rows.Scan(&data); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, nil
}
