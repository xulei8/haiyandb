all:
	go build -o app  main.go 
run:
	go build -o app  main.go
	./app
clean:
	rm app -f
push:
	rm app -f
	git push
