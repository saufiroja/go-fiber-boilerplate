DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id VARCHAR(100) PRIMARY KEY,
    full_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_male BOOLEAN,
    created_at BIGINT NOT NULL,
    updated_at BIGINT,
    deleted_at BIGINT DEFAULT NULL
);
