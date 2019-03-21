# Deposit
## API: Fee Calculate
### URL: api/v1/deposit/limit/transaction
### METHOD: GET
#### Response
```json
// Success (status code: 200)
{
    "limit": "30000.00"
}
```
#### Response Data
| ชื่อ Field | type | อธิบาย |
|---------|----------|--------|
| limit | string | จำนวนเงินที่ limit ต่อ transaction |


## API: Fee Calculate
### URL: api/v1/deposit/fee
### METHOD: POST

#### Request
##### Body
```json
{
    "account_number": "2000000012541",
    "amount": "500.00",
}
```

#### Response
```json
// Success (status code: 200)
{
    "account_number": "12345678",
    "amount": "500.00",
    "fee": "15.00",
    "net_amount": "485.00"
}
```

#### Error response
```json
// Failed (status code: 400)
{
    "error_code":"E12002",
    "error_message":"wrong input validation"
}
```
#### Request Data
| ชื่อ Field | type | อธิบาย |
|---------|----------|--------|
| account_number | string | เลขที่บัญชี |
| amount | string | จำนวนเงินที่ฝาก |


#### Response Data
| ชื่อ Field | type | อธิบาย |
|---------|----------|--------|
| account_number | string | เลขที่บัญชี |
| amount | string | จำนวนเงินที่ฝาก |
| fee | string | ค่าธรรมเนียม |
| net_amount | string | จำนวนเงินฝากที่เข้าบัญชี |

