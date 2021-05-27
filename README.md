# Mini Forum App

```
https://forum-app-go.herokuapp.com
```

## List of available endpoints

### users

- `GET /users`
- `POST /users/register`
- `POST /users/login`
- `GET /users/:id`
- `PUT /users/:id`
- `DELETE /users/:id`

### questions

- `GET /questions`
- `GET /questions/:id`
- `POST /questions/ask`
- `PATCH /questions/:id/edit`
- `DELETE /questions/:id`

### answers

- `POST /questions/:id`
- `PATCH /answers/:id`
- `DELETE /answers/:id`

### categories

- `GET /categories`
- `POST /categories`
- `PUT /categories/:category_name`

### userprofile

- `GET /user_profile`
- `POST /user_profile`
- `PUT /user_profile`

## RESTful endpoints users

### GET /users

> Get All users

*Request Header*

```
{
   "Authorization": "<your Authorization>"
}
```

*Request Body*

```
not needed
```

*Response (200)*

```
{
    "meta": {
        "message": "get all users succeed",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "user_name": "marwanjuna",
            "first_name": "marwan",
            "last_name": "juna",
            "email": "marwan@mail.com"
        },
        {
            "id": 2,
            "user_name": "uchihasasuke",
            "first_name": "uchiha",
            "last_name": "saske",
            "email": "uchiha@mail.com"
        }
    ]
}
```

*Response (500 - Internal Server Error)*

```
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  }, 
  "data": 
      {
        "error": ""
      }
}
```

------

### POST /users/register

> Create new user

*Request Header*

```
not needed
```

*Request Body*

```
{
  "first_name" : "<first name to get insert into>",
  "last_name" : "<last name to get insert into>",
  "user_name": "<user name to get insert into>",
  "email": "<email to get insert into>",
  "password": "<password to get insert into>"
}
```

*Response (201)*

```
{
    "meta": {
        "message": "insert user data succeed",
        "code": 201,
        "status": "success"
    },
    "data": {
        "id": 3,
        "user_name": "boruto",
        "first_name": "uzumaki",
        "last_name": "boruto",
        "email": "boruto@mail.com"
    }
}
```

*Response (400 - Bad Request)*

```
{
  "meta": {
      "message": "input data required",
      "code": 400,
      "status": "bad request"
  }, 
  "data": 
      {
        "errors": []
      }
}
```

*Response (500 - Internal Server Error)*

```
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  }, 
  "data": 
      {
        "error": ""
      }
}
```

------

### POST /users/login

> Compare data login on database with data inputed

*Request Header*

```
not needed
```

*Request Body*

```
{
  "email": "<email to get compare>",
  "password": "<password to get compare>"
}
```

*Response (200)*

```
{
    "meta": {
        "message": "login user succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "token": "<your token>"
    }
}
```

*Response (400 - Bad Request)*

```
{
  "meta": {
      "message": "input data required",
      "code": 400,
      "status": "bad request"
  }, 
  "data": 
      {
        "errors": []
      }
}
```

*Response (401 - Unauthorized)*

```
{
    "meta": {
      "message": "input data error",
      "code": 401,
      "status": "error"
  }, 
  "data": 
      {
        "error": ""
      }
}
```

*Response (500 - Internal Server Error)*

```
{
  "meta": {
      "message": "Internal Server error",
      "code": 500,
      "status": "error"
  }, 
  "data": 
      {
        "error": ""
      }
}
```

------

### GET /users/:user_id

> Get user by user ID

*Request Header*

```
{
   "Authorization": "<your Authorization>"
}
```

*Request Body*

```
not needed
```

*Response (200)*

```
{
    "meta": {
        "message": "get user succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 3,
        "user_name": "boruto",
        "first_name": "uzumaki",
        "last_name": "boruto",
        "email": "boruto@mail.com"
    }
}
```

*Response (500 - Internal Server Error)*

```
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  }, 
  "data" : {
      "error": ""
  }
}
```

------

### PUT /users/:user_id

> Update user by User iD

*Request Header*

```
{
   "Authorization": "<your Authorization>"
}
```

*Request Body*

```
{
    "last_name": "boruto"
}
```

*Response (200)*

```
{
    "meta": {
        "message": "update user succeed",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 3,
        "user_name": "boruto",
        "first_name": "boruto",
        "last_name": "boruto",
        "email": "boruto@mail.com"
    }
}
```

*Response (500 - Internal Server Error)*

```
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  }, 
  "data" : {
      "error": ""
  }
}
```

------

### DELETE /users/:user_id

> Delete user by ID

*Request Header*

```
{
   "Authorization": "<your Authorization>"
}
```

*Request Body*

```
not needed
```

*Response (200)*

```
{
    "meta": {
        "message": "user was deleted successfully",
        "code": 200,
        "status": "success"
    },
    "data": {
        "message": "delete user id 3 succeed",
        "delete_time": "2021-05-27T05:55:25.705047677Z"
    }
}
```

*Response (500 - Internal Server Error)*

```go
{
  "meta" : {
      "message": "Internal server error",
      "code": 500,
      "status": "error"
  }, 
  "data" : {
      "error": ""
  }
}
```
