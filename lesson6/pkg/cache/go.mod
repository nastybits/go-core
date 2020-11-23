module go.core/lesson6/pkg/cache

go 1.13

replace (
	go.core/lesson6/pkg/crawler v0.0.0 => ../crawler
	go.core/lesson6/pkg/crawler/spider v0.0.0 => ../crawler/spider
	go.core/lesson6/pkg/crawler/membot v0.0.0 => ../crawler/membot
)

require (
	go.core/lesson6/pkg/crawler v0.0.0
)
