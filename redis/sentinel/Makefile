up:
	docker compose --env-file .env up
leader:
	docker exec -it sentinel-2 redis-cli -p 26379 sentinel get-master-addr-by-name mymaster
swap:
	docker-compose start redis-master