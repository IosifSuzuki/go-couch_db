package connection

import (
	"bytes"
	"encoding/json"
	"github.com/IosifSuzuki/go-couch_db/model"
	"github.com/IosifSuzuki/go-couch_db/tool"
	"net/http"
)

type Operation interface {
	Ping() (bool, error)
	Len() (int, error)
	Find(model any, selector string) error
	ExecuteView(designDoc, view string, target any, params model.QueryViewParams) error
	CreateView(designDoc, definition string) error
	CreateDB(db string) (bool, error)
}

func (c *connection) Ping() (bool, error) {
	path := c.pathBuilder.Build()
	resp, err := c.httpClient.Execute(http.MethodHead, path, nil)
	if err != nil {
		return false, err
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if !tool.IsSuccessStatusCode(resp.StatusCode) {
		var respErr = model.ErrorResponse{
			Code: resp.StatusCode,
		}
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return false, err
		}
		return false, respErr
	}
	return resp.StatusCode == http.StatusOK, nil
}

func (c *connection) Len() (int, error) {
	path := c.pathBuilder.AddSegment(c.db).Build()
	resp, err := c.httpClient.Execute(http.MethodGet, path, nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if !tool.IsSuccessStatusCode(resp.StatusCode) {
		var respErr = model.ErrorResponse{
			Code: resp.StatusCode,
		}
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return 0, err
		}
		return 0, respErr
	}
	var dbInfo model.DBInfo
	if err := json.NewDecoder(resp.Body).Decode(&dbInfo); err != nil {
		return 0, err
	}
	return dbInfo.DocCount, nil
}

func (c *connection) Find(target any, selector string) error {
	path := c.pathBuilder.AddSegment(c.db).AddSegment("_find").Build()
	reader := bytes.NewReader([]byte(selector))
	resp, err := c.httpClient.Execute(http.MethodPost, path, reader)
	if err != nil {
		return err
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	var findResp model.FindResponse
	if err := json.NewDecoder(resp.Body).Decode(&findResp); err != nil {
		return err
	}
	docsBytes, err := json.Marshal(findResp.Docs)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(bytes.NewReader(docsBytes)).Decode(target); err != nil {
		return err
	}
	return nil
}

func (c *connection) CreateView(designDoc, definition string) error {
	path := c.pathBuilder.AddSegment(c.db).AddSegment("_design").AddSegment(designDoc).Build()
	reader := bytes.NewReader([]byte(definition))
	resp, err := c.httpClient.Execute(http.MethodPut, path, reader)
	if err != nil {
		return err
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if !tool.IsSuccessStatusCode(resp.StatusCode) {
		var respErr = model.ErrorResponse{
			Code: resp.StatusCode,
		}
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return err
		}
		return respErr
	}
	return nil
}

func (c *connection) ExecuteView(designDoc, view string, target any, params model.QueryViewParams) error {
	path := c.pathBuilder.AddSegment(c.db).
		AddSegment("_design").
		AddSegment(designDoc).
		AddSegment("_view").
		AddSegment(view).
		AddQueryModel(params).
		Build()
	resp, err := c.httpClient.Execute(http.MethodGet, path, nil)
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		return err
	}
	if !tool.IsSuccessStatusCode(resp.StatusCode) {
		var respErr = model.ErrorResponse{
			Code: resp.StatusCode,
		}
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return err
		}
		return respErr
	}
	var viewResp model.ViewResponse
	if err := json.NewDecoder(resp.Body).Decode(&viewResp); err != nil {
		return err
	}
	docsBytes, err := json.Marshal(viewResp.Rows)
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(docsBytes)).Decode(&target)
}

func (c *connection) CreateDB(db string) (bool, error) {
	path := c.pathBuilder.AddSegment(c.db).Build()
	resp, err := c.httpClient.Execute(http.MethodPut, path, nil)
	if err != nil {
		return false, err
	}
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	if !tool.IsSuccessStatusCode(resp.StatusCode) {
		var respErr = model.ErrorResponse{
			Code: resp.StatusCode,
		}
		if err := json.NewDecoder(resp.Body).Decode(&respErr); err != nil {
			return false, err
		}
		return false, respErr
	}
	return resp.StatusCode == http.StatusCreated, nil
}
