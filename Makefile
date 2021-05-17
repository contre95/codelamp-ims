grafana:
	docker run -d -p 3000:3000 -v $$(pwd)/db:/usr/share/grafana/db  -e "GF_INSTALL_PLUGINS=frser-sqlite-datasource" --user 104 --name grafana grafana/grafana:latest
clean:
	rm -rf $$(pwd)/db
db_dir:
	mkdir $$(pwd)/db
samples: clean db_dir
	go run cmd/sample-data/*.go
run:
	go run cmd/api/main.go
