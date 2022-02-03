-- name: GetAccount :one
select
    acc.id,
    balance,
    currency
from public.account acc
inner join public.account_balance acb
    on acc.id = acb.id
where
    acc.id = @id
LIMIT 1;

-- name: ListAccounts :many
select
    acc.id,
    balance,
    currency
from
    public.account acc
inner join
    public.account_balance acb
on acc.id = acb.id
LIMIT @number
OFFSET @page;

-- name: CreateAccount :one
insert into public.account (
    id,
    name,
    created_date,
    modified_date)
values(
    @id,
    @name,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
RETURNING id;

-- name: TopupAccountBalance :one
insert into public.account_balance (id, balance, currency)
values($1, $2, $3)
    on conflict (id)
do
update set balance = public.account_balance.balance + excluded.balance
returning id;

