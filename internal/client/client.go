package client

import (
	"github.com/pinpt/esp/internal/common"
	"github.com/pinpt/esp/internal/ssm"
)

type Backend string

type Client interface {
	Save(p common.EspParamInput) common.SaveOutput
	GetOne(p common.GetOneInput) common.EspParam
	GetMany(p common.ListParamInput) []common.EspParam
	Copy(cc common.CopyCommand) common.SaveOutput
	Delete(p common.DeleteInput) string
}

type EspClient struct {
	Backend Backend
	Client Client
}

func New(c EspClient) *EspClient {
	if c.Backend == "ssm" {
		svc := ssm.New()
		svc.Init()
		c.Client = svc
	} else {
		panic("Currently only the ssm backend is valid.")
	}
	return &c
}

// GetParam Queries the ssm param
func (c *EspClient) GetParam(i common.GetOneInput) common.EspParam {
	in := common.GetOneInput{
		Name: i.Name,
		Decrypt: i.Decrypt,
	}
	return c.Client.GetOne(in)
}

// ListParams takes a path and returns all of the parameters under it
func (c *EspClient) ListParams(p common.ListParamInput) []common.EspParam {
	return c.Client.GetMany(p)
}

// Save stores the parameter in the configured backend
func (c *EspClient) Save(p common.EspParamInput) common.SaveOutput {
	return c.Client.Save(p)
}

// Delete removes a parameter from the backend
func (c *EspClient) Delete(p common.DeleteInput) string {
	return c.Client.Delete(p)
}

func (c *EspClient) Copy(cc common.CopyCommand) common.EspParam {
	_ = c.Client.Copy(cc)

	query := common.GetOneInput{
		Name:    cc.Destination,
		Decrypt: true,
	}
	return c.GetParam(query)
}
