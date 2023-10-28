package luogu

type UserSearchResponse struct {
	Users []UserSummary `json:"users"`
}

func SearchUser(keyword string) (UserSearchResponse, error) {
	return Request[UserSearchResponse]("GET", "https://www.luogu.com.cn/api/user/search?keyword="+keyword, nil)
}
