-- name: GetFruit :one
SELECT * FROM fruit
WHERE id = $1 LIMIT 1;

-- name: ListFruits :many
SELECT * FROM fruit
ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateFruit :one
INSERT INTO fruit (
  id, name, color, price, quantity
) VALUES (
  uuid_if_empty(sqlc.arg(id))::uuid, $1, $2, $3, $4
) ON CONFLICT (id) DO UPDATE 
  SET name = $2, color = $3, price = $4, quantity = $5
RETURNING *;

-- name: UpdateFruit :one
UPDATE fruit
  set price = $2
WHERE id = $1
RETURNING *;

-- name: DeleteFruit :exec
DELETE FROM fruit
WHERE id = $1;