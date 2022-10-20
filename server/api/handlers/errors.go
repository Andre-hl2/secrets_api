package handlers

type MissingStore struct{}

func (ms MissingStore) Error() string {
	return "Missing store in context object"
}
