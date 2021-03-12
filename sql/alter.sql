-- name: insert-or-update-storage
INSERT INTO storage(name, articleId, stock)
  VALUES(?, ?, ?) 
  ON CONFLICT(name) 
  DO UPDATE SET articleId=excluded.articleId, stock=excluded.stock;

-- name: delete-storage
DELETE FROM storage WHERE name=?;

-- name: insert-purchase
INSERT INTO purchase(vendor, articleId, bulk) VALUES(?, ?, ?);