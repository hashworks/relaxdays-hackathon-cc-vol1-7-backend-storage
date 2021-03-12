-- name: select-storage-by-name
SELECT articleId, stock FROM storage WHERE name=?;

-- name: select-storage-by-cursor
SELECT name, articleId, stock FROM storage WHERE name > ? LIMIT ?;