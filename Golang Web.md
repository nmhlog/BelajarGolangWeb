Handler
- Server hanya bertugas sebagai web serber, sedangkan untuk menerima HTTP request yang masuk ke server, kita butuh yang namanya handler.
- Handler diGo-Lang direpresentasian dalam interface, dimana dalam kontraknya terdapat sebuh function bernama serveHTTP() yang digunakan sebagai funciton yang akan dieksekusi ketika menerima HTTP request.

HandlerFunc
- Salah satu implementasi dari interface Handler adalah HandlerFunc
- Kita bisa menggunakan HandlerFunc untuk membuat function handler HTTP

URL Pattern
- URL Pattern dalam serveMux sederhana, kita tinggal menambhakan string yang ingin kita gunakan sebagai endpoint, tanpa perlu mamasukkan domain web kita.
- Jika URL pattern dalam serveMux kita tambahakan diakhirnya dengan garis miring, artinya semua url tersebut akan menerima path dengan awalan tersebut, misalnya /images/ artinya akan dieksekusi jika endpointnya /images/contoh, /images/contoh/lagi
- Namun jika terdapat URL pattern yang lebih panjang, maka akan diprioritaskan yang lebih panjang, misalakan terdapat URL /images/ dan images/thumbsnails/, maka jika mengakses /images/thumbnails/ akan mengakses /images/thumsnails/ bukan images/