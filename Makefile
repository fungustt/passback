.PHONY: docker_build
docker_build:
	docker build -f ./build/package/Dockerfile -t pass_backend .

.PHONY: docker_run
docker_run:
	docker run -it pass_backend

.PHONY: env_up
env_up: env_down mod_vendor
	docker-compose -f build/package/docker-compose.yml -p dev pull
	docker-compose -f build/package/docker-compose.yml -p dev build passback
	docker-compose -f build/package/docker-compose.yml -p dev up -d passback

.PHONY: env_down
env_down:
	docker-compose -f build/package/docker-compose.yml -p dev down -v --remove-orphans


mod_vendor:
	go mod vendor

generate:
	go generate -v ./...