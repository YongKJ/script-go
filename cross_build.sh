# windows build

export GOARCH=amd64
export GOOS=windows
export CGO_ENABLED=1

rm -rf ./dist

# export CC=D:/Software/scoop/apps/gcc/current/bin/gcc.exe
# export CXX=D:/Software/scoop/apps/gcc/current/bin/g++.exe

go build -o ./dist/script-go.exe

# linux build

#export GOARCH=amd64
#export GOOS=linux
#export CGO_ENABLED=1

#rm -rf ./dist

# export CC=D:/Software/scoop/apps/mingw/current/bin/x86_64-w64-mingw32-gcc.exe
# export CXX=D:/Software/scoop/apps/mingw/current/bin/x86_64-w64-mingw32-g++.exe

#go build -o ./dist/script-go

# android 9 build

# export GOARCH=arm64
# export GOOS=android
# export CGO_ENABLED=1

# rm -rf ./dist

# export ANDROID_API=android28
# export CC=D:/Software/Android/Sdk/ndk/27.0.11902837/toolchains/llvm/prebuilt/windows-x86_64/bin/aarch64-linux-%ANDROID_API%-clang.cmd
# export CXX=D:/Software/Android/Sdk/ndk/27.0.11902837/toolchains/llvm/prebuilt/windows-x86_64/bin/aarch64-linux-%ANDROID_API%-clang++.cmd

# go build -o ./dist/script-go
