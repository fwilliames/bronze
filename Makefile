BUILD_DIR=build
BINARY_NAME=cadastro.exe
SOURCE=./cmd/main.go
INNOSETUP=wine 'C:/Program Files (x86)/Inno Setup 6/ISCC.exe'
SETUP_FILE=setup.iss
GOOS=GOOS=windows
GOARCH=GOARCH=amd64
CGO_ENABLED=CGO_ENABLED=1
CC=CC=x86_64-w64-mingw32-gcc
CXX= CXX=x86_64-w64-mingw32-g++
FLAGS=-ldflags
WINDOWSGUI="-H windowsgui"
RELEASE_NAME=CadastroSetup.exe

# Regra para compilar o código para Windows
build: FORCE
	mkdir -p $(BUILD_DIR)
	$(GOOS) $(GOARCH) $(CGO_ENABLED) $(CC) $(CXX) go build $(FLAGS) $(WINDOWSGUI) -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE)

# Regra para criar o instalador (Apenas se tiver o Innosteup e wine instalados)
setup: 
	$(INNOSETUP) $(SETUP_FILE)

# Regra para compilar o binário e gerar o instalador
release: build setup

# Build para Linux
gobuild:
	go build cmd/main.go
	go run cmd/main.go

goclean:
	rm -rf \CadastroUsuarios
	rm -r \CadastroUsuarios\users.db

# Regra para limpar os arquivos gerados
clean:
	rm -rf $(BUILD_DIR)/$(BINARY_NAME)
	rm -rf $(BUILD_DIR)/$(RELEASE_NAME)

FORCE: