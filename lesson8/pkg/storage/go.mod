module go.core/lesson8/pkg/storage

go 1.15

replace go.core/lesson8/pkg/fixtures v0.0.0 => ../fixtures
replace go.core/lesson8/pkg/crawler v0.0.0 => ../crawler
replace go.core/lesson8/pkg/crawler/membot v0.0.0 => ../crawler/membot

require (
	go.core/lesson8/pkg/crawler v0.0.0
)
