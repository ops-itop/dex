// Package mock implements connectors which help test various server components.
package learn

import (
	"context"
//	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dexidp/dex/connector"
	"github.com/dexidp/dex/pkg/log"
)

// NewCallbackConnector returns a mock connector which requires no user interaction. It always returns
// the same (fake) identity.
func NewCallbackConnector(logger log.Logger) connector.Connector {
	return &Callback{
		Identity: connector.Identity{
			UserID:        "1000",
			Username:      "lisi",
			Email:         "lisi@qq.com",
			EmailVerified: true,
			Groups:        []string{"authors"},
			ConnectorData: connectorData,
		},
		Logger: logger,
	}
}

var (
	_ connector.CallbackConnector = &Callback{}
)

// Callback is a connector that requires no user interaction and always returns the same identity.
type Callback struct {
	// The returned identity.
	Identity connector.Identity
	Logger   log.Logger
}

// LoginURL returns the URL to redirect the user to login with.
func (m *Callback) LoginURL(s connector.Scopes, callbackURL, state string) (string, error) {
	u, err := url.Parse(callbackURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse callbackURL %q: %v", callbackURL, err)
	}
	v := u.Query()
	v.Set("state", state)
	u.RawQuery = v.Encode()
	return u.String(), nil
}

var connectorData = []byte("foobar")

// HandleCallback parses the request and returns the user's identity
func (m *Callback) HandleCallback(s connector.Scopes, r *http.Request) (connector.Identity, error) {
	return m.Identity, nil
}

// Refresh updates the identity during a refresh token request.
func (m *Callback) Refresh(ctx context.Context, s connector.Scopes, identity connector.Identity) (connector.Identity, error) {
	return m.Identity, nil
}

// CallbackConfig holds the configuration parameters for a connector which requires no interaction.
type CallbackConfig struct{}

// Open returns an authentication strategy which requires no user interaction.
func (c *CallbackConfig) Open(id string, logger log.Logger) (connector.Connector, error) {
	return NewCallbackConnector(logger), nil
}
