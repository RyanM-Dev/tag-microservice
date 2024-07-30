package http

import (
	"net/http"
	"tagMicroservice/internal/adapters/controllers/requests"
	"tagMicroservice/internal/adapters/controllers/response"
	"tagMicroservice/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagUsecases usecases.TagUsecase
}

func (h *TagHandler) NewTagHandler(tagUsecases usecases.TagUsecase) *TagHandler {
	return &TagHandler{
		tagUsecases: tagUsecases,
	}
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var createTagReq requests.CreateTagReq

	if err := c.BindJSON(&createTagReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := requests.CreateTagReqToTagEntity(createTagReq)
	if err := h.tagUsecases.CreateTag(tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "Tag created successfully"})

}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	var updateTagReq requests.UpdateTagReq
	if err := c.BindJSON(&updateTagReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag := requests.UpdateTagReqToTagEntity(updateTagReq)

	if err := h.tagUsecases.UpdateTag(tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Tag updated successfully"})

}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	var deleteTagReq requests.TagIDReq
	if err := c.BindJSON(&deleteTagReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tagUsecases.DeleteTag(deleteTagReq.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "Tag deleted successfully"})
}

func (h *TagHandler) GetTagByID(c *gin.Context) {
	var tagIDReq requests.TagIDReq
	if err := c.BindJSON(&tagIDReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tag, err := h.tagUsecases.GetTagByID(tagIDReq.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tagRes := response.DomainToTagRes(*tag)
	c.JSON(200, gin.H{"tag": tagRes})
}
