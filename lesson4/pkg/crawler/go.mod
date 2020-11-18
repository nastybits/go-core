module go.core/lesson4/pkg/crawler

go 1.13

replace go.core/lesson4/pkg/document v0.0.0 => ./../document
replace go.core/lesson4/pkg/crawler/stub v0.0.0 => ./stub

require (
	go.core/lesson4/pkg/document v0.0.0
	go.core/lesson4/pkg/crawler/stub v0.0.0
)
