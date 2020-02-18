-- name: InsertCounter :exec
insert into
    counter (id, val)
values ($1, $2);

-- name: UpdateCounter :exec
update
    counter
set
    val = $1
where
    id = $2;

-- name: SelectCounter :one
select
    val
from
    counter
where
    id = $1;

-- name: IncrementCounter :exec
update
    counter
set
    val = val + 1
where
    id = $1;
