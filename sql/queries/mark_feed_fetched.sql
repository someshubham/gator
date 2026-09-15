-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $1,
    last_fetched_at = $1
WHERE id = $2;