package gql

// Article is object model of article
type Article struct {
	ID    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Rooms []Room `json:"rooms"`
}

// Room is object model of room
type Room struct {
	ID        int64  `json:"id" db:"id"`
	ArticleID int64  `json:"articleId" db:"article_id"`
	Index     int64  `json:"index" db:"index"`
	Name      string `json:"name" db:"name"`
	Rent      int64  `json:"rent" db:"rent"`
	Tenant    Tenant `json:"tenant"`
}

// Tenant is object model of Tenant
type Tenant struct {
	ID     int64  `json:"id" db:"id"`
	RoomID int64  `json:"roomId" db:"room_id"`
	Name   string `json:"name" db:"name"`
	Rent   int64  `json:"rent" db:"rent"`
	Since  string `json:"since" db:"since"`
	Until  string `json:"until" db:"until"`
}

// Bill 領収書
type Bill struct {
	ID            int64  `json:"id" db:"id"`
	TenantID      int64  `json:"tenantId" db:"tenant_id"`
	BillingTermID int64  `json:"billingTermId" db:"billing_term_id"`
	Rent          int64  `json:"rent" db:"rent"`
	Since         string `json:"since" db:"since"`
	Until         string `json:"until" db:"until"`
}
