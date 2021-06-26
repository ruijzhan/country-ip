GO=go

.PHONY: test
test:
	@$(GO) test -v ./...

.PHONY: data
data:
	@curl https://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest -o apnic.txt
	@go-bindata -pkg country_ip -o data.go apnic.txt
	@rm apnic.txt 
