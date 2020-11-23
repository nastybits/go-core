package engine

import (
	"encoding/json"
	"go.core/lesson6/pkg/cache/membot"
	"go.core/lesson6/pkg/crawler"
	"go.core/lesson6/pkg/index"
	"go.core/lesson6/pkg/storage"
	"go.core/lesson6/pkg/storage/bstree"
	"testing"
)

func TestEngine_Load(t *testing.T) {
	var bt bstree.Tree
	e := Engine{
		Index:   index.New(),
		Storage: storage.New(&bt),
		Cache:   &membot.Cache{},
	}

	t.Run("Load cached docs", func(t *testing.T) {
		if err := e.Load(); err != nil {
			t.Errorf("Load() error = %v, want %v", err, nil)
		}
	})
}

func TestEngine_Save(t *testing.T) {
	var bt bstree.Tree
	e := Engine{
		Index:   index.New(),
		Storage: storage.New(&bt),
		Cache:   &membot.Cache{},
	}

	data, err := e.Cache.Load()
	if err != nil {
		t.Errorf("Ошибка в тесте Load")
	}

	var docs []crawler.Document
	err = json.Unmarshal(data, &docs)
	if err != nil {
		t.Errorf("Ошибка в тесте Load")
	}

	t.Run("Save cached docs", func(t *testing.T) {
		if err := e.Save(docs); err != nil {
			t.Errorf("Save() error = %v, want %v", err, nil)
		}
	})
}
