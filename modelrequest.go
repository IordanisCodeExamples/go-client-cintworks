package client

// SurveySettingsRequest holds the configuration for a survey request.
// It includes settings related to respondent identification, webcam use,
// quota groups, and general survey details.
// endpoint /ordering/FeasibilityEstimates?version=2
type SurveySettingsRequest struct {
	RespondentIdentifiableInfo bool                `json:"respondentIdentifiableInfo"`
	WebcamStudies              bool                `json:"webcamStudies"`
	QuotaGroups                []QuotaGroupRequest `json:"quotaGroups"`
	IncidenceRate              int32               `json:"incidenceRate"`
	LengthOfInterview          int32               `json:"lengthOfInterview"`
	CountryId                  int32               `json:"countryId"`
	Limit                      int32               `json:"limit"`
	FieldPeriod                int32               `json:"fieldPeriod"`
}

// QuotaGroupRequest defines a group of quotas within a survey.
// Each group can target specific respondent segments with its own completion limit.
type QuotaGroupRequest struct {
	Quotas []QuotaRequest `json:"quotas"`
	Limit  int32          `json:"limit"`
}

// QuotaRequest describes a quota within a quota group, targeting a specific segment
// of respondents with defined criteria like demographics or behaviors.
type QuotaRequest struct {
	TargetGroup      TargetGroupRequest `json:"targetGroup"`
	Limit            int32              `json:"limit"`
	PricingStrategy  int32              `json:"pricingStrategy"`
	FeasibilityCount int32              `json:"feasibilityCount"`
}

// TargetGroupRequest specifies the criteria for targeting a group of respondents.
// This includes demographics, regions, and other variables for segmentation.
type TargetGroupRequest struct {
	MaxAge           int32   `json:"maxAge"`
	MinAge           int32   `json:"minAge"`
	Gender           int32   `json:"gender"`
	RegionIds        []int32 `json:"regionIds"`
	VariableIds      []int32 `json:"variableIds"`
	PanelVariableIds []int32 `json:"panelVariableIds"`
	PostalCodes      []int32 `json:"postalCodes"`
	PanelIds         []int32 `json:"panelIds"`
}
