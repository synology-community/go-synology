package client

// get performs an HTTP request to remote Synology instance.
//
// Returns error in case of any transport errors.
// For API-level errors, check response object.
// func (c APIClient) get(r api.Request, response api.Response) error {
// 	u := c.BaseURL

// 	caller, err := util.GetCaller()
// 	if err != nil {
// 		return err
// 	}
// 	method, err := LookupMethod(caller)
// 	if err != nil {
// 		return err
// 	}
// 	aq, err := query.Values(method.AsApiParams())
// 	if err != nil {
// 		return err
// 	}
// 	dq, err := query.Values(r)
// 	if err != nil {
// 		return err
// 	}
// 	log.Infof("Query from data: %v", dq.Encode())

// 	for k := range dq {
// 		aq.Add(k, dq.Get(k))
// 	}

// 	u.RawQuery = aq.Encode()

// 	//data := api.NewRequest(method.AsApiParams(), r)

// 	// if q, err := query.Values(data); err != nil {
// 	// 	return err
// 	// } else {
// 	// 	u.RawQuery = q.Encode()
// 	// }

// 	log.Infoln(u.String())

// 	// Set context value for caller function
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
// 	if err != nil {
// 		return err
// 	}

// 	return Do(c.httpClient, req, &response)
// }
// _, params, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))

// 	if boundary, ok := params["boundary"]; ok {

// 		mr := multipart.NewReader(resp.Body, boundary)
// 		fr, err := mr.ReadForm(1024 * 1024 * 16) // max 16 MB
// 		if err != nil {
// 			return nil, err
// 		}

// 		for k, v := range fr.Value {
// 			log.Info(k, v)
// 		}

// 		for k, v := range fr.File {
// 			log.Info(k, v)
// 		}

// 		// for part, err := mr.NextPart(); err == nil; part, err = mr.NextPart() {
// 		// 	value, _ := ioutil.ReadAll(part)
// 		// 	log.Printf("Value: %s", value)
// 		// }

// 		return nil, nil
// 	} else {
