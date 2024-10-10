SELECT 'CREATE DATABASE social '
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'social');

ALTER USER 'admin' WITH PASSWORD 'adminpassword';