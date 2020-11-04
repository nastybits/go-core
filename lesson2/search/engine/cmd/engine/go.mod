module go.core/lesson2/search/engine/cmd/engine

go 1.13

replace go.core/lesson2/search/crawler/pkg/spider v0.0.0 => ../../../crawler/pkg/spider

replace go.core/lesson2/search/engine/pkg/search v0.0.0 => ../../../engine/pkg/search

require (
	go.core/lesson2/search/crawler/pkg/spider v0.0.0
	go.core/lesson2/search/engine/pkg/search v0.0.0
)
