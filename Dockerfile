from postgres:13.3

label description "postgresql instance"

copy migrations/init.sql /docker-entrypoint-initdb.d/