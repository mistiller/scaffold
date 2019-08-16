package news

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Source struct {
	ID string `json:"id"`
	Name string `json:"name"`
}
type Headline struct {
	Source Source `json:"source"`
	Author string `json:"author"`
	Title string `json:"title"`
	Description string `json:"description"`
	URL string `json:"url"`
	URLToImage string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content string `json:"content"`
}
type Response struct {
	Status string `json:"status"`
	TotalResults int `json:"totalResults"`
	Articles []Headline `json:"articles"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type NewsClient struct {
	host string
	apiKey string
	endpoints []string
}

func NewNewsClient(apiKey string) NewsClient{
	return NewsClient {
		host: "https://newsapi.org/v2/",
		apiKey: apiKey,
		endpoints: []string{
			"top-headlines",
			"everything",
			"sources",
		},
	}
}

func (n *NewsClient) GetHeadlines(country string) (headlines []Headline, err error){
	res, err := n.Get(
		"top-headlines",
		map[string]string{
			"country": country,
		},
	)
	if err != nil {
		return headlines, fmt.Errorf("Fetching headlines - %v", err)
	}

	r := new(Response)
	err = json.Unmarshal(res, r)
	if err != nil {
		return headlines, fmt.Errorf("Fetching headlines - %v", err)
	}
	if r.Status == "error"{
		return headlines, fmt.Errorf("Fetching headlines - %s - %s", r.Code, r.Message)
	}

	headlines = r.Articles

	return r.Articles, nil
}

func (n *NewsClient) getPage(params map[string]string, nPage int) (articles []Headline, totalNumber int, err error) {
	params["page"] = fmt.Sprintf("%d", nPage)

	res, err := n.Get(
		"everything",
		params,
	)
	if err != nil {
		return articles, totalNumber, fmt.Errorf("Fetching headlines - %v", err)
	}

	r := new(Response)
	err = json.Unmarshal(res, r)
	if err != nil {
		return articles, totalNumber, fmt.Errorf("Fetching headlines - %v", err)
	}
	if r.Code == "maximumResultsReached"{
		return articles, totalNumber, nil
	}
	if r.Status == "error"{
		return articles, totalNumber, fmt.Errorf("Fetching headlines - %s - %s", r.Code, r.Message)
	}

	articles = r.Articles
	totalNumber = r.TotalResults

	return articles, totalNumber, nil
}

func (n *NewsClient) RunQuery(query Query) (articles []Headline, err error) {
	params := query.ToMap()

	batch, num, err := n.getPage(params, 1)
	if err != nil {
		return articles, fmt.Errorf("Fetching headlines - %v", err)
	}

	articles = make([]Headline, num)
	for i := range batch {
		articles[i] = batch[i]
	}

	pages, offset := num/query.PageSize, 0
	for pageNum := 2; pageNum <= pages; pageNum++{
		offset = query.PageSize * pageNum
		
		batch, _, err = n.getPage(params, pageNum)
		if err != nil {
			return articles, fmt.Errorf("Fetching headlines - %v", err)
		}
		if len(batch) == 0{
			break
		}

		for i := range batch {
			articles[offset+i] = batch[i]
		}
	}
	return articles, nil
}

func (n *NewsClient) Get(endpoint string, params map[string]string) (response []byte, err error) {
	if n.checkInEndpoints(endpoint) == false {
		return response, fmt.Errorf("Endpoint %s not available", endpoint)
	}
	req, err := http.NewRequest("GET", n.host + endpoint, nil)
	if err != nil {
		return response, fmt.Errorf("Get Request: %s - %v", endpoint, err)
	}

	q := req.URL.Query()
	q.Add("apiKey", n.apiKey)
	for k := range params {
		if params[k] == "" {
			continue
		}
		q.Add(k, params[k])
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return response, fmt.Errorf("Get Request: %s - %v", endpoint, err)
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("Get Request: %s - %v", endpoint, err)
	}

	return response, nil
}

func (n *NewsClient) checkInEndpoints(endpoint string) bool {
	for i := range n.endpoints {
		if endpoint == n.endpoints[i]{
			return true
		}
	}
	return false
}


