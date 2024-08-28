-- name: GetQuotes :many
SELECT * FROM tb_quotes;

-- name: GetRandomQuote :one
SELECT * FROM tb_quotes ORDER BY RANDOM() LIMIT 1;

-- name: CreateQuote :one
INSERT INTO tb_quotes (quote) VALUES (?) RETURNING *;
