module go.core/lesson5/pkg/storage

go 1.15

replace go.core/lesson5/pkg/crawler v0.0.0 => ../crawler
replace go.core/lesson5/pkg/crawler/membot v0.0.0 => ../crawler/membot

require (
	go.core/lesson5/pkg/crawler v0.0.0
)
