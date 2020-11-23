module go.core/lesson7/cmd

go 1.13

replace (
	go.core/lesson7/pkg/engine v0.0.0 => ../pkg/engine
	go.core/lesson7/pkg/cache v0.0.0 => ../pkg/cache
	go.core/lesson7/pkg/cache/local v0.0.0 => ../pkg/cache/local
	go.core/lesson7/pkg/cache/membot v0.0.0 => ../pkg/cache/membot
	go.core/lesson7/pkg/index v0.0.0 => ../pkg/index
	go.core/lesson7/pkg/storage v0.0.0 => ../pkg/storage
	go.core/lesson7/pkg/storage/bstree v0.0.0 => ../pkg/storage/bstree
	go.core/lesson7/pkg/crawler v0.0.0 => ../pkg/crawler
	go.core/lesson7/pkg/crawler/spider v0.0.0 => ../pkg/crawler/spider
	go.core/lesson7/pkg/crawler/membot v0.0.0 => ../pkg/crawler/membot
)

require (
	go.core/lesson7/pkg/cache v0.0.0
	go.core/lesson7/pkg/cache/local v0.0.0
	go.core/lesson7/pkg/engine v0.0.0
	go.core/lesson7/pkg/index v0.0.0
	go.core/lesson7/pkg/storage v0.0.0
	go.core/lesson7/pkg/storage/bstree v0.0.0
	go.core/lesson7/pkg/crawler v0.0.0
	go.core/lesson7/pkg/crawler/spider v0.0.0
)
