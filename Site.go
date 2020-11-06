package msgraph

import (
	"fmt"
	"time"
)

// Site represents a SharePoint site from the Graph API
type Site struct {
	ID                   string      `json:"id"`
	CreatedDateTime      time.Time   `json:"createdDateTime"`
	Description          string      `json:"description"`
	DisplayName          string      `json:"displayName"`
	ETag                 string      `json:"eTag"`
	LastModifiedDateTime time.Time   `json:"lastModifiedDateTime"`
	Name                 string      `json:"name"`
	Root                 interface{} `json:"root"`
	WebURL               string      `json:"webUrl"`
	SharepointIds        struct {
		ListID           string `json:"listId"`
		ListItemID       string `json:"listItemId"`
		ListItemUniqueID string `json:"listItemUniqueId"`
		SiteID           string `json:"siteId"`
		SiteURL          string `json:"siteUrl"`
		TenantID         string `json:"tenantId"`
		WebID            string `json:"webId"`
	} `json:"sharepointIds"`
	SiteCollection struct {
		Hostname         string      `json:"hostname"`
		DataLocationCode string      `json:"dataLocationCode"`
		Root             interface{} `json:"root"`
	} `json:"siteCollection"`

	graphClient *GraphClient
}

func (s *Site) setGraphClient(gC *GraphClient) {
	s.graphClient = gC
}

// ListSiteLists returns a list of Lists in a Site
func (s Site) ListSiteLists() (Lists, error) {
	if s.graphClient == nil {
		return Lists{}, ErrNotGraphClientSourced
	}

	resource := fmt.Sprintf("/sites/%v/lists", s.ID)

	var lists Lists
	return lists, s.graphClient.makeGETAPICall(resource, nil, &lists)
}
