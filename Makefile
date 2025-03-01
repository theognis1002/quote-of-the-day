build:
	@echo "Building and starting containers with docker compose..."
	docker compose up --build

run:
	@echo "Starting containers with docker compose..."
	docker compose up

stop:
	@echo "Stopping containers..."
	docker compose down

destroy:
	@echo "Stopping containers and removing volumes..."
	docker compose down -v

destroy-hard: stop destroy
	@echo "Performing hard destroy - removing all containers and volumes..."
	docker compose rm -v
	docker rm -f $(docker ps -a -q)
	docker volume rm $(docker volume ls -q)

rebuild: destroy build
r: rebuild

nuke:
	@echo "WARNING: Nuclear option - removing all Docker containers, images, networks and volumes..."
	docker stop $$(docker ps -a -q)
	docker rm $$(docker ps -a -q)
	docker network prune -f
	docker rmi -f $$(docker images --filter dangling=true -qa)
	docker volume rm $$(docker volume ls --filter dangling=true -q)
	docker rmi -f $$(docker images -a -q)
	docker system prune -a

clear-cache:
	@echo "Clearing quote cache..."
	curl -X POST http://localhost:8080/clear-cache
