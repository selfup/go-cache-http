cd handlers \
  && go test -coverprofile=coverage.out \
  && cd ../ \
  && go tool cover -html=./handlers/coverage.out
