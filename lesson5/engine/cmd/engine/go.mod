module go.core/lesson5/engine/cmd/engine

go 1.13

replace (
	go.core/lesson5/engine/pkg/index v0.0.0 => ../../pkg/index
	go.core/lesson5/engine/pkg/storage v0.0.0 => ../../pkg/storage
	go.core/lesson5/engine/pkg/storage/bstree v0.0.0 => ../../pkg/storage/bstree
	go.core/lesson5/engine/pkg/crawler v0.0.0 => ../../pkg/crawler
	go.core/lesson5/engine/pkg/crawler/spider v0.0.0 => ../../pkg/crawler/spider
	go.core/lesson5/engine/pkg/crawler/membot v0.0.0 => ./../../pkg/crawler/membot
)

require (
	go.core/lesson5/engine/pkg/index v0.0.0
	go.core/lesson5/engine/pkg/storage v0.0.0
	go.core/lesson5/engine/pkg/storage/bstree v0.0.0
	go.core/lesson5/engine/pkg/crawler v0.0.0
	go.core/lesson5/engine/pkg/crawler/spider v0.0.0
)
