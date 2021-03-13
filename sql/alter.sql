-- name: insert-or-update-storage
INSERT INTO storage(name, articleId, stock, capacity)
  VALUES(?, ?, ?, ?)
  ON CONFLICT(name)
  DO UPDATE SET articleId=excluded.articleId, stock=excluded.stock, capacity=excluded.capacity;

-- name: delete-storage
DELETE FROM storage WHERE name=?;