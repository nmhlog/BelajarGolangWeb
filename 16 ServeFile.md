# Notes

## Serve File

- Kadang ada kasus misal kita hanya ingin menggunakan static file sesui dengan yang kita inginkan 
- Hal ini bisa dilakukan menggunakna function http.ServerFile()
- Dengan menggunakan function ini, kita bisa menentukan file mana yang ingin kita tulis ke HTTP response

## Integrasi Dengan Go-lang Embed

- Parameter function http.ServeFile hanya berisi string file name sehingga tidak bisa menggunakna go-lang embed.
- Namun bukan berarti kita tidak bisa menggunakan Go-Lang embed, karena jika untuk melakukan laod file, kita hanya butuh menggunakna package fmt dan responseWriter Saja