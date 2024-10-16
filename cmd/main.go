package main

import (
	"fmt"
	"log"
	"os"
	myhttp "tagMicroservice/internal/adapters/controllers/http"
	"tagMicroservice/internal/adapters/databases/mysql"
	"tagMicroservice/internal/application/usecases"
	"tagMicroservice/internal/domain/services"
)

func main() {
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlRootPassword := os.Getenv("MYSQL_ROOT_PASSWORD")

	if mysqlHost == "" || mysqlDatabase == "" || mysqlUser == "" || mysqlPassword == "" || mysqlRootPassword == "" {
		log.Fatal("Error: One or more required environment variables are missing")
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)

	var mysqlStruct mysql.Mysql

	if err := mysqlStruct.NewMysqlDatabase(dsn); err != nil {
		log.Fatal(err)
	}

	if err := mysqlStruct.AutoMigrate(); err != nil {
		log.Fatal(err)
	}
	db := mysqlStruct.GetDB()
	tagRepo := mysql.NewTagRepository(db)
	taxonomyRepo := mysql.NewTaxonomyRepository(db)
	tagServices := services.NewTagService(tagRepo, taxonomyRepo)
	taxonomyServices := services.NewTaxonomyService(tagRepo, taxonomyRepo)
	tagUsecases := usecases.NewTagUsecases(*tagServices, *taxonomyServices)
	tagHandlers := myhttp.NewTagHandler(tagUsecases)
	ginWebserver := myhttp.NewGinWebServer(*tagHandlers)
	addr := ":8080"
	if err := ginWebserver.RunWebServer(addr); err != nil {
		log.Fatalf("failed to run web server:%v", err)
	}
}
