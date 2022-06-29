# Gin-MongoDB API
- an api built with Gin framework and MongoDB as the database
- CRUD operations related to characters of Stranger Things tv series
- MVC pattern

### endpoints
- GET /characters
- GET /character/:characterId
- POST /character
- PUT /character/:characterId
- DELETE /character/:characterId

### project structure
- configs: project configuration files
- controllers: application logics
- models: data and database logics
- responses: files describing the response we want our API to give
- routes: URL pattern and handler information# stranger-things-Gin-api

### deployment
- run `go build -o bin/stranger-things-gin -v .`
- create Procfile to deploy to Heroku
- run `heroku local` to test endpoints locally
- run `heroku create appName`
- 
