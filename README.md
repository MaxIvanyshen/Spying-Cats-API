# Spying Cats API

### Stack
Application is written using Go programming language with use of [Gin](https://github.com/gin-gonic/gin) framework.
I decided to use [SQLite](https://www.sqlite.org/) database to simplify sharing this application's code. Used [GORM](https://gorm.io/index.html) to perform transactions between go application and database

### Running Application
To run this API just go to root directory and run ```go run main.go``` and go to http://localhost:8080 \nIf application requires some modules (which should be already installed), try running these commands before running application:
```bash
    go get github.com/gin-gonic/gin
    go get github.com/gin-gonic/gin
    go get github.com/go-resty/resty/v2
```

### Packages And Directory Files
- [main.go](https://github.com/MaxIvanyshen/Spying-Cats-API/blob/master/main.go) - main file of the application
- [db](https://github.com/MaxIvanyshen/Spying-Cats-API/blob/master/db) - package with reposities and database initalization
- [controllers](https://github.com/MaxIvanyshen/Spying-Cats-API/blob/master/controllers) - package with functions responsible for handling requests to the server
- [models](https://github.com/MaxIvanyshen/Spying-Cats-API/blob/master/models) - package with models (Cat, Mission, Target, Note)
- [logger](https://github.com/MaxIvanyshen/Spying-Cats-API/blob/master/logger) - package with logger initalization 
- [validation](https://github.com/MaxIvanyshen/Spying-Cats-API/blob/master/validation) - package with functions responsible for handling validation with [TheCatAPI](https://api.thecatapi.com/v1/breeds)
- [cats.db](https://github.com/MaxIvanyshen/Spying-Cats-API/blob/master/cats.db) - file with SQLite database
- [spyingCats.postman_collection.json](https://github.com/MaxIvanyshen/Spying-Cats-API/blob/master/spyingCats.postman_collection.json) - postman requests collection


### Endpoints
##### Cats
| Method        | Endpoint      | Action|
| ------------- |:-------------:| -----:|
| GET           | /cats         | list all cats      |
| POST      | /cats      |   create new cat |
| GET      | /cats/:id      |   get information about specific cat by id |
| DELETE      | /cats/:id      |   remove cat by id|
| PATCH     | /cats/:id      |   update cat's salary|

##### Missions
| Method        | Endpoint      | Action|
| ------------- |:-------------:| -----:|
| GET           | /missions         | list all missions      |
| POST      | /missions      |   create new mission |
| GET      | /missions/:id      |   get information about specific mission by id |
| DELETE      | /missions/:id      |   remove mission by id|
| PATCH     | /missions/:id      |   mark mission as complete|
| PATCH     | /missions/assign/:missionId/:catId      |   assing mission to cat using their id's|

##### Targets
| Method        | Endpoint      | Action|
| ------------- |:-------------:| -----:|
| POST           | /targets/:missionId         | add target to mission by id      |
| PATCH     | /targets/:id      |   mark target as complete|
| PATCH     | /targets/notes/:id      |   add note to target by id|
| DELETE     | /targets/:missionId/:targetId      |   remove target from mission |

### Object Examples
- creating a cat:
```javascript
{
    "name": "Max",
    "yearsOfExperience": 3,
    "salary": 2000,
    "breed": "Abyssinian"
}
```

- updating cat's salary:
```javascript
{
    "salary": 5000
}
```

- creating a mission:
```javascript
{
    "cat": null,
    "targets": [{
        "name": "target 1",
        "country": "Ukraine",
        "notes": [],
        "complete": false
    }],
    "complete": false
}
```

- creating a target:
```javascript
{
    "name": "target 2",
    "country": "Spain",
    "notes": [{
        "content": "this is a nice country"
    }],
    "complete": false
}
```

- adding a note to a target:
```javascript
{
    "notes": "this is a note"
}
```
