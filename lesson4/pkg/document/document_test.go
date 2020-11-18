package document

import (
	"reflect"
	"testing"
)

func TestWebPage_Words(t *testing.T) {
	type fields struct {
		id    int
		url   string
		title string
	}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"1", fields{1, "url", "title"}, []string{"url", "title"}},
		{"2", fields{1, "url2", "title2"}, []string{"url2", "title2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := WebPage{
				id:    tt.fields.id,
				url:   tt.fields.url,
				title: tt.fields.title,
			}
			if got := d.Words(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Words() = %v, want %v", got, tt.want)
			}
		})
	}
}
