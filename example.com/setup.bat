@echo off

set GOOS=linux
cd copyobject
go build -ldflags="-s -w" -o ..\bin\copyobject\main main.go
cd ..\encrypt
go build -ldflags="-s -w" -o ..\bin\encrypt\main main.go
cd ..\property
go build -ldflags="-s -w" -o ..\bin\property\main main.go
cd ..\sign
go build -ldflags="-s -w" -o ..\bin\sign\main main.go
cd ..\watermark
go build -ldflags="-s -w" -o ..\bin\watermark\main main.go
cd %~dp0
copy sign\TanukiMagic.ttf bin\sign\