package call

type Status int32

const (
	READY      Status = 0
	PROCESSING Status = 1
	FINISH     Status = 2
)

type RTCType int32

const (
	Undefined RTCType = 0
	Audio     RTCType = 1
	Video     RTCType = 2
)

type Session struct {
	TraceId int64
	RoomId  int64
	RTCType RTCType
	//Deadline 超出后对方未接就结束通话
	Deadline int64
	//Status 0=对方未接通, 1=双方正在通话中, 2=通话结束
	Status     Status
	Invitees   []string
	Caller     string
	Timeout    int64
	CreateTime int64
	//GroupId 0=私聊，^0=群id
	GroupId int64
}

func (s *Session) IsReady() bool {
	return s.Status == READY
}

func (s *Session) IsTimeout(now int64) bool {
	return s.Deadline < now
}

func (s *Session) Finish() {
	s.Status = FINISH
}

func (s *Session) Processing() {
	s.Status = PROCESSING
}

func (s *Session) IsPrivate() bool {
	return s.GroupId == 0
}

func (s *Session) GetPrivateInvitee() string {
	if len(s.Invitees) < 1 {
		return ""
	}
	return s.Invitees[0]
}
