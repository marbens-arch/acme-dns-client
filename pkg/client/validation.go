package client

import (
	"context"
	"fmt"

	"github.com/acme-dns/acme-dns-client/pkg/integration"

	"github.com/nrdcg/goacmedns"
	"github.com/nrdcg/goacmedns/storage"
)

func (c *AcmednsClient) Validation() bool {
	ctx := context.Background()

	token := c.FindValidationToken()
	c.Debug(fmt.Sprintf("Got validation token: %s", token))
	domain := c.FindValidationDomain()
	c.Debug(fmt.Sprintf("Got validation domain: %s", domain))
	if domain == "" || token == "" {
		return false
	}
	acct, err := c.Storage.Fetch(ctx, domain)
	if err != nil && err != storage.ErrDomainNotFound {
		PrintError(fmt.Sprintf("Validation failed: %s", err), 0)
		return false
	} else if err == storage.ErrDomainNotFound {
		PrintError(fmt.Sprintf("Domain %s does not have acme-dns account registered for it. Validation failed.", domain),0)
		return false
	}
	client, err := goacmedns.NewClient(acct.ServerURL, nil)
	err = client.UpdateTXTRecord(ctx, acct, token)
	if err != nil {
		PrintError(fmt.Sprintf("Validation failed: %s", err), 0)
		return false
	}
	return true
}

func (c *AcmednsClient) FindValidationToken() string {
	intgrs := integration.GetIntegrations()
	for _, i := range intgrs {
		token, err := i.FindValidationToken()
		if err != nil {
			c.Debug(fmt.Sprintf("%s", err))
		} else {
			return token
		}
	}
	return ""
}

func (c *AcmednsClient) FindValidationDomain() string {
	intgrs := integration.GetIntegrations()
	for _, i := range intgrs {
		domain, err := i.FindValidationDomain()
		if err != nil {
			c.Debug(fmt.Sprintf("%s", err))
		} else {
			return domain
		}
	}
	return ""
}