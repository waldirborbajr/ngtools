build:
	go build -v -a -ldflags "-X 'localhost/ngtools/cmd.version=v0.1.0'" -o bin/ngtools main.go

run: build
	./bin/ngtools start -protocol http -port 3031
