package v121100

import (
	mtCore "github.com/TheMMD-X/multitunnel/core"
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121100/minecraft"
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v121100/minecraft/protocol/login"
)

func Connect(host string, mtDialer mtCore.Dialer) (*minecraft.Conn, error) {
	dialer := minecraft.Dialer{
		ErrorLog: mtDialer.ErrorLog,
		//HTTPClient: mtDialer.HTTPClient,
		ClientData: login.ClientData{
			CapeData: mtDialer.ClientData.CapeData,
			CapeID: mtDialer.ClientData.CapeID,
			CapeImageHeight: mtDialer.ClientData.CapeImageHeight,
			CapeImageWidth: mtDialer.ClientData.CapeImageWidth,
			CapeOnClassicSkin: mtDialer.ClientData.CapeOnClassicSkin,
			ClientRandomID: mtDialer.ClientData.ClientRandomID,
			CurrentInputMode: mtDialer.ClientData.CurrentInputMode,
			DefaultInputMode: mtDialer.ClientData.DefaultInputMode,
			GameVersion: mtDialer.ClientData.GameVersion,
			GUIScale: mtDialer.ClientData.GUIScale,
			//FilterProfanity: mtDialer.ClientData.FilterProfanity,
			//ClientEditorConnectionIntent: mtDialer.ClientData.ClientEditorConnectionIntent,
			//ClientIsEditorCapable: mtDialer.ClientData.ClientIsEditorCapable,
			LanguageCode: mtDialer.ClientData.LanguageCode,
			PersonaSkin: mtDialer.ClientData.PersonaSkin,
			PlatformOfflineID: mtDialer.ClientData.PlatformOfflineID,
			PlatformOnlineID: mtDialer.ClientData.PlatformOnlineID,
			PlatformUserID: mtDialer.ClientData.PlatformUserID,
			PremiumSkin: mtDialer.ClientData.PremiumSkin,
			SelfSignedID: mtDialer.ClientData.SelfSignedID,
			ServerAddress: mtDialer.ClientData.ServerAddress,
			SkinAnimationData: mtDialer.ClientData.SkinAnimationData,
			SkinData: mtDialer.ClientData.SkinData,
			SkinGeometry: mtDialer.ClientData.SkinGeometry,
			SkinGeometryVersion: mtDialer.ClientData.SkinGeometryVersion,
			SkinID: mtDialer.ClientData.SkinID,
			PlayFabID: mtDialer.ClientData.PlayFabID,
			SkinImageHeight: mtDialer.ClientData.SkinImageHeight,
			SkinImageWidth: mtDialer.ClientData.SkinImageWidth,
			SkinResourcePatch: mtDialer.ClientData.SkinResourcePatch,
			SkinColour: mtDialer.ClientData.SkinColour,
			ArmSize: mtDialer.ClientData.ArmSize,
			ThirdPartyName: mtDialer.ClientData.ThirdPartyName,
			ThirdPartyNameOnly: mtDialer.ClientData.ThirdPartyNameOnly,
			UIProfile: mtDialer.ClientData.UIProfile,
			TrustedSkin: mtDialer.ClientData.TrustedSkin,
			OverrideSkin: mtDialer.ClientData.OverrideSkin,
			CompatibleWithClientSideChunkGen: mtDialer.ClientData.CompatibleWithClientSideChunkGen,
			MaxViewDistance: mtDialer.ClientData.MaxViewDistance,
			MemoryTier: mtDialer.ClientData.MemoryTier,
			PlatformType: mtDialer.ClientData.PlatformType,
			GraphicsMode: mtDialer.ClientData.GraphicsMode,
			//PartyID: mtDialer.ClientData.PartyID,
			//PartyLeader: mtDialer.ClientData.PartyLeader,
			//ProfileHash: mtDialer.ClientData.ProfileHash,
			//Nonce: mtDialer.ClientData.Nonce,
		},
		IdentityData: login.IdentityData{
			XUID: mtDialer.IdentityData.XUID,
			Identity: mtDialer.IdentityData.Identity,
			DisplayName: mtDialer.IdentityData.DisplayName,
			TitleID: mtDialer.IdentityData.TitleID,
			//PlayFabTitleID: mtDialer.IdentityData.PlayFabTitleID,
			//PlayFabID: mtDialer.IdentityData.PlayFabID,
		},
		TokenSource: mtDialer.TokenSource,
		//XBLToken: mtDialer.XBLToken,
		//XBLClient: mtDialer.XBLClient,
		//PlayFabClient: mtDialer.PlayFabClient,
		//PacketFunc: mtDialer.PacketFunc,
		//DownloadResourcePack: mtDialer.DownloadResourcePack,
		//ResourcePackCache: mtDialer.ResourcePackCache,
		DisconnectOnUnknownPackets: mtDialer.DisconnectOnUnknownPackets,
		DisconnectOnInvalidPackets: mtDialer.DisconnectOnInvalidPackets,
		//Protocol: mtDialer.Protocol,
		FlushRate: mtDialer.FlushRate,
		EnableClientCache: mtDialer.EnableClientCache,
		KeepXBLIdentityData: mtDialer.KeepXBLIdentityData,
		EnableLegacyAuth: mtDialer.EnableLegacyAuth,
	}

	conn, err := dialer.Dial("raknet", host)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
