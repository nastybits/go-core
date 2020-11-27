module go.core/lesson8/pkg/cache/membot

go 1.13

replace (
		go.core/lesson8/pkg/fixtures v0.0.0 => ../../fixtures
		go.core/lesson8/pkg/crawler v0.0.0 => ../../crawler
    	go.core/lesson8/pkg/crawler/membot v0.0.0 => ../../crawler/membot
)

require (
	go.core/lesson8/pkg/fixtures v0.0.0
)
