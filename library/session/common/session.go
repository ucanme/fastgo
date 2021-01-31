package common
type Session struct {
	Sid string `json:"sid"`
	Value string `json:"value"`
	LastAccessTime int64 `json:"last_access_time"`
}