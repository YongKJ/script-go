:: windows build

:: SET GOARCH=amd64
:: SET GOOS=windows
:: SET CGO_ENABLED=1

:: rm -rf D:/Document/MyCodes/Worker/Go/script-go/dist
:: cd D:/Document/MyCodes/Worker/Go/script-go

:: SET CC=D:\Software\scoop\apps\gcc\current\bin\gcc.exe
:: SET CXX=D:\Software\scoop\apps\gcc\current\bin\g++.exe

:: go build -o ./dist/script-go.exe

:: linux build

SET GOARCH=amd64
SET GOOS=linux
SET CGO_ENABLED=1

rm -rf D:/Document/MyCodes/Worker/Go/script-go/dist
cd D:/Document/MyCodes/Worker/Go/script-go

:: SET CC=D:\Software\scoop\apps\mingw\current\bin\x86_64-w64-mingw32-gcc.exe
:: SET CXX=D:\Software\scoop\apps\mingw\current\bin\x86_64-w64-mingw32-g++.exe

go build -o ./dist/script-go

:: android 9 build

:: SET GOARCH=arm64
:: SET GOOS=android
:: SET CGO_ENABLED=1

:: rm -rf D:/Document/MyCodes/Worker/Go/script-go/dist
:: cd D:/Document/MyCodes/Worker/Go/script-go

:: SET ANDROID_API=android28
:: SET CC=D:\Software\Android\Sdk\ndk\27.0.11902837\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-%ANDROID_API%-clang.cmd
:: SET CXX=D:\Software\Android\Sdk\ndk\27.0.11902837\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-%ANDROID_API%-clang++.cmd

:: go build -o ./dist/script-go
