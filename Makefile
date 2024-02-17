include cmd/system/.env

createmigration:
	migrate create -ext=sql -dir=internal/infra/database/migrations -seq init

migrateup:
	migrate -path=internal/infra/database/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose up

migratedown:
	migrate -path=internal/infra/database/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" -verbose down

.PHONY: migrate