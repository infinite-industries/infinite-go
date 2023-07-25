package infinite

// VersionService queries the version of the II API.
type VersionService service

// verify license for an organization
func (v *VersionService) Get() (Version, error) {

	var version Version
	errorResponse := new(ErrorResponse)

	resp, err := v.sling.Receive(&version, errorResponse)

	return version, newError(resp, err, errorResponse)

}
