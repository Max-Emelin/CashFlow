run:
	docker-compose up --build -d       
	docker-compose run --rm goose       

down:
	docker-compose down

