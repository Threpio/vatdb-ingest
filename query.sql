-- name: CreateDataInstance :one
INSERT INTO data_instance (
                         value
) VALUES (
             $1
) RETURNING id;

-- name: ListDataInstanceTimestampDesc :many
SELECT * FROM data_instance
ORDER BY timestamp DESC;

-- name: GetDataInstanceById :one
SELECT * FROM data_instance
WHERE id = $1;

-- name: GetDataInstancesByTimestamp :many
SELECT * FROM data_instance
WHERE
    timestamp >= $1
  AND
      timestamp <= $2;
