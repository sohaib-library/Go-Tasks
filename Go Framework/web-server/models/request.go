package models

type File struct {
	ID               int  `json:"id"`
	TotalWords       int `json:"total_words"`
	TotalLetters     int `json:"total_letters"`
	TotalSpaces      int `json:"total_spaces"`
	TotalLines       int `json:"total_lines"`
	TotalSpecial int `json:"total_special_char"`
}



type Users struct {
	ID         int `json:"id"`
	NAME      string `json:"name"`
	EMAIL     string `json:"email"`
	PASSWORD string `json:"password"`
}
