#!/bin/bash
cd docker || exit
docker compose cp ./mysql_data/sql mysql:/var/lib/mysql/
docker compose exec mysql bash -c "mysql -uroot -proot < /var/lib/mysql/sql/reset_database.test.sql"
docker compose exec mysql bash -c "mysql -uroot -proot < /var/lib/mysql/sql/safe_updates.test.sql"
