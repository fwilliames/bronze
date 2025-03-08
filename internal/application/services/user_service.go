package services

import (
	"bronze/internal/application/usecases"
	"bronze/internal/config/utils"
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

func (s *UserService) SaveProduct(name, data, market string, value, totalValue float64, quantity int64) error {
	return s.repo.SaveProduct(name, data, market, value, totalValue, quantity)
}

func (s *UserService) SaveMarket(name string) error {
	if name == "" {
		err := errors.New("")
		return err
	}
	return s.repo.SaveMarket(name)
}

func (s *UserService) SaveData(data string) error {
	if data == "" {
		err := errors.New("")
		return err
	}
	return s.repo.SaveData(data)
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

func (s *UserService) GetDates() ([]string, error) {
	rawDatas, err := s.repo.GetAllDates()
	if err != nil {
		return nil, err
	}

	var datas = make([]string, len(rawDatas))
	for i, rawData := range rawDatas {
		datas[i] = rawData.Name
	}
	return datas, nil
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

	createPDFDocument(filters.Market, filters.Data, products)

	return nil
}

func (s *UserService) GetUniqueDates() ([]string, error) {
	return s.repo.GetUniqueDates()
}

func createPDFDocument(market, data string, products []domain.Product) {

	var fileName string

	if runtime.GOOS == "windows" {
		fileName = "reports/report.pdf"
	} else {
		fileName = "./temp/report.pdf"
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(66, 10, "Lista de Produtos")
	pdf.Cell(66, 10, market)
	pdf.Cell(66, 10, data)
	pdf.Ln(20)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(66, 10, "Produto")
	pdf.Cell(66, 10, "Quantidade")
	pdf.Cell(66, 10, "Valor")
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(200, 10, "________________________________________________________________________________")
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 12)

	for _, product := range products {
		pdf.Cell(77, 10, product.Name)
		pdf.Cell(55, 10, fmt.Sprintf("%d", product.Quantity))
		pdf.Cell(66, 10, fmt.Sprintf("%.2f", product.TotalValue))
		pdf.Ln(7)
	}

	var values = make([]float64, len(products))
	for i, product := range products {
		values[i] = product.TotalValue
	}

	totalValue := utils.Sum(values)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(200, 10, "________________________________________________________________________________")
	pdf.Ln(5)
	pdf.Cell(66, 10, "Total")
	pdf.Cell(66, 10, "")
	pdf.Cell(66, 10, fmt.Sprintf("%.2f", totalValue))
	pdf.Ln(10)

	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		log.Printf("erro ao criar o PDF: %v", err)
	}

}
