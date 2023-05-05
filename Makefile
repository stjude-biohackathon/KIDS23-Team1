DB_URL=postgresql://genomeMaster@genomefileflipper.postgres.database.azure.com:Hackathon23Team1Cat@genomefileflipper:5432/fileflipperdb?sslmode=disable

createuser:
	su - postgres -c "createuser fileadmin"

password:
	su - postgres -c "psql -c \"alter user fileadmin with password 'test1234'\""
createdb:
	createdb --username=fileadmin --owner=fileadmin fileflipperdb

dropdb:
	dropdb fileflipperdb

createmigration:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "$(DB_URL)" --verbose up

loaddata:
	PGPASSWORD=Hackathon23Team1Cat psql -h genomefileflipper.postgres.database.azure.com -d fileflipperdb -U genomeMaster@genomefileflipper.postgres.database.azure.com -f - fileflipperdb -t < ./db/migration/load_data.sql

doover:
	dropdb fileflipperdb && createdb --username=fileadmin --owner=fileadmin fileflipperdb && migrate -path db/migration -database "$(DB_URL)" --verbose up
