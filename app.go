package vivo

import (
	"encoding/json"
)

func (c *VivoClient) GetAppDetail(query *GetAppDetailParams) (*GetAppDetailRes, error) {
	p := c.NewBaseParams()

	data, err := HandleParams(query, p, c.accessSecret)
	if err != nil {
		return nil, err
	}
	reqBody := ParamsToSortQuery(data)
	body, err := c.Post(reqBody)
	if err != nil {
		return nil, err
	}
	getAppDetailRes := &GetAppDetailRes{}
	if err := json.Unmarshal(body, getAppDetailRes); err != nil {
	}
	return getAppDetailRes, err
}

func (c *VivoClient) PublishVersion(params *PublishVersionParams) (*PublishVersionRes, error) {
	p := c.NewBaseParams()

	data, err := HandleParams(params, p, c.accessSecret)
	if err != nil {
		return nil, err
	}
	reqBody := ParamsToSortQuery(data)
	body, err := c.Post(reqBody)
	if err != nil {
		return nil, err
	}
	publishVersionRes := &PublishVersionRes{}
	if err := json.Unmarshal(body, publishVersionRes); err != nil {
	}
	return publishVersionRes, err
}
