[Setup]
AppName=Guia SuperMarket
AppVersion=1.0
DefaultDirName={localappdata}\GuiaSuperMarket
DefaultGroupName=GuiaSuperMarket
OutputDir=./
OutputBaseFilename=SuperMarketSetup
Compression=lzma
SolidCompression=yes

[Files]
Source: "superMarket.exe"; DestDir: "{app}"
Source: "../products.db"; DestDir: "{app}"

[Dirs]
Name: "{app}\reports"; Permissions: everyone-full

[Icons]
Name: "{group}\Guia SuperMarket"; Filename: "{app}\guiaSuperMarket.exe"

[Run]
Filename: "{app}\superMarket.exe"; Description: "Executar o Guia SuperMarket"; Flags: nowait postinstall

[UninstallDelete]
Type: files; Name: "{userappdata}\GuiaSuperMarket\products.db";
Type: files; Name: "{userappdata}\GuiaSuperMarket\GuiaSuperMarket\reports\report";
Type: files; Name: "{userappdata}\GuiaSuperMarket";