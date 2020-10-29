module education/lesson3/engine/cmd/engine

go 1.13

replace education/lesson3/engine/pkg/spider v0.0.0 => ./../../pkg/spider
replace education/lesson3/engine/pkg/stub v0.0.0 => ./../../pkg/stub

require (
	education/lesson3/engine/pkg/spider v0.0.0
	education/lesson3/engine/pkg/stub v0.0.0
)
