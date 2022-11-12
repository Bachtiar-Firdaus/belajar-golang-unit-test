# **BELAJAR GO-LANG UNIT TEST**

## **PERSYARATAN YANG WAJIB**

- Kalian telah mempelajari Materi :
  - CAPTER 1 Golang Dasar [LINK CATPER 1](https://github.com/Bachtiar-Firdaus/belajar-golang-dasar)
  - CAPTER 2 Golang Modules [LINK CATPER 2](https://github.com/Bachtiar-Firdaus/belajar-golang-modules)

## **DESKRIPSI SINGKAT MATERI CAPTER 3**

> Ini adalah repositori yang berisi materi materi Go-Lang Unit Test

- Agenda
  - Pengenalan Software Testing
  - Pengenalan Unit Test
  - Testing Package
  - Unit Test
  - Assertion
  - Mock
  - Benchmark

## **PENGENALAN SOFTWARE TESTING**

> Software Testing adalah salah satu disiplin ilmu dalam software engineering, tujuan utama dari testing adalah memastikan kualitas kode dan aplikasi kita baik, ilmu untuk software testing sendiri sangatlah luas, pada materi ini kita hanya fokus ke unit testing.

> Secara garis besar ada namanya test pyramid yang terdiri atas UI (end to end test), Service (integration testing), dan Unit (unit testing).

## **PENGENALAN UNIT TEST**

![image](https://user-images.githubusercontent.com/33290851/200884253-295adf85-c989-43cf-b95c-ab4eb9204ae7.png)

> Secara singkat unit test bisa di gambarkan seperti gambar berikut, unit test dilakukan pada unit terkecil dari kode program.

> Unit test akan fokus menguji bagian kode program terkecil, biasanya menguji sebuah method, unit test biasanya di buat kecil dan cepat, oleh karna itu biasanya kadang kode unit test lebih banyak dari kode program aslinya, karna semua skenario pengujian akan dicoba di unit test, unit test bisa digunakan untuk meningkatkan kualitas kode program kita

## **PENGENALAN TESTING PACKAGE**

> Dibahasa pemrograman lain, biasanya untuk implementasi unit test, kita butuh library atau framework khusus untuk unit test, berbeda dengan GO-Lang, di GO-Lang untuk unit test sudah disediakan sebuah package khusus bernama testing, selain itu untuk menjalankan unit test, di GO-Lang juga sudah disediakan perintah nya, hal ini membuat implementasi unit testing di golang sangat mudah dibandingkan dengan bahasa pemrograman yang lain

- https://golang/pkg/testing

- testing.T

  > Golang menyediakan sebuah struct yang bernama testing.T, struct ini digunakan untuk unit test di golang

- testing.M

  > testing.M adalah stuct yang disediakan Golang untuk mengatur lifecycle testing Materi ini nanti akan kita bahas di capter Main

- testing.B
  > testing.B adalah struct yang disediakan Golang untuk melakukan benchmarking, Di Golang untuk melakukan benchmark (mengukur kecepatan kode program) pun sudah disediakan sehingga kita tidak perlu menggunakan library atau framework tambahan

## **ASSERTION**

- Assertion

  > Melakukan pengecekan di unit test secara manual menggunakan if else sangatlah menyebalkan. apalagi jika result data yang harus di cek banyak. oleh karna itu sangat tidak disarankan menggunakan assertion untuk melakukan pengecekan. sayangnya go-lang tidak menyediakan package assertion sehingga kita butuh menggunakan library untuk melakukan assertion

- Testify

  > Salah satu library yang paling populer di golang adalh Testify. kita bisa menggunakan library ini untuk melakukan assertion terhadap result data unit test.

  - https://github.com/stretchr/testify

  > kita bisa menggunakan di GO module kita dengan mengetikan comand berikut :

  - go get github.com/stretchr/testify

- Assert vs Require
  > Testify menyediakan dua package untuk assertion yaitu assertion dan require, saat kita menggunakan assert, jika pengecekan gagal maka akan memanggil Fail(), artinya eksekusi unit test akan tetap dilanjutkan, seangkan jika kita menggunakan require, jika pengecekan gagal maka require akan memanggil FailNow(), artinya unit test tidak akan dilanjutkan

## **SKIP TEST**

- Skip Test
  > Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test, di golang juga kita bisa membatalkan eksekusi unit test jika kita mau, untuk membatalkan unit test kita bisa menggunakan function Skip()

## **BEFORE AND AFTER TEST**

- Before and After Test
  > Biasanya dalam unit test, kita kadang ingin melakukan sesuatu sebelum dan setelah sebuah unit test di eksekusi, jika kode yang kita lakukan sebelumnya dan setelahnya selalu sama antar unit test function maka membuat manual di unit test functionnya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya, untungnya di golangterdapat fitur yang bernama testing.M, fitur ini bernama main dimana digunakan untuk mengatur eksekusi unit test, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di unit test

## **SUB TEST**

- Sub Test

  > Go-Lang mendukung fitur pembuatan function unit tet didalam function unit test, fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman lainnya, Untuk membuat sub test, kita bisa menggunakan function Run()

- Menjalankan Hanya Sub Test

  > Kita sudah tahu jika ingin menjalankan sebuah untit test function, kita bisa gunakan perintah :

  - go test -run TestNamaFunction

  > Jika kita ingin menjalankan hanya salah satu sub test, kita bisa menggunakan perintah :

  - go test -run TestNamaFunction/NamaSubTest

  > Atau untuk semua test sub test di semua function, kita bisa gunakan perintah :

  - go test -run /NamaSubTest

## **TABLE TEST**

- Table Test
  > Sebelumnya kita sudah belajar tentang sub test, jika diperhatikan sebenarnya dengan sub test, kita bisa membuat test secara dinamis, dan fitur sub test ini, bisa digunakan oleh programmer golang untuk membuat test dengan konsep table test, table test yaitu dimana kita menyediakan data berupa slice yang berisi parameter dan ekspetasi hasil dari unit test, lalu slice tersebut kita iterasi menggunkan sub test

## **MOCK**

- Mock

  > Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yang sudah kita program di awal, Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing, misal kita ingin membuat unit testing, misal kita ingin membuat unit test, namunternyata ada kode program kita yang harus memanggil API Call ke Third party service, dan belum tentu response nya sesuai dengan apa yang kita mau, pada kasus seperti ini, cocok sekali untuk menggunakan mock object

- Testify Mock

  > Untuk membuat mock object, tidak ada fitur bawaan Go-Lang namun kita bisa menggunakan library testify yang sebelumnya kita gunakan untuk assertion, Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object, namun perlu di perhatikan, jika desain kode program kita kurang baik, akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik, mari kita buat contoh kasus :

  - Aplikasi Query ke database
    > Dimana kita akan buat layer service sebagai business logic, dan layer Repository sebagai jembatan ke database, agar kode mudah untuk di test disarankan agar membuat kontrak berupa interface

## **BENCHMARK**

- Benchmark

  > Selain unit test, Go-Lang testing package juga mendukung melakukan benchmark, benchmark adalah mekanisme menghitungkecepatan performa kode aplikasi kita, benchmark di golang dilakukan dengan cara otomatis melakukan iterasi kode yang kita panggil berkali kali sampai waktu tertentu, kita tidak perlu menentukan jumlah iterasi dan lamanya, karna itu sudah diatur oleh testing.B bawaan dari testing package

- testing.B

  > testing.B adalah struct yang di gunakan untuk melakukan benchmark, testing.B mirip dengan testing.T terdapat dunction Fail(), FailNow(), Error(), Fatal() dan lain lain, yang membedakan ada beberapa attribut dan function tambahan yang di gunakan untuk melakukan benchmark, salah satunya adalah attribut N, ini digunakan untuk melakukan total iterasi sebuah benchmark

- Cara Kerja Benchmark

  > Cara kerja benchmark di Go-Lang sangat sederhana, dimana kita hanya perlu membuat perulangan sejumlah N attribut, Nanti secara otomatis Go-Lang akan melakukan eksekusi sejumlah perulangan yang ditentukan secara otomatis, lalu mendeteksi beberapa lama proses tersebut berjalan, dan disimpulkan performa benchmarknya dalam waktu

- Benchmark Function

  > Mirip seperti unit test, untuk benchmark pun, di golang sudah ditentukan nama function nya, harus diawali dengan kata Benchmark, misal BenchmarkHelloWorl, BenchmarkXxx, selain itu harus memiliki parameter (b \*testing.B), dan tidak boleh mengembalikan Value, untuk nama file benchmark, sama seperti unit test, diakhiri dengan \_test, misal hello_world_test.go

- Menjalankan Benchmark

  > Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namun ditambahkan parameter bench :

      - go test -v-bench=.

  > Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah :

      - go test -v-run=NotMathUnitTest -bench=.

  > Kode diatas selain menjalankan benchmark akan menjalankan unit test juga jika kita hanya ingin menjalankan benchmark tertentu, kita bisa menggunakan perintah berikut

      - go test -v-run=NotMathUnitTest-bench=BenchmarkTest

  > Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah :

      - go test -v bench=./...

## **SUB BENCHMARK**

- Sub Benchmark

  > Sama seperti testing.T di testing.B juga kita bisa membuat sub benchmark menggunakan function Run()

- Menjalankan Hanya Sub Benchmark

  > Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan, namun jika kita ingin menjalankan salah satu sub benchmark saja kita bisa gunakan perintah :

      - go test -v -bench=BenchmakNama/NamaSub

## **TABLE BENCHMARK**

- Table Benchmark
  > Sama seperti di unit test, Programmer Go-Lang terbiasa membuat table benchmark juga, ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda beda tanpa harus membuat banyak benchmark function
