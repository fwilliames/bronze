[Setup]
AppName=SuperMarketTracker
AppVersion=1.0
DefaultDirName={localappdata}\SuperMarketTracker
DefaultGroupName=SuperMarketTracker
OutputDir=./
OutputBaseFilename=SuperMarketSetup
Compression=lzma
SolidCompression=yes
SetupIconFile=./setupicon.ico
UninstallIconFile=./uninstallicon.ico

[Files]
Source: "superMarket.exe"; DestDir: "{app}"
Source: "../icon.ico"; DestDir: "{app}"

[Dirs]
Name: "{app}\reports"; Permissions: everyone-full

[Icons]
Name: "{autoprograms}\SuperMarketTracker"; Filename: "{app}\superMarket.exe"; IconFilename: "{app}\icon.ico"
Name: "{autodesktop}\SuperMarketTracker"; Filename: "{app}\superMarket.exe"; IconFilename: "{app}\icon.ico" ;Tasks: desktopicon

[Tasks]
Name: "desktopicon"; Description: "Criar ícone na área de trabalho"; GroupDescription: "Tarefas adicionais"

[Run]
Filename: "{app}\superMarket.exe"; Description: "Executar o Guia SuperMarket"; Flags: nowait postinstall

[UninstallDelete]
Type: files; Name: "{userappdata}\SuperMarketTracker\products.db";
Type: files; Name: "{userappdata}\SuperMarketTracker\reports\report";
Type: files; Name: "{userappdata}\SuperMarketTracker";