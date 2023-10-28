package luogu

// https://github.com/0f-0b/luogu-api-docs/blob/main/luogu-api.d.ts

type DataResponse[T any] struct {
	Code            int    `json:"code"`
	CurrentTemplate string `json:"currentTemplate"`
	CurrentData     T      `json:"currentData"`
	CurrentTitle    string `json:"currentTitle"`
	CurrentTheme    string `json:"currentTheme"`
	CurrentTime     int    `json:"currentTime"`
	CurrentUser     User   `json:"currentUser"`
}

type UserSummary struct {
	Uid        int    `json:"uid"`
	Name       string `json:"name"`
	Slogan     string `json:"slogan"`
	Badge      string `json:"badge"`
	IsAdmin    bool   `json:"isAdmin"`
	IsBanned   bool   `json:"isBanned"`
	Color      string `json:"color"`
	CcfLevel   int    `json:"ccfLevel"`
	Background string `json:"background"`
}

type User struct {
	FollowingCount int    `json:"followingCount"`
	FollowerCount  int    `json:"followerCount"`
	Ranking        int    `json:"ranking"`
	EloValue       int    `json:"eloValue"`
	BlogAddress    string `json:"blogAddress"`
	UserSummary
}

type UserDetails struct {
	RegisterTime int    `json:"registerTime"`
	Introduction string `json:"introduction"`
	Prize        []struct {
		Year        int    `json:"year"`
		ContestName string `json:"contestName"`
		Prize       string `json:"prize"`
	} `json:"prize"`
	Elo struct{} `json:"elo"`
	User
}

type UserData struct {
	User              UserDetails `json:"user"`
	EloMax            struct{}    `json:"eloMax"`
	PassedProblems    []struct{}  `json:"passedProblems"`
	SubmittedProblems []struct{}  `json:"submittedProblems"`
}
