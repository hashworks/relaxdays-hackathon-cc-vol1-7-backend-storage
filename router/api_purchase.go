package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashworks/relaxdays-hackathon-vol1-backend/models"
)

// API endpoint that returns all saved purchases
//
// @Summary Returns all saved purchases
// @Produce json
// @Success 200 {object} models.Purchase
// @Router /purchase [get]
// @Tags Purchase
func (s Server) PurchaseGet(c *gin.Context) {
	purchaseRows, err := s.DotSelect.Query(s.DB, "select-purchase")
	defer purchaseRows.Close()

	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	allPurchases := make([]models.Purchase, 0)

	for purchaseRows.Next() {
		var purchase models.Purchase
		err := purchaseRows.Scan(&purchase.Vendor, &purchase.ArticleID, &purchase.Bulk)
		if err != nil {
			s.internalServerError(c, err.Error())
			return
		}
		allPurchases = append(allPurchases, purchase)
	}

	c.JSON(http.StatusOK, allPurchases)
}

// API endpoint that saves a purchase
//
// @Summary Save a purchase
// @Success 204
// @Failure 400 {} {} "Invalid purchase"
// @Accept json
// @Param purchase body models.Purchase true "Purchase to save"
// @Router /purchase [post]
// @Tags Purchase
func (s Server) PurchaseSave(c *gin.Context) {
	var purchase models.Purchase
	c.BindJSON(&purchase)

	if !purchase.IsValid() {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := s.DotAlter.Exec(s.DB, "insert-purchase", purchase.Vendor, purchase.ArticleID, purchase.Bulk)
	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
