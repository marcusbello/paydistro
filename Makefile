run_a:
	cd a && go run a.go

run_b:
	cd b && go run b.go

run_c:
	cd c && go run c.go

kafka:
	docker-compose up