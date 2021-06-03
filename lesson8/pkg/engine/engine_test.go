package engine

import (
	"encoding/json"
	"go.core/lesson8/pkg/cache/membot"
	"go.core/lesson8/pkg/crawler"
	"go.core/lesson8/pkg/fixtures"
	"go.core/lesson8/pkg/index"
	"go.core/lesson8/pkg/storage"
	"os"
	"reflect"
	"testing"
)

var engine *Service

func TestMain(m *testing.M) {
	engine = New(index.New(), storage.New(), &membot.Cache{})
	os.Exit(m.Run())
}

func TestEngine_LoadSave(t *testing.T) {
	err := engine.Load()

	if err != nil {
		t.Errorf("Load() error = %v, want %v", err, nil)
	}

	var docs []crawler.Document

	data, _ := engine.cache.Load()
	err = json.Unmarshal(data, &docs)

	if err != nil {
		t.Logf("Ошибка при раскодировании данных")
	}

	if err := engine.Save(docs); err != nil {
		t.Errorf("Save() error = %v, want %v", err, nil)
	}
}

func TestService_Search(t *testing.T) {
	docs := fixtures.Documents()
	err := engine.Load()
	if err != nil {
		t.Errorf("Не удалось загрузить тестовые данные")
	}

	tests := []struct {
		name     string
		arg      string
		wantDocs []crawler.Document
	}{
		{name: "Find one", arg: "one", wantDocs: []crawler.Document{docs[0]}},
		{name: "Find in8and2", arg: "in8and2", wantDocs: []crawler.Document{docs[1], docs[7]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDocs := engine.Search(tt.arg); !reflect.DeepEqual(gotDocs, tt.wantDocs) {
				t.Errorf("Search() = %v, want %v", gotDocs, tt.wantDocs)
			}
		})
	}
}
