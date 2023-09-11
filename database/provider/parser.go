package provider

// Relation  `provider:"aweme_info,local:aweme_id,foreign:ids,exclude:foo,bar,select:foo,bar"`
type Relation struct {
	Alias      string
	Api        string
	LocalKey   string
	ForeignKey string
	Exclude    []string
	Select     []string
}

func NewRelation(r string) *Relation {
	return &Relation{}
}
