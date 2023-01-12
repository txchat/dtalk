package call

type Ticket struct {
	RoomId        int64
	UserSig       string
	PrivateMapKey string
	SDKAppID      int32
}

type TicketCreator func(user string, roomId int64) (*Ticket, error)
