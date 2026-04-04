-- name: CreateTodo :one
INSERT INTO
  todos (id, user_id, title, description, status)
VALUES
  ($1, $2, $3, $4, $5)
RETURNING
  *;

-- name: GetTodoByID :one
SELECT
  *
FROM
  todos
WHERE
  id = $1
  AND user_id = $2
LIMIT
  1;

-- name: ListTodosByUserID :many
SELECT
  *
FROM
  todos
WHERE
  user_id = $1
ORDER BY
  created_at DESC;

-- name: UpdateTodoStatus :exec
UPDATE todos
SET
  title = $3,
  description = $4,
  status = $5,
  updated_at = NOW()
WHERE
  id = $1
  AND user_id = $2
RETURNING
  *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE
  id = $1
  AND user_id = $2;
