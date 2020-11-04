module go.core/lesson3/engine/cmd/engine

go 1.13

replace go.core/lesson3/engine/pkg/spider v0.0.0 => ./../../pkg/spider
replace go.core/lesson3/engine/pkg/stub v0.0.0 => ./../../pkg/stub

require (
	go.core/lesson3/engine/pkg/spider v0.0.0
	go.core/lesson3/engine/pkg/stub v0.0.0
)
