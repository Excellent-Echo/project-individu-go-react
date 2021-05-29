# API Konsultasi Psikolog

```
https://konsultasi-psikolog.herokuapp.com/
```

Konsultasi psikolog app is an application to schedule a psychologist consultation, This app has : 
* User Friendly
* Easy to Use
* RESTful endpoint for product CRUD operation
* JSON formatted response

&nbsp;

## List of available endpoints
### users
- `GET /users`
- `POST /users/register`
- `POST /users/login`
- `GET /users/:user_id`
- `PUT /users/:user_id`
- `DELETE /users/:user_id`

### psikologs
- `GET /psikologs`
- `POST /psikologs`
- `GET /psikologs/:psikolog_id`
- `PUT /psikologs/:psikolog_id`
- `DELETE /psikologs/:psikolog_id`

### booking
- `GET /booking`
- `POST /booking/order`
- `GET /booking/:booking_id`
- `DELETE /booking/:booking_id`

### booking detail
- `GET /booking-detail`
- `POST /booking-detail`
- `GET /booking-detail/:booking_detail_id`

### roles
- `GET /roles`
- `POST /roles`
- `GET /roles/:role_id`
- `DELETE /roles/:role_id`


## RESTful endpoints users
### GET /users

> Get All users

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get all user",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : [
      {
          "id" : 1,
          "role_id" : 1,
          "firstname" : "admin",
          "lastname" : "admin",
          "email" : "admin@mail.com"
      }, {
          "id" : 2,
          "role_id" : 2
          "firstname" : "muhamad",
          "lastname" : "aziz",
          "email" : "muhamadaziz@mail.com"
      }
      , {
          "id" : 3,
          "role_id" : 3,
          "firstname" : "Psikolog Harrista",
          "lastname" : "Adiati M.Psi",
          "email" : "psikologharrista@mail.com"
      }
  ]
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### POST /users/register

> Create new user 

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "role_id" : "<Your role admin is 1, user is 2, or psikolog is 3>",
  "firstname" : "<first name to get insert into>",
  "lastname" : "<last name to get insert into>",
  "email": "<email to get insert into>",
  "password": "<password to get insert into>"
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success create new user",
      "code" : 201,
      "status" : "success"
  }, 
  "data" : 
      {
        "id" : <given id by system>,
        "role_id" : "<posted role_id>",
        "firstname" : "<posted firstname>",
        "lastname" : "<posted lastname>",
        "email" : "<posted email>"
      }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---

### POST /users/login

> Compare data login on database with data inputed

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "email": "<email to get compare>",
  "password": "<password to get compare>"
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success login user",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : 
      {
        "token" : "<your access token>"
      }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "input data error",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---

### GET /users/:user_id

> Get user by user ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get all user",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 2,
        "role_id" : 2,
        "first_name" : "muhamad",
        "last_name" : "aziz",
        "email" : "muhamadaziz@mail.com"
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### PUT /users/:user_id

> Update user by User iD

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
    "role_id" : "2",
    "firstname" : "iron",
    "lastname" : "man",
    "email" : ""
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success update user by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 1,
        "first_name" : "afista",
        "last_name" : "pratama",
        "email" : "pratama@mail.com"
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### DELETE /users/:user_id

> Delete user by ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success delete user by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : "",
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

## RESTful endpoints psikologs
### GET /psikologs

> Get All psikologs

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get all psikolog",
      "code" : 200,
      "status" : "status OK"
  }, 
  "data" : [
      {
        "id" : 1,
        "firstname" : "muhamad",
        "lastname" : "aziz",
        "title" : "Muhamad Aziz, S.Psi, M.Psi, Psikolog",
        "price" : "muhamadaziz@mail.com",
        "jenis_konsultasi" : "konsultasi mental",
        "description" : "Muhamad Aziz, S.Psi, M.Psi, Psikolog adalah seorang psikolog 							yang berpraktik di Rumah Sakit Columbia Asia Semarang. Beliau 							dapat menangani tindakan berupa : Konsultasi terkait Psikologis."
      },  
      {
        "id" : 2,
        "firstname" : "iron",
        "lastname" : "man",
        "title" : "Iron Man, S.Psi, M.Psi, Psikolog",
        "price" : "ironman@mail.com",
        "jenis_konsultasi" : "konsultasi teknologi",
        "description" : "Iron man, S.Psi, M.Psi, Psikolog adalah seorang psikolog 							yang berpraktik di Gedung Stark. Beliau dapat menangani tindakan 						berupa : Konsultasi terkait Psikologis teknologi."
      }
  ]
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### POST /psikologs

> Create new psikologs

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "firstname" : "<string first name to get insert into>",
  "lastname" : "<string last name to get insert into>",
  "title" : "<string title to get insert into>",
  "price": <integer price to get insert into>,
  "jenis_konsultasi": "<string jenis konsultasi to get insert into>",
  "description" : "<string description to get insert into>"
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success create new Psikolog",
      "code" : 201,
      "status" : "status Created"
  }, 
  "data" : 
      {
        "id" : <given id by system>,
        "firstname" : "<posted firstname>",
        "lastname" : "<posted lastname>",
        "title" : "<posted title>",
        "price" : "<posted price>",
        "jenis_konsultasi" : "<posted jenis_konsultasi>",
        "description" : "<posted description>",
      }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---


### GET /psikologs/:psikolog_id

> Get user by user ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get psikolog by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 1,
        "firstname" : "muhamad",
        "lastname" : "aziz",
        "title" : "Muhamad Aziz, S.Psi, M.Psi, Psikolog",
        "price" : "muhamadaziz@mail.com",
        "jenis_konsultasi" : "konsultasi mental",
        "description" : "Muhamad Aziz, S.Psi, M.Psi, Psikolog adalah seorang psikolog 							yang berpraktik di Rumah Sakit Columbia Asia Semarang. Beliau 							dapat menangani tindakan berupa : Konsultasi terkait Psikologis."
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### PUT /psikologs/:psikolog_id

> Update user by Psikologs iD

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
        "firstname" : "muhamad",
        "lastname" : "aziz",
        "title" : "Muhamad Aziz, S.Psi, M.Psi, Psikolog",
        "price" : "muhamadaziz@mail.com",
        "jenis_konsultasi" : "konsultasi mental",
        "description" : "Muhamad Aziz, S.Psi, M.Psi, Psikolog adalah seorang psikolog 							yang berpraktik di Rumah Sakit Columbia Asia Semarang. Beliau 							dapat menangani tindakan berupa : Konsultasi terkait Psikologis."
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success update psikolog by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 1,
        "firstname" : "muhamad",
        "lastname" : "aziz",
        "title" : "Muhamad Aziz, S.Psi, M.Psi, Psikolog",
        "price" : "muhamadaziz@mail.com",
        "jenis_konsultasi" : "konsultasi mental",
        "description" : "Muhamad Aziz, S.Psi, M.Psi, Psikolog adalah seorang psikolog 							yang berpraktik di Rumah Sakit Columbia Asia Semarang. Beliau 							dapat menangani tindakan berupa : Konsultasi terkait Psikologis."
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### DELETE /psikologs/:psikolog_id 

> Delete user by ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success delete psikolog by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : "",
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

