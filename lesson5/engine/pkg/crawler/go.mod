module go.core/lesson5/engine/pkg/crawler

go 1.13

replace go.core/lesson5/engine/pkg/document v0.0.0 => ../document
replace go.core/lesson5/engine/pkg/crawler/membot v0.0.0 => ./membot

require (
	go.core/lesson5/engine/pkg/document v0.0.0
	go.core/lesson5/engine/pkg/crawler/membot v0.0.0
)
