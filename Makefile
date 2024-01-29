run:
	go run cmd/api/main.go

build:
	go build cmd/api/main.go

doc:
	touch doc
	echo "welcome" > doc
	echo "Done."