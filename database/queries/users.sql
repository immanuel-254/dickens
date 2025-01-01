-- name: UserCreate :one
INSERT INTO users (
    surname, 
    first_name, 
    last_name, 
    password, 
    email, 
    created_at, 
    updated_at
    ) 
    VALUES (?, ?, ?, ?, ?, ?, ?)
    RETURNING id, surname, first_name, last_name, email, created_at, updated_at;

-- name: UserList :many
SELECT id, surname, first_name, last_name, email, created_at, updated_at FROM users
ORDER BY id ASC;

-- name: UserRead :one
SELECT id, surname, first_name, last_name, email, created_at, updated_at FROM users
WHERE id = ?;

-- name: UserLoginRead :one
SELECT email, password FROM users
WHERE email = ?;

-- name: UserUpdate :one
UPDATE users
SET surname = ?, first_name = ?, last_name = ?, updated_at = ?
WHERE id = ?
RETURNING id, surname, first_name, last_name, email, created_at, updated_at;

-- name: UserUpdatePassword :one
UPDATE users SET password = ?, updated_at = ? WHERE id = ? 
RETURNING id, surname, first_name, last_name, email, created_at, updated_at;

-- name: UserUpdateEmail :one
UPDATE users SET email = ?, updated_at = ? WHERE id = ? 
RETURNING id, surname, first_name, last_name, email, created_at, updated_at;

-- name: UserDelete :exec
DELETE FROM users WHERE id = ?;
