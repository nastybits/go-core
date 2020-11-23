module go.core/lesson6/pkg/engine

go 1.13

replace (
	go.core/lesson6/pkg/cache v0.0.0 => ../cache
	go.core/lesson6/pkg/cache/local v0.0.0 => ../cache/local
	go.core/lesson6/pkg/cache/membot v0.0.0 => ../cache/membot
	go.core/lesson6/pkg/crawler v0.0.0 => ../crawler
	go.core/lesson6/pkg/crawler/membot v0.0.0 => ../crawler/membot
	go.core/lesson6/pkg/crawler/spider v0.0.0 => ../crawler/spider
	go.core/lesson6/pkg/index v0.0.0 => ../index
	go.core/lesson6/pkg/storage v0.0.0 => ../storage
	go.core/lesson6/pkg/storage/bstree v0.0.0 => ../storage/bstree
)

require (
	go.core/lesson6/pkg/cache v0.0.0
	go.core/lesson6/pkg/cache/membot v0.0.0
	go.core/lesson6/pkg/crawler v0.0.0
	go.core/lesson6/pkg/index v0.0.0
	go.core/lesson6/pkg/storage v0.0.0
	go.core/lesson6/pkg/storage/bstree v0.0.0
)
