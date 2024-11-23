package models

type Produk struct {
	ID        int     `json:"id"`
	Nama      string  `json:"nama"`
	Deskripsi string  `json:"deskripsi"`
	Harga     float64 `json:"harga"`
	Kategori  string  `json:"kategori"`
}
