package sso

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	d_identity "github.com/aygumov-g/service-users-go/internal/domain/identity"
)

type SSO struct {
	URL    string
	client *http.Client
}

func NewSSO(URL string, timeout time.Duration) *SSO {
	return &SSO{
		URL: URL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (s *SSO) Me(ctx context.Context, token string) (*d_identity.Identity, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		s.URL,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return nil, errors.New("unauthorized")
	}

	var data d_identity.Identity
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
