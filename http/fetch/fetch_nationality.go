package fetch

import "fmt"

type Nationality struct {
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func FetchNationality(name string) (string, error) {
	var n Nationality

	if err := getJson("https://api.nationalize.io/?name="+name, &n); err != nil {
		return err.Error(), err
	}

	fmt.Println(n)

	// sort.Slice(n, func(i, j int) bool {
	// 	return n[i].Probability < n[j].Probability
	// })

	return n.Country[0].CountryID, nil
}