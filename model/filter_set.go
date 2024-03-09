package model

// FilterSet holds information about the filter
type FilterSet struct {
	ID                       string   `json:"id"`
	Name                     string   `json:"name"`
	OrganisationID           string   `json:"organisation_id"`
	WorkspaceID              string   `json:"workspace_id"`
	Version                  int      `json:"version"`
	IsDeleted                bool     `json:"is_deleted"`
	IsLocked                 bool     `json:"is_locked"`
	EligibleParticipantCount int      `json:"eligible_participant_count"`
	Filters                  []Filter `json:"filters"`
}

// Filter holds information about the filter that makes up a filter set
type Filter struct {
	FilterID       string   `json:"filter_id"`
	SelectedValues []string `json:"selected_values"`
}
