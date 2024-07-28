# Frontend React untuk "News" Project

![React Logo](https://reactjs.org/logo-og.png)

## URL Live

Anda dapat mengakses proyek ini secara live di [https://news-front-end-five.vercel.app/](https://news-front-end-five.vercel.app/).

## Pengenalan

Assalamualaikum, nama saya Muhammad Ilham, peserta kelas Fullstack Sanber Super Bootcamp Jabar 2024 dengan fokus pada Golang dan Next.js. Saya senang memperkenalkan frontend React untuk proyek "news" yang telah saya buat. Proyek ini berfungsi sebagai antarmuka pengguna untuk mengelola berita dengan menggunakan React dan berbagai teknologi modern.

## Fitur-fitur

- **Autentikasi Pengguna**: Pengguna dapat mendaftar, login, dan mengubah kata sandi mereka.
- **Manajemen Berita**: Pengguna dapat melihat daftar berita dan detail berita.
- **Profil Pengguna**: Pengguna dapat melihat dan memperbarui profil mereka.
- **Manajemen Pengguna**: Halaman untuk mengelola pengguna (dengan akses admin).
- **Routing**: Navigasi yang mudah antar halaman menggunakan React Router.
- **State Management**: Pengelolaan state global dengan Context API.
- **HTTP Requests**: Mengambil dan mengirim data ke API menggunakan Axios.
- **Proteksi Rute**: Halaman tertentu hanya dapat diakses oleh pengguna yang terautentikasi.

## Struktur Proyek

- **`src/context/AuthContext.jsx`**: Menyediakan Context untuk otentikasi pengguna.
- **`src/context/NewsContext.jsx`**: Menyediakan Context untuk data berita.
- **`src/context/CommentsContext.jsx`**: Menyediakan Context untuk komentar berita.
- **`src/context/UserContext.jsx`**: Menyediakan Context untuk data profil pengguna.
- **`src/components/ProtectedRoute.jsx`**: Komponen untuk melindungi rute dari akses pengguna yang tidak terautentikasi.
- **`src/pages/LoginPage.jsx`**: Halaman untuk login pengguna.
- **`src/pages/RegisterPage.jsx`**: Halaman untuk pendaftaran pengguna baru.
- **`src/pages/ChangePasswordPage.jsx`**: Halaman untuk mengubah kata sandi pengguna.
- **`src/pages/Home.jsx`**: Halaman utama yang menampilkan daftar berita.
- **`src/pages/NewsPage.jsx`**: Halaman untuk menampilkan berita.
- **`src/pages/NewsDetail.jsx`**: Halaman untuk menampilkan detail berita.
- **`src/pages/ProfilePage.jsx`**: Halaman profil pengguna.
- **`src/pages/UserManagementPage.jsx`**: Halaman untuk mengelola pengguna.
- **`src/components/News/NewsDetail.jsx`**: Komponen untuk menampilkan detail berita.

## Setup Proyek

Untuk menjalankan proyek ini, pastikan Anda memiliki Node.js dan npm atau yarn terinstal. Clone repositori ini dan jalankan `npm install` atau `yarn install` untuk menginstal dependensi. Kemudian, jalankan `npm start` atau `yarn start` untuk memulai server pengembangan lokal.

```bash
git clone [URL_REPOSITORI]
cd [NAMA_FOLDER]
npm install
npm start
