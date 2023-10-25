# FINAL PROJECT 3

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)


Berikut ini adalah final project ke-3 dari hacktiv8, aplikasi ini bernama Kanban Board, Aplikasi ini akan dilengkapi dengan proses CRUD.

## Endpoint
Berikut ini adalah seluruh endpoint yang dapat diakses melalui client

### Users
 
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Users
 
| Method | URL |
| ------ | ------ |
| POST | [https://fp3-hacktiv8-production.up.railway.app/users/register] |
| POST | [https://fp3-hacktiv8-production.up.railway.app/users/login |
| PUT | [https://fp3-hacktiv8-production.up.railway.app/users] |
| DELETE | [https://fp3-hacktiv8-production.up.railway.app/users] |

###### Daftar request users

### Categories
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Categories

| Method | URL |
| ------ | ------ |
| POST | [https://fp3-hacktiv8-production.up.railway.app/categories] |
| GET | [https://fp3-hacktiv8-production.up.railway.app/categories] |
| PUT | [https://fp3-hacktiv8-production.up.railway.app/categories/:categoryId] |
| DELETE | [https://fp3-hacktiv8-production.up.railway.app/categories/:categoryId] |

###### Daftar request categories

### Tasks
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Tasks

| Method | URL |
| ------ | ------ |
| POST | [https://fp3-hacktiv8-production.up.railway.app/tasks] |
| GET | [https://fp3-hacktiv8-production.up.railway.app/tasks] |
| PUT | [https://fp3-hacktiv8-production.up.railway.app/tasks/:taskId] |
| DELETE | [https://fp3-hacktiv8-production.up.railway.app/tasks/:taskId] |

###### Daftar request tasks

> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login. Untuk methode PUT dan DELETE diperlukan parameter Id pada URL.
