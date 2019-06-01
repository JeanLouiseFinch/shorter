package shorter

import (
	"fmt"
	"net/url"
	"strings"
)

type Shortener interface {
	Shorten(urle string) string
	Resolve(urle string) string
}

// Short. поле max - это наш инкремент на шортурл,collect - таблица соответсвий, ключ - короткая ссылка
type Short struct {
	collect map[string]string
	max     int
}

//NewShort - возвращаем указатель на нашу структуру с инициализированной мапой
func NewShort() *Short {
	return &Short{
		collect: make(map[string]string),
		max:     0,
	}
}
func (s *Short) Shorten(urle string) string {
	var (
		path     string
		shortURL string
	)
	path = getPath(urle)
	if path == "" {
		return ""
	}
	for key := range s.collect {
		if s.collect[key] == urle {
			return key
		}
	}
	shortURL = urle[:strings.Index(urle, path)] + "/" + fmt.Sprintf("%d", s.max)
	s.collect[shortURL] = urle
	s.max = s.max + 1
	return shortURL
}

func (s *Short) Resolve(urle string) string {
	if _, ok := s.collect[urle]; ok {
		return s.collect[urle]
	}
	return ""
}

// getPath - пытается распарсить наш url и проверить на валидность, возвращает путь или пустую строку при ошибках
func getPath(u string) string {
	uParse, err := url.Parse(u)
	switch {
	case err != nil:
		return ""
	case uParse.Scheme != "":
		return uParse.Path
	case strings.Index(uParse.Path, "/") != -1 && !strings.HasPrefix(uParse.Path, "/") && strings.Index(uParse.Path, ".") != -1 && !strings.HasPrefix(uParse.Path, ".") && strings.Index(uParse.Path, ".")+2 < strings.Index(uParse.Path, "/"):
		return uParse.Path[strings.Index(uParse.Path, "/"):]
	default:
		return ""
	}
}
