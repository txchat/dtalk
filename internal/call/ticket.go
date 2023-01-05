package call

type Ticket []byte

type TicketCreator func(user string, roomId int64) (Ticket, error)
