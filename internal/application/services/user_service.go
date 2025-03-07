package services

import (
	"bronze/internal/application/usecases"
	"bronze/internal/domain"
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/jung-kurt/gofpdf"
)

type UserService struct {
	repo usecases.UserRepository
}

func NewUserService(repo usecases.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SaveProduct(name, data, market string, value float64) error {
	return s.repo.SaveProduct(name, data, market, value)
}

func (s *UserService) SaveMarket(name string) error {
	if name == "" {
		err := errors.New("")
		return err
	}
	return s.repo.SaveMarket(name)
}

func (s *UserService) GetProducts() ([]domain.Product, error) {
	users, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetMarkets() ([]string, error) {
	markets, err := s.repo.GetAllMarkets()
	if err != nil {
		return nil, err
	}
	return markets, nil
}

func (s *UserService) GetProductsByFilter(filter string) ([]domain.Product, error) {
	users, err := s.repo.GetAllProductsbyFilter(filter)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GenerateReport(filters Filters) error {

	productsByData, err := s.GetProductsByFilter(filters.Data)
	if err != nil {
		return err
	}

	var products []domain.Product

	for _, product := range productsByData {
		if product.Market == filters.Market {
			products = append(products, product)
		}
	}

	createPDFDocument(filters.Market, products)

	return nil
}

func (s *UserService) GetUniqueDates() ([]string, error) {
	return s.repo.GetUniqueDates()
}

func createPDFDocument(market string, products []domain.Product) {

	var fileName string

	if runtime.GOOS == "windows" {
		fileName = "reports/report.pdf"
	} else {
		fileName = "./temp/report.pdf"
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(100, 10, "Lista de Produtos")
	pdf.Cell(100, 10, market)
	pdf.Ln(20)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(66, 10, "Produto")
	pdf.Cell(66, 10, "Valor")
	pdf.Cell(66, 10, "Mes/Ano")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)

	for _, product := range products {
		pdf.Cell(66, 10, product.Name)
		pdf.Cell(66, 10, fmt.Sprintf("%.2f", product.Value))
		pdf.Cell(66, 10, product.Data)
		pdf.Ln(10)
	}

	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		log.Printf("erro ao criar o PDF: %v", err)
	}

}
