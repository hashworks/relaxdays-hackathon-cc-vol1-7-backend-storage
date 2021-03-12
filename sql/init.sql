-- name: create-table-storage
CREATE TABLE IF NOT EXISTS storage(
  name VARCHAR NOT NULL PRIMARY KEY,
  articleId INTEGER,
  stock INTEGER
);

-- name: create-table-purchase
CREATE TABLE IF NOT EXISTS purchase(
  purchaseId INTEGER PRIMARY KEY AUTOINCREMENT,
  vendor VARCHAR NOT NULL,
  articleId INTEGER,
  bulk INTEGER
);