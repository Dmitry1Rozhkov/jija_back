package handler

import (
	"github.com/gin-gonic/gin"
	"jija_back/internal/config"
	"jija_back/internal/domain"
	"net/http"
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

func (h *Handler) GetAtmsInfo(c *gin.Context) {

	//train := [][]float64{}
	//labels := []string{}
	//knn
	//k := 3
	//dm := DMT_EulerMethod
	//w := []float64{}
	//
	//var test [][]float64
	//test = [][]float64{}
	//
	//knn := NewKNNClassifier(k, dm, w)
	//res := knn.Classify(train, test, labels)
	//log.Println(res)

	c.JSONP(http.StatusOK, h.atms)
}

func (h *Handler) GetOfficesInfo(c *gin.Context) {

	c.JSONP(http.StatusOK, h.offices)
}
