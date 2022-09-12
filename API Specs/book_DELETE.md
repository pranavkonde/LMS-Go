## Delete	Book 
Description: this will delete the book details

### HTTP Request
`DELETE/book/{id}`

### URL Parameters
Book/delete

### Query Parameters


### Request Headers

```
Content-Type: application/x-www-form-urlencoded
```


### Request Body
| Parameter | Format | Description                                |
|-----------|--------|--------------------------------------------|
| Name    | String | Name of book|
| id  | String | id of book which is unique  |


### Sample cURL request
```

```

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
    "Message": “book successfully deleted”
}
```
### Bad Request Response when Password validation failed
```
{
    "Message": "Invalid password."
}
```


### Bad Request Response when book does not exist 
```
{
    "Message": "Book not present"
}
```

### Forbidden Response when Email not present in request
```
{
    "Message": "Unable to fetch. Please contact your admin"
}
```


















