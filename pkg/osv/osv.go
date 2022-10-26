package osv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/vuln/osv"
	"zotregistry.io/go-osv/errors"
)

/*
curl -X POST -d \
  '{"version": "2.4.1",
    "package": {"name": "jinja2", "ecosystem": "PyPI"}}' \
  "https://api.osv.dev/v1/query"
*/

const (
	QueryHost  = "api.osv.dev"
	V1QueryURL = "/v1/query"
)

type V1Query struct {
	Commit  string      `json:"commit,omitempty"`
	Version string      `json:"version,omitempty"`
	Package osv.Package `json:"package,omitempty"`
}

type V1Response struct {
	Vulns []osv.Entry `json:"vulns,omitempty"`
}

func lookup(ctx context.Context, osvData *V1Query) (*V1Response, error) {
	body, err := json.Marshal(osvData)
	if err != nil {
		return nil, err
	}

	requestURL := "https://" + QueryHost + V1QueryURL

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.ErrBadParam
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\v", string(data))

	var resp V1Response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func LookupPackage(ctx context.Context, name, version string, ecosystems ...string) ([]osv.Entry, error) {
	// input validation
	if name == "" || version == "" {
		return nil, errors.ErrBadParam
	}

	if len(ecosystems) > 1 {
		return nil, errors.ErrBadParam
	}

	var ecosystem osv.Ecosystem
	if len(ecosystems) == 1 {
		ecosystem = osv.Ecosystem(ecosystems[0])
	}

	osvData := &V1Query{Package: osv.Package{Name: name, Ecosystem: ecosystem}, Version: version}

	resp, err := lookup(ctx, osvData)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%v\n", resp)

	return resp.Vulns, nil
}

func LookupCommitHash(ctx context.Context, commit string) ([]osv.Entry, error) {
	// input validation
	if commit == "" {
		return nil, errors.ErrBadParam
	}

	osvData := &V1Query{Commit: commit}

	resp, err := lookup(ctx, osvData)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%v\n", resp)

	return resp.Vulns, nil
}
