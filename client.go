func (c *Client) execute(req *Request) (*Response, error) {
	// ... existing setup ...
	for i := 0; i <= c.RetryCount; i++ {
		// Re-apply headers from Request and Client to the http.Request before each attempt
		// to ensure modifications made in RetryCondition hooks are reflected.
		for k, v := range req.Header {
			req.rawRequest.Header[k] = v
		}
		for k, v := range c.Header {
			if _, ok := req.rawRequest.Header[k]; !ok {
				req.rawRequest.Header[k] = v
			}
		}

		resp, err := c.httpClient.Do(req.rawRequest)
		// ... existing retry logic ...
	}
	// ...
}