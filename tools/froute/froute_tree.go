package froute

import (
	"net/http"
	"strings"
	"sync"
)

var ErrorHandlers ErrorHandler

type ErrorHandler struct {
}

func (e ErrorHandler) Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

// MiddlewareResult 中间件处理结果
type MiddlewareResult struct {
	// 中间件是否中止后续的handler处理
	Suspend bool
}

// FrouteGroup 路由组
type FrouteGroup struct {
	// 路由是否能覆盖重写,用于防止因错误路由覆盖导致问题
	RouteRewritable bool
	// 路由组前缀
	prefix string
	// 是否启用,控制节点
	pause bool
	// 是否运行简短的运行,例如/处理定义在/user的处理器
	ShortPatch bool
	// 分组拦截处理器
	middeleware func(http.ResponseWriter, *http.Request, *FrouteNode) MiddlewareResult
	// 分区组的使用锁
	lock sync.RWMutex
	// 所有分支
	branches map[string]*FrouteBranch
}

// FrouteBranch 路由分支
type FrouteBranch struct {
	// 分支处理的方法
	method string
	// 分支所属分组
	group *FrouteGroup
	// 层次的处理器
	handler func(http.ResponseWriter, *http.Request)
	// 当前分支下的所有节点
	nodes map[string]*FrouteNode
}

// FrouteNode 路由分组的节点
type FrouteNode struct {
	// 层次
	floor int
	// 上一个层次的节点
	preFloorNode *FrouteNode
	// 节点所属Froute分支
	branch *FrouteBranch
	// 层次的字符值,可使用*通配
	value string
	// 当前分支下的所有节点
	nodes map[string]*FrouteNode
	// 校验内容是否走该路由,比如必须要求为数字
	check func(string) bool
	// 层次的处理器
	handler func(http.ResponseWriter, *http.Request)
}

// Pause 暂停组的使用
func (group *FrouteGroup) Pause() {
	group.lock.Lock()
	defer group.lock.Unlock()
	group.pause = true
}

// Resume 恢复组的使用
func (group *FrouteGroup) Resume() {
	group.lock.Lock()
	defer group.lock.Unlock()
	group.pause = false
}

// SetMiddleware 设置当前分组的拦截处理器
func (group *FrouteGroup) SetMiddleware(middeleware func(http.ResponseWriter, *http.Request, *FrouteNode) MiddlewareResult) {
	group.lock.Lock()
	defer group.lock.Unlock()
	group.middeleware = middeleware
}

// fetch 查找某个路径下面使用特定方法的请求的对应的处理器
func (group *FrouteGroup) fetch(uri, method string) (handler func(http.ResponseWriter, *http.Request), node *FrouteNode, fetched bool) {
	group.lock.RLock()
	defer group.lock.RUnlock()
	if strings.HasPrefix(uri, group.prefix) || group.prefix == "" {
		branch := group.branches[method]
		if branch == nil {
			return
		}
		path := uri[len(group.prefix):]
		var processedPath string
		if strings.HasSuffix(path, "/") && len(path) > 2 {
			processedPath = path[1 : len(path)-2]
		} else {
			processedPath = path[1:]
		}
		nodeValues := strings.Split(processedPath, "/")
		var floor int
		var value string
		for floor, value = range nodeValues {
			var temp *FrouteNode
			if floor == 0 {
				temp = branch.nodes[value]
			} else if node != nil {
				temp = node.nodes[value]
			}
			if temp == nil {
				break
			}
			node = temp
		}
		fetched = floor == len(nodeValues)-1
		if node == nil {
			handler = branch.handler
		} else {
			handler = node.handler
		}
	}
	return
}

/*addHandler 在路由组内添加路由
method 路由使用的方法
router 路由的路径
check 检查路径中的某个值是否符合要求,如果不符合则不使用该路由
*/
func (group *FrouteGroup) AddHandler(method, path string, check func(floor int) func(string) bool, handler func(http.ResponseWriter, *http.Request)) {
	group.lock.Lock()
	defer group.lock.Unlock()
	if group.branches == nil {
		group.branches = make(map[string]*FrouteBranch, 0)
	}
	if method == "" {
		panic("method not assined")
	}
	if !strings.HasPrefix(path, "/") {
		panic("path need to start with /")
	}
	branch := group.branches[method]
	if branch == nil {
		branch = &FrouteBranch{
			method: method,
			group:  group,
		}
	}
	if path == "/" {
		branch.handler = handler
	} else {
		var processedPath string
		if strings.HasSuffix(path, "/") {
			processedPath = path[1 : len(path)-2]
		} else {
			processedPath = path[1:]
		}
		nodeValues := strings.Split(processedPath, "/")
		var preFloorNode *FrouteNode
		floorSize := len(nodeValues)
		for floor, value := range nodeValues {
			if value == "" {
				panic("path parsed with error: //")
			}
			floorNode := &FrouteNode{
				floor:        floor,
				preFloorNode: preFloorNode,
				branch:       branch,
				value:        value,
			}
			floorNode.check = check(floor)
			if preFloorNode == nil {
				if branch.nodes == nil {
					branch.nodes = make(map[string]*FrouteNode)
				}
				oldNode := branch.nodes[value]
				if oldNode != nil && !group.RouteRewritable && floor == floorSize-1 {
					panic("handler could not be rewritable!")
				}
				if oldNode != nil {
					floorNode = oldNode
				}
				branch.nodes[value] = floorNode
			} else {
				if preFloorNode.nodes == nil {
					preFloorNode.nodes = make(map[string]*FrouteNode)
				}
				oldNode := preFloorNode.nodes[value]
				if oldNode != nil && !group.RouteRewritable && floor == floorSize-1 {
					panic("handler could not be rewritable!")
				}
				if oldNode != nil {
					floorNode = oldNode
				}
				preFloorNode.nodes[value] = floorNode
			}
			if floor == floorSize-1 {
				floorNode.handler = handler
			}
			preFloorNode = floorNode
		}
	}
	group.branches[method] = branch
}