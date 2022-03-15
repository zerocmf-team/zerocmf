/**
** @创建时间: 2021/3/14 2:50 下午
** @作者　　: return
** @描述　　:
 */
package cmfWebsocket

type Client struct {
	Conn   *Connection
	Token  string `json:"token"`  // 唯一标识
	Params Params `json:"params"` // 额外参数
}

type Params struct {
	Token string `json:"token"`
	Mid   int    `json:"mid"`
}

var clientsMap = make(map[string]Client, 0)

func (c *Client) SetClient(userId string) {
	clientsMap[userId] = *c
}

func (c *Client) GetClient(userId string) Client {
	return clientsMap[userId]
}

func (c *Client) DelClient(userId string) {
	delete(clientsMap, userId)
}
