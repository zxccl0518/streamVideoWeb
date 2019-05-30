package session

import (
	"log"
	"sync"
	"time"

	"github.com/zxccl0518/streamVideoWeb/video_server/api/defs"
	"github.com/zxccl0518/streamVideoWeb/video_server/api/utils"

	"github.com/zxccl0518/streamVideoWeb/video_server/api/dbops"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

// 返回当前时间的一个时间戳(纳秒)
func nowInMiliion() int64 {
	return time.Now().UnixNano() / 1000000
}

// 装载所有的sessions
func LoadSessionsFromDB() {
	res, err := dbops.RetrieveAllSessions()
	if err != nil {
		log.Printf("LoadSessionsFromDB() error = %v", err)
		return
	}

	res.Range(func(k, v interface{}) bool {
		ss, ok := v.(*defs.SimpleSession)
		if ok == true {
			sessionMap.Store(k, ss)
		}
		return true
	})
}

func GenerateNewSessionId(un string) string {
	// 创建 一个uuid
	sid, err := utils.NewUUID()
	if err != nil {
		log.Printf("创建uuid 出错, err = %v", err)
		return ""
	}
	// 创建一个64位的 当前时间的时间戳
	ct := nowInMiliion()
	// 创建一个ttl 超时时间. 30min
	ttl := ct + 30*60*1000
	// 创建一个sessionid
	ss := &defs.SimpleSession{Username: un, TTL: ttl}

	// 将新创建的sessionid  分别存储到cache缓存和数据库中.
	sessionMap.Store(sid, ss)

	err = dbops.InsertSession(sid, ttl, un)
	if err != nil {
		return "nil"
	}

	return sid
}

// 判断是不是 超时了. 如果超时了,删除对应sessionid 的内容, 如果没有超时 返回用户名.
func IsSessionExpired(sid string) (string, bool) {
	// Load()方法返回的值是一个接口类型的, 所以要做断言 判断其类型.
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMiliion()
		if ss.(*defs.SimpleSession).TTL < ct {
			// delete expire session
			// session 超时了.
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}

// 删除过期的sessionid
func DeleteExpiredSession(sid string) error {
	// 分别从 cache缓存 和 数据库中删除.
	sessionMap.Delete(sid)
	err := dbops.DeleteSession(sid)

	return err
}
