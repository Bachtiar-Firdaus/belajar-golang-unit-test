package helper

import (
	"fmt"
	"testing"
)

func TestOld(t *testing.T) {
	name := "bachtiar"
	old := 20
	rName, rOld := Old(name+"fail", old) // kita coba gagalkan test
	if name != rName || old != rOld {
		// panic("Result is not Bachtiar 20")
		/*
			Sangat tidak disarankan membuat panic di unit test karna akan menghentikan program ketika program fail
		*/
		t.Fail()
	}
	// jika menggunakan fail masih akan melanjutkan ke proses selanjutnya
	fmt.Println(rName, rOld)

}
func TestOldV2(t *testing.T) {
	name := "bachtiar"
	old := 20
	rName, rOld := Old(name+"fail", old) // kita coba gagalkan test
	if name != rName || old != rOld {
		t.FailNow()
	}
	// jika menggunakan failnow tidak melanjutkan ke proses selanjutnya
	fmt.Println(rName, rOld)
}
func TestOldV3(t *testing.T) {
	name := "bachtiar"
	old := 20
	rName, rOld := Old(name+"fail", old) // kita coba gagalkan test
	if name != rName || old != rOld {
		t.Fatal("Harusnya bachtiar 20")
	}
	// jika menggunakan Fatal tidak melanjutkan ke proses selanjutnya
	fmt.Println(rName, rOld)
}
func TestOldV4(t *testing.T) {
	name := "bachtiar"
	old := 20
	rName, rOld := Old(name+"fail", old) // kita coba gagalkan test
	if name != rName || old != rOld {
		t.Error("Harusnya bachtiar 20")
	}
	// jika menggunakan Error masih melanjutkan ke proses selanjutnya
	fmt.Println(rName, rOld)
}
