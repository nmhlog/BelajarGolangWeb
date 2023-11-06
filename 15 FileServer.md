# Notes

## File Server

- Golang Meliki sebuah fiture yang bernama fileserver
- Dnengan ini, kita bisa membuat handler di Golang web yang digunakan sebagai static file server
- Dengan menggunakan filte Server kita tidak perlu manual me-load file lagi
- File Server adalah handler, jadi bisa kitambahakan kedalam http.server atau server.mux

Kenapa keluaranya 404 Not Found?

- Kenapa ini terjadi
- Hal ini dikarenakan FileServer akan membaca url, lalu mencari file berdasarkan urlnya jadi ketika kita membuat /static/index.js, maka fileserver akan mencari ke file /resources/static/index.js
- Hal ini menyebabakan 404 not found karena memang filenya tidak bisa ditemukan
- oleh karena itu, kita bisa menggunakna function http.StripPrefix

Go-lang Embed

- Digolang 1.16 terdapat fiture baru yang bernama golang embed
- Dalam golang embed kita bisa embed file kedalam binar distribution file, hal ini memudahkan kita tidak perlu meng-copy static file lagi
- Golang embed juga memiliki fiture yang bernama embed.Fs, fiture ini bisa diintegrasikan dengan file server.