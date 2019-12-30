SHELL   			:= /bin/bash

BACKUP_FILENAME		:= $(shell date '+%Y%m%d%H%M%S')
BACKUP_FILEPATH		:= db/db_backup_${BACKUP_FILENAME}.sql

export BACKUP_FILEPATH

.PHONY: db-drop
db-drop:
	docker-compose exec db mysql -u root -ppassword -e 'DROP DATABASE snippetbox'

.PHONY: db-create
db-setup:
	docker-compose exec -T db mysql -u root -ppassword < db/setup.sql

.PHONY: db-reset
db-reset: db-drop db-create

.PHONY: db-backup
db-backup:
	docker-compose exec -T db mysqldump snippetbox -u root -ppassword > "${BACKUP_FILEPATH}"


