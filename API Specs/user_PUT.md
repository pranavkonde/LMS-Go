# Change details of the user

Description: will change the required details of the user as per his need

### HTTP Request
`PUT/user`

### URL Parameters
user/id

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
| Parameter | Format | Description|
|-----------|--------|--------------------------------------------|
| Email     | String | Email Id of the user requesting password reset |
| field | String | the field the user wants to update |
| password | String | password of the user |



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
    "Message": “Field edited successfully"
}
```

### Bad Request Response when Password validation failed
```
{
    "Message": "Invalid password."
}
```



### Forbidden Response when Passwords don't match
```
{
    "Message": "Passwords don't match!"
}
```

### Forbidden Response when Email not present in the request
```
{
    "Message": "Unable to verify the token. Please contact your administrator"
}
```

