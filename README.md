# backend-afmi-ruri-fandho

Repository ini berisi implementasi backend API untuk marketplace "merah kuning hijau", yang menghubungkan merchant dan customer, serta menyediakan fitur diskon dan bebas ongkir berdasarkan ketentuan tertentu.

## ✨ Fitur

### 🔐 Authentication
- Login & Register dengan JWT (untuk Merchant dan Customer)

### 🛍️ Merchant
- CRUD Produk (Create, Read, Update, Delete)
- Lihat daftar customer yang membeli produk

### 👥 Customer
- Lihat daftar produk
- Melakukan pembelian
- Otomatis mendapat:
  - **Bebas Ongkir** jika total produk > Rp15.000
  - **Diskon 10%** jika total produk > Rp50.000

---

## 🚀 Cara Instalasi & Menjalankan Aplikasi

### 1. Clone Repository
git clone https://github.com/your-username/backend-nama-lengkap.git
cd backend-nama-lengkap

### 2. Install Depedency
go mod init

### 3. Menjalankan Project
go run cmd/main.go


```bash