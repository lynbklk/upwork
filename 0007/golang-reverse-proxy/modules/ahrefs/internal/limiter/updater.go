package limiter

var (
	ReportUsageRemarks = []string{
		"target=",
		"keyword=",
		"content-explorer/results",
		"link-intersect/subdomains",
		"domain-comparison?targets[]=",
		"/v4/seBacklinks?input=",
		"/v4/seGetOrganicKeywords?input=",
		"/v4/seRefdomains?input=",
		"/v4/seAnchors?input=",
		"/v4/seInternalBacklinks?input=",
		"/v4/seGetTopPagesHistory?input=",
		"/v4/keKeywordOverview",
		"/v4/ceSearchResultsPublishType?input=",
	}
	ReportUsageExceptRemarks = []string{
		"/site-explorer/export/v2/",
		"/batch-analysis?export=1",
	}
	ExportUsageRemarks = []string{
		"export=",
		"Export",
		"/export",
	}
)
