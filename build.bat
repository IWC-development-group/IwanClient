@echo off

set "OUTPUT_DIR=./out/"
mkdir "%OUTPUT_DIR%" 2>nul
go build -o "%OUTPUT_DIR%iwan.exe" ./src/ && echo Build completed!