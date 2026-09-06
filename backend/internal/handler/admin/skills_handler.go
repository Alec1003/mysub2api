package admin

import (
	"github.com/Wei-Shaw/sub2api/internal/pkg/response"
	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SkillHandler struct{ svc *service.SkillService }

func NewSkillHandler(svc *service.SkillService) *SkillHandler { return &SkillHandler{svc: svc} }

type skillRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	EffectImage string `json:"effect_image"`
	DriveURL    string `json:"drive_url"`
	Status      string `json:"status"`
}

func (h *SkillHandler) List(c *gin.Context) {
	items, err := h.svc.AdminList(c.Request.Context())
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, items)
}
func (h *SkillHandler) Create(c *gin.Context) {
	var req skillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	s := &service.Skill{Name: req.Name, Description: req.Description, EffectImage: req.EffectImage, DriveURL: req.DriveURL, Status: req.Status}
	if err := h.svc.Create(c.Request.Context(), s); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, s)
}
func (h *SkillHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req skillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	s := &service.Skill{ID: id, Name: req.Name, Description: req.Description, EffectImage: req.EffectImage, DriveURL: req.DriveURL, Status: req.Status}
	if err := h.svc.Update(c.Request.Context(), s); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	response.Success(c, s)
}
func (h *SkillHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, gin.H{"deleted": true})
}
