package structs

type PhpStorm struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Releases []struct {
		Date      string `json:"date"`
		Type      string `json:"type"`
		Downloads struct {
			Linux struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"linux,omitempty"`
			Windows struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windows,omitempty"`
			Mac struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"mac,omitempty"`
			MacM1 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"macM1,omitempty"`
			ThirdPartyLibrariesJson struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"thirdPartyLibrariesJson,omitempty"`
			WindowsZip struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windowsZip,omitempty"`
		} `json:"downloads"`
		Patches struct {
			Win []struct {
				FromBuild    string `json:"fromBuild"`
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"win,omitempty"`
			Mac []struct {
				FromBuild    string `json:"fromBuild"`
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"mac,omitempty"`
			Unix []struct {
				FromBuild    string `json:"fromBuild"`
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"unix,omitempty"`
		} `json:"patches"`
		NotesLink              *string `json:"notesLink"`
		LicenseRequired        bool    `json:"licenseRequired"`
		Version                string  `json:"version"`
		MajorVersion           string  `json:"majorVersion"`
		Build                  string  `json:"build"`
		Whatsnew               *string `json:"whatsnew"`
		UninstallFeedbackLinks struct {
			Linux                   string `json:"linux"`
			ThirdPartyLibrariesJson string `json:"thirdPartyLibrariesJson"`
			Windows                 string `json:"windows"`
			WindowsZip              string `json:"windowsZip"`
			Mac                     string `json:"mac"`
			MacM1                   string `json:"macM1"`
		} `json:"uninstallFeedbackLinks"`
		PrintableReleaseType *string `json:"printableReleaseType"`
	} `json:"releases"`
	Distributions struct {
		Linux struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"linux"`
		ThirdPartyLibrariesJson struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"thirdPartyLibrariesJson"`
		Windows struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windows"`
		WindowsZip struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windowsZip"`
		Mac struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"mac"`
		MacM1 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macM1"`
	} `json:"distributions"`
}
