go.build:
	rm -rf ./build
	mkdir build
	go build -o ./build .

docker.build:
	cd dist
	docker build

push:
	git add .
	git commit -m "$m"
	git push