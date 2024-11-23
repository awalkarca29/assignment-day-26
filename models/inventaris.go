package models

type Inventaris struct {
	ID       int    `json:"id"`
	IDProduk int    `json:"id_produk"`
	Jumlah   int    `json:"jumlah"`
	Lokasi   string `json:"lokasi"`
}
