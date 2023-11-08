package xmlsvcwrapper

// Call function execute the request
func (r *Request) Call() (*Response, error) {

	//Build the request
	err := r.build()
	if err != nil {
		return nil, err
	}

	//Send the petition
	resp, err := r.send()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
