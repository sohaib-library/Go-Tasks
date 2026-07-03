package analyzer_impl

func (input *AnalyzerImpl) Create(userID, words, letters, spaces, lines, specialchar int) error {

	_, err := input.db.Exec(

		`INSERT INTO result (
		user_id, 
		total_words,
		total_letters,
		total_spaces,
		total_lines,
		total_special_char
		) 
		
		VALUES ($1, $2, $3, $4, $5, $6)`,


		userID,
		words,
		letters,
		spaces,
		lines,
	    specialchar)
		
	if err != nil {
		return err
	}
	return err
}
