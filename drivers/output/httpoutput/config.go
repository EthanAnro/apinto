package httpoutput

import (
	"encoding/json"
	"github.com/eolinker/eosc"
	"net/http"
)

type Config struct {
	Method    string               `json:"method" yaml:"method" enum:"GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS"`
	Url       string               `json:"url" yaml:"url" format:"uri"`
	Headers   map[string]string    `json:"headers" yaml:"headers"`
	Type      string               `json:"type" yaml:"type" enum:"json,line"`
	Formatter eosc.FormatterConfig `json:"formatter" yaml:"formatter"`
}

func (h *Config) isConfUpdate(conf *Config) bool {
	if h.Method != conf.Method || h.Url != conf.Url || !compareTwoMapStringEqual(h.Headers, conf.Headers) {
		return true
	}
	return false
}

func compareTwoMapStringEqual(mapA, mapB map[string]string) bool {
	if len(mapA) != len(mapB) {
		return false
	}
	length := len(mapA)
	keySlice := make([]string, 0, length)
	dataValueA := make([]string, 0, length)
	dataValueB := make([]string, 0, length)
	for k, v := range mapA {
		keySlice = append(keySlice, k)
		dataValueA = append(dataValueA, v)
	}

	for _, key := range keySlice {
		if vb, has := mapB[key]; has {
			dataValueB = append(dataValueB, vb)
		} else {
			return false
		}
	}

	strValueA, _ := json.Marshal(dataValueA)
	strValueB, _ := json.Marshal(dataValueB)

	return string(strValueA) == string(strValueB)
}

func toHeader(items map[string]string) http.Header {
	header := make(http.Header)
	for k, v := range items {
		header.Set(k, v)
	}
	return header
}
