package fetch

type Gender struct {
	Gender string `json:"gender"`
}

func FetchGender(name string) (string, error) {
	var g Gender

	if err := getJson("https://api.genderize.io/?name=" + name, &g); err != nil {
		return err.Error(), err
	}

	return g.Gender, nil
}