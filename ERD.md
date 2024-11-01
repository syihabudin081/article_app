# Penjelasan Entity Relationship Diagram (ERD) - "Article Online"

Diagram ini menggambarkan struktur database untuk aplikasi **Article Online** dengan beberapa tabel utama yaitu: `articles`, `comments`, `users`, dan `roles`.

## Tabel dan Relasi

### 1. Tabel `articles`
- **Deskripsi**: Tabel ini menyimpan informasi tentang artikel yang dibuat oleh pengguna.
- **Kolom**:
    - `uuid`: Kunci utama unik untuk setiap artikel.
    - `created_at`: Waktu artikel dibuat.
    - `updated_at`: Waktu artikel diperbarui.
    - `deleted_at`: Waktu artikel dihapus (jika dihapus).
    - `title`: Judul artikel.
    - `category`: Kategori artikel.
    - `content`: Isi artikel.
    - `author_id`: Referensi ke tabel `users` untuk menunjukkan penulis artikel.
    - `status`: Status artikel (misalnya, dipublikasikan, disimpan sebagai draft, dll.).
- **Relasi**:
    - `author_id` berelasi dengan `users(uuid)`, artinya setiap artikel memiliki seorang penulis.

### 2. Tabel `comments`
- **Deskripsi**: Tabel ini menyimpan komentar yang diberikan oleh pengguna pada artikel.
- **Kolom**:
    - `uuid`: Kunci utama unik untuk setiap komentar.
    - `created_at`: Waktu komentar dibuat.
    - `updated_at`: Waktu komentar diperbarui.
    - `deleted_at`: Waktu komentar dihapus (jika dihapus).
    - `article_id`: Referensi ke tabel `articles` untuk menunjukkan artikel yang dikomentari.
    - `content`: Isi komentar.
    - `user_id`: Referensi ke tabel `users` untuk menunjukkan pengguna yang membuat komentar.
- **Relasi**:
    - `article_id` berelasi dengan `articles(uuid)`, artinya setiap komentar terhubung ke sebuah artikel.
    - `user_id` berelasi dengan `users(uuid)`, artinya setiap komentar dibuat oleh seorang pengguna.

### 3. Tabel `users`
- **Deskripsi**: Tabel ini menyimpan data pengguna yang terdaftar di aplikasi.
- **Kolom**:
    - `uuid`: Kunci utama unik untuk setiap pengguna.
    - `created_at`: Waktu pengguna terdaftar.
    - `updated_at`: Waktu pengguna diperbarui.
    - `deleted_at`: Waktu pengguna dihapus (jika dihapus).
    - `username`: Nama pengguna yang unik.
    - `email`: Alamat email pengguna.
    - `password`: Kata sandi pengguna yang dienkripsi.
    - `role_id`: Referensi ke tabel `roles` untuk menentukan peran pengguna.
- **Relasi**:
    - `role_id` berelasi dengan `roles(id)`, artinya setiap pengguna memiliki peran tertentu (seperti admin atau user biasa).

### 4. Tabel `roles`
- **Deskripsi**: Tabel ini menyimpan daftar peran yang dapat dimiliki oleh pengguna.
- **Kolom**:
    - `id`: Kunci utama unik untuk setiap peran.
    - `name`: Nama peran, seperti "admin" atau "user".
- **Relasi**:
    - Tabel `users` memiliki kolom `role_id` yang merujuk ke tabel ini untuk menentukan hak akses pengguna.

## Ringkasan Relasi

- **Pengguna** dapat membuat **artikel**, sehingga `articles.author_id` mengacu pada `users.uuid`.
- **Artikel** dapat memiliki banyak **komentar** dari **pengguna** yang berbeda, sehingga `comments.article_id` mengacu pada `articles.uuid` dan `comments.user_id` mengacu pada `users.uuid`.
- **Pengguna** memiliki **peran** yang didefinisikan dalam tabel `roles`, sehingga `users.role_id` mengacu pada `roles.id`.
