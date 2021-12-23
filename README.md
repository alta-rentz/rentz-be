# RENTZ.ID
<div id="top"></div>
<!-- PROJECT LOGO -->
<br/>
<div align="center">
  <a href="https://github.com/alta-be4-andri/Project-2">
    <img src="image/welcome.gif" alt="Logo" width="700" height="350">
  </a>

  <h3 align="center">Project "RENTZ.ID" Rent Product App </h3>

  <p align="center">
    Projek Capstone Pembangunan RESTful API Program Immersive Back End Batch 4
    <br />
    <a href="https://github.com/alta-rentz/rentz-be"><strong>Kunjungi kami »</strong></a>
    <br />
  </p>
</div>

### 🛠 &nbsp;Build App & Database

![JSON](https://img.shields.io/badge/-JSON-05122A?style=flat&logo=json&logoColor=000000)&nbsp;
![GitHub](https://img.shields.io/badge/-GitHub-05122A?style=flat&logo=github)&nbsp;
![Visual Studio Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=007ACC)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=4479A1)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=4479A1)&nbsp;
![AWS](https://img.shields.io/badge/-AWS-05122A?style=flat&logo=amazon)&nbsp;
![Postman](https://img.shields.io/badge/-Postman-05122A?style=flat&logo=postman)&nbsp;
![Docker](https://img.shields.io/badge/-Docker-05122A?style=flat&logo=docker)&nbsp;
![Ubuntu](https://img.shields.io/badge/-Ubuntu-05122A?style=flat&logo=ubuntu)&nbsp;
![GDC](https://img.shields.io/badge/-GoogleCloud-05122A?style=flat&logo=google)&nbsp;

<!-- ABOUT THE PROJECT -->
### 💻 &nbsp;About The Project

RENTZ.ID merupakan projek Capstone untuk membangun sebuah RESTful API Rental App dengan menggunakan bahasa Golang.    
dilengkapi dengan berbagai fitur yang memungkinkan user untuk mengakses data yang ada didalam server. mulai dari membuat akun hingga hosting produk yang ingin disewakan. Adapun fitur yang ada dalam RESTful API kami antara lain :
<div>
      <details>
<summary>🙎 Users</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
 Di User terdapat fitur untuk membuat Akun dan Login agar mendapat legalitas untuk mengakses berbagai fitur lain di aplikasi, 
 terdapat juga fitur Update untuk mengedit data yang berkaitan dengan user, serta fitur delete berfungsi jika user menginginkan hapus akun.
 
<div>
  
| Feature User | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /signup  | - | NO | Melakukan proses registrasi user |
| POST | /signin | - | NO | Melakukan proses login user |
| GET | /users | - | YES | Mendapatkan informasi user yang sedang login |
| PUT | /users | - | YES | Melakukan update informasi user yang sedang login | 
| DEL | /users | - | YES | Menghapus user yang sedang login |

</details>  

<details>
<summary>🛍 &nbsp;Product</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
User dapat mem-posting berbagai product untuk disewakan kepada user lain, terdapat beberapa fitur seperti melihat seluruh product, mencari product sesuai dengan id product, melihat product yang dipost user, menambahkan dan meng-update product dengan detail harga, stok untuk memudahkan user lain yang akan membeli productnya, serta fitur delete yang memungkinkan user menghapus product yang sudah tidak dijual.
  
| Feature Products | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /products  | - | YES | Membuat Product sewa baru |
| GET | /products | - | NO | Mendapatkan informasi seluruh product |
| GET | /products | - | YES | Mendapatkan informasi seluruh product user yang sedang login |
| GET | /products/:id | id | NO | Mendapatkan informasi product berdasarkan product id |
| GET | /products/subcategory/:id | id | NO | Mendapatkan informasi product berdasarkan subcategories |
| DEL | /products/:id | id | YES | Melakukan delete product tertentu berdasarkan id product |

</details>


<details>
<summary>🛒 &nbsp;Cart</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Cart merupakan fitur untuk menampung berbagai product yang akan dibeli oleh user, adapun fiturnya ada GET dimana user bisa melihat barang apa aja yang ada di dalam keranjang, ada fitur history dimana user bisa melihat jumlah product yang sudah dibayar.
  
| Feature cart | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| GET | /cart | - | YES | Mendapatkan informasi booking yang ada didalam cart |
| GET | /history | - | YES | Mendapatkan informasi booking yang telah dibayar |

</details>

<details>
<summary>🗓&nbsp;Booking</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Setelah user melakukan pencarian product yang dibutuhkan dengan berbagai jaminan yang dibutuhkan, user melakukan booking dengan melakukan pengecekan tanggal diawal, jika sistem merespon product yang dimaksud "avalaible", user baru dapat melakukan booking.  
  
| Feature booking | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /booking/check/:id | id | YES | Melakukan cek ketersediaan product tertentu berdasarkan tanggal time-in dan time-out |
| POST | /booking | - | YES | Membuat booking product |
| GET | /booking/:id | id | YES | Mendapatkan informasi booking berdasarkan booking id |
| DEL | /booking/:id | id | YES | Melakukan cancel booking berdasarkan booking id |

</details>

<details>
<summary>💳&nbsp;CheckOut</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
 Merupakan fitur untuk dimana user melakukan pembayaran sesuai product yang dipilih dari cart, adapun payment gateway yang digunakan adalah xendit, payment_method yang digunakan ewallet dengan 4 channel, DANA, OVO, LINKAJA, dan SHOPEEPAY
  
| Feature Reservaton | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /checkout | - | YES | Melakukan Checkout |
| POST | /checkout/ovo | - | YES | Melakukan Checkout melalui channel OVO |

</details>
      

<!-- IMAGES -->
### 🖼&nbsp;Images

<details>
<summary>📈&nbsp;ERD</summary>
<img src="images/Project2 (3).jpg">
</details>

<details>
<summary>📖&nbsp;User Stories</summary>
<img src="images/Project2 (3).jpg">
</details>

<details>
<summary>📨&nbsp;Workflow User</summary>
<img src="images/Project2 (3).jpg">
</details>

<!-- CONTACT -->
### Contact

[![Gmail: Alfy](https://img.shields.io/badge/-Alfy-maroon?style=flat&logo=gmail)](https://mail.google.com/mail/u/0/#inbox?compose=CllgCHrjmjRlSpLttDDmhqnRQTQVTSQCjFvQxCSSqGDHvQjrjJvvzKMvnlWTrWwkcGdSzfJPXnV)
[![GitHub Alfy](https://img.shields.io/badge/-Alfy-white?style=flat&logo=github&logoColor=black)](https://github.com/alfiancikoa)

[![Gmail: Andri](https://img.shields.io/badge/-Andri-maroon?style=flat&logo=gmail)](https://mail.google.com/mail/u/0/#inbox?compose=DmwnWslzCnrLrhrlnrRWdpHqsBmRtbbtZSKxXFrdGHmhLVLjLDmVfNRxdBShrxQNTBBHFgDdLfKQ)
[![GitHub Andri](https://img.shields.io/badge/-Andri-white?style=flat&logo=github&logoColor=black)](https://github.com/DylanRipper)

[![Gmail: Fafa](https://img.shields.io/badge/-Fafa-maroon?style=flat&logo=gmail)](https://mail.google.com/mail/u/0/#inbox?compose=DmwnWslzCnrLrhrlnrRWdpHqsBmRtbbtZSKxXFrdGHmhLVLjLDmVfNRxdBShrxQNTBBHFgDdLfKQ)
[![GitHub FAfa](https://img.shields.io/badge/-Fafa-white?style=flat&logo=github&logoColor=black)](https://github.com/DylanRipper)


<p align="center">:copyright: 2021 | AAF</p>
</h3>
