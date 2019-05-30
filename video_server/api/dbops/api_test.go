package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var (
	tempvid string
)

// init(dblogin, truncate tables )->run tests-> clear data(truncate tables)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("Truncate video_info")
	dbConn.Exec("truccate conments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// 使用subtest 的方式, 可以控制test 的执行执行顺序.
func TestUserWordFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredentTial("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser:%v \n", err)
	}
}

// 测试获取用户.
func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser:%v \n", err)
	}
	fmt.Printf(" getUser success pwd = %v \n", pwd)
}

// 测试删除用户
func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of testDeleteUser:%v \n", err)
	}
}

// 查看是否真正的被删除掉.
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of testRegetUser:%v \n", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}

	// fmt.Println(" 执行 testRegetUser() 这个方法, 最后的结果 pwd = ", pwd)
	// 这里的结果 pwd 是空的字符串. "", 因为去 获取已经删除的一个用户, 是无效的, 返回的是一个空串.
}

func testVideoInfoWordFlow(t *testing.T) {
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "myVideo")
	if err != nil {
		t.Errorf("Error of AddVideoInfo :%v", err)
	}

	tempvid = vi.Id
}

// 测试获取视频信息.
func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo err:%v", err)
	}
}

// 测试删除视频信息.
func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo err :%v", err)
	}
}

// 测试 再次获取视频信息.
func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo err :%v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

// 测试 增加新的评论
func testAddComments(t *testing.T) {
	vid := "12345"
	err := AddNewComments(vid, 1, "content_1 first comment")
	if err != nil {
		t.Errorf("Error of AddNewComments err = %v", err)
	}
}

// 测试 展示所有的评论
func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, err := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	if err != nil {
		t.Errorf("Error of listComments strconv.Atoi err= %v", err)
		return
	}

	commentList, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments err = %v", err)
		return
	}
	for i, ele := range commentList {
		fmt.Printf("commnet := %d, %v\n", i, ele)
	}
}
