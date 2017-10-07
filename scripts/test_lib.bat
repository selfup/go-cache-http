cd lib
go test -coverprofile=coverage.out
cd ..
go tool cover -html=./lib/coverage.out
