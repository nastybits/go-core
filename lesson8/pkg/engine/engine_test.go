package engine

import (
	"encoding/json"
	"go.core/lesson8/pkg/cache/membot"
	"go.core/lesson8/pkg/crawler"
	"go.core/lesson8/pkg/index"
	"go.core/lesson8/pkg/storage"
	"testing"
)

func TestEngine_Load(t *testing.T) {
	e := New(index.New(), storage.New(), &membot.Cache{})
	t.Run("Load cached docs", func(t *testing.T) {
		if err := e.Load(); err != nil {
			t.Errorf("Load() error = %v, want %v", err, nil)
		}
	})
}

func TestEngine_Save(t *testing.T) {
	c := membot.Cache{}
	e := New(index.New(), storage.New(), &c)

	data, err := c.Load()
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
