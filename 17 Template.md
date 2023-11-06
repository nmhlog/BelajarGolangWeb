# Template

## HTML Template

- Fiture HTML template terdapat dipackage html/template
- Sebelum menggunakan HTML template, kita perlu terlebih dahulu mebuat templatenya
- Tmepaltenya berupa file atau string
- Bagian dinamin dari html template, adalah bagian yang menggunakna tanda "{{}}"

## Membuat Template

- Saat membuat template dengan string, Kita perlu memberi tahu nama templatenya
- Dan untuk mebuat text template, cukup buat text html, dan untuk konten yang dinamis kita bisa gunakan tandat {{.}}, contoh :

```
<html><body>{{}}</body></html>

```

## Template Directory

- Kadang bisanya kita jarang sekali menyebutkan file template satu persatu
- Alangkah baiknya untuk template kita simpan di satu directory
- Go-lang template mendukung process load tamplate dari directory
- Hal ini memudahkan kita, sehingga tidak perlu menyebutkan nama filenya satu per satu

## Template Dari Go-lang Embed

- Sejak Go-Lang 1.16 karena sudah ada Go-lang embed, jadi direkomendasikan menggunakan Go-lang embed untuk menyimpan data template.
- Menggunakan go-lang embed menjadi kita tidak perlu ikut meng-copy tempalte file lagi, karena sudah otomatis di embed didalam distribusi file