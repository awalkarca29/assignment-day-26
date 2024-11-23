package models

type Pesanan struct {
	ID             int    `json:"id"`
	IDProduk       int    `json:"id_produk"`
	Jumlah         int    `json:"jumlah"`
	TanggalPesanan string `json:"tanggal_pesanan"`
}
