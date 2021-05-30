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
- `GET https://konsultasi-psikolog.herokuapp.com/users`
- `POST https://konsultasi-psikolog.herokuapp.com/users/register`
- `POST https://konsultasi-psikolog.herokuapp.com/users/login`
- `GET https://konsultasi-psikolog.herokuapp.com/users/:user_id`
- `PUT https://konsultasi-psikolog.herokuapp.com/users/:user_id`
- `DELETE https://konsultasi-psikolog.herokuapp.com/users/:user_id`

### psikologs
- `GET https://konsultasi-psikolog.herokuapp.com/psikologs`
- `POST https://konsultasi-psikolog.herokuapp.com/psikologs`
- `GET https://konsultasi-psikolog.herokuapp.com/psikologs/:psikolog_id`
- `PUT https://konsultasi-psikolog.herokuapp.com/psikologs/:psikolog_id`
- `DELETE https://konsultasi-psikolog.herokuapp.com/psikologs/:psikolog_id`

### booking
- `GET https://konsultasi-psikolog.herokuapp.com/booking`
- `POST https://konsultasi-psikolog.herokuapp.com/booking/order`
- `GET https://konsultasi-psikolog.herokuapp.com/booking/:booking_id`
- `DELETE https://konsultasi-psikolog.herokuapp.com/booking/:booking_id`

### booking detail
- `GET https://konsultasi-psikolog.herokuapp.com/booking-detail`
- `POST https://konsultasi-psikolog.herokuapp.com/booking-detail`
- `GET https://konsultasi-psikolog.herokuapp.com/booking-detail/:booking_detail_id`

### roles
- `GET https://konsultasi-psikolog.herokuapp.com/roles`
- `POST https://konsultasi-psikolog.herokuapp.com/roles`
- `GET https://konsultasi-psikolog.herokuapp.com/roles/:role_id`
- `DELETE https://konsultasi-psikolog.herokuapp.com/roles/:role_id`

### userProfiles
- `GET https://konsultasi-psikolog.herokuapp.com/users_profile`
- `POST https://konsultasi-psikolog.herokuapp.com/users_profile`
- `PUT https://konsultasi-psikolog.herokuapp.com/users_profile`

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
## RESTful endpoints booking
### GET /booking

> Get All booking

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
      "message" : "Success Get All Bookings Data",
      "code" : 200,
      "status" : "Status OK"
  }, 
  "data" : [			
      {
          "id" : 1,
          "user_id" : 1,
          "psikolog_id" : 2,
          "booking_date" : 2021-05-25,
          "booking_time" : 07:27:39.907,
      }, {
          "id" : 2,
          "user_id" : 2,
          "psikolog_id" : 4,
          "booking_date" : 2021-05-25,
          "booking_time" : 07:27:39.907,
      }
      , {
          "id" : 3,
          "user_id" : 3,
          "psikolog_id" : 6,
          "booking_date" : 2021-05-25,
          "booking_time" : 07:27:39.907,
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

### POST /booking/order  

> Create new booking

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "user_id" : "<Your id user>",
  "psikolog_id" : "<Psikolog id to get insert into>",
  "booking_date" : "<Booking Date to get insert into>",
  "booking_time": "<Booking Time to get insert into>"
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "Success Create new Bookings Data",
      "code" : 201,
      "status" : "Status Created"
  }, 
  "data" : 
      {
        "id" : <given id by system>,
        "user_id" : "<posted user id>",
        "psikolog_id" : "<posted psikolog id>",
        "booking_date" : "<posted booking date>",
        "booking_time" : "<posted booking time>"
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

### GET /booking/:booking_id 

> Get booking by booking ID

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
      "message" : "Success get booking by id",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 2,
        "user_id" : 2,
        "psikolog_id" : 4,
        "booking_date" : 2021-05-25,
        "booking_time" : 07:27:39.907
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

### DELETE /booking/:booking_id 

> Delete booking by ID

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
      "message" : "Success delete booking id",
      "code" : 200,
      "status" : "Delete OK"
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
---
## RESTful endpoints roles
### GET /roles

> Get All roles

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
      "message" : "Success Get All Roles Data",
      "code" : 200,
      "status" : "Status OK"
  }, 
  "data" : [			
      {
          "id" : 1,
          "nama_role" : admin,
          "description" : CRUD frontend function from dashboard,
      }, {
          "id" : 2,
          "nama_role" : user,
          "description" : looking for psikolog consulting services,
      }
      , {
          "id" : 3,
          "nama_role" : psikolog,
          "description" : provide consulting services to users,
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

### POST /roles

> Create new roles

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "nama_role" : "<Nama role  to get insert into>",
  "description" : "<Description to get insert into>",
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "Success Create new Roles Data",
      "code" : 201,
      "status" : "Status Created"
  }, 
  "data" : 
      {
        "id" : <given id by system>,
        "nama_role" : "<posted name role>",
        "description" : "<posted description>"
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

### GET /roles/:role_id 

> Get roles by role ID

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
      "message" : "Success get role by id",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 1,
        "nama_role" : admin,
        "description" : CRUD frontend function from dashboard
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

### DELETE /roles/:role_id   

> Delete role by ID

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
      "message" : "Success delete role id",
      "code" : 200,
      "status" : "Delete OK"
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
---
## RESTful endpoints booking detail
### GET /booking-detail  

> Get All booking detail  

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
      "message" : "Success Get All Booking Details Data",
      "code" : 200,
      "status" : "Status OK"
  }, 
  "data" : [			
      {
          "id" : 1,
          "booking_id" : 2,
          "psikolog_id" : 4,
      }, {
          "id" : 2,
          "booking_id" : 3,
          "psikolog_id" : 5,
      }
      , {
          "id" : 3,
          "booking_id" : 5,
          "psikolog_id" : 5,
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

### POST /booking-detail 

> Create new booking detail

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "booking_id" : "<booking id to get insert into>",
  "psikolog_id" : "<psikolog id to get insert into>",
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "Success Create new Booking Detail Data",
      "code" : 201,
      "status" : "Status Created"
  }, 
  "data" : 
      {
        "id" : <given id by system>,
        "booking_id" : "<posted name role>",
        "psikolog_id" : "<posted description>"
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

### GET /booking-detail/:booking_detail_id 

> Get booking detail by booking detail ID

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
      "message" : "Success get booking detail by id",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 1,
        "booking_id" : 2,
        "psikolog_i" : 4
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

## RESTful endpoints userProfile
- `GET /users_profile` 
- `POST /users_profile`
- `PUT /users_profile`

### GET /users_profile
> get user profile by user ID login 

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
      "message" : "success get user profile by user ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "profile_user" : "https://konsultasi-psikolog.herokuapp.com/images/profile-7-google.com.jpg",
        "user_id" : 2 
    }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "Unauthorize",
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

### POST /users_profile
> update user profile by user id login

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
file upload (google.com.jpg)
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success get user profile by user ID",
      "code" : 201,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "profile_user" : "https://konsultasi-psikolog.herokuapp.com/images/profile-7-google.com.jpg",
        "user_id" : 2 
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
  "data" : {
      "error" : ""
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

### PUT /users_profile
> update user profile by user id login

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
Upload new image file in postman form data
```


_Response (201)_
```json
{
  "meta" : {
      "message" : "success update user profile image",
      "code" : 200,
      "status" : "success update"
  }, 
  "data" : {
        "id": 1,
        "image_user": "https://konsultasi-psikolog.herokuapp.com/images/profile-7-google.com.jpg,
        "user_id": 4,
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
  "data" : {
      "errors" : []
  }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "Unauthorize",
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
