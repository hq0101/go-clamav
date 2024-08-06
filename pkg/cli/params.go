package cli

import "time"

type OutType string

func (o OutType) String() string {
	return string(o)
}

const (
	Json OutType = "json"
	Text OutType = "text"
)

type ClamdParams struct {
	Address     string
	NetworkType string
	ConnTimeout time.Duration
	ReadTimeout time.Duration
	Out         OutType
}

var _ Params = (*ClamdParams)(nil)

func (c *ClamdParams) SetAddress(address string) {
	c.Address = address
}

func (c *ClamdParams) SetNetworkType(netWorkType string) {
	c.NetworkType = netWorkType
}

func (c *ClamdParams) SetConnTimeout(timeout time.Duration) {
	c.ConnTimeout = timeout
}

func (c *ClamdParams) SetReadTimeout(timeout time.Duration) {
	c.ReadTimeout = timeout
}

func (c *ClamdParams) SetOut(out OutType) {
	c.Out = out
}

func (c *ClamdParams) GetAddress() string {
	return c.Address
}

func (c *ClamdParams) GetNetworkType() string {
	return c.NetworkType
}

func (c *ClamdParams) GetConnTimeout() time.Duration {
	return c.ConnTimeout
}

func (c *ClamdParams) GetReadTimeout() time.Duration {
	return c.ReadTimeout
}

func (c *ClamdParams) GetOut() OutType {
	return c.Out
}

type Params interface {
	SetAddress(address string)
	SetNetworkType(netWorkType string)
	SetConnTimeout(timeout time.Duration)
	SetReadTimeout(timeout time.Duration)
	SetOut(out OutType)
	GetAddress() string
	GetNetworkType() string
	GetConnTimeout() time.Duration
	GetReadTimeout() time.Duration
	GetOut() OutType
}
