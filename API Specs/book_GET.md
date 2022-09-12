## View the details of the Book

Description: This will show the details of the book

### HTTP Request
`GET/book`

### URL Parameters
/book/{id}

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

### Request Body
| Parameter | Format | Description                                
|-----------|--------|--------------------------------------------|
| publisher   | String |publisher of book|
| Name   | String | name of book|
| no_of_copies | int | total no. Of copies of the book |
| availabe_copies  | int | remaining number of books|
| status | string |to know whether the book is available or not|


### Response Headers
N/A

### Success Response Body
```
{
    "Message": "book info"
}
```

### Bad Request Response when invalid book id
```
{
    "Message": "book not found. Must be a valid book id"
}
```
