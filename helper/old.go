package helper

func Old(name string, old int) (string, int) {
	// Menggagalkan Unit Test
	/*
		Menggagalkan unit test menggunkan panic bukanlah hal yang bagus
		Golang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
		Terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test
		- Fail() akan menggagalkan unit test namun tetap melanjutkan eksekusi unit test namun di akhir ketika selesai maka unit test tersebut dianggap gagal
		- FailNow() akan menggagalkan unit test saat ini juga tanpa melanjutkan eksekusi unit test
		- Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
		- Fatal() mirip dengan Error(), hanya saja setelah melakukan log error dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti
	*/
	return name, old
}
