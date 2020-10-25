package obj_pool

import (
	"errors"
	"time"
)

/*建议使用不同的池来缓冲不同的对象*/

type ReuseableObj struct {

}

type ObjPool struct {
	bufChan chan *ReuseableObj //用于缓冲可重用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReuseableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReuseableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReuseableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <- time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReuseableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
