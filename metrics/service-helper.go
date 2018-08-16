package metrics

import "github.com/ovh/go-ovh/ovh"

// Tokens return all tokens linked to the Metrics service
func (s *Service) Tokens(c *ovh.Client) (Tokens, error) {
	tokenIDs, err := ListTokens(c, s.Name)
	if err != nil {
		return nil, err
	}

	tokens := make(Tokens, len(tokenIDs))
	for i, tokenID := range tokenIDs {
		t, err := GetToken(c, s.Name, tokenID)
		if err != nil {
			return nil, err
		}
		tokens[i] = t
	}
	return tokens, nil
}

// Len implement sort.Sort()
func (s Services) Len() int {
	return len(s)
}

// Swap implement sort.Sort()
func (s Services) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less implement sort.Sort()
func (s Services) Less(i, j int) bool {
	return s[i].Description < s[j].Description
}
