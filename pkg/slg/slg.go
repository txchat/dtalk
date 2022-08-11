package slg

type BackendRPCClient interface {
	GroupPermissionVerification([]*UserCondition) (GroupPermission, error)
}

type Client struct {
	client BackendRPCClient
}

func NewClient(client BackendRPCClient) *Client {
	return &Client{client: client}
}

func (c *Client) LoadGroupPermission(ucs []*UserCondition) (GroupPermission, error) {
	return c.client.GroupPermissionVerification(ucs)
}

type UserCondition struct {
	UID        string
	HandleType int32
	Conditions []string
}

type GroupPermission map[string]bool

func (gp GroupPermission) IsPermission(uid string) bool {
	b, ok := gp[uid]
	return b && ok
}
