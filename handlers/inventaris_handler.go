package handlers

import (
	"assignment-day-26/database"
	"assignment-day-26/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventaris(c *gin.Context) {
	id := c.Param("id")
	query := `SELECT * FROM Inventaris WHERE id_produk = ?`
	row := database.DB.QueryRow(query, id)

	var inventaris models.Inventaris
	if err := row.Scan(&inventaris.ID, &inventaris.IDProduk, &inventaris.Jumlah, &inventaris.Lokasi); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	c.JSON(http.StatusOK, inventaris)
}

func UpdateInventaris(c *gin.Context) {
	id := c.Param("id")
	var inventaris models.Inventaris
	if err := c.ShouldBindJSON(&inventaris); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE Inventaris SET jumlah = ? WHERE id_inventaris = ?`
	_, err := database.DB.Exec(query, inventaris.Jumlah, id)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
}
