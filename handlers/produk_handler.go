package handlers

import (
	"assignment-day-26/database"
	"assignment-day-26/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CreateProduk(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO Produk (nama, deskripsi, harga, kategori) VALUES (?, ?, ?, ?)`
	result, err := database.DB.Exec(query, produk.Nama, produk.Deskripsi, produk.Harga, produk.Kategori)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Product created successfully"})
}

func GetProdukByID(c *gin.Context) {
	id := c.Param("id")
	query := `SELECT * FROM Produk WHERE id_produk = ?`
	row := database.DB.QueryRow(query, id)

	var produk models.Produk
	if err := row.Scan(&produk.ID, &produk.Nama, &produk.Deskripsi, &produk.Harga, &produk.Kategori); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, produk)
}

func GetProdukByKategori(c *gin.Context) {
	kategori := c.Param("kategori")
	query := `SELECT * FROM Produk WHERE kategori = ?`
	rows, err := database.DB.Query(query, kategori)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	defer rows.Close()

	var produkList []models.Produk
	for rows.Next() {
		var produk models.Produk
		if err := rows.Scan(&produk.ID, &produk.Nama, &produk.Deskripsi, &produk.Harga, &produk.Kategori); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning product"})
			return
		}
		produkList = append(produkList, produk)
	}

	c.JSON(http.StatusOK, produkList)
}

func UpdateProduk(c *gin.Context) {
	id := c.Param("id")
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE Produk SET nama = ?, deskripsi = ?, harga = ?, kategori = ? WHERE id_produk = ?`
	_, err := database.DB.Exec(query, produk.Nama, produk.Deskripsi, produk.Harga, produk.Kategori, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduk(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM Produk WHERE id_produk = ?`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

const imageFolder = "./uploads/"

func UploadProdukImage(c *gin.Context) {
	productID := c.Param("id")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	if _, err := os.Stat(imageFolder); os.IsNotExist(err) {
		err := os.Mkdir(imageFolder, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload folder"})
			return
		}
	}

	fileName := fmt.Sprintf("product_%s%s", productID, filepath.Ext(file.Filename))
	filePath := filepath.Join(imageFolder, fileName)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "file_path": filePath})
}

func DownloadProdukImage(c *gin.Context) {
	productID := c.Param("id")

	fileName := fmt.Sprintf("product_%s.jpg", productID)
	filePath := filepath.Join(imageFolder, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(filePath)
}
