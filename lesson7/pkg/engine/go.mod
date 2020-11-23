module go.core/lesson7/pkg/engine

go 1.13

replace (
	go.core/lesson7/pkg/cache v0.0.0 => ../cache
	go.core/lesson7/pkg/cache/local v0.0.0 => ../cache/local
	go.core/lesson7/pkg/cache/membot v0.0.0 => ../cache/membot
	go.core/lesson7/pkg/crawler v0.0.0 => ../crawler
	go.core/lesson7/pkg/crawler/membot v0.0.0 => ../crawler/membot
	go.core/lesson7/pkg/crawler/spider v0.0.0 => ../crawler/spider
	go.core/lesson7/pkg/index v0.0.0 => ../index
	go.core/lesson7/pkg/storage v0.0.0 => ../storage
	go.core/lesson7/pkg/storage/bstree v0.0.0 => ../storage/bstree
)

require (
	go.core/lesson7/pkg/cache v0.0.0
	go.core/lesson7/pkg/cache/membot v0.0.0
	go.core/lesson7/pkg/crawler v0.0.0
	go.core/lesson7/pkg/index v0.0.0
	go.core/lesson7/pkg/storage v0.0.0
	go.core/lesson7/pkg/storage/bstree v0.0.0
)
