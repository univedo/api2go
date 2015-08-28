package api2go


// The CRUD interface MUST be implemented in order to use the api2go api.
type CRUD interface {
	// FindOne returns an object by its ID
	FindOne(ID string, req Request) (Responder, error)

	// Create a new object. Newly created object/struct must be in Responder.
	// Possible status codes are:
	// - 201 Created: Resource was created and needs to be returned
	// - 202 Accepted: Processing is delayed, return nothing
	// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
	//   the server
	Create(obj interface{}, req Request) (Responder, error)

	// Delete an object
	// Possible status codes are:
	// - 200 OK: Update successful, however some field(s) were changed, returns updates source
	// - 202 Accepted: Processing is delayed, return nothing
	// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
	Delete(id string, req Request) (Responder, error)

	// Update an object
	// Possible status codes are:
	// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
	// - 202 Accepted: Processing is delayed, return nothing
	// - 204 No Content: Deletion was successful, return nothing
	Update(obj interface{}, req Request) (Responder, error)
}

type CRUDHooks interface {
	BeforeHandle(*Request) *Responder
	AfterHandle(*Request, Responder)
}

// ContentMarshaler controls how requests from clients are unmarshaled
// and responses from the server are marshaled. The content marshaler
// is in charge of encoding and decoding data to and from a particular
// format (e.g. JSON). The encoding and decoding processes follow the
// rules of the standard encoding/json package.
type ContentMarshaler interface {
	Marshal(i interface{}) ([]byte, error)
	Unmarshal(data []byte, i interface{}) error
	MarshalError(error) (int, []byte)
}

// The PaginatedFindAll interface can be optionally implemented to fetch a subset of all records.
// Pagination query parameters must be used to limit the result. Pagination URLs will automatically
// be generated by the api. You can use a combination of the following 2 query parameters:
// page[number] AND page[size]
// OR page[offset] AND page[limit]
type PaginatedFindAll interface {
	PaginatedFindAll(req Request) (totalCount uint, response Responder, err error)
}

// The FindAll interface can be optionally implemented to fetch all records at once.
type FindAll interface {
	// FindAll returns all objects
	FindAll(req Request) (Responder, error)
}

// The Responder interface is used by all Resource Methods as a container for the Response.
// Metadata is additional Metadata. You can put anything you like into it, see jsonapi spec.
// Result returns the actual payload. For FindOne, put only one entry in it.
// StatusCode sets the http status code.
type Responder interface {
	Metadata() map[string]interface{}
	Result() interface{}
	StatusCode() int
}
