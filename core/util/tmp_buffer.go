package util

import (
	"sync"
	"time"
)

var _tmpBuf sync.Map

var _tmpBufCleanOnce sync.Once

var _bufferTTL = 5 * time.Second

type queryResult struct {
	ret interface{}
	t   time.Time
}

//根据sql加载缓存
func LoadFromBuffer(strSQL string) (interface{}, bool) {

	tmp, ok := _tmpBuf.Load(strSQL)
	if !ok || nil == tmp {
		return nil, false
	}

	qr, ok := tmp.(*queryResult)
	if !ok || nil == qr {
		return nil, false
	}

	if time.Since(qr.t) > _bufferTTL {
		return nil, false
	}
	return qr.ret, true
}

//StoreToBuffer 保存sql-结果缓存
func StoreToBuffer(strSQL string, resp interface{}) {

	tmp, ok := _tmpBuf.Load(strSQL)

	if ok {
		qr, ok := tmp.(*queryResult)
		if ok && nil != qr {
			qr.ret = resp
			qr.t = time.Now()
			return
		}
	}
	_tmpBuf.Store(strSQL, &queryResult{
		resp, time.Now(),
	})
	_tmpBufCleanOnce.Do(func() {

		go func() {
			for {
				time.Sleep(_bufferTTL)

				_tmpBuf.Range(func(key, val interface{}) bool {
					qr, ok := val.(*queryResult)
					if ok && nil != qr {
						if time.Since(qr.t) > _bufferTTL {
							_tmpBuf.Delete(key)
						}
					}
					return true
				})
			}
		}()
	})
}
