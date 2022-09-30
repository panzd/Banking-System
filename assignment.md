1.Enhance getAllCustomers API

The API should provide an option to fetch the customers by status.
/customers?status=active
/customers?status=inactive
/customers

2.Make a transaction in bank account

transaction /træn'zækʃ(ə)n/ 处理

create a new transaction for an existing customer

transaction can only be "withdrawal" or "deposit"

amount cannot be negative // 

withdrawal amount should be available in the account

successful transaction,shoule return the updated balance with transaction id response

error handling should be done for bad request,

validation and unexpected errors from the server side and should 

return the appropriate http status code with message




