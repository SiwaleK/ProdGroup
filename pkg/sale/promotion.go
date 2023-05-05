package sale

import (
	"net/http"

	"github.com/SiwaleK/ProdGroup/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// // quiriesได้หมดโดยที่ถ้าโดนเปลี่ยนชื่อdatabaseไม่ต้องแก้ได้หมด
// func GetPromotionByID(c *gin.Context) {
// 	id := c.Param("id")

// 	var Promotion []view.Promotion // ให้promotionid มาจาก struct เผื่อเปลี่ยนfield
// 	// go ควรใช้class ไปลง database ได้
// 	// controller code น้อยๆ
// 	result := config.DB.Where("promotionid = ?", id).Find(&Promotion)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			c.JSON(404, gin.H{"error": "Promotion not found"})
// 			return
// 		}
// 		c.JSON(500, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.JSON(200, Promotion)
// }

// type Service struct {
// 	queries *db.Queries
// }

// func NewService(queries *db.Queries) *Service {
// 	return &Service{queries: queries}
// }

// func (service *Service) RegisterHandlers(router *gin.Engine) {
// 	router.GET("/Promotion", service.GetPromotion2)
// }

// type apiPromotion struct {
// 	PromotionID   string          `json:"PromotionID,omitempty"`
// 	PromotionType int32           `json:"PromotionType"`
// 	StartDate     time.Time       `json:"StartDate,omitempty"`
// 	EndDate       time.Time       `json:"EndDate,omitempty"`
// 	Description   string          `json:"Description,omitempty"`
// 	Conditions    json.RawMessage `json:"Conditions,omitempty"`
// }

// func fromDB(promotion db.Promotion) *apiPromotion {
// 	return &apiPromotion{
// 		PromotionID:   promotion.PromotionID.String,
// 		PromotionType: promotion.PromotionType,
// 		StartDate:     promotion.StartDate.(time.Time),
// 		EndDate:       promotion.EndDate.(time.Time),
// 		Description:   promotion.Description.String,
// 		Conditions:    promotion.Conditions,
// 	}
// }

// func (service *Service) GetPromotion2(ctx *gin.Context) {
// 	Get Promotion
// 	promotions, err := service.queries.GetPromotion(context.Background())
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}
// 		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
// 		return
// 	}

//		Build Response
//		response := make([]*apiPromotion, len(promotions))
//		for i, p := range promotions {
//			response[i] = fromDB(p)
//		}
//		ctx.IndentedJSON(http.StatusOK, response)
//	}

// type PromotionHandler struct {
// 	repo *repository.PromotionRepository
// }

// func NewPromotionHandler(repo *repository.PromotionRepository) *PromotionHandler {
// 	return &PromotionHandler{repo: repo}
// }

// func (h *PromotionHandler) GetPromotionByID(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.ParseUint(idStr, 10, 32)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
// 		return
// 	}

// 	promotion, err := h.repo.GetPromotionByID(c.Request.Context(), uint(id))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if promotion == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "promotion not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, promotion)
// }

// func (h *PromotionHandler) GetPromotionByID(c *gin.Context) {
// 	id := c.Param("id")
// 	parsedID, err := uuid.Parse(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
// 		return
// 	}

// 	idStr := parsedID.String()
// 	idUint, err := strconv.ParseUint(idStr, 10, 32)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID123"})
// 		return
// 	}

// 	promotion, err := h.repo.GetPromotionByID(c.Request.Context(), uint(idUint))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get promotion"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"promotion": promotion})
// }

type PromotionHandler struct {
	repo repository.PromotionRepository
}

func NewPromotionHandler(repo repository.PromotionRepository) *PromotionHandler {
	return &PromotionHandler{repo: repo}
}

func (h *PromotionHandler) GetPromotionByID(c *gin.Context) {
	id := c.Param("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid promotion ID123"})
		return
	}

	promotion, err := h.repo.GetPromotionByID(c.Request.Context(), parsedID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get promotion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"promotion": promotion})
}
