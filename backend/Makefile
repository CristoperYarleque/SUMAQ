.PHONY: run
run :
	@if [ -z "$(env)" ]; then \
		echo "Por favor, proporciona un valor para 'env'. Ejemplo: make run env=development"; \
		exit 1; \
	fi
	go run -ldflags "-w -s -X main.env=$(env)" server.go

