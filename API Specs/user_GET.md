## View the details of the User
 
Description: Will show the details of the User who logged in

### HTTP Request
`GET/user`

### URL Parameters
/user/{id}

### Query Parameters
N/A


### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```

### Request Body
Empty

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
 "Message" : "User info"
}
```

### Bad Request Response when invalid user-id
```
{
    "Message": "user not found. Must be a valid user"
}
```

###  Forbidden Response when attempting unauthorized access
```
{
    "Message": "Unable to verify the token. Please contact your administrator"
}
```

### Gone Response when Reset Password link for Existing User is expired
```
{
    "Message": "Your link to reset your password has expired. Please use “Forgot your password?” to get a new link."
}
```