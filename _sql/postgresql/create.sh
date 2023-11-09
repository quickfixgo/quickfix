dropdb -U postgres --if-exists quickfix
createdb -U postgres quickfix
psql -U postgres -d quickfix -f postgresql.sql
