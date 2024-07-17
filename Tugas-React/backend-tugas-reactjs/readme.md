# REST API Buku menggunakan Golang

![Logo Golang](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)

## URL Live

Anda dapat mengakses proyek ini secara live di [https://tugas-sb-sanbercode-go-next-2024-732bc1iwr-xnnms-projects.vercel.app/](https://tugas-sb-sanbercode-go-next-2024-732bc1iwr-xnnms-projects.vercel.app/).

## Pengenalan

Assalamualaikum, nama saya Muhammad Ilham, peserta kelas Fullstack Sanber Super Bootcamp Jabar 2024 dengan fokus pada Golang dan Next.js. Saya senang memperkenalkan proyek REST API "Buku" yang telah saya buat menggunakan Golang. Proyek ini memungkinkan Anda untuk mengelola data buku dengan fitur-fitur terstruktur.

## Fitur-fitur

- **Manajemen Buku**: Pengguna dapat menambahkan, melihat, mengedit, dan menghapus buku.
- **Validasi Input**: Validasi data input sebelum disimpan ke database.
- **Swagger API Documentation**: Dokumentasi API untuk memudahkan integrasi dan pengujian.

## Setup Proyek

1. Clone repositori ini.
2. Pastikan Anda memiliki lingkungan pengembangan Golang yang sudah siap.
3. Konfigurasi file `.env` sesuai dengan kebutuhan Anda.
4. Jalankan perintah `go run main.go` untuk memulai server lokal.

## API Routes

### Manajemen Buku

- **POST /books**
  - Deskripsi: Menambahkan buku baru.
  
- **GET /books**
  - Deskripsi: Mendapatkan semua buku yang tersedia.
  
- **GET /books/:id**
  - Deskripsi: Mendapatkan detail buku berdasarkan ID.
  
- **PATCH /books/:id**
  - Deskripsi: Mengupdate buku yang ada.
  
- **DELETE /books/:id**
  - Deskripsi: Menghapus buku.

## Dokumentasi API

Dokumentasi API menggunakan Swagger dan dapat diakses melalui endpoint `/swagger`. Ini memberikan informasi lebih lanjut tentang setiap endpoint dan parameter yang dibutuhkan.

Link Swagger: [https://tugas-sb-sanbercode-go-next-2024-732bc1iwr-xnnms-projects.vercel.app/swagger/index.html](https://tugas-sb-sanbercode-go-next-2024-732bc1iwr-xnnms-projects.vercel.app/swagger/index.html)

## Catatan

Pastikan untuk mengonfigurasi database dan lingkungan proyek sesuai dengan kebutuhan Anda sebelum menjalankannya di lingkungan produksi.

## Kontribusi

Anda dapat berkontribusi pada proyek ini dengan mengirimkan pull request atau melaporkan masalah yang ditemukan.

Terima kasih telah membaca tentang proyek REST API "Buku" saya yang dibangun dengan Golang. Semoga bermanfaat!
