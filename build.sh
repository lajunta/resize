echo "--------------------------------------------------------------"
echo "|---------   First make 386 version for windows...  ---------|"
echo "--------------------------------------------------------------"

GOOS=windows GOARCH=386  CC=/usr/bin/i686-w64-mingw32-gcc CXX=/usr/bin/i686-w64-mingw32-g++ CGO_ENABLED=1 go  build -ldflags="-w -s " -o "dist/resize_32.exe" 

upx "dist/resize_32.exe"

# build amd64

echo "--------------------------------------------------------------"
echo "|--------- Second make amd64 version for windows... ---------|"
echo "--------------------------------------------------------------"

GOOS=windows GOARCH=amd64  CC=/usr/bin/x86_64-w64-mingw32-gcc CXX=/usr/bin/x86_64-w64-mingw32-g++ CGO_ENABLED=1 go  build  -ldflags="-w -s" -o "dist/resize_64.exe" 

upx "dist/resize_64.exe"

# build for linux 

echo "--------------------------------------------------------------"
echo "|------------------- Making linux binary ---------------------|"
echo "--------------------------------------------------------------"

GOOS=linux GOARCH=amd64  CGO_ENABLED=1  go build  -ldflags="-w -s" -o "dist/resize" 

upx "dist/resize"


cd dist

chmod 0755 resize