package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"app-penjualan/internal/dto"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
)

type AuthHandler struct {
	svc *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{svc: s}
}

// Login handles user authentication.
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := utils.Validate.Struct(req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	tok, name, role, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}
	utils.OK(c, dto.LoginResponse{Token: tok, Name: name, Role: role})
}

// CreateUser handles creating a new user (OWNER only).
type createUserReq struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   uint   `json:"role_id" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := utils.Validate.Struct(req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.CreateUser(req.Name, req.Email, req.Password, req.RoleID, req.Status); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	utils.Created(c, gin.H{"message": "user created"})
}

// ListUsers retrieves and returns a list of all users.
func (h *AuthHandler) ListUsers(c *gin.Context) {
	users, err := h.svc.ListUsers()
	if err != nil {
		utils.ServerError(c, err)
		return
	}
	utils.OK(c, users)
}

// UpdateUser handles updating an existing user.
func (h *AuthHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid user ID")
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := utils.Validate.Struct(req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.svc.UpdateUser(uint(id), req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	utils.OK(c, gin.H{"message": "user updated"})
}

// DeleteUser handles deleting a user.
func (h *AuthHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.BadRequest(c, "Invalid user ID")
		return
	}

	if err := h.svc.DeleteUser(uint(id)); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	utils.OK(c, gin.H{"message": "user deleted"})
}

// ListRoles retrieves and returns a list of all roles.
func (h *AuthHandler) ListRoles(c *gin.Context) {
	roles, err := h.svc.ListRoles()
	if err != nil {
		utils.ServerError(c, err)
		return
	}
	utils.OK(c, roles)
}