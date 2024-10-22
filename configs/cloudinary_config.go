package configs

type cdr struct {
	cloudName         string
	apiKey            string
	apiSecret         string
	cloudinaryBaseUrl string
}

func (c *cdr) CloudName() string { return c.cloudName }
func (c *cdr) ApiKey() string    { return c.apiKey }
func (c *cdr) ApiSecret() string { return c.apiSecret }
func (c *cdr) CloudinaryBaseURL() string {
	return c.cloudinaryBaseUrl
}
