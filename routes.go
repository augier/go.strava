package strava

import "encoding/json"

// RouteDetails contain all of the information about a route
type RouteDetails struct {
	RouteSummary
}

// RouteSummary is returned from the API
type RouteSummary struct {
	Athlete     AthleteMeta `json:"athlete"`
	Description string      `json:"description"`
}

// RouteService wraps a client to access some routes
type RouteService struct {
	client *Client
}

// NewRouteService initliases a RouteService
func NewRouteService(client *Client) *RouteService {
	return &RouteService{client}
}

// RouteGetCall is a helper for requesting a route
type RouteGetCall struct {
	service *RouteService
	id      string
}

// Get sets up a route to get
func (s *RouteService) Get(routeID string) *RouteGetCall {
	return &RouteGetCall{
		service: s,
		id:      routeID,
	}
}

// Do does the request to the API and gets a route
func (c *RouteGetCall) Do() (*RouteDetails, error) {
	data, err := c.service.client.run("GET", "/routes/"+c.id, nil)
	if err != nil {
		return nil, err
	}

	var route RouteDetails
	err = json.Unmarshal(data, &route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}
