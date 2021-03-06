module go.core/lesson4/cmd/engine

go 1.13

replace (
	go.core/lesson4/pkg/index v0.0.0 => ./../../pkg/index
	go.core/lesson4/pkg/document v0.0.0 => ./../../pkg/document
	go.core/lesson4/pkg/crawler v0.0.0 => ./../../pkg/crawler
	go.core/lesson4/pkg/crawler/spider v0.0.0 => ./../../pkg/crawler/spider
	go.core/lesson4/pkg/crawler/stub v0.0.0 => ./../../pkg/crawler/stub
)

require (
	go.core/lesson4/pkg/document v0.0.0
	go.core/lesson4/pkg/crawler v0.0.0
	go.core/lesson4/pkg/index v0.0.0
	go.core/lesson4/pkg/crawler/spider v0.0.0
	go.core/lesson4/pkg/crawler/stub v0.0.0
)
