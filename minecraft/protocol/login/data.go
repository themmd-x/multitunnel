package login

type IdentityData struct {
        XUID string
        Identity string `json:"identity"`
        DisplayName string `json:"displayName"`
        TitleID string `json:"titleId,omitempty"`
        PlayFabTitleID string `json:"-"`
        PlayFabID string `json:"-"`
}
