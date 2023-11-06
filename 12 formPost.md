# Note

## Form Post

- Saat kita belajar HTML, kita tahu bahwa saat kita membuat form, kita bisa subit datanya dengan method GET atau POST.
- Jika menggukana method Get, maka hasil data diForm akan menjadi query parameter.
- Sedangkan jika mneggunakan POST, maka semua data diForm akan dikirim via body HTTP request
- Di Go-lang untuk mengambil data Form-post sangatlah mudah

## Request.PostForm

- Semau data form post yang dikirm dari client, secara otomatis akan disimpan dalam attribut Request.PostForm
- Namun sebelum kita bisa mengambil data di attribute postform, kita wajib memanggil method Requset.PaseForm() terlebih dahulu, method ini digunakan untuk malakukan parsing data body apakah bisa di parsing menjadi form data atau tidak, jika kita bisa diparsing, maka akan menyababkan error.
