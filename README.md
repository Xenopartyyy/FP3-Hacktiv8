# FINAL PROJECT 3

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)


Berikut ini adalah final project ke-3 dari hacktiv8, aplikasi ini bernama Kanban Board, Aplikasi ini akan dilengkapi dengan proses CRUD.

## Anggota kelompok
 - Willyawan Maulana - GLNG-KS07-014
 - Dirham Triyadi - GLNG-KS07-025

## Endpoint
Berikut ini adalah seluruh endpoint yang dapat diakses melalui client.

### Users
 
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Users
 
| Method | URL |
| ------ | ------ |
| POST | [https://fp3-hacktiv8-production.up.railway.app/users/register] |
| POST | [https://fp3-hacktiv8-production.up.railway.app/users/login |
| PUT | [https://fp3-hacktiv8-production.up.railway.app/users] |
| DELETE | [https://fp3-hacktiv8-production.up.railway.app/users] |

###### Daftar request users

POST Register User
 ```sh
{
    "full_name": "string",
    "email": "string",
    "password": "string",
}
```
#
POST Login User
 ```sh
{
    "email": "string",
    "password": "string",
}
```
#

PUT User

-Bearer Token <br />
 ```sh
{
    "full_name": "string",
    "email": "string"
}
```
#
DELETE User

-Bearer Token

> Note: Untuk method PUT dan DELETE diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login
#



### Categories
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Categories

| Method | URL |
| ------ | ------ |
| POST | [https://fp3-hacktiv8-production.up.railway.app/categories] |
| GET | [https://fp3-hacktiv8-production.up.railway.app/categories] |
| PUT | [https://fp3-hacktiv8-production.up.railway.app/categories/:categoryId] |
| DELETE | [https://fp3-hacktiv8-production.up.railway.app/categories/:categoryId] |

###### Daftar request categories

POST Categories

-Bearer Token <br />
 ```sh
{
    "type": "string"
}
```
#

GET Categories 

-Bearer Token <br />

#

PATCH Categories

-Bearer Token <br />
-Param categoryID <br />

 ```sh
{
    "type": "string"
}
```
#
DELETE Categories

-Bearer Token <br />
-Param categoryID <br />

> Note: Untuk method POST, PATCH dan DELETE hanya bisa diakses oleh akun dengan role admin serta diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login. Untuk methode PATCH dan DELETE diperlukan parameter Id pada URL

#
### Tasks
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Tasks

| Method | URL |
| ------ | ------ |
| POST | [https://fp3-hacktiv8-production.up.railway.app/tasks] |
| GET | [https://fp3-hacktiv8-production.up.railway.app/tasks] |
| PUT | [https://fp3-hacktiv8-production.up.railway.app/tasks/:taskId] |
| DELETE | [https://fp3-hacktiv8-production.up.railway.app/tasks/:taskId] |

###### Daftar request tasks

POST Tasks

-Bearer Token <br />

 ```sh
{
    "title": "string",
    "description": "string",
    "category_id": integer
}
```
#

GET Tasks 

-Bearer Token <br />

#

PUT Tasks

-Bearer Token <br />
-Param categoryID <br />

 ```sh
{
    "title": "string",
    "description": "string
}
```

#
PATCH Tasks status

-Bearer Token <br />
-Param categoryID <br />

 ```sh
{
    "status": boolean
}
```
#

PATCH Tasks categoryId

-Bearer Token <br />
-Param categoryID <br />

 ```sh
{
    "category_id": integer
}
```
#

DELETE Tasks

-Bearer Token <br />
-Param categoryID <br />

> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login. Untuk methode PUT, PATCH dan DELETE diperlukan parameter Id pada URL. User hanya dapat melakukan GET, PUT, PATCH, dan DELETE pada tasks yang dibuat oleh user itu sendiri 
