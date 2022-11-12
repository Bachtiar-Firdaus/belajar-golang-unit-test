package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// example before and after testing
func TestMain(m *testing.M) {
	fmt.Println("Sebelum Dilakkukan Testing")
	// ini untuk menjalankan semua unit test
	m.Run()
	fmt.Println("Setelah Dilakkukan Testing")
}

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

// example Assertion
func TestOldV5(t *testing.T) {
	name := "bachtiar"
	old := 20
	rName, rOld := Old(name, old)
	assert.Equal(t, "bachtiar", rName, "Hasilnya harusnya bachtiar")
	assert.Equal(t, 20, rOld, "Hasilnya harusnya 20")
	fmt.Println("Dieksekusi")
}

// example Require
func TestOldV6(t *testing.T) {
	name := "bachtiar"
	old := 20
	rName, rOld := Old(name, old)
	require.Equal(t, "bachtiar", rName, "Hasilnya harusnya bachtiar")
	require.Equal(t, 20, rOld, "Hasilnya harusnya 20")
	fmt.Println("Dieksekusi")
}

// example Skip Test
func TestSkip(t *testing.T) {
	// runtime disini untuk mengetahui sistem operasi apa yang kita gunakan
	if runtime.GOOS == "windows" {
		t.Skip("Unit test tidak bisa jalan di windows")
	}
	result := HelloWord("daus")
	require.Equal(t, "hello daus", result)
}

// example Subtest
func TestSubTest(t *testing.T) {
	t.Run("sub1", func(t *testing.T) {
		result := HelloWord("bachtiar")
		require.Equal(t, "Hello bachtiar", result)
	})
	t.Run("sub2", func(t *testing.T) {
		result := HelloWord("firdaus")
		require.Equal(t, "Hello firdaus", result)
	})
}

// example Table Test
func TestHelloWorldTable(t *testing.T) {
	// example struct untuk table test
	tests := []struct {
		name    string
		request string
		expect  string
	}{
		{
			name:    "sub1",
			request: "daus",
			expect:  "Hello daus",
		}, {
			name:    "sub2",
			request: "bachtiar",
			expect:  "Hello bachtiar",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWord(test.request)
			assert.Equal(t, test.expect, result)
		})
	}
}
