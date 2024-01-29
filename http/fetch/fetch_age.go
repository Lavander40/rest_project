package fetch

type Age struct {
	Age int `json:"age"`
}

func FetchAge(name string) (int, error) {
	var a Age

	if err := getJson("https://api.agify.io/?name=" + name, &a); err != nil {
		return -1, err
	}

	return a.Age, nil
}