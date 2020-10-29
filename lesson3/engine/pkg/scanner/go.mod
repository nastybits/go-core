module education/lesson3/engine/pkg/scanner

go 1.13

replace education/lesson3/crawler/pkg/spider v0.0.0 => ./../../../crawler/pkg/spider

require (
	education/lesson3/crawler/pkg/spider v0.0.0
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
)
