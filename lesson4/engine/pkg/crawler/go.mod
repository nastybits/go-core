module go.core/lesson4/engine/pkg/crawler

go 1.13

replace go.core/lesson4/engine/pkg/document v0.0.0 => ../document
replace go.core/lesson4/engine/pkg/crawler/stub v0.0.0 => ../crawler/stub

require (
	go.core/lesson4/engine/pkg/document v0.0.0
	go.core/lesson4/engine/pkg/crawler/stub v0.0.0
)
