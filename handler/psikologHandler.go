package handler

import (
	"github.com/gin-gonic/gin"
	"project-individu-go-react/auth"
	"project-individu-go-react/entity"
	"project-individu-go-react/helper"
	"project-individu-go-react/layer/psikolog"
	"strconv"
)

type psikologHandler struct {
	psikologService psikolog.Service
	authService     auth.Service
}

func NewPsikologHandler(psikologService psikolog.Service, authService auth.Service) *psikologHandler {
	return &psikologHandler{psikologService, authService}
}

func (h *psikologHandler) ShowPsikologHandler(c *gin.Context) {
	psikologs, err := h.psikologService.GetAllPsikolog()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all psikolog", 200, "status OK", psikologs)
	c.JSON(200, response)
}

func (h *psikologHandler) CreatePsikologHandler(c *gin.Context) {
	var inputPsikolog entity.PsikologInput

	if err := c.ShouldBindJSON(&inputPsikolog); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newPsikolog, err := h.psikologService.SaveNewPsikolog(inputPsikolog)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new Psikolog", 201, "status Created", newPsikolog)
	c.JSON(201, response)
}

func (h *psikologHandler) GetPsikologByIDHandler(c *gin.Context) {
	id := c.Params.ByName("psikolog_id")

	psikologs, err := h.psikologService.GetPsikologByID(id)
	if err != nil {
		responseError := helper.APIResponse("error bad request psikolog ID", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get psikolog by ID", 200, "success", psikologs)
	c.JSON(200, response)
}

func (h *psikologHandler) DeletePsikologByIDHandler(c *gin.Context) {
	id := c.Params.ByName("psikolog_id")

	psikologs, err := h.psikologService.DeletePsikologByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete psikolog", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete psikolog by ID", 200, "success", psikologs)
	c.JSON(200, response)
}

func (h *psikologHandler) UpdatePsikologByIDHandler(c *gin.Context) {
	id := c.Params.ByName("psikolog_id")

	var updatePsikologInput entity.UpdatePsikologInput

	if err := c.ShouldBindJSON(&updatePsikologInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	idParam, _ := strconv.Atoi(id)

	// authorization userid dari params harus sama dengan user id yang login
	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	psikologs, err := h.psikologService.UpdatePsikologByID(id, updatePsikologInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update psikolog by ID", 200, "success", psikologs)
	c.JSON(200, response)
}
