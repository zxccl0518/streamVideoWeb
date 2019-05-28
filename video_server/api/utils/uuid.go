package utils

import uuid "github.com/satori/go.uuid"

func NewUUID() (string, error) {
	// uuid := make([]byte, 16)
	// n, err := io.ReadFull(rand.Reader, uuid)
	// if n != len(uuid) || err != nil {
	// 	return "", err
	// }
	// uuid[8] = uuid[8]&^0xc0 | 0x80
	// uuid[6] = uuid[6]&^0xf0 | 0x40
	// return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:12], uuid[12:16]), err
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	// 将字节切片形式的uuid 转换为 string
	uuidStr := string(u.Bytes())

	return uuidStr, err
}
