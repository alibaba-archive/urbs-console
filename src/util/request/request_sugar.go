package request

// Get ...
func Get(url string) *Request {
	return New().Get(url)
}

// Post ...
func Post(url string) *Request {
	return New().Post(url)
}

// Put ...
func Put(url string) *Request {
	return New().Put(url)
}

// Delete ...
func Delete(url string) *Request {
	return New().Delete(url)
}
