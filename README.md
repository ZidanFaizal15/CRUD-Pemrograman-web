# CRUD-Pemrograman-web
# E-Commerce Barang Bekas

Aplikasi web sederhana untuk jual beli barang bekas. Dibuat menggunakan bahasa pemrograman **Go (Golang)**, dengan framework **Gin**, database **MySQL**, dan tampilan antarmuka menggunakan **Bootstrap**.

---

## 🧱 Fitur Utama

* ✅ **CRUD Produk** (Tambah, Lihat, Edit, Hapus)
* ✅ **Upload Gambar Produk**
* ✅ **Login & Registrasi User**
* ✅ **Export ke PDF**
* ✅ **Bootstrap untuk Tampilan Responsive**

---

## ⚙️ Struktur Folder

```
barang-bekas/
├── config/             # Koneksi database
│   └── db.go
├── controllers/        # Logic aplikasi (produk, login)
│   ├── auth_controller.go
│   └── product_controller.go
├── models/             # Struktur tabel database
│   ├── product.go
│   └── user.go
├── routes/             # Daftar rute aplikasi
│   └── routes.go
├── static/             # Asset statis (CSS, JS, gambar)
│   └── css/
│       └── style.css
├── templates/          # HTML template
│   ├── form.html
│   ├── index.html
│   ├── login.html
│   └── register.html
├── uploads/            # Gambar produk yang diunggah
├── main.go             # Entry point aplikasi
└── go.mod              # Dependensi Go
```

---

## 💡 Cara Menjalankan

### 1. Setup Database

Buat database baru di **phpMyAdmin**:

```sql
CREATE DATABASE db_barang_bekas;
```

Buat dua tabel: `users` dan `products`. Atau biarkan Go membuatnya otomatis melalui `AutoMigrate()`.

### 2. Konfigurasi Koneksi

Edit file `config/db.go`:

```go
dsn := "root:@tcp(127.0.0.1:3306)/db_barang_bekas?parseTime=true"
```

### 3. Jalankan Aplikasi

```
go run main.go
```

Buka di browser: [http://localhost:8080](http://localhost:8080)

---

## 🖼 Contoh Halaman

* **Halaman Login** dengan opsi Register
* **Daftar Produk** dengan tombol Edit, Hapus, Export
* **Form Produk** untuk Tambah & Edit

---

## 📦 Dependensi

* [Gin](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io)
* [GoFPDF](https://github.com/jung-kurt/gofpdf) - untuk cetak PDF
* [Excelize](https://github.com/xuri/excelize) - untuk export Excel (Tombol tidak tersedia untuk sekarang)
* [Bootstrap CDN](https://getbootstrap.com)

---

## 🙌 Kontributor

* Zidan Faizal

---

## 📄 Lisensi

Project ini bebas digunakan untuk keperluan belajar dan tugas akademik.
