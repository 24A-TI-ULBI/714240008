# Panduan Kontribusi

Dokumen ini wajib dibaca sebelum mulai mengerjakan modul.

---

## 1. Setup Awal

```bash
# Clone repo
git clone https://github.com/24A-TI-ULBI/siakad-2A.git
cd siakad-2A

# Taruh file .env yang dibagikan via WA di root folder
# Jangan commit file .env ke GitHub

# Jalankan aplikasi
go run main.go
```

---

## 2. Workflow Git

```bash
# 1. Selalu sync dari main sebelum mulai
git pull origin main

# 2. Buat branch sesuai modulmu
git checkout -b feature/modul-XX-namamu
# Contoh: git checkout -b feature/modul-02-keyla

# 3. Kerjakan modulmu

# 4. Stage hanya file milikmu — JANGAN git add .
git add controller/dosenController.go
git add model/dosen.go
git add url/dosenRoute.go
git add frontend/dosen/index.html
git add url/url.go   # hanya kalau kamu sudah tambah baris route-mu

# 5. Commit
git commit -m "feat: tambah CRUD dosen dan jabatan"

# 6. Push ke branch-mu
git push origin feature/modul-XX-namamu

# 7. Buat Pull Request di GitHub ke branch main
```

> ⚠️ Tidak ada yang boleh push langsung ke `main`. Wajib lewat Pull Request.

---

## 3. File yang Kamu Buat

Setiap mahasiswa hanya membuat **4 file baru**:

```
controller/[modul]Controller.go   ← logika handler
model/[modul].go                  ← struct data
url/[modul]Route.go               ← definisi route
frontend/[modul]/index.html       ← tampilan frontend
```

Dan **edit 1 baris** di file ini:

```
url/url.go   ← tambah pemanggilan fungsi route-mu
```

---

## 4. File yang TIDAK Boleh Disentuh

```
main.go
config/
helper/
go.mod
go.sum
.env
```

File-file ini adalah domain maintainer. Kalau ada yang perlu diubah, hubungi maintainer.

---

## 5. Standar Penulisan Kode

### Model
```go
// model/dosen.go
package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dosen struct {
    ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    NIDN string             `bson:"nidn" json:"nidn"`
    Nama string             `bson:"nama" json:"nama"`
    // ... field lainnya
}
```

### Controller
```go
// controller/dosenController.go
package controller

import (
    "backend/helper"
    "backend/model"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllDosen(c *fiber.Ctx) error {
    col := helper.GetCollection("dosen")   // ← wajib pakai helper
    ctx, cancel := helper.GetContext()
    defer cancel()

    // ... logika
    return helper.SuccessResponse(c, list) // ← wajib pakai helper
}
```

### Route
```go
// url/dosenRoute.go
package url

import (
    "backend/controller"
    "github.com/gofiber/fiber/v2"
)

func DosenRoute(app *fiber.App) {
    app.Get("/dosen", controller.GetAllDosen)
    app.Get("/dosen/:nidn", controller.GetDosenByNIDN)
    app.Post("/dosen", controller.CreateDosen)
    app.Put("/dosen/:nidn", controller.UpdateDosen)
    app.Delete("/dosen/:nidn", controller.DeleteDosen)
}
```

### Daftarkan di url/url.go
```go
// Tambah satu baris ini di url/url.go
DosenRoute(app) // ← Keyla tambah ini
```

---

## 6. Aturan Wajib

| Aturan | Keterangan |
|---|---|
| Akses MongoDB | Wajib pakai `helper.GetCollection("nama_collection")` |
| Format response sukses | Wajib pakai `helper.SuccessResponse(c, data)` |
| Format response error | Wajib pakai `helper.ErrorResponse(c, status, "pesan")` |
| ID dokumen | Wajib pakai `primitive.ObjectID` |
| Nama collection | Pakai nama resource dalam huruf kecil, contoh: `"dosen"`, `"jabatan"` |

---

## 7. Contoh Format Response

**Sukses:**
```json
{
  "status": "success",
  "data": { ... }
}
```

**Error:**
```json
{
  "status": "error",
  "message": "deskripsi error"
}
```
