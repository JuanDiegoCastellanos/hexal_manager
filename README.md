
![Logo](https://github.com/JuanDiegoCastellanos/hexal_manager/blob/main/banner.png)


##
| [![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/) | ![GitHub commit activity (branch)](https://img.shields.io/github/commit-activity/t/JuanDiegoCastellanos/hexal_manager/main) |  ![Static Badge](https://img.shields.io/badge/go-main_language-blue) | ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JuanDiegoCastellanos/hexal_manager) | ![GitHub Tag](https://img.shields.io/github/v/tag/JuanDiegoCastellanos/hexal_manager) | ![GitHub Release](https://img.shields.io/github/v/release/JuanDiegoCastellanos/hexal_manager)  | ![GitHub Repo stars](https://img.shields.io/github/stars/JuanDiegoCastellanos/hexal_manager?style=flat-square)  |
|--------------|--------------|--------------|--------------|--------------|--------------|--------------|


## Appendix

This is the backend layer of hexal manager system, hexal manager is a enterprise software to manage all the things about residential management, things like issues, tasks, incidents, meets, notes and everything that someone wants to manage with this software are possible 


## Features

- Concurrency
- Mutex/ RWLock and RWUnlock
- Multiple persistence drivers
- Highly available services  
- Clean Architecture
- SQLC techniques
- Myslq, Postgresql and MongoDB
- Microservices
- Docker


## Dependecies
- require
- fake "only for tests"
- testify
- lib/pq
- net
- encoding/json
- gin
- gorilla/websocket
- gorilla/mux
## Deployment

To deploy this project run

1- First download all the dependecies:
```bash
  go mod download
```
2- Compile the app:
```bash
  go build main.go
```
3- Run the app:
```bash
  ./hexal_manager
```

And always make sure about the dependecies version, sometimes one project might fail because of a deprecated dependency.
## API Reference
I am going to show you just 2 enpoints to consume the API.
#### Get all issues

```http
  GET /api/${company_name}/${project}/issues
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `x-token` | `string` | **Required**. Your Auth token |
| `company_name` | `string` | **Required**. Name of the company or enterprise |
| `project` | `string` | **Required**. It is required to give a project |


#### Get issue

```http
  GET  /api/${company_name}/${project}/issues/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `x-token` | `string` | **Required**. Your Auth token |
| `company_name` | `string` | **Required**. Name of the company or enterprise |
| `project` | `string` | **Required**. It is required to give a project |
| `id`      | `string` | **Required**. Id of issue to fetch |




## Support

For support, email juandiego.castellanosjerez@gmail.com

