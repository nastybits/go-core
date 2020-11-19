module go.core/lesson6/pkg/index

go 1.15

replace go.core/lesson6/pkg/crawler v0.0.0 => ../crawler
replace go.core/lesson6/pkg/crawler/membot v0.0.0 => ../crawler/membot

require (
	go.core/lesson6/pkg/crawler v0.0.0
)
