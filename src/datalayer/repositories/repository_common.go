package repositories

func getQuerySuffix(onlyFavourites bool) string {
	if onlyFavourites {
		return "AND Favourite = True"
	}
	return ""
}
