PODMAN ?= podman

.PHONY: generate
generate:
	go build -v -o ./cmd/generate ./cmd/generate && ./cmd/generate/generate

.PHONY: pg-test
pg-test:
	$(PODMAN) rm -f pg-test || true
	$(PODMAN) run --name pg-test -e POSTGRES_PASSWORD=postgres -d --rm -p 5432:5432 docker.io/library/postgres:13
	$(PODMAN) exec -it pg-test bash -c 'while ! pg_isready; do sleep 1; done;' 