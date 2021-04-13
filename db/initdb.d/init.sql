CREATE DATABASE IF NOT EXISTS mydb CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
CREATE USER IF NOT EXISTS 'user'@'%' IDENTIFIED BY 'root';
GRANT ALL PRIVILEGES ON mydb.* TO 'user'@'%';
SET session wait_timeout=24 * 60 * 60;
SET global max_allowed_packet=100 * 1024 * 1024;

FLUSH PRIVILEGES;