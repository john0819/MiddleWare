package freecache

import (
	"sync"

	"github.com/coocood/freecache"
)

// freecache 做成缓存管理器, 不同的应用可以共用一个缓存管理器

var (
	CacheManagerInstance CacheManager
	CacheManagerOnce     sync.Once
)

type CacheManager interface {
	GetOrCreateCache(name string, size int) *freecache.Cache
	GetCache(name string) (*freecache.Cache, bool)
	DeleteCache(name string) bool
}

func InitCacheManager() {
	GetCacheManager()
}

func GetCacheManager() CacheManager {
	CacheManagerOnce.Do(func() {
		CacheManagerInstance = &cacheManager{}
	})
	return CacheManagerInstance
}

type cacheManager struct {
	caches sync.Map
}

// 根据不同应用name获取对应的cache
func (cm *cacheManager) GetCache(name string) (*freecache.Cache, bool) {
	cache, ok := cm.caches.Load(name)
	if !ok {
		return nil, false
	}
	return cache.(*freecache.Cache), true
}

func (cm *cacheManager) DeleteCache(name string) bool {
	if _, exists := cm.caches.Load(name); exists {
		cm.caches.Delete(name)
		return true
	}
	return false
}

func (cm *cacheManager) GetOrCreateCache(name string, size int) *freecache.Cache {
	cache, ok := cm.GetCache(name)
	if !ok {
		cache = freecache.NewCache(size)
		cm.caches.Store(name, cache)
	}
	return cache
}
