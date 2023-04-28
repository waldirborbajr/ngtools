bin:
	go build -o bin/ngtools main.go

run: bin
	./bin/ngtools start -protocol http -port 3031
