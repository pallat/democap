up:
	docker compose up -d
down:
	docker compose down
info:
	docker exec mongo1 sh -c "mongosh --eval 'rs.status()'"
disconnect:
	docker network disconnect treenodes_mongo_network mongo1
connect:
	docker network connect treenodes_mongo_network mongo1
