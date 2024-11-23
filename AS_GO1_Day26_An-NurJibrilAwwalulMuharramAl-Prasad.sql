CREATE TABLE Produk (
    id_produk INT PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR(255) NOT NULL,
    deskripsi TEXT,
    harga DECIMAL(10, 2) NOT NULL,
    kategori VARCHAR(100)
);

CREATE TABLE Inventaris (
    id_inventaris INT PRIMARY KEY AUTO_INCREMENT,
    id_produk INT,
    jumlah INT NOT NULL,
    lokasi VARCHAR(255),
    FOREIGN KEY (id_produk) REFERENCES Produk(id_produk)
);

CREATE TABLE Pesanan (
    id_pesanan INT PRIMARY KEY AUTO_INCREMENT,
    id_produk INT,
    jumlah INT NOT NULL,
    tanggal_pesanan DATE NOT NULL,
    FOREIGN KEY (id_produk) REFERENCES Produk(id_produk)
);

INSERT INTO Produk (nama, deskripsi, harga, kategori) VALUES
('Laptop ABC', 'Laptop dengan spesifikasi tinggi untuk profesional', 15000000.00, 'Elektronik'),
('Smartphone XYZ', 'Smartphone dengan kamera 108 MP', 8000000.00, 'Elektronik'),
('Meja Kerja', 'Meja kerja kayu untuk kantor', 1200000.00, 'Furniture');

INSERT INTO Inventaris (id_produk, jumlah, lokasi) VALUES
(1, 50, 'Gudang A'),
(2, 100, 'Gudang B'),
(3, 30, 'Gudang C');

INSERT INTO Pesanan (id_produk, jumlah, tanggal_pesanan) VALUES
(1, 2, '2024-11-01'),
(2, 5, '2024-11-02'),
(3, 1, '2024-11-03');

-- Menampilkan Semua Produk
SELECT * FROM Produk;

-- Menampilkan Stok Produk di Setiap Lokasi
SELECT p.nama, i.lokasi, i.jumlah
FROM Inventaris i
JOIN Produk p ON i.id_produk = p.id_produk;

-- Menampilkan Detail Pesanan
SELECT p.nama, ps.jumlah, ps.tanggal_pesanan
FROM Pesanan ps
JOIN Produk p ON ps.id_produk = p.id_produk;

-- Agregasi Total Pesanan per Produk
SELECT p.nama, SUM(ps.jumlah) AS total_pesanan
FROM Pesanan ps
JOIN Produk p ON ps.id_produk = p.id_produk
GROUP BY p.id_produk;

-- Menampilkan Stok per Lokasi
SELECT lokasi, SUM(jumlah) AS total_stok
FROM Inventaris
GROUP BY lokasi;