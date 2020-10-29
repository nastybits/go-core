// Package stub - пакет-заглушка для реализации поискового робота
// Имитирует паука обходящего web-страницы при помощи метода Scan
package stub

type Spider int

// Scan - метод типа Spider возвращающий предустановленные данные
func (s *Spider) Scan(url string, depth int) (map[string]string, error) {
	pages := make(map[string]string)
	pages["http://test1.url"]  = "Some test 1 url title"
	pages["http://test2.url"]  = "Some test 2 url title"
	pages["http://test3.url"]  = "Some test 3 url title"
	pages["http://test4.url"]  = "Some test 4 url title"
	pages["http://test5.url"]  = "Some test 5 url title"
	pages["http://test6.url"]  = "Some test 6 url title"
	pages["http://test7.url"]  = "Some test 7 url title"
	pages["http://test8.url"]  = "Some test 8 url title"
	pages["http://test9.url"]  = "Some test 9 url title"

	return pages, nil
}

