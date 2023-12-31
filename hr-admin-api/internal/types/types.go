// Code generated by goctl. DO NOT EDIT.
package types

type PingResp struct {
	Ping string `json:"ping"`
}

type LoginReq struct {
	Code string `path:"code"`
}

type AdminGroupPivot struct {
	UserId  int64 `json:"user_id"`
	GroupId int64 `json:"group_id"`
	Power   int32 `json:"power"`
}

type AdminGroup struct {
	Id    int64           `json:"id"`
	Name  string          `json:"name"`
	Pivot AdminGroupPivot `json:"pivot"`
}

type LoginResp struct {
	Id     int64        `json:"id"`
	Name   string       `json:"name"`
	UUID   string       `json:"uuid"`
	Email  string       `json:"email"`
	Avatar string       `json:"avatar"`
	Groups []AdminGroup `json:"groups"`
	Token  string       `json:"token"`
}

type OAuth struct {
	ClientId    string `json:"client_id"`
	RedirectUri string `json:"redirect_uri"`
}

type Url struct {
	BaseUri   string `json:"base_uri"`
	AdminUri  string `json:"admin_uri"`
	WechatUri string `json:"wechat_uri"`
}

type FrontendConfig struct {
	Debug bool  `json:"debug"`
	OAuth OAuth `json:"oauth"`
	Url   Url   `json:"url"`
}

type GetAdminResp struct {
	Id                int64        `json:"id"`
	UUID              string       `json:"uuid"`
	Name              string       `json:"name"`
	Email             string       `json:"email"`
	Avatar            string       `json:"avatar"`
	Groups            []AdminGroup `json:"groups"`
	DefaultStandardId int64        `json:"default_standard_id"`
	SmsEnabled        bool         `json:"sms_enabled"`
}

type SetDefaultStandardReq struct {
	StandardId int64 `form:"standard_id"`
}

type OAuthConfig struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AdminUri     string `json:"admin_uri"`
	WechatUri    string `json:"wechat_uri"`
	RedirectUri  string `json:"redirect_uri"`
}

type GetInfoReq struct {
	Id int64 `path:"id"`
}

type Intent struct {
	Group       string `json:"group"`
	ApplicantId int64  `json:"applicant_id"`
	Rank        int32  `json:"rank"`
	Parallel    int32  `json:"parallel"`
}

type Pivot struct {
	ApplicantId int64 `json:"applicant_id"`
	TimeId      int64 `json:"time_id"`
}

type SavedTime struct {
	Id              int64                  `json:"id"`
	Group           string                 `json:"group"`
	Date            string                 `json:"date"`
	Time            string                 `json:"time"`
	Rank            int32                  `json:"rank"`
	Location        string                 `json:"location"`
	TotalCnt        int32                  `json:"total_cnt"`
	Grade           string                 `json:"grade"`
	Campus          string                 `json:"campus"`
	Pivot           Pivot                  `json:"pivot"`
	ApplicantId     int64                  `json:"applicant_id"`
	Extend          map[string]interface{} `json:"extend"`
	ApplicantTimeId int64                  `json:"applicant_time_id"`
}

type Admin struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	DefaultStandardId int64  `json:"default_standard_id"`
}

type Admit struct {
	AdminId     int64  `json:"admin_id"`
	ApplicantId int64  `json:"applicant_id"`
	Group       string `json:"group"`
	Admin       Admin  `json:"admin"`
	UpdatedAt   string `json:"updated_at"`
}

type Applicant struct {
	Id        int64                  `json:"id"`
	WechatId  string                 `json:"wechat_id"`
	Name      string                 `json:"name"`
	Gender    string                 `json:"gender"`
	Phone     string                 `json:"phone"`
	Avatar    string                 `json:"avatar"`
	Form      map[string]interface{} `json:"form"`
	Intents   []Intent               `json:"intents"`
	Times     []SavedTime            `json:"times"`
	Admits    []Admit                `json:"admits"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

type GetInfoResp struct {
	Applicant Applicant `json:"applicant"`
}

type GetScoresReq struct {
	Id int64 `path:"id"`
}

type ScoreItem struct {
	AdminId           int64   `json:"admin_id"`
	ApplicantId       int64   `json:"applicant_id"`
	Score             float32 `json:"score"`
	Group             string  `json:"group"`
	StandardId        int64   `json:"standard_id"`
	EvaluationDetails string  `json:"evaluation_details"`
	Admin             Admin   `json:"admin"`
}

type GetScoresResp struct {
	Scores []ScoreItem `json:"scores"`
}

type GetGroupReq struct {
	Group string `form:"group"`
	State string `form:"state"`
	Page  int32  `form:"page"`
}

type GetGroupData struct {
	Id          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Gender      string                 `json:"gender"`
	Avatar      string                 `json:"avatar"`
	ScoresCount int32                  `json:"scores_count"`
	AvgScore    float32                `json:"avg_score"`
	Date        string                 `json:"date"`
	Time        string                 `json:"time"`
	Location    string                 `json:"location"`
	Extend      map[string]interface{} `json:"extend"`
	Admits      []Admit                `json:"admits"`
}

type Link struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

type GetGroupResp struct {
	CurrentPage  int32          `json:"current_page"`
	Data         []GetGroupData `json:"data"`
	FirstPageUrl string         `json:"first_page_url"`
	LastPage     int32          `json:"last_page"`
	LastPageUrl  string         `json:"last_page_url"`
	Links        []Link         `json:"links"`
	NextPageUrl  string         `json:"next_page_url"`
	Path         string         `json:"path"`
	PerPage      int32          `json:"per_page"`
	PrevPageUrl  string         `json:"prev_page_url"`
	From         int32          `json:"from"`
	To           int32          `json:"to"`
	Total        int32          `json:"total"`
}

type NameListItem struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GetNameListResp struct {
	List []NameListItem `json:"list"`
}

type AdmitSetReq struct {
	ApplicantId int64  `form:"applicant_id"`
	Group       string `form:"group"`
}

type AdmitResetReq struct {
	ApplicantId int64  `form:"applicant_id"`
	Group       string `form:"group"`
}

type DeleteApplicantReq struct {
	Id int64 `path:"id"`
}

type SetApplicantTimeExtendReq struct {
	AppicantId int64                  `path:"applicant_id"`
	Group      string                 `path:"group"`
	Extend     map[string]interface{} `json:"extend"`
}

type GetConfigReq struct {
	Key string `path:"key"`
}

type GetConfigResp struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SetConfigReq struct {
	Key   string `path:"key"`
	Value string `form:"value"`
}

type SetConfigResp struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetSmsServiceConfigResp struct {
	Debug    bool   `json:"debug"`
	Enabled  bool   `json:"enabled"`
	SendUrl  string `json:"sendUrl"`
	UserName string `json:"userName"`
}

type Standard struct {
	Name   string `json:"name"`
	Weight int32  `json:"weight"`
}

type EvaluationStandard struct {
	Id              int64           `json:"id"`
	Name            string          `json:"name"`
	ScoresCount     int32           `json:"scores_count"`
	LastEditAdminId int64           `json:"last_edit_admin_id"`
	LastEditAdmin   EvaluationAdmin `json:"last_edit_admin"`
	Standard        []Standard      `json:"standard"`
	UpdatedAt       string          `json:"updated_at"`
}

type GetEvaluationInfoReq struct {
	Id int64 `path:"id"`
}

type GetEvaluationInfoResp struct {
	EvaluationStandard EvaluationStandard `json:"evaluation_standard"`
}

type EvaluationAdmin struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	DefaultStandardId int64  `json:"default_standard_id"`
}

type GetEvaluationListResp struct {
	List []EvaluationStandard `json:"list"`
}

type AddEvaluationReq struct {
	Name     string     `json:"name"`
	Standard []Standard `json:"standard"`
}

type AddEvaluationResp struct {
	EvaluationStandard EvaluationStandard `json:"evaluation_standard"`
}

type UpdateEvaluationReq struct {
	Id       int64      `json:"id"`
	Name     string     `json:"name"`
	Standard []Standard `json:"standard"`
}

type UpdateEvaluationResp struct {
	EvaluationStandard EvaluationStandard `json:"evaluation_standard"`
}

type DeleteEvaluationReq struct {
	Id int64 `json:"id"`
}

type GetExamGroupReq struct {
	Group string `path:"group"`
}

type GetExamGroupQuestion struct {
	Id       int64  `json:"id"`
	Question string `json:"question"`
}

type GetExamGroupResp struct {
	Questions []GetExamGroupQuestion `json:"questions"`
}

type CreateExamReq struct {
	Group    string `json:"group"`
	Question string `json:"question"`
}

type Question struct {
	Id       int64  `json:"id"`
	Group    string `json:"group"`
	Question string `json:"question"`
}

type CreateExamResp struct {
	Question Question `json:"question"`
}

type DeleteExamReq struct {
	Id int64 `path:"id"`
}

type DeleteExamResp struct {
	Id int64 `json:"id"`
}

type UpdateExamReq struct {
	Id       int64  `json:"id"`
	Group    string `json:"group"`
	Question string `json:"question"`
}

type UpdateExamResp struct {
	Id int64 `json:"id"`
}

type GetExamByApplicantIdReq struct {
	ApplicantId int64 `form:"applicantId"`
}

type QuestionAndAnswer struct {
	QuestionId int64  `json:"questionId"`
	Group      string `json:"group"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
}

type GetExamByApplicantIdResp struct {
	ApplicantId int64               `json:"applicantId"`
	Questions   []QuestionAndAnswer `json:"questions"`
}

type GetRemarksReq struct {
	Id int64 `path:"id"`
}

type RemarksResp struct {
	AdminId     int64  `json:"admin_id"`
	ApplicantId int64  `json:"applicant_id"`
	Remark      string `json:"remark"`
}

type PostRemarksReq struct {
	Id     int64  `path:"id"`
	Remark string `form:"remark"`
}

type DeleteRemarksReq struct {
	Id int64 `path:"id"`
}

type GetRemarksListReq struct {
	Id int64 `path:"id"`
}

type RemarkItem struct {
	AdminId     int64  `json:"admin_id"`
	ApplicantId int64  `json:"applicant_id"`
	Remark      string `json:"remark"`
	AdminName   string `json:"admin_name"`
	UpdatedAt   string `json:"updated_at"`
}

type GetRemarksListResp struct {
	Remarks []RemarkItem `json:"remarks"`
}

type GetScoreReq struct {
	Id    int64  `path:"id"`
	Group string `form:"group"`
}

type ScoresAdmin struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	DefaultStandardId int64  `json:"default_standard_id"`
}

type EvaluationDetail struct {
	Name   string  `json:"name"`
	Weight int32   `json:"weight"`
	Score  float32 `json:"score"`
}

type Score struct {
	AdminId           int64              `json:"admin_id"`
	ApplicantId       int64              `json:"applicant_id"`
	Score             float32            `json:"score"`
	Group             string             `json:"group"`
	StandardId        int64              `json:"standard_id"`
	EvaluationDetails []EvaluationDetail `json:"evaluation_details"`
	Admin             ScoresAdmin        `json:"admin"`
}

type PostScoreReq struct {
	Id                int64              `path:"id"`
	Group             string             `json:"group"`
	Score             float32            `json:"score"`
	StandardId        int64              `json:"standard_id"`
	EvaluationDetails []EvaluationDetail `json:"evaluation_details"`
}

type DeleteScoreReq struct {
	Id    int64  `path:"id"`
	Group string `form:"group"`
}

type DeleteScoreResp struct {
}

type GetRankReq struct {
	Group string `path:"group"`
	Id    int64  `path:"id"`
}

type GetRankResp struct {
	Rank  int32 `json:"rank"`
	Total int32 `json:"total"`
}

type DailyNewApplicant struct {
	Date   string `json:"date"`
	Counts int32  `json:"counts"`
}

type DailyNewResp struct {
	Applicants []DailyNewApplicant `json:"applicants"`
}

type ClassReq struct {
	Group string `form:"group"`
	Key   string `form:"key"`
}

type ClassItem struct {
	Key   string `json:"key"`
	Count int32  `json:"count"`
}

type ClassResp struct {
	Class []ClassItem `json:"class"`
}

type TimeConfigItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PostTimeConfigReq struct {
	TimeConfigs []TimeConfigItem `json:"timeConfigs"`
}

type PostTimeConfigResp struct {
	TimeConfigs []TimeConfigItem `json:"timeConfigs"`
}

type StatisticsItem struct {
	Group string `json:"group"`
	Count int32  `json:"count"`
}

type GetStatisticsResp struct {
	Times []StatisticsItem `json:"times"`
}

type ClearTimeReq struct {
	Groups []string `json:"groups"`
}

type GetExportReq struct {
	Groups []string `form:"groups"`
}

type PostUploadReq struct {
	Groups []string            `json:"groups"`
	Data   []map[string]string `json:"data"`
}

type GetScheduleReq struct {
	Group           string `form:"group"`
	Date            string `form:"date,optional"`
	Type            string `form:"type"`
	ShowNotSelected bool   `form:"showNotSelected,optional"`
}

type TimeTableItem struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Color    string `json:"color"`
	Category string `json:"category"`
}

type GetScheduleResp struct {
	TimeTable  []TimeTableItem `json:"timeTable"`
	Categories []string        `json:"categories"`
	Focus      string          `json:"focus"`
}

type ReceiveSmsReq struct {
	Phone   string `form:"m"`
	Content string `form:"c"`
}

type ReceiveSmsResp struct {
	Code int32 `json:"code"`
}

type GetSmsConfigResp struct {
	Types  []string `json:"types"`
	Alerts []string `json:"alerts"`
}

type GetApplicantSmsReq struct {
	Type     string `path:"type"`
	Page     int32  `form:"page"`
	PageSize int32  `form:"pageSize"`
}

type ApplicantSms struct {
	ApplicantId int64    `json:"id"`
	Name        string   `json:"name"`
	Phone       string   `json:"phone"`
	Status      int32    `json:"status"`
	Time        string   `json:"time"`
	Args        []string `json:"args"`
	Content     string   `json:"content"`
	ReplyCnt    int32    `json:"reply_cnt"`
}

type GetApplicantSmsResp struct {
	ApplicantSms []ApplicantSms `json:"applicantSms"`
	Total        int32          `json:"total"`
}

type SendSmsReq struct {
	Typ          string  `json:"typ"`
	ApplicanrIds []int64 `json:"applicantIds"`
}

type GetHistorySmsReq struct {
	ApplicantId int64 `path:"id"`
}

type HistorySms struct {
	SenderName    string `json:"senderName"`
	SenderPhone   string `json:"senderPhone"`
	ReceiverName  string `json:"receiverName"`
	ReceiverPhone string `json:"receiverPhone"`
	Content       string `json:"content"`
	Time          string `json:"time"`
	Right         bool   `json:"right"`
}

type GetHistorySmsResp struct {
	HistorySms []HistorySms `json:"historySms"`
}

type Room struct {
	Id                 int64    `json:"id"`
	Name               string   `json:"name"`
	Location           string   `json:"location"`
	Status             int32    `json:"status"`
	Admins             []string `json:"admins"`
	Applicant          string   `json:"applicant"`
	Group              string   `json:"group"`
	Time               string   `json:"time"`
	UpdatedAt          string   `json:"updatedAt"`
	ApplicantId        int64    `json:"applicantId"`
	UpdatedBy          string   `json:"updatedBy"`
	InterviewerComment string   `json:"interviewerComment"`
	ReceiverComment    string   `json:"receiverComment"`
	GroupLabel         string   `json:"groupLabel"`
}

type GetRoomResp struct {
	Rooms      []Room `json:"rooms"`
	MyRoomName string `json:"myRoomName"`
	MyRoomId   int64  `json:"myRoomId"`
}

type SetMyRoomReq struct {
	RoomId int64 `form:"roomId"`
}

type PushableApplicant struct {
	ApplicantId     int64  `json:"applicantId"`
	Name            string `json:"name"`
	Group           string `json:"group"`
	Time            string `json:"time"`
	ApplicantTimeId int64  `json:"applicantTimeId"`
}

type GetPushableListResp struct {
	Applicants []PushableApplicant `json:"applicants"`
}

type PushApplicantReq struct {
	RoomId          int64 `form:"roomId"`
	ApplicantId     int64 `form:"applicantId"`
	ApplicantTimeId int64 `form:"applicantTimeId"`
}

type SetRoomStatusReq struct {
	RoomId int64 `form:"roomId"`
	Status int32 `form:"status"`
}

type SetRoomCommentReq struct {
	RoomId  int64  `path:"roomId"`
	Type    int32  `json:"type"`
	Comment string `json:"comment"`
}

type SetRoomGroupReq struct {
	RoomId int64  `form:"roomId"`
	Group  string `form:"group,optional"`
}
