set -x
GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -a -ldflags "-extldflags \"-static\"" -o "randomclass.exe"
GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -ldflags "-extldflags \"-static\"" -o "randomclass-linux"
GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 go build -a -ldflags "-extldflags \"-static\"" -o "randomclass-darwin"
