# Run

1. Setup DB using docker: `docker-compose -f docker-compose.db.yml up -d`

2. Run db migration: `go run migrate\migrate.go` 

3. Run the project: `CompileDaemon -command="./go-gin"`


# Used Libraries

**[GIN](https://gin-gonic.com/docs/quickstart/)**

**[GORM](https://gorm.io/docs/#Install)**

**[CompileDaemon](https://github.com/githubnemo/CompileDaemon#installation)**

**[godotenv](https://github.com/joho/godotenv)**


