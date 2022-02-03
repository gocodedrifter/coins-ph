-- name: ListPaymentsById :many
select account_id,
       case when direction = 'outgoing' then credit else debit end amount,
       from_account_id,
       to_account_id,
       direction
from public.payment_data
where
    account_id = @account
LIMIT @number
OFFSET @page;

-- name: ListPayments :many
select account_id,
       case when direction = 'outgoing' then credit else debit end amount,
       from_account_id,
       to_account_id,
       direction
from public.payment_data
LIMIT @number
OFFSET @page;