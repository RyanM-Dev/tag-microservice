package main

import (
	"log"
	myhttp "tagMicroservice/internal/adapters/controllers/http"
	"tagMicroservice/internal/adapters/databases/mysql"
	"tagMicroservice/internal/application/usecases"
	"tagMicroservice/internal/domain/services"
)

func main() {
	var dsn = "root:831374@tcp(db:3306)/tag-microservice?charset=utf8mb4&parseTime=True&loc=Local"
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
