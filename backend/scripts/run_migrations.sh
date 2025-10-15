set -e
export $(grep -v '^#' .env | xargs)
PGPASSWORD=${DB_PASSWORD} psql -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER} -d ${DB_NAME} -f migrations/0001_create_tasks.sql
echo "Migrations applied"