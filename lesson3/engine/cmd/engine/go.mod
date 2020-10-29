module education/lesson3/engine/cmd/engine

go 1.13

replace education/lesson3/crawler/pkg/spider v0.0.0 => ./../../../crawler/pkg/spider
replace education/lesson3/engine/pkg/scanner v0.0.0 => ./../../pkg/scanner

require (
	education/lesson3/crawler/pkg/spider v0.0.0
	education/lesson3/engine/pkg/scanner v0.0.0
)
