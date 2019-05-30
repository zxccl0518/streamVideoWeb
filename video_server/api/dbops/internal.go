package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"

	"github.com/zxccl0518/streamVideoWeb/video_server/api/defs"
)

// 插入一个 session
func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions(session_id, TTL, login_name) VALUE (?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	return nil
}

// 获取所有的 sessions
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	// var ttlint int64
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	}
	defer stmtOut.Close()

	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELETE * FROM sessions")
	if err != nil {
		log.Printf("RetrieveAllSessions() prepare err = %s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("RetrieveAllSessions() Query err = %s", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err = rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrive sessions error : %s", err)
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 != nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id :%s, ttl:%d", id, ss.TTL)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("DeleteSession Prepare error = %s", err)
		return err
	}

	if _, err = stmtOut.Exec(sid); err != nil {
		log.Printf("DeleteSession Exec error = %v", err)
		return err
	}

	return nil
}
