package luogu

import "encoding/json"

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

type List[T any] struct {
	Result  []T `json:"result"`
	PerPage int `json:"perPage"`
	Count   int `json:"count"`
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
	User              UserDetails      `json:"user"`
	EloMax            struct{}         `json:"eloMax"`
	PassedProblems    []ProblemSummary `json:"passedProblems"`
	SubmittedProblems []ProblemSummary `json:"submittedProblems"`
}

type TeamSummary struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsPremium bool   `json:"isPremium"`
}

type ProblemSummary struct {
	Pid        string `json:"pid"`
	Title      string `json:"title"`
	Difficulty int    `json:"difficulty"`
	FullScore  int    `json:"fullScore"`
	Type       string `json:"type"`
}

type ContestSummary struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	StartTime int    `json:"startTime"`
	EndTime   int    `json:"endTime"`
}

type Contest[T map[string]any] struct {
	RuleType           int  `json:"ruleType"`
	VisibilityType     int  `json:"visibilityType"`
	InvitationCodeType int  `json:"invitationCodeType"`
	Rated              bool `json:"rated"`
	EloThreshold       int  `json:"eloThreshold"`
	Host               T    `json:"host"`
	ProblemCount       int  `json:"problemCount"`
	ContestSummary
}

type ContestDetails[T map[string]any] struct {
	Description       string `json:"description"`
	TotalParticipants int    `json:"totalParticipants"`
	EloDone           bool   `json:"eloDone"`
	CanEdit           bool   `json:"canEdit"`
	Contest[T]
}

type ContestData[T map[string]any] struct {
	Contest         ContestDetails[T] `json:"contest"`
	ContestProblems []struct {
		Score     int            `json:"score"`
		Problem   ProblemSummary `json:"problem"`
		Submitted bool           `json:"submitted"`
	} `json:"contestProblems"`
	IsScoreboardFrozen bool     `json:"isScoreboardFrozen"`
	AccessLevel        int      `json:"accessLevel"`
	Joined             bool     `json:"joined"`
	UserElo            struct{} `json:"userElo"`
}

type ScoreDetails map[string]struct {
	Score       int `json:"score"`
	RunningTime int `json:"runningTime"`
}

type Score struct {
	Details     json.RawMessage `json:"details"`
	User        UserSummary     `json:"user"`
	Score       int             `json:"score"`
	RunningTime int             `json:"runningTime"`
}

type GetScoreboardResponse struct {
	Scoreboard    List[Score]    `json:"scoreboard"`
	UserScore     Score          `json:"userScore"`
	UserRank      int            `json:"userRank"`
	FirstBloodUID map[string]int `json:"firstBloodUID"`
}
