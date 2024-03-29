package structs

type WebStorm struct {
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
			WindowsJBR8 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windowsJBR8,omitempty"`
			MacJBR8 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"macJBR8,omitempty"`
			LinuxJBR8 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"linuxJBR8,omitempty"`
			MacWithJBR struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"macWithJBR,omitempty"`
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
		NotesLink              *string     `json:"notesLink"`
		LicenseRequired        bool        `json:"licenseRequired"`
		Version                string      `json:"version"`
		MajorVersion           string      `json:"majorVersion"`
		Build                  string      `json:"build"`
		Whatsnew               *string     `json:"whatsnew"`
		UninstallFeedbackLinks interface{} `json:"uninstallFeedbackLinks"`
		PrintableReleaseType   *string     `json:"printableReleaseType"`
	} `json:"releases"`
	Distributions struct {
		MacWithJBR struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macWithJBR"`
		WindowsJBR8 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windowsJBR8"`
		WindowsZipJBR8 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windowsZipJBR8"`
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
		LinuxJBR8 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"linuxJBR8"`
		Mac struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"mac"`
		MacJBR8 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macJBR8"`
		MacM1 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macM1"`
	} `json:"distributions"`
}
