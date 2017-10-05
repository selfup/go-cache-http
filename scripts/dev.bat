IF EXIST gin (
    gin run main.go 
) ELSE (
    go get github.com/codegangsta/gin
    gin run main.go
)
