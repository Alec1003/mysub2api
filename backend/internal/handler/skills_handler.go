package handler

import (
	"github.com/Wei-Shaw/sub2api/internal/pkg/response"
	"github.com/Wei-Shaw/sub2api/internal/server/middleware"
	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SkillHandler struct{ svc *service.SkillService }

func NewSkillHandler(svc *service.SkillService) *SkillHandler { return &SkillHandler{svc: svc} }
func (h *SkillHandler) List(c *gin.Context) {
	sub, ok := middleware.GetAuthSubjectFromContext(c)
	if !ok {
		response.Unauthorized(c, "user not found")
		return
	}
	items, ent, err := h.svc.List(c.Request.Context(), sub.UserID)
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, gin.H{"items": items, "entitlement": ent})
}
func (h *SkillHandler) Redeem(c *gin.Context) {
	sub, ok := middleware.GetAuthSubjectFromContext(c)
	if !ok {
		response.Unauthorized(c, "user not found")
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		response.BadRequest(c, "invalid skill id")
		return
	}
	s, err := h.svc.Redeem(c.Request.Context(), sub.UserID, id)
	if err != nil {
		response.ErrorFrom(c, err)
		return
	}
	response.Success(c, s)
}

var _ = http.StatusOK
