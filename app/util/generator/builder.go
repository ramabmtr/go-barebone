package generator

import "strings"

const (
	cacheKeyDivider = "::"
)

func BuildCacheKey(keys ...string) string {
	key := "cache" + cacheKeyDivider
	for _, s := range keys {
		key = key + s + cacheKeyDivider
	}
	return strings.TrimSuffix(key, cacheKeyDivider)
}
