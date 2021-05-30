client:
	curl -H "Content-Type: application/json" -X POST -d "@sample-client.json" http://127.0.0.1:3000/api/v1/clients | jq
invalid_client:
	curl -H "Content-Type: application/json" -X POST -d "@invalid-client.json" http://127.0.0.1:3000/api/v1/clients | jq
	#curl -H "Content-Type: application/json" -X POST http://127.0.0.1:3000/api/v1/clients | jq
get:
	curl -H 'accept: application/dns-json' http://localhost:3000/api/v1/clients/1 | jq
ping:
	curl -H "Content-Type: application/json" -X GET http://127.0.0.1:3000/ping | jq
