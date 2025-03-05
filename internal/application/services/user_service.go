package services

import (
	"bronze/internal/application/usecases"
	"bronze/internal/domain"
	"fmt"
	"runtime"

	"github.com/jung-kurt/gofpdf"
)

// UserService implementa os casos de uso
type UserService struct {
	repo usecases.UserRepository
}

// NewUserService cria uma nova instância do serviço
func NewUserService(repo usecases.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// SaveUser salva um usuário no banco
func (s *UserService) SaveProduct(name, data string, value float64) error {
	return s.repo.SaveProduct(name, data, value)
}

// GetUsers retorna a lista de usuários cadastrados
func (s *UserService) GetProducts() ([]domain.Product, error) {
	users, err := s.repo.GetAllProducts() // Este método deve ser implementado no seu repositório
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GenerateReport() error {
	println("Generate Report")
	products, err := s.GetProducts()
	if err != nil {
		return err
	}

	var fileName string

	if runtime.GOOS == "windows" {
		fileName = "reports/report.pdf"
	} else {
		fileName = "./temp/report.pdf"
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Título principal
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(200, 10, "Lista de Produtos")
	pdf.Ln(20) // Pula uma linha

	// Títulos das colunas
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(66, 10, "Produto")
	pdf.Cell(66, 10, "Valor")
	pdf.Cell(66, 10, "Mes/Ano")
	pdf.Ln(10) // Pula uma linha

	// Definir a fonte para os dados
	pdf.SetFont("Arial", "", 12)

	for _, product := range products {
		pdf.Cell(66, 10, product.Name)
		pdf.Cell(66, 10, fmt.Sprintf("%.2f", product.Value))
		pdf.Cell(66, 10, product.Data)
		pdf.Ln(10)
	}

	err = pdf.OutputFileAndClose(fileName)
	if err != nil {
		return fmt.Errorf("erro ao criar o PDF: %v", err)
	}

	return nil
}
