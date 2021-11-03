build:
	docker compose up --build

run:
	docker compose up

stop:
	docker compose down

destroy:
	docker compose down -v

destroy-hard: stop destroy
	docker compose rm -v
	docker rm -f $(docker ps -a -q)
	docker volume rm $(docker volume ls -q)

rebuild: destroy build
r: rebuild

nuke:
	docker stop $$(docker ps -a -q)
	docker rm $$(docker ps -a -q)
	docker network prune -f
	docker rmi -f $$(docker images --filter dangling=true -qa)
	docker volume rm $$(docker volume ls --filter dangling=true -q)
	docker rmi -f $$(docker images -a -q)
	docker system prune -a
