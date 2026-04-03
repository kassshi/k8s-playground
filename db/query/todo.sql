-- name: GetTodoById :one
SELECT
  *
FROM
  todos
WHERE
  id = $1
LIMIT
  1;

-- name: ListTodosByUserId :many
SELECT
  *
FROM
  todos
WHERE
  user_id = $1
ORDER BY
  created_at DESC;

-- name: CreateTodo :one
INSERT INTO
  todos (id, user_id, title, description)
VALUES
  ($1, $2, $3, $4)
RETURNING
  *;

-- name: UpdateTodo :exec
UPDATE todos
set
  user_id = $2,
  title = $3,
  description = $4
WHERE
  id = $1;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE
  id = $1;
