package bloomgo

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	bloomFilt := New(10000, 0.01)
	bloomFilt.Add([]byte("test-1"))
	exists := bloomFilt.Exists([]byte("test-1"))
	if !exists {
		t.Fatalf("test-1 not present in the bloom filter")
	}

	notexists := bloomFilt.Exists([]byte("test-2"))
	if notexists {
		t.Fatalf("test-2 present in the bloom filter")
	}
}

func TestMultiple(t *testing.T) {
	bloomFilt := New(10000, 0.001)
	for i := 0; i < 500; i++ {
		data := fmt.Sprintf("test-%d", i)
		bloomFilt.Add([]byte(data))
	}

	for i := 0; i < 500; i++ {
		data := fmt.Sprintf("test-%d", i)
		exists := bloomFilt.Exists([]byte(data))
		if !exists {
			t.Fatalf("test-1 not present in the bloom filter")
		}
	}

	for i := 500; i < 1000; i++ {
		data := fmt.Sprintf("test-%d", i)
		exists := bloomFilt.Exists([]byte(data))
		if exists {
			t.Fatalf("%s not present in the bloom filter", data)
		}
	}
}
