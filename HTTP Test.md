# HTTP test
- Go-lang sudah menyediakan package khusus untuk membuat unit test terhadap fitur web yang ktia buat
- Semuannya ada didalam package net/http/httptest https://golang.org/pkg/net/http/httptest/
- Dengan menggunakan package ini, kita bisa melakukan testing handler web di Go-lang tnapa harus menjalankan aplikasi webnya.
- Kita bisa langsung fokus terhadapt function yang ingin kita test.

# http.NewRequest()
- NewRequest(Method,url,body) merupakan function yang digunakan untuk membuat http.Request
- Kita bisa menentukan method,url,body yang kita akan kirmakna sebagai simulasi unit test.
- Selai itu, kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim, seperti header, cookie dan lain-lain

# http.NewRecorder()
- httptest.NewRecorder() merupakan function yang digunakan untuk membuat ResponseRecorder
- ResponseRecorder merupakn struct bantun untuk merekam HTTP resposnse dari hasil testing yang akan kita lakukan.
