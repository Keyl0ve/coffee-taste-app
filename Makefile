dup:
	docker-compose up

dbuild:
	docker-compose up -d --build

ddown:
	docker-compose down

dremove:
	docker-compose down --volume

drestart:
	docker-compose restart