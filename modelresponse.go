package client

// SurveyResponse represents the response structure for a survey request .
// It includes various metrics and configurations related to the survey.
// endpoint /ordering/FeasibilityEstimates?version=2
type SurveyResponse struct {
	IncidenceRate                   int                              `json:"incidenceRate"`
	LengthOfInterview               int                              `json:"lengthOfInterview"`
	CountryID                       int                              `json:"countryId"`
	Limit                           int                              `json:"limit"`
	FieldPeriod                     int                              `json:"fieldPeriod"`
	QuotaGroups                     []QuotaGroupResponse             `json:"quotaGroups"`
	DedupeTags                      []string                         `json:"dedupeTags"`
	Feasibility                     int                              `json:"feasibility"`
	FeasibilityCount                int                              `json:"feasibilityCount"`
	ParticipationExclusionOverrides []ParticipationExclusionResponse `json:"participationExclusionOverrides"`
	RespondentIdentifiableInfo      bool                             `json:"respondentIdentifiableInfo"`
	WebcamStudies                   bool                             `json:"webcamStudies"`
	Categories                      []int                            `json:"categories"`
	Errors                          []ErrorResponse                  `json:"errors"`
}

// QuotaGroupResponse represents a response for a specific quota group in the survey.
type QuotaGroupResponse struct {
	Quotas []QuotaResponse `json:"quotas"`
	Limit  int             `json:"limit"`
	Errors []ErrorResponse `json:"errors"`
}

// QuotaResponse contains details about a specific quota within a quota group.
type QuotaResponse struct {
	Limit            int                 `json:"limit"`
	TargetGroup      TargetGroupResponse `json:"targetGroup"`
	PricingStrategy  int                 `json:"pricingStrategy"`
	FeasibilityCount float64             `json:"feasibilityCount"`
	Errors           []ErrorResponse     `json:"errors"`
}

// TargetGroupResponse describes the criteria for targeting respondents within a quota.
type TargetGroupResponse struct {
	MaxAge      int             `json:"MaxAge"`
	MinAge      int             `json:"MinAge"`
	Gender      int             `json:"Gender"`
	RegionIds   []int           `json:"RegionIds"`
	VariableIds []int           `json:"VariableIds"`
	Errors      []ErrorResponse `json:"Errors"`
}

// ParticipationExclusionResponse provides details on participation exclusion rules.
type ParticipationExclusionResponse struct {
	ResponseStatus int `json:"responseStatus"`
	DurationInDays int `json:"durationInDays"`
}

// ErrorResponse captures details of an error encountered in the survey configuration.
type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
