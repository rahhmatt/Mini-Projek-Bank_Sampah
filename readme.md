Mini proyek ini adalah aplikasi REST API yang dibangun menggunakan Golang, Echo Framework, dan MySQL, dengan dukungan Docker untuk containerisasi dan deployment di AWS.

Fitur
CRUD Operasi: Mendukung operasi Create, Read, Update, Delete untuk data tertentu.
Autentikasi JWT: Login dan registrasi dengan token JWT.
Deployment dengan Docker: Memudahkan pengelolaan aplikasi di lingkungan produksi.
Database Terintegrasi: Menggunakan AWS RDS sebagai database.
Prasyarat
Sebelum menjalankan aplikasi, pastikan Anda telah menginstal:

Docker
Docker Compose
MySQL Client (jika ingin mengakses database secara manual)
Instalasi
Ikuti langkah-langkah berikut untuk menjalankan proyek ini:

Clone Repository

bash
Copy code
git clone https://github.com/rahhmatt/mini-proyek.git
cd mini-proyek

Dokumentasi API
Berikut adalah daftar endpoint yang tersedia:

HTTP Method	Endpoint	Deskripsi
POST	/login	Login pengguna
POST	/register	Registrasi pengguna baru
GET	/data	Mendapatkan data tertentu
POST	/data	Menambahkan data baru
PUT	/data/:id	Mengupdate data berdasarkan ID
DELETE	/data/:id	Menghapus data berdasarkan ID


Teknologi yang Digunakan
Golang: Bahasa pemrograman utama.
Echo Framework: Framework untuk REST API.
MySQL: Database management system.
Docker: Containerization.
AWS: Platform hosting (EC2 dan RDS).
