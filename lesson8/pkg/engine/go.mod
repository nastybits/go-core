module go.core/lesson8/pkg/engine

go 1.13

replace (
	go.core/lesson8/pkg/cache v0.0.0 => ../cache
	go.core/lesson8/pkg/cache/local v0.0.0 => ../cache/local
	go.core/lesson8/pkg/cache/membot v0.0.0 => ../cache/membot
	go.core/lesson8/pkg/crawler v0.0.0 => ../crawler
	go.core/lesson8/pkg/crawler/membot v0.0.0 => ../crawler/membot
	go.core/lesson8/pkg/crawler/spider v0.0.0 => ../crawler/spider
	go.core/lesson8/pkg/fixtures v0.0.0 => ../fixtures
	go.core/lesson8/pkg/index v0.0.0 => ../index
	go.core/lesson8/pkg/storage v0.0.0 => ../storage
	go.core/lesson8/pkg/storage/bstree v0.0.0 => ../storage/bstree
)

require (
	go.core/lesson8/pkg/cache v0.0.0
	go.core/lesson8/pkg/cache/membot v0.0.0
	go.core/lesson8/pkg/crawler v0.0.0
	go.core/lesson8/pkg/fixtures v0.0.0
	go.core/lesson8/pkg/index v0.0.0
	go.core/lesson8/pkg/storage v0.0.0
	go.core/lesson8/pkg/storage/bstree v0.0.0
)
