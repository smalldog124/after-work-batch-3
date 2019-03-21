*** Settings ***
Library  Collections
Library  String
Library  RequestsLibrary

*** Test Cases ***
Get Limit
    Create Session  banking  http://localhost:3000
    ${response}=  Get Request  banking  /api/v1/deposit/limit/transaction
    Should Be Equal As Strings  ${response.status_code}  200
    ${jsondata}=  To Json  ${response.content}
    Should Be Equal  ${jsondata['limit']}  30000.00

Fee Calculate
    Create Session  google  http://localhost:3000
    &{data}=    Create Dictionary   account_number=12345678  amount=500.00
    ${response}=  Post Request  banking  /api/v1/deposit/fee  json=${data}
    Should Be Equal As Strings  ${response.status_code}  200
    ${jsondata}=  To Json  ${response.content}
    Should Be Equal  ${jsondata['fee']}  15.00
    Should Be Equal  ${jsondata['net_amount']}  485.00
