package resolver

// Null represent Null
func (r *RootResolver) Null() (string, error) {
	resutl := "Just another Null mutation!"
	return resutl, nil
}
