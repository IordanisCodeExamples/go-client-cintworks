# go-client-cintworks

## Getting started

### how to use it

```
cintAPI, err := client.New("https://api.cintworks.net", "some-api-key", http.DefaultClient)
if err != nil {
	//handle error
}

resp, err := s.CintAPI.GetFeasibilityEstimates(
	client.SurveySettingsRequest{
		Limit:             limit,
		LengthOfInterview: lengthOfInterview,
		FieldPeriod:       period,
		IncidenceRate:     DefaultIncidenceRate,
		CountryId:         DefaultCountryID,
		QuotaGroups: []client.QuotaGroupRequest{
			{
				Quotas: []client.QuotaRequest{
					{
						Limit: limit,
						TargetGroup: client.TargetGroupRequest{
							MaxAge: DefaultMaxAge,
							MinAge: DefaultMinAge,
						},
					},
				},
				Limit: limit,
			},
		},
	},
)


fmt.Println(resp)
```