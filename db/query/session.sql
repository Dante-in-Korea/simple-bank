-- name: CreateSession :one
INSERT INTO sessions (
                      id,
                      username,
                      refresh_token,
                      user_agent,
                      client_ip,
                      is_blocked,
                      expires_at
) values (
          $1, $2, $3, $4, $5, $6, $7
         ) RETURNING *;

-- name: GetSession :one
select * from sessions
where id = $1 limit 1;