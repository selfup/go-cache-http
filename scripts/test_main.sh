# coverage will never be 100% because of main()
# as well as forcing json.Marshall to fail (unclear how to do that)
go test -coverprofile=coverage.out \
&& go tool cover -html=coverage.out