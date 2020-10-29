module education/lesson2/search/engine/cmd/engine

go 1.13

replace education/lesson2/search/crawler/pkg/spider v0.0.0 => ../../../crawler/pkg/spider

replace education/lesson2/search/engine/pkg/search v0.0.0 => ../../../engine/pkg/search

require (
	education/lesson2/search/crawler/pkg/spider v0.0.0
	education/lesson2/search/engine/pkg/search v0.0.0
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
)
