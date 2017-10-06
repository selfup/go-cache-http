gin_version=$(gin -v | grep version)

if [[ $gin_version ]]
then
  gin run main.go 
else
  go get github.com/codegangsta/gin
  gin run main.go
fi
