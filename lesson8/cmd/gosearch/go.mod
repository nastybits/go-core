module go.core/lesson8/cmd

go 1.13

replace (
	go.core/lesson8/pkg/engine v0.0.0 => ../../pkg/engine
	go.core/lesson8/pkg/cache v0.0.0 => ../../pkg/cache
	go.core/lesson8/pkg/cache/local v0.0.0 => ../../pkg/cache/local
	go.core/lesson8/pkg/cache/membot v0.0.0 => ../../pkg/cache/membot
	go.core/lesson8/pkg/index v0.0.0 => ../../pkg/index
	go.core/lesson8/pkg/storage v0.0.0 => ../../pkg/storage
	go.core/lesson8/pkg/storage/bstree v0.0.0 => ../../pkg/storage/bstree
	go.core/lesson8/pkg/crawler v0.0.0 => ../../pkg/crawler
	go.core/lesson8/pkg/crawler/spider v0.0.0 => ../../pkg/crawler/spider
	go.core/lesson8/pkg/crawler/membot v0.0.0 => ../../pkg/crawler/membot
)

require (
	go.core/lesson8/pkg/cache v0.0.0
	go.core/lesson8/pkg/cache/local v0.0.0
	go.core/lesson8/pkg/engine v0.0.0
	go.core/lesson8/pkg/index v0.0.0
	go.core/lesson8/pkg/storage v0.0.0
	go.core/lesson8/pkg/storage/bstree v0.0.0
	go.core/lesson8/pkg/crawler v0.0.0
	go.core/lesson8/pkg/crawler/spider v0.0.0
)
