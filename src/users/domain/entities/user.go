package entities

type User struct {
    Id              int32  `json:"id"`
    FullName        string `json:"full_name"`
    Email           string `json:"email"`
    PasswordHash    string `json:"password_hash"`
    Gender          string `json:"gender"`
    MatchPreference string `json:"match_preference"`
    City            string `json:"city"`
    State           string `json:"state"`
    Interests       string `json:"interests"`
    StatusMessage   string `json:"status_message"`
    ProfilePicture  string `json:"profile_picture"`
}
func NewUser(fullName, email, passwordHash, gender, matchPreference, city, state, interests, statusMessage, profilePicture string) *User {
	return &User{
		FullName:        fullName,
		Email:           email,
		PasswordHash:    passwordHash,
		Gender:          gender,
		MatchPreference: matchPreference,
		City:            city,
		State:           state,
		Interests:       interests,
		StatusMessage:   statusMessage,
		ProfilePicture:  profilePicture,
	}
}


