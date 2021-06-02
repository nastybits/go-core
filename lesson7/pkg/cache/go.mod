module go.core/lesson7/pkg/cache

go 1.13

replace (
	go.core/lesson7/pkg/crawler v0.0.0 => ../crawler
	go.core/lesson7/pkg/crawler/spider v0.0.0 => ../crawler/spider
	go.core/lesson7/pkg/crawler/membot v0.0.0 => ../crawler/membot
)

require (
	go.core/lesson7/pkg/crawler v0.0.0
)
