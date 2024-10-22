package configs

type jwt struct {
	secertKey        string
	adminKey         string
	apiKey           string
	accessExpiresAt  int //sec
	refreshExpiresAt int //sec
}

func (j *jwt) SecretKey() []byte          { return []byte(j.secertKey) }
func (j *jwt) AdminKey() []byte           { return []byte(j.adminKey) }
func (j *jwt) ApiKey() []byte             { return []byte(j.apiKey) }
func (j *jwt) AccessExpiresAt() int       { return j.accessExpiresAt }
func (j *jwt) RefreshExpiresAt() int      { return j.refreshExpiresAt }
func (j *jwt) SetJwtAccessExpires(t int)  { j.accessExpiresAt = t }
func (j *jwt) SetJwtRefreshExpires(t int) { j.refreshExpiresAt = t }
