-- name: create-table-storage
CREATE TABLE IF NOT EXISTS storage(
  name VARCHAR NOT NULL PRIMARY KEY,
  articleId INTEGER,
  stock INTEGER
);