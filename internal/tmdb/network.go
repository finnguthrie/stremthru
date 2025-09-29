package tmdb

import "strconv"

type Network struct {
	Headquarters  string `json:"headquarters"`
	Homepage      string `json:"homepage"`
	Id            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type FetchNetworkData struct {
	ResponseError
	Network
}

type FetchNetworkParams struct {
	Ctx
	Id int
}

func (c APIClient) FetchNetwork(params *FetchNetworkParams) (APIResponse[Network], error) {
	response := FetchNetworkData{}
	res, err := c.Request("GET", "/3/network/"+strconv.Itoa(params.Id), params, &response)
	return newAPIResponse(res, response.Network), err
}
