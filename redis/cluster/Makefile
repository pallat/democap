.PHONY: init
init: # To initialize the Redis cluster
	docker exec -it redis-1 redis-cli --cluster create \
  $(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-1):6380 \
  $(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-2):6380 \
  $(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-3):6380 \
  --cluster-replicas 0