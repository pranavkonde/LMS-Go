## Delete User 
Description: This will delete the user details

### HTTP Request
`DELETE/user`

### URL Parameters
user/id

### Query Parameters


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Email     | String | Email Id of user|
| Password  | String | Password of user |


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
 "Message": “User successfully deleted"
}
```

### Bad Request Response when Password validation failed
```
{
    "Message": "Invalid password. "
}
```


### Bad Request Response when a user does not exist 
```
{
    "Message": "Try logging with correct email-id"
}
```

### Forbidden Response when Passwords don't match
```
{
    "Message": "Passwords don't match!"
}
```

### Forbidden Response when Email not present in a request
```
{
    "Message": "Unable to verify the token. Please contact your administrator"
}
```