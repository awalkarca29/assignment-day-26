package main

import (
	"assignment-day-26/database"
	"assignment-day-26/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.POST("/produk", handlers.CreateProduk)
	r.GET("/produk/:id", handlers.GetProdukByID)
	r.GET("/produk/kategori/:kategori", handlers.GetProdukByKategori)
	r.PUT("/produk/:id", handlers.UpdateProduk)
	r.DELETE("/produk/:id", handlers.DeleteProduk)
	r.POST("/produk/:id/upload-image", handlers.UploadProdukImage)
	r.GET("/produk/:id/download-image", handlers.DownloadProdukImage)

	r.GET("/inventaris/:id", handlers.GetInventaris)
	r.PUT("/inventaris/:id", handlers.UpdateInventaris)

	r.POST("/pesanan", handlers.CreatePesanan)
	r.GET("/pesanan/:id", handlers.GetPesananByID)

	r.Run(":8080")
}
