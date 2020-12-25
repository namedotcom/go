package namecom

import (
	"testing"
)

func TestNameCom_ListDomains(t *testing.T) {
	repo, err := namecom.ListDomains(&ListDomainsRequest{
		PerPage: 100,
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range repo.Domains {
		t.Logf("%v, %v\n", v.DomainName, v.ExpireDate)
	}
}
