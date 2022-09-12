## Create user
Description: to create a new user

### HTTP Request
`POST/user`

### URL Parameters
/user/create

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                
|-----------|--------|--------------------------------------------|
| Email     | String | Email Id of user |
| Name   | String | name of user|
| role id   | int | select from 1,2,3 as super admin,admin, and end user |
| contact_no   | long_int | contact details of user|


### Sample cURL request


### Status codes and errors

| Value | Description           |
|-------|-----------------------|
| 200   | OK                    |
| 400   | Bad Request           |
| 403   | Forbidden             |
| 410   | Gone                  |
| 500   | Internal Server Error |

### Response Headers
N/A

### Success Response Body
```
{
    "Message": “User Created"
}
```

### Bad Request Response when Password validation failed
```
{
    "Message": "Invalid information. Please enter the valid information"
}
```

### Bad Request Response when user already exits
```
{
    "Message": "Your account is already created Try logging in."
}
```


### Forbidden Response when the field is empty 
```
{
    "Message": “Fields cannot be empty"
}
```