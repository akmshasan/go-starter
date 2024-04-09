-- name: GetFruit :one
SELECT * FROM fruit
WHERE id = $1 LIMIT 1;

-- name: ListFruits :many
SELECT * FROM fruit
ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateFruit :one
INSERT INTO fruit (
  name, color, price
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateFruit :one
UPDATE fruit
  set price = $2
WHERE id = $1
RETURNING *;

-- name: DeleteFruit :exec
DELETE FROM fruit
WHERE id = $1;