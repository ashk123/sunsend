setup: config
	go run cmd/gen_key.go

run:
	go run cmd/api/main.go

build:
	go build cmd/api/main.go

doc:
	touch doc
	echo "welcome" > doc
	echo "Done."

restart:
	rm -rf Storage/SunSend.db
	# And remove other configuration files
config:
	mkdir -p Config/
	touch .env
	touch Config/config.json
	echo "Done.""
