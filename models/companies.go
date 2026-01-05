package models

// Company represents a company record.
type Company struct {
	ActiveDate           string         `json:"activeDate,omitempty"`
	Aliases              []Alias        `json:"aliases,omitempty"`
	Country              string         `json:"country,omitempty"`
	ID                   int64          `json:"id,omitempty"`
	InactiveDate         string         `json:"inactiveDate,omitempty"`
	Name                 string         `json:"name,omitempty"`
	NameTranslations     []string       `json:"nameTranslations,omitempty"`
	OverviewTranslations []string       `json:"overviewTranslations,omitempty"`
	PrimaryCompanyType   *int64         `json:"primaryCompanyType,omitempty"`
	Slug                 string         `json:"slug,omitempty"`
	ParentCompany        *ParentCompany `json:"parentCompany,omitempty"`
	TagOptions           []TagOption    `json:"tagOptions,omitempty"`
}

// ParentCompany represents a parent company record.
type ParentCompany struct {
	ID       *int                 `json:"id,omitempty"`
	Name     string               `json:"name,omitempty"`
	Relation *CompanyRelationship `json:"relation,omitempty"`
}

// CompanyRelationship represents a company relationship.
type CompanyRelationship struct {
	ID       *int   `json:"id,omitempty"`
	TypeName string `json:"typeName,omitempty"`
}

// CompanyType represents a company type record.
type CompanyType struct {
	CompanyTypeId   int    `json:"companyTypeId,omitempty"`
	CompanyTypeName string `json:"companyTypeName,omitempty"`
}

// Companies represents companies grouped by type.
type Companies struct {
	Studio         []Company `json:"studio,omitempty"`
	Network        []Company `json:"network,omitempty"`
	Production     []Company `json:"production,omitempty"`
	Distributor    []Company `json:"distributor,omitempty"`
	SpecialEffects []Company `json:"special_effects,omitempty"`
}

// StudioBaseRecord represents a studio record.
type StudioBaseRecord struct {
	ID           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ParentStudio int    `json:"parentStudio,omitempty"`
}
