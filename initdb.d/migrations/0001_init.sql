-- Buat user 'root' jika belum ada
CREATE USER root WITH PASSWORD 'root';

-- Berikan semua hak akses ke database 'boilerplate' kepada user 'root'
GRANT ALL PRIVILEGES ON DATABASE boilerplate TO root;