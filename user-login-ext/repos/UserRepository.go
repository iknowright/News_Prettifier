package repos

func UserIsValid(uName, pwd string) bool {
    // DB simulation

    // n := app.news{ID: id}
	// if err := n.getNews(a.DB); err != nil {
	// 	switch err {
	// 	case sql.ErrNoRows:
	// 		respondWithError(w, http.StatusNotFound, "Account not found")
	// 	default:
	// 		respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	}
	// 	return
	// }

    _uName, _pwd, _isValid := "cihanozhan", "1234!*.", false
 
    if uName == _uName && pwd == _pwd {
        _isValid = true
    } else {
        _isValid = false
    }
 
    return _isValid
}