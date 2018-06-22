build:
	docker-compose build

up:
	docker-compose up -d

gen:
	docker-compose run --rm slidegen go run main.go $(F)

stop:
	docker-compose stop

clean:
	sudo rm *.pdf
