DB_URL=postgresql://dkennetz:harrison40@localhost:5432/fileflipperdb?sslmode=disable

createuser:
	su - postgres -c "createuser fileadmin"

password:
	su - postgres -c "psql -c \"alter user fileadmin with password 'test1234'\""
createdb:
	createdb --username=dkennetz --owner=dkennetz fileflipperdb

dropdb:
	dropdb fileflipperdb

createmigration:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "$(DB_URL)" --verbose up

loaddata:
	psql -f - fileflipperdb -t < ./db/migration/load_data.sql

doover:
	dropdb fileflipperdb && createdb --username=dkennetz --owner=dkennetz fileflipperdb && migrate -path db/migration -database "$(DB_URL)" --verbose up