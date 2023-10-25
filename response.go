package luogu

// https://github.com/0f-0b/luogu-api-docs/blob/main/luogu-api.d.ts

type DataResponse[T any] struct {
	Code            int    `json:"code"`
	CurrentTemplate string `json:"currentTemplate"`
	CurrentData     T      `json:"currentData"`
	CurrentTitle    string `json:"currentTitle"`
	CurrentTheme    string `json:"currentTheme"`
	CurrentTime     int    `json:"currentTime"`
}

type User struct {
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

type UserData struct {
	User              User       `json:"user"`
	EloMax            struct{}   `json:"eloMax"`
	PassedProblems    []struct{} `json:"passedProblems"`
	SubmittedProblems []struct{} `json:"submittedProblems"`
}
