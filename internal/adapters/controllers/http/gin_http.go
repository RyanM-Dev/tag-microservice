package http

import "github.com/gin-gonic/gin"

type GinWebServer struct {
	router     *gin.Engine
	tagHandler TagHandler
}

func NewGinWebServer() *GinWebServer {
	var tagHandler TagHandler
	router := gin.New()
	router.POST("/api/tags", tagHandler.CreateTag)
	router.PUT("/api/tags", tagHandler.UpdateTag)
	router.DELETE("/api/tags", tagHandler.DeleteTag)
	router.GET("/api/tags/id", tagHandler.GetTagByID)
	router.POST("/api/tags/approve", tagHandler.ApproveTag)
	router.POST("/api/tags/reject", tagHandler.RejectTag)
	router.POST("/api/tags/merge", tagHandler.MergeTags)

	router.POST("/api/taxonomies", tagHandler.AddTaxonomy)
	router.PUT("/api/taxonomies", tagHandler.SetTaxonomy)

	router.GET("/api/related-tags/key", tagHandler.GetRelatedTagsByKey)
	router.GET("/api/related-tags/id", tagHandler.GetRelatedTagsByID)
	router.POST("/api/related-tags/search", tagHandler.GetRelatedTagsByTitleAndKey)

	return &GinWebServer{router: router, tagHandler: tagHandler}
}

func (s *GinWebServer) RunWebServer(addr string) error {
	err := s.router.Run(addr)
	return err
}
