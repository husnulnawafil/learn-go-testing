package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Husnul",
			request: "Husnul",
		},
		{
			name:    "Andi",
			request: "Andi",
		},
		{
			name:    "Benzema",
			request: "Benzema",
		},
		{
			name:    "Kevin",
			request: "Kevin",
		},
		{
			name:    "Modric",
			request: "Modric",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkHelloWorldHusnul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Husnul")
	}
}

func BenchmarkHelloWorldAfif(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Afif")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Lidya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Lidya")
		}
	})
	b.Run("Sarimi Isi Dua", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sarimi Isi Dua")
		}
	})
}
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorldHusnul(t *testing.T) {
	result := HelloWorld("Husnul")
	assert.Equal(t, "Hello Husnul", result, "Result must be 'Hello Husnul'")
}

func TestHelloWorldKhannedy(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("can not run on mac os")
	}
	result := HelloWorld("Khannedy")
	require.Equal(t, "Hello Khannedy", result, "Result must be 'Hello Khannedy'")

}

func TestSubTest(t *testing.T) {
	t.Run("Abil", func(t *testing.T) {
		result := HelloWorld("Abil")
		assert.Equal(t, "Hello Abil", result, "Result must be 'Hello Abil'")
	})
	t.Run("Kamal", func(t *testing.T) {
		result := HelloWorld("Kamal")
		assert.Equal(t, "Hello Kamal", result, "Result must be 'Hello Kamal'")
	})
}

// TABLE UNIT TEST
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Husnul",
			request:  "Husnul",
			expected: "Hello Husnul",
		},
		{
			name:     "Alif",
			request:  "Alif",
			expected: "Hello Alif",
		},
		{
			name:     "Yunus",
			request:  "Yunus",
			expected: "Hello Yunus",
		},
		{
			name:     "Daffa",
			request:  "Daffa",
			expected: "Hello Daffa",
		},
		{
			name:     "Riyadh",
			request:  "Riyadh",
			expected: "Hello Riyadh",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
