package login

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

// checkOfflineUsername is used to check if a username is valid for normal Minecraft client,
// it validates usernames only for unauthenticated clients.
var checkOfflineUsername = regexp.MustCompile(`[ \p{L}]`).MatchString

// checkOnlineUsername is used to check if a username is valid according to the Microsoft specification: "You can
// use up to 15 characters: Aa-Zz, 0-9, and single spaces. It cannot start with a number and cannot start or
// end with a space."
var checkOnlineUsername = regexp.MustCompile("[A-Za-z0-9 ]").MatchString

// checkVersion is used to check if a version is an actual valid version. It must only contain numbers and
// dots.
var checkVersion = regexp.MustCompile("[0-9.]").MatchString

// Validate validates the identity data. It returns an error if any data contained in the IdentityData is
// invalid.
func (data IdentityData) Validate() error {
	if _, err := strconv.ParseInt(data.XUID, 10, 64); err != nil && len(data.XUID) != 0 {
		return fmt.Errorf("XUID must be parseable as an int64, but got %v", data.XUID)
	}
	if id, err := uuid.Parse(data.Identity); err != nil || id == uuid.Nil {
		return fmt.Errorf("UUID must be parseable as a valid UUID, but got %v", data.Identity)
	}
	nameLimit := 15
	if data.XUID == "" {
		// Non-authenticated clients can have up to 16 characters in their name.
		nameLimit = 16
	}
	if len(data.DisplayName) == 0 || len(data.DisplayName) > nameLimit {
		return fmt.Errorf("DisplayName must not be empty or longer than %d characters, but got %v characters", nameLimit, len(data.DisplayName))
	}
	if data.DisplayName[0] == ' ' || data.DisplayName[len(data.DisplayName)-1] == ' ' {
		return fmt.Errorf("DisplayName may not have a space as first/last character, but got %v", data.DisplayName)
	}
	if data.DisplayName[0] >= '0' && data.DisplayName[0] <= '9' {
		return fmt.Errorf("DisplayName may not have a number as first character, but got %v", data.DisplayName)
	}
	if data.XUID != "" {
		if !checkOnlineUsername(data.DisplayName) {
			return fmt.Errorf("DisplayName for authorized client must only contain numbers, Latin letters and spaces, but got %v", data.DisplayName)
		}
	} else {
		if !checkOfflineUsername(data.DisplayName) {
			return fmt.Errorf("DisplayName for unauthorized client must only contain numbers, letters and spaces, but got %v", data.DisplayName)
		}
	}
	// The name is only allowed to have single spaces.
	if strings.Contains(data.DisplayName, "  ") {
		return fmt.Errorf("DisplayName must only have single spaces, but got %v", data.DisplayName)
	}
	return nil
}

// Validate validates the client data. It returns an error if any of the fields checked did not carry a valid
// value. The DeviceOS and AnimatedImageData checks of gophertunnel are skipped, as those fields are not part
// of this version-independent ClientData.
func (data ClientData) Validate() error {
	if !checkVersion(data.GameVersion) {
		return fmt.Errorf("GameVersion must only contain dots and numbers, but got %v", data.GameVersion)
	}
	if _, err := language.Parse(strings.Replace(data.LanguageCode, "_", "-", 1)); err != nil {
		return fmt.Errorf("LanguageCode must be a valid BCP-47 ISO language code, but got %v", data.LanguageCode)
	}
	if _, err := uuid.Parse(data.PlatformOfflineID); err != nil && len(data.PlatformOfflineID) != 0 {
		return fmt.Errorf("PlatformOfflineID must be parseable as a valid UUID or empty, but got %v", data.PlatformOfflineID)
	}
	if _, err := strconv.ParseUint(data.PlatformOnlineID, 10, 64); err != nil && len(data.PlatformOnlineID) != 0 {
		return fmt.Errorf("PlatformOnlineID must be parseable as an int64 or empty, but got %v", data.PlatformOnlineID)
	}
	if _, err := uuid.Parse(data.SelfSignedID); data.SelfSignedID != "" && err != nil {
		return fmt.Errorf("SelfSignedID must be parseable as a valid UUID, but got %v", data.SelfSignedID)
	}
	if strings.Contains(data.ServerAddress, "://") {
		// The server address for NetherNet connections has the following format:
		// https://<host>:<port>:<port>
		ind := strings.LastIndex(data.ServerAddress, ":")
		u, err := url.Parse(data.ServerAddress[:ind])
		if err != nil {
			return fmt.Errorf("ServerAddress must be a URL, but got %v", data.ServerAddress)
		}
		if u.Host == "" || u.Port() == "" || u.Port() != data.ServerAddress[ind+1:] ||
			(u.Scheme != "https" && u.Scheme != "http") {
			return fmt.Errorf("ServerAddress is invalid: %v", data.ServerAddress)
		}
	} else if _, err := net.ResolveUDPAddr("udp", data.ServerAddress); err != nil {
		return fmt.Errorf("ServerAddress must be resolveable as a UDP address, but got %v", data.ServerAddress)
	}
	if err := base64DecLength(data.SkinData, data.SkinImageHeight*data.SkinImageWidth*4); err != nil {
		return fmt.Errorf("SkinData is invalid: %w", err)
	}
	if err := base64DecLength(data.CapeData, data.CapeImageHeight*data.CapeImageWidth*4); err != nil {
		return fmt.Errorf("CapeData is invalid: %w", err)
	}
	for _, s := range data.PlayFabID {
		if (s < '0' || s > '9') && (s < 'a' || s > 'f') {
			return fmt.Errorf("PlayFabID must consist of hex characters, got %v", data.PlayFabID)
		}
	}
	if geomData, err := base64.StdEncoding.DecodeString(data.SkinGeometry); err != nil {
		return fmt.Errorf("SkinGeometry was not a valid base64 string: %w", err)
	} else if len(geomData) != 0 {
		m := make(map[string]any)
		if err := json.Unmarshal(geomData, &m); err != nil {
			return fmt.Errorf("SkinGeometry base64 decoded was not a valid JSON string: %w", err)
		}
	}
	b, err := base64.StdEncoding.DecodeString(data.SkinResourcePatch)
	if err != nil {
		return fmt.Errorf("SkinResourcePatch was not a valid base64 string: %w", err)
	}
	m := make(map[string]any)
	if err := json.Unmarshal(b, &m); err != nil {
		return fmt.Errorf("SkinResourcePatch base64 decoded was not a valid JSON string: %w", err)
	}
	if data.SkinID == "" {
		return fmt.Errorf("SkinID must not be an empty string")
	}
	if data.UIProfile < 0 || data.UIProfile > 2 {
		return fmt.Errorf("UIProfile must be between 0-2, but got %v", data.UIProfile)
	}
	return nil
}

// base64DecLength decodes the base64 data passed and checks if its length is one of the valid lengths
// passed. If either of these checks fails, an error is returned.
func base64DecLength(base64Data string, validLengths ...int) error {
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return fmt.Errorf("decode base64 data: %w", err)
	}
	actualLength := len(data)
	for _, length := range validLengths {
		if length == actualLength {
			return nil
		}
	}
	return fmt.Errorf("invalid size: got %v, expected one of %v", actualLength, validLengths)
}
