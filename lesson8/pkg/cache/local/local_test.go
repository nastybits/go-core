package local

import (
	"log"
	"os"
	"reflect"
	"testing"
)

var cache *Cache

func TestMain(m *testing.M) {
	path := "./test.txt"
	cache = New(path)
	exitCode := m.Run()
	err := os.Remove(path)
	if err != nil {
		log.Fatalf("Временный тестовый файл %s не был удален", path)
	}
	os.Exit(exitCode)
}

func TestCache_SaveLoad(t *testing.T) {
	s := []byte("Test string for saving")
	test := struct {
		name string
		str  []byte
		want []byte
	}{
		name: "Save and Load string",
		str:  s,
		want: s,
	}

	err := cache.Save(test.str)
	if err != nil {
		t.Errorf("Save() error = %v", err)
		return
	}

	got, err := cache.Load()
	if !reflect.DeepEqual(got, test.want) {
		t.Errorf("Load() got = %v, want %v", got, test.want)
	}
}
