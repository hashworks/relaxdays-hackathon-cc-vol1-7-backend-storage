-- name: select-storage-by-name
SELECT articleId, stock, capacity FROM storage WHERE name=?;

-- name: select-storage-by-cursor
SELECT name, articleId, stock, capacity FROM storage WHERE name > ? LIMIT ?;

-- name: select-storage-by-articleID
SELECT name, articleId, stock, capacity FROM storage WHERE articleId=?;

-- name: select-storage-by-cursor-and-place
SELECT name, articleId, stock, capacity FROM storage WHERE name LIKE ? || '%' AND name > ? LIMIT ?;