package miro

type OEmbedServices struct {
	client     *Client
	apiVersion string
	resource   string
}

// Get information to embed a Miro board as a live embed.
// The URL is the resource to return as oEmbed data. Currently, it supports only URLs pointing to Miro boards.
// OEmbed params: OEmbedParams{}
func (o *OEmbedServices) Get(URL string, queryParams ...OEmbedParams) (*OEmbed, error) {
	response := &OEmbed{}

	if url, err := constructURL(o.client.BaseURL, o.apiVersion, o.resource); err != nil {
		return response, err
	} else {
		var searchParams []Parameter
		if len(queryParams) > 0 {
			searchParams = parseQueryTags(queryParams[0])
		}
		searchParams = append(searchParams, Parameter{"url": URL})

		err = o.client.Get(o.client.ctx, url, response, searchParams...)
		return response, err
	}
}
