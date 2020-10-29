package scanner

import "education/lesson3/crawler/pkg/spider"

type Web int

func (s Web) Scan(url string, depth int) (map[string]string, error) {
	return spider.Scan(url, depth)
}
