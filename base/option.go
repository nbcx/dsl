package base

type H map[string]interface{}

func (h H) Get(k string, defauleValue ...interface{}) interface{} {
	if value, ok := h[k]; ok {
		return value
	}
	return defauleValue[0]
}

func (h H) GetString(k string, defauleValue ...string) string {
	return h.Get(k, defauleValue[0]).(string)
}

func (h H) GetInt(k string, defauleValue ...int) int {
	return h.Get(k, defauleValue[0]).(int)
}

func (h H) GetUint64(k string, defauleValue ...uint64) uint64 {
	return h.Get(k, defauleValue[0]).(uint64)
}
