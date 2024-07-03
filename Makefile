clean:
	docker compose down

	docker volume rm -f backend_scorify-postgres
	docker volume rm -f backend_scorify-redis
