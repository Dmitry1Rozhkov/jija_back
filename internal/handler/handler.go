package handler

import (
	"cmp"
	"github.com/gin-gonic/gin"
	"jija_back/internal/config"
	"jija_back/internal/domain"
	"jija_back/pkg/nearest"
	"net/http"
	"slices"
	"strconv"
)

func New(cfg *config.Config, atms []domain.Atm, offices []domain.Office) *Handler {
	return &Handler{
		config:  cfg,
		atms:    atms,
		offices: offices,
	}
}

type Handler struct {
	config  *config.Config
	atms    []domain.Atm
	offices []domain.Office
}

//type Distances struct {
//	distance
//	atm domain.
//}

func (h *Handler) GetAtmsInfo(c *gin.Context) {
	c.JSON(http.StatusOK, h.atms)
}

type node struct {
	index    int
	distance float64
}

func (h *Handler) GetAtmsNearest(c *gin.Context) {
	resultAtms := make([]domain.Atm, 0, 50)

	clientX, err := strconv.ParseFloat(c.Param("x"), 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	clientY, err := strconv.ParseFloat(c.Param("y"), 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var distancesArray = make([]node, 0, len(h.atms))

	for i := 0; i < len(h.atms); i++ {
		distancesArray = append(distancesArray, node{
			index:    i,
			distance: nearest.GetSortedDistances(clientX, clientY, &h.atms[i]),
		})
	}

	slices.SortFunc(distancesArray, func(a, b node) int {
		return cmp.Compare(a.distance, b.distance)
	})

	//fmt.Println(distancesArray)

	for i := 0; i < 50 && i < len(distancesArray); i++ {
		resultAtms = append(resultAtms, h.atms[distancesArray[i].index])
	}

	c.JSON(http.StatusOK, resultAtms)
}

func (h *Handler) GetOfficesInfo(c *gin.Context) {
	c.JSON(http.StatusOK, h.offices)
}

func (h *Handler) GetOfficesNearest(c *gin.Context) {
	resultOffices := make([]domain.Office, 0, 50)

	clientX, err := strconv.ParseFloat(c.Param("x"), 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	clientY, err := strconv.ParseFloat(c.Param("y"), 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var distancesArray = make([]node, 0, len(h.atms))

	for i := 0; i < len(h.offices); i++ {
		distancesArray = append(distancesArray, node{
			index:    i,
			distance: nearest.GetSortedDistances(clientX, clientY, &h.offices[i]),
		})
	}

	slices.SortFunc(distancesArray, func(a, b node) int {
		return cmp.Compare(a.distance, b.distance)
	})

	//fmt.Println(distancesArray)

	for i := 0; i < 50 && i < len(distancesArray); i++ {
		resultOffices = append(resultOffices, h.offices[distancesArray[i].index])
	}

	c.JSON(http.StatusOK, resultOffices)
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{msg})
}
