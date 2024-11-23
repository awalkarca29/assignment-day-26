package handlers

import (
	"assignment-day-26/database"
	"assignment-day-26/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePesanan(c *gin.Context) {
	var pesanan models.Pesanan
	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO Pesanan (id_produk, jumlah, tanggal_pesanan) VALUES (?, ?, ?)`
	result, err := database.DB.Exec(query, pesanan.IDProduk, pesanan.Jumlah, pesanan.TanggalPesanan)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Order created successfully"})
}

func GetPesananByID(c *gin.Context) {
	id := c.Param("id")
	query := `SELECT * FROM Pesanan WHERE id_pesanan = ?`
	row := database.DB.QueryRow(query, id)

	var pesanan models.Pesanan
	if err := row.Scan(&pesanan.ID, &pesanan.IDProduk, &pesanan.Jumlah, &pesanan.TanggalPesanan); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, pesanan)
}
