# Change details of the book
Description: will change the book details as per his need

### HTTP Request
`PUT/book/id`

### URL Parameters
book/id

### Query Parameters
N/A

### Request Headers
```
Content-Type: application/x-www-form-urlencoded
```


### Request Body
| Parameter | Format | Description                                
|-----------|--------|--------------------------------------------|
| publisher   | String |publisher of book|
| Name   | String | name of book|
| no_of_copies | int | total no. Of copies of the book |
| available_copies  | int | remaining number of books|
| status | string |to know whether the book is available or not|



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


