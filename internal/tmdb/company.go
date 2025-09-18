package tmdb

import "strconv"

type Company struct {
	Description   string `json:"description"`
	Headquarters  string `json:"headquarters"`
	Homepage      string `json:"homepage"`
	Id            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
	ParentCompany struct {
		Name     string `json:"name"`
		Id       int    `json:"id"`
		LogoPath string `json:"logo_path"`
	} `json:"parent_company"`
}

type FetchCompanyData struct {
	ResponseError
	Company
}

type FetchCompanyParams struct {
	Ctx
	Id int
}

func (c APIClient) FetchCompany(params *FetchCompanyParams) (APIResponse[Company], error) {
	response := FetchCompanyData{}
	res, err := c.Request("GET", "/3/company/"+strconv.Itoa(params.Id), params, &response)
	return newAPIResponse(res, response.Company), err
}
