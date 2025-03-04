[Setup]
AppName=Cadastro de Usuários
AppVersion=1.0
DefaultDirName={localappdata}\CadastroUsuarios
DefaultGroupName=CadastroUsuarios
OutputDir=/build
OutputBaseFilename=CadastroSetup
Compression=lzma
SolidCompression=yes

[Files]
Source: "/build/cadastro.exe"; DestDir: "{app}"
Source: "users.db"; DestDir: "{app}"

[Icons]
Name: "{group}\Cadastro de Usuários"; Filename: "{app}\cadastro.exe"

[Run]
Filename: "{app}\cadastro.exe"; Description: "Executar o Cadastro"; Flags: nowait postinstall

[UninstallDelete]
Type: files; Name: "{userappdata}\CadastroUsuarios\usuarios.db";