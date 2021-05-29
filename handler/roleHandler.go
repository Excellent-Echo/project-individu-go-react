package handler

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/auth"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/layer/role"
)

type roleHandler struct {
	roleService role.Service
	authService auth.Service
}

func NewRoleHander(roleService role.Service, authService auth.Service) *roleHandler {
	return &roleHandler{roleService, authService}
}

func (h *roleHandler) ShowRoleHandler(c *gin.Context) {
	roles, err := h.roleService.GetAllRole()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error internal server",
			"error":   err.Error(),
		})
		return
	}

	response := helper.APIResponse(
		"Success Get All Roles Data",
		200,
		"Status OK",
		roles,
	)

	c.JSON(200, response)
}

func (h *roleHandler) CreateRoleHandler(c *gin.Context) {
	var NewRoleInput entity.RoleInput

	if err := c.ShouldBindJSON(&NewRoleInput); err != nil {
		splitError := helper.SplitErrorInformation(err)

		responseError := helper.APIResponse(
			"Input data required",
			400,
			"bad request",
			gin.H{
				"error": splitError,
			},
		)

		c.JSON(400, responseError)
		return
	}

	newRole, err := h.roleService.SaveNewRole(NewRoleInput)

	if err != nil {
		responseError := helper.APIResponse(
			"Internal server error",
			500,
			"error",
			gin.H{
				"error": err.Error(),
			},
		)

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(
		"Success Create new Roles Data",
		201,
		"Status Created",
		newRole,
	)

	c.JSON(201, response)
}

func (h *roleHandler) GetRoleByIDHandler(c *gin.Context) {
	id := c.Params.ByName("role_id")

	roleByID, err := h.roleService.GetRoleByID(id)

	if err != nil {
		responseError := helper.APIResponse("Error bad request get role id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("Success get role by id", 200, "success", roleByID)
	c.JSON(200, response)
}

func (h *roleHandler) GetandDeleteRoleByIDHandler(c *gin.Context) {
	id := c.Params.ByName("role_id")

	roleDelete, err := h.roleService.GetandDeleteRoleByID(id)

	if err != nil {
		responseError := helper.APIResponse("Error delete role id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	responseSuccess := helper.APIResponse("Success delete role id", 200, "Delete OK", roleDelete)
	c.JSON(200, responseSuccess)
}

func (h *roleHandler) GetandUpdateRoleByIDHandler(c *gin.Context) {
	id := c.Params.ByName("role_id")

	var roleUpdate entity.RoleInputUpdate

	if err := c.ShouldBindJSON(&roleUpdate); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": splitError})

		c.JSON(400, responseError)
		return
	}

	roles, err := h.roleService.GetandUpdateRoleByID(id, roleUpdate)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update role by id", 200, "success", roles)
	c.JSON(200, response)
}
