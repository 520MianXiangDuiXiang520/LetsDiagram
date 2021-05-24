package cooperate

import (
    `lets_diagram/src`
    "math"
    "math/rand"
    "strconv"
    "sync"
)

// 用来维持所有 canvas 对应的协作码

func getCode() string {
	size := src.GetSetting().DefaultSetting.CooperateCodeLen
	code := rand.Int63n(int64(math.Pow(10,
		float64(size)))-1-int64(math.Pow(10,
		float64(size-1)))) + int64(math.Pow(10, float64(size-1)))
	return strconv.Itoa(int(code))
}

type cooperateCodeTable struct {
	codeDepot   sync.Map
	canvasDepot sync.Map
}

// 生成一个新的协作码
func (c *cooperateCodeTable) newCode(canvasID uint) string {
	for {
		code := getCode()
		_, ok := c.codeDepot.Load(code)
		if !ok {
			return code
		}
	}
}

func (c *cooperateCodeTable) GetCode(canvasID uint) string {
	code := c.newCode(canvasID)
	v, stored := c.canvasDepot.LoadOrStore(canvasID, code)
	if stored {
		return v.(string)
	}
	c.codeDepot.Store(code, canvasID)
	return code
}

func (c *cooperateCodeTable) Remove(canvasID uint) {
    code, ok := c.canvasDepot.Load(canvasID)
    if !ok {
        return
    }
    c.canvasDepot.Delete(canvasID)
    c.codeDepot.Delete(code)
}

func (c *cooperateCodeTable) GetCanvasID(code string) (uint, bool) {
	v, ok := c.codeDepot.Load(code)
	if !ok {
		return 0, false
	}
	return v.(uint), true
}

var cct *cooperateCodeTable
var initCCTMutex sync.Mutex

func InitCooperateCodeTable() {
	if cct == nil {
		initCCTMutex.Lock()
		if cct == nil {
			cct = &cooperateCodeTable{codeDepot: sync.Map{}, canvasDepot: sync.Map{}}
		}
		initCCTMutex.Unlock()
	}
}

func GetCooperateCodeTable() *cooperateCodeTable {
	if cct == nil {
		InitCooperateCodeTable()
	}
	return cct
}
