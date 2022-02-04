<a name="top"></a>
# MICROSERVICES (DDD-GOKIT)

Coins-PH API

# Table of contents

- [Account](#Account)
    - [Create Account](#Create-Account)
    - [Topup Balance](#Topup-Balance)
    - [Get Account Based On ID](#Get-Account)
    - [Get All Account](#Get-All-Account)
- [Payment](#Payment)
    - [Transfer Amount to other Account](#Transfer-Amount)
    - [Get Payment History By Account](#History-Account)
    - [Get All Payment History](#Get-All-Payment)

___


# <a name='Account'></a> Account

## <a name='Create-Account'></a> Create Account
[Back to top](#top)

<p>Create the account</p>

```
POST /wallet/v1/account/proc
```

### Parameters - `Request Body Parameters`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| id | `String` | <p>Account id</p> |

```
Example Request
{
    "id": "michael"
}
```

```
Example Response
{
    "account": {
        "id": "michael",
        "balance": 0,
        "currency": "USD"
    }
}
```

## <a name='Topup-Balance'></a> Topup Balance
[Back to top](#top)

<p>To Add the balance for the account</p>

```
POST /wallet/v1/account/proc/topup
```

### Parameters - `Request Body Parameters`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| id | `string` | <p>account id</p> |
| amount | `numeric` | <p>amount to add to the account balance</p> |

```
Example Request : 
{
    "id": "michael",
    "amount": 25000
}
```

```
Example Response : 
{
    "account": {
        "id": "michael",
        "balance": 25000,
        "currency": "USD"
    }
}
```
## <a name='Get-Account'></a> Get Account Based On ID
[Back to top](#top)

<p>Get The Account Based On Account ID</p>

```
GET /wallet/v1/account/proc/{id}
```

```
Example Request :
/wallet/v1/account/proc/michael
```

```
Example Response :
{
    "account": {
        "id": "michael",
        "balance": 25000,
        "currency": "USD"
    }
}
```

## <a name='Get-All-Account'></a> Get All Account
[Back to top](#top)

<p>Get All Existing Account</p>

```
POST /wallet/v1/account/proc/all
```

### Parameters - `Parameter`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| offset | `Number` | <p>offset of the record</p> |
| limit | `Number` | <p>limit of the record</p> |

```
Example Request : 
{
    "offset": 0,
    "limit": 20
}
```

```
Example Response :
{
    "accounts": [
        {
            "id": "bjb45",
            "balance": 12000,
            "currency": "USD"
        },
        {
            "id": "michael123",
            "balance": 0,
            "currency": "USD"
        },
        {
            "id": "bkb47",
            "balance": 22500,
            "currency": "USD"
        },
        {
            "id": "william51",
            "balance": 12500,
            "currency": "USD"
        },
        {
            "id": "tumhiho",
            "balance": 91500,
            "currency": "USD"
        },
        {
            "id": "michael",
            "balance": 25000,
            "currency": "USD"
        }
    ]
}
```

# <a name='Payment'></a> Payment

## <a name='Transfer-Amount'></a> Transfer Amount to other Account
[Back to top](#top)

<p>Transfer Amount from one account to another account</p>

```
POST /wallet/v1/payment/proc/transfer
```

### Parameters - `Request Body Parameters`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| account | `String` | <p>Account source transfer</p> |
| to_account | `String` | <p>Account destination transfer</p> |
| amount | `Number` | <p>Amount </p> |

```
Example Request : 
{
    "account": "william",
    "to_account": "tumhiho",
    "amount": 7000
}
```

```
Example Response : 
{
    "data": [
        {
            "account": "william51",
            "amount": 7000,
            "to_account": "tumhiho",
            "direction": "outgoing"
        },
        {
            "account": "tumhiho",
            "amount": 7000,
            "from_account": "william51",
            "direction": "incoming"
        }
    ]
}
```
## <a name='History-Account'></a> Get Payment History By Account
[Back to top](#top)

<p>Get Payment History Based on Account</p>

```
POST /wallet/v1/payment/proc/{id}
```

### Parameters - `Request Body Parameters`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| offset | `Number` | <p>offset of the record</p> |
| limit | `Number` | <p>limit of the record</p> |

```
Example Request : 
/wallet/v1/payment/proc/william51

{
    "limit": 20,
    "offset": 0
}
```

```
Example Response : 
{
    "data": [
        {
            "account": "william51",
            "amount": 7000,
            "to_account": "tumhiho",
            "direction": "outgoing"
        },
        {
            "account": "william51",
            "amount": 5000,
            "to_account": "tumhiho",
            "direction": "outgoing"
        },
        {
            "account": "william51",
            "amount": 7500,
            "to_account": "tumhiho",
            "direction": "outgoing"
        },
        {
            "account": "william51",
            "amount": 12500,
            "to_account": "bkb47",
            "direction": "outgoing"
        }
    ]
}
```

## <a name='Get-All-Payment'></a> Get All Payment History
[Back to top](#top)

<p>Get All Payment History</p>

```
POST /wallet/v1/payment/proc
```

### Parameters - `Request Body Parameters`

| Name     | Type       | Description                           |
|----------|------------|---------------------------------------|
| offset | `Number` | <p>offset of the record</p> |
| limit | `Number` | <p>limit of the record</p> |

```
Example Request : 
/wallet/v1/payment/proc

{
    "limit": 20,
    "offset": 0
}
```

```
Example Response : 
{
    "data": [
        {
            "account": "tumhiho",
            "amount": 7000,
            "from_account": "william51",
            "direction": "incoming"
        },
        {
            "account": "william51",
            "amount": 7000,
            "to_account": "tumhiho",
            "direction": "outgoing"
        },
        {
            "account": "tumhiho",
            "amount": 5000,
            "from_account": "william51",
            "direction": "incoming"
        },
        {
            "account": "william51",
            "amount": 5000,
            "to_account": "tumhiho",
            "direction": "outgoing"
        },
        {
            "account": "tumhiho",
            "amount": 7500,
            "from_account": "william51",
            "direction": "incoming"
        },
        {
            "account": "william51",
            "amount": 7500,
            "to_account": "tumhiho",
            "direction": "outgoing"
        },
        {
            "account": "bkb47",
            "amount": 12500,
            "from_account": "william51",
            "direction": "incoming"
        },
        {
            "account": "william51",
            "amount": 12500,
            "to_account": "bkb47",
            "direction": "outgoing"
        },
        {
            "account": "bkb47",
            "amount": 2000,
            "from_account": "bjb45",
            "direction": "incoming"
        },
        {
            "account": "bjb45",
            "amount": 2000,
            "to_account": "bkb47",
            "direction": "outgoing"
        }
    ]
}
```