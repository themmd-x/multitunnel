package login

type IdentityData struct {
        XUID string
        Identity string `json:"identity"`
        DisplayName string `json:"displayName"`
        TitleID string `json:"titleId,omitempty"`
        PlayFabTitleID string `json:"-"`
        PlayFabID string `json:"-"`
}
type ClientData struct {
        //AnimatedImageData []SkinAnimation
        CapeData string
        CapeID string `json:"CapeId"`
        CapeImageHeight, CapeImageWidth int
        CapeOnClassicSkin bool
        ClientRandomID int64 `json:"ClientRandomId"`
        CurrentInputMode int
        DefaultInputMode int
        //DeviceModel string
        //DeviceOS protocol.DeviceOS
        //DeviceID DeviceID `json:"DeviceId"`
        GameVersion string
        GUIScale int `json:"GuiScale"`
        FilterProfanity bool
        ClientEditorConnectionIntent int
        ClientIsEditorCapable bool
        LanguageCode string
        PersonaSkin bool
        PlatformOfflineID string `json:"PlatformOfflineId"`
        PlatformOnlineID string `json:"PlatformOnlineId"`
        PlatformUserID string `json:"PlatformUserId,omitempty"`
        PremiumSkin bool
        SelfSignedID string `json:"SelfSignedId"`
        ServerAddress string
        SkinAnimationData string
        SkinData string
        SkinGeometry string `json:"SkinGeometryData"`
        SkinGeometryVersion string `json:"SkinGeometryDataEngineVersion"`
        SkinID string `json:"SkinId"`
        PlayFabID string `json:"PlayFabId"`
        SkinImageHeight, SkinImageWidth int
        SkinResourcePatch string
        SkinColour string `json:"SkinColor"`
        ArmSize string
        //PersonaPieces []PersonaPiece
        //PieceTintColours []PersonaPieceTintColour `json:"PieceTintColors"`
        ThirdPartyName string
        ThirdPartyNameOnly *bool `json:"ThirdPartyNameOnly,omitempty"`
        UIProfile int
        TrustedSkin bool
        OverrideSkin bool
        CompatibleWithClientSideChunkGen bool
        MaxViewDistance int
        MemoryTier int
        PlatformType int
        GraphicsMode int
        PartyID string `json:"PartyId"`
        PartyLeader bool `json:"IsPartyLeader"`
        ProfileHash string `json:"ProfileHash"`
        Nonce string `json:",omitempty"`
}
