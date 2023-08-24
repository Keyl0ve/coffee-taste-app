dup:
	docker-compose up -d

dbuild:
	docker-compose up -d --build

ddown:
	docker-compose down

dremove:
	docker-compose down --volume

drestart:
	docker-compose restart