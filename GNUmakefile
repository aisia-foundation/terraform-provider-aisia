default: testacc

build:
	go build -o terraform-provider-aisia

fmt:
	gofmt -s -w .

validate:
	go vet ./...
	go build ./...

docs:
	tfplugindocs generate

# Tests d'acceptation (nécessitent une instance AISIA + AISIA_TOKEN ; TF_ACC=1).
testacc:
	TF_ACC=1 go test ./... -v -timeout 120m
