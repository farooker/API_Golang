package model
type Profiles struct {
    ID              interface{}     `json:"_id,omitempty" bson:"_id,omitempty"`
    Name            string          `json:"name"`
    ProfileImgURL   string          `json:"profileImgURL "`
    Team            string          `json:"team"`
    Region          string          `json:"region "`
    ServiceInfo     []*ServiceInfo  `json:"serviceInfo"`
}

type ServiceInfo struct {
    Name        string      `json:"name"`
    Id          string      `json:"id"`
    OrgScopeId  string      `json:"orgScopeId"`
    Timestamp   int64       `json:"timestamp"`
}