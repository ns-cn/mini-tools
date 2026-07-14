package froute

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// Frouters 聚合所有的接口
type Frouters struct {
	// 所有的组聚合
	routers []*FrouteGroup
	//
	lock sync.RWMutex
}

func (routers *Frouters) AddGroup(group *FrouteGroup) {
	routers.lock.Lock()
	defer routers.lock.Unlock()
	if routers.routers == nil {
		routers.routers = make([]*FrouteGroup, 0)
	}
	routers.routers = append(routers.routers, group)
	fmt.Println(len(routers.routers))
}

func (f *Frouters) ServeHttp(w http.ResponseWriter, r *http.Request) {
	f.lock.RLock()
	defer f.lock.RUnlock()
	uri := r.RequestURI
	method := r.Method
	fmt.Println(method, uri)
	if len(f.routers) == 0 {
		ErrorHandlers.Handle404(w, r)
		return
	}
	handled := false
	for _, group := range f.routers {
		if strings.HasPrefix(uri, group.prefix) {
			group.lock.RLock()
			defer group.lock.RUnlock()
			fetchHandler, node, fetched := group.fetch(uri, method)
			if fetchHandler != nil && (fetched || group.ShortPatch) {
				middleware := MiddlewareResult{Suspend: false}
				if group.middeleware != nil {
					middleware = group.middeleware(w, r, node)
				}
				if !middleware.Suspend {
					fetchHandler(w, r)
					handled = true
				}
				break
			}
		}
	}
	if !handled {
		ErrorHandlers.Handle404(w, r)
	}
}