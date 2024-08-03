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

func NewTagHandler(tagUsecases usecases.TagUsecase) *TagHandler {
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

func (h *TagHandler) ApproveTag(c *gin.Context) {
	var tagIDReq requests.TagIDReq
	if err := c.BindJSON(&tagIDReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tagUsecases.ApproveTag(tagIDReq.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "requested tag successfully approved")

}

func (h *TagHandler) RejectTag(c *gin.Context) {
	var tagIDReq requests.TagIDReq
	if err := c.BindJSON(&tagIDReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tagUsecases.RejectTag(tagIDReq.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "requested tag successfully Rejected")

}

func (h *TagHandler) MergeTags(c *gin.Context) {
	var mergeTagsReq requests.MergeTagsReq
	if err := c.BindJSON(&mergeTagsReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tagUsecases.MergeTags(mergeTagsReq.FromTagID, mergeTagsReq.ToTagID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "requested tags successfully merged")
}

func (h *TagHandler) AddTaxonomy(c *gin.Context) {
	var addTaxonomyReq requests.AddTaxonomyReq
	if err := c.BindJSON(&addTaxonomyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tagUsecases.AddTaxonomy(addTaxonomyReq.FromTagID, addTaxonomyReq.ToTagID, addTaxonomyReq.RelationshipKind, addTaxonomyReq.State); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "requested taxonomy was added successfully")

}

func (h *TagHandler) SetTaxonomy(c *gin.Context) {
	var setTaxonomyReq requests.SetTaxonomyReq
	if err := c.BindJSON(&setTaxonomyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.tagUsecases.SetTaxonomy(setTaxonomyReq.TaxonomyID, setTaxonomyReq.RelationshipKind); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "requested taxonomy was set successfully")
}

func (h *TagHandler) GetRelatedTagsByKey(c *gin.Context) {
	var getRelatedTagsByKeyReq requests.GetRelatedTagsByKeyReq
	if err := c.BindJSON(&getRelatedTagsByKeyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tags, err := h.tagUsecases.GetRelatedTagsByKey(getRelatedTagsByKeyReq.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"tags": tags,
	})

}

func (h *TagHandler) GetRelatedTagsByID(c *gin.Context) {
	var tagIDReq requests.TagIDReq
	if err := c.BindJSON(&tagIDReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tags, err := h.tagUsecases.GetRelatedTagsByID(tagIDReq.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"tags": tags,
	})
}

func (h *TagHandler) GetRelatedTagsByTitleAndKey(c *gin.Context) {
	var getRelatedTagsByTitleAndKey requests.GetRelatedTagsByTitleAndKeyReq
	if err := c.BindJSON(&getRelatedTagsByTitleAndKey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tags, err := h.tagUsecases.GetRelatedTagsByTitleAndKey(getRelatedTagsByTitleAndKey.Title, getRelatedTagsByTitleAndKey.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"tags": tags,
	})

}
