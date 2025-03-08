package database

import (
	"bronze/internal/application/usecases"
	"bronze/internal/domain"
	"database/sql"
	"os"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

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

func NewSQLiteRepository() (usecases.UserRepository, error) {

	dbPath := getDBPath()

	if runtime.GOOS == "windows" {
		os.MkdirAll(os.Getenv("APPDATA")+"\\SuperMarketTracker", os.ModePerm)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	statements := []string{
		`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY, 
		name TEXT, 
		data TEXT, 
		value FLOAT,
		market TEXT,
		quantity INT,
		totalValue FLOAT
	)`,
		`CREATE TABLE IF NOT EXISTS markets (
		id INTEGER PRIMARY KEY, 
		name TEXT UNIQUE
	)`,
		`CREATE TABLE IF NOT EXISTS datas (
		id INTEGER PRIMARY KEY, 
		data TEXT UNIQUE
	)`,
	}

	for _, stmt := range statements {
		if _, err := db.Exec(stmt); err != nil {
			db.Close()
			return nil, err
		}
	}

	return &SQLiteRepository{db: db}, nil
}

func (r *SQLiteRepository) SaveProduct(name, data, market string, value, totalValue float64, quantity int64) error {
	_, err := r.db.Exec("INSERT INTO products (name, data, value, market, quantity, totalValue) VALUES (?, ?, ?, ?, ?, ?)", name, data, value, market, quantity, totalValue)
	return err
}

func (r *SQLiteRepository) SaveMarket(name string) error {
	_, err := r.db.Exec("INSERT INTO markets (name) VALUES (?)", name)
	return err
}

func (r *SQLiteRepository) SaveData(data string) error {
	_, err := r.db.Exec("INSERT INTO datas (data) VALUES (?)", data)
	return err
}

func (r *SQLiteRepository) GetAllProducts() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, name, data, market, quantity, totalValue value FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Data, &product.Value, &product.Market, &product.Quantity, &product.TotalValue); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *SQLiteRepository) GetAllProductsbyFilter(filter string) ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT * FROM products WHERE data = ?", filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Data, &product.Value, &product.Market, &product.Quantity, &product.TotalValue); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

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

func (r *SQLiteRepository) GetAllDates() ([]domain.Data, error) {
	rows, err := r.db.Query("SELECT DISTINCT * FROM datas ORDER BY data DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datas []domain.Data
	for rows.Next() {
		var data domain.Data
		if err := rows.Scan(&data.ID, &data.Name); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return datas, nil
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
