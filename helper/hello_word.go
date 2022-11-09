package helper

func HelloWord(name string) string {
	// MEMBUAT UNIT TEST

	/*
		step example mebuat unit test
		jalan kan perintah ini di terminal : go mod init belajar-golang-unit-test
		kemudian buat sebuah package helper
		buat sebuah file hello_word.go
	*/

	// Aturan File Test
	/*
		Golang memiliki aturan cara membuat file khusus untuk unit test
		Untuk membuat file unit test, kita harus menggunakan akhiran _test
		Jadi kita misal kita membaut file hello_world.go artinya untuk membuat unit testnya kita harus membuat file hello_eorld_test.go
	*/

	// Aturan Function Unit Test
	/*
		Selain aturan nama file, di golang juga sudah diatur untuk nama function unit test
		Nama Function untuk unit test harus diawali dengan nama Test
		Misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test dengan nama TestHelloWorld
		Tak ada aturan untuk nama belakang funtion unit test harus sama dengan function yang akan di test, yang penting adalah diawali dengan kata Test
		selanjutnya harus memiliki parameters (t*testing.T) dan tidak memiliki return value
	*/

	// Menjalankan Unit Test
	/*
		Untuk menjalankan unit test kita bisa menggunakan perintah :
		- go test
		jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa gunakan perintah :
		- go test -v
		dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah :
		go test -v -run TestNamaFunction
	*/

	// Menjalankan Semua unit test
	/*
		jika kita ingin menjalankan semua unit test dari top folder modulenya kita bisa gunakan perintah :
		go test ./...
		go test -v ./...
	*/
	return "Hello " + name
}
