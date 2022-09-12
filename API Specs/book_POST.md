## Add book

Description: to add a new book

### HTTP Request
`POST/book`

### URL Parameters
/user/book/create

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
| availabe_copies  | int | remaining number of books|
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
"Message": “Book added successfully”
}
```


### Bad Request Response when book already present
```
{
    "Message": "Book already exists"
}
```

### Forbidden Response when the field is empty 
```
{
    "Message": “Fields cannot be empty"
}
```
