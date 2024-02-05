setup:
	go run cmd/gen_key.go

run:
	go run cmd/api/main.go

build:
	go build cmd/api/main.go

doc:
	touch doc
	echo "welcome" > doc
	echo "Done."

config:
	mkdir -p Config/
	touch Config/config.json
	echo "Done.""