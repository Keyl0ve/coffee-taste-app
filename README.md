# coffee-taste-app

## How to start
- docker-compose up -d

## How to see DB
- mysql -h dockerMySQL -u admin -p app
- show databases;
- show tables from app;

- localhost:4000



## How to down
- docker-compose down

## Example request
- curl http://localhost:8080/api/coffee/get
- curl -X POST -H 'userName: kyo' -H 'password: password' http://localhost:8080/api/user/create

## Example response
- name1

## front: https://github.com/Keyl0ve/coffee-taste-app-front
