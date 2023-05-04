package sale

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/SiwaleK/ProdGroup/model"
	"github.com/gin-gonic/gin"
)

// func GetProdGroup(c *gin.Context) {
// 	var prodGroups []view.Prodgroup
// 	result := config.DB.Find(&prodGroups)
// 	if result.Error != nil {
// 		c.JSON(500, gin.H{"error": "Failed to retrieve prodgroups"})
// 		return
// 	}

// 	if len(prodGroups) == 0 {
// 		c.Status(204)
// 		return
// 	}

// 	c.JSON(200, prodGroups)
// }

// func GetProdGroupByID(c *gin.Context) {
// 	id := c.Param("id")

// 	var Prodgroup []view.Prodgroup
// 	result := config.DB.Where("id = ?", id).First(&Prodgroup)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			c.JSON(404, gin.H{"error": "Prodgroup not found"})
// 			return
// 		}
// 		c.JSON(500, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.JSON(200, Prodgroup)
// }

type ProdgroupRepository interface {
	GetProdgroup(ctx context.Context) (*[]model.Prodgroup, error)
}

type DBProdgroupRepository struct {
	db *sql.DB
}

func NewProdgroupRepository(db *sql.DB) ProdgroupRepository {
	return &DBProdgroupRepository{
		db: db,
	}
}

func (r *DBProdgroupRepository) GetProdgroup(ctx context.Context) (*[]model.Prodgroup, error) {
	prodgroups := &[]model.Prodgroup{}
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM prodgroup")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pg model.Prodgroup
		err := rows.Scan(&pg.Prodgroupid, &pg.Th_name, &pg.En_name)
		if err != nil {
			return nil, err
		}
		*prodgroups = append(*prodgroups, pg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return prodgroups, nil
}

type ProdgroupHandlerImpl struct {
	repo ProdgroupRepository
}

func NewProdgroupHandler(repo ProdgroupRepository) *ProdgroupHandlerImpl {
	return &ProdgroupHandlerImpl{
		repo: repo,
	}
}

func (h *ProdgroupHandlerImpl) GetProdgroup(c *gin.Context) {
	prodgroups, err := h.repo.GetProdgroup(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get prodgroups"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"prodgroups": prodgroups})
}
