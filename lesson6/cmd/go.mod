module go.core/lesson6/cmd

go 1.13

replace (
	go.core/lesson6/pkg/engine v0.0.0 => ../pkg/engine
	go.core/lesson6/pkg/cache v0.0.0 => ../pkg/cache
	go.core/lesson6/pkg/cache/local v0.0.0 => ../pkg/cache/local
	go.core/lesson6/pkg/cache/membot v0.0.0 => ../pkg/cache/membot
	go.core/lesson6/pkg/index v0.0.0 => ../pkg/index
	go.core/lesson6/pkg/storage v0.0.0 => ../pkg/storage
	go.core/lesson6/pkg/storage/bstree v0.0.0 => ../pkg/storage/bstree
	go.core/lesson6/pkg/crawler v0.0.0 => ../pkg/crawler
	go.core/lesson6/pkg/crawler/spider v0.0.0 => ../pkg/crawler/spider
	go.core/lesson6/pkg/crawler/membot v0.0.0 => ../pkg/crawler/membot
)

require (
	go.core/lesson6/pkg/cache v0.0.0
	go.core/lesson6/pkg/cache/local v0.0.0
	go.core/lesson6/pkg/engine v0.0.0
	go.core/lesson6/pkg/index v0.0.0
	go.core/lesson6/pkg/storage v0.0.0
	go.core/lesson6/pkg/storage/bstree v0.0.0
	go.core/lesson6/pkg/crawler v0.0.0
	go.core/lesson6/pkg/crawler/spider v0.0.0
)
