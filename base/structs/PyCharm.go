package structs

type PyCharm struct {
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
			WindowsAnaconda struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windowsAnaconda,omitempty"`
			MacAnaconda struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"macAnaconda,omitempty"`
			LinuxAnaconda struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"linuxAnaconda,omitempty"`
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
			WindowsAnaconda         string `json:"windowsAnaconda"`
			MacAnaconda             string `json:"macAnaconda"`
			LinuxAnaconda           string `json:"linuxAnaconda"`
			ThirdPartyLibrariesJson string `json:"thirdPartyLibrariesJson"`
			Windows                 string `json:"windows"`
			Mac                     string `json:"mac"`
			MacM1                   string `json:"macM1"`
			MacM1Anaconda           string `json:"macM1Anaconda"`
		} `json:"uninstallFeedbackLinks"`
		PrintableReleaseType interface{} `json:"printableReleaseType"`
	} `json:"releases"`
	Distributions struct {
		Linux struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"linux"`
		WindowsAnaconda struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windowsAnaconda"`
		MacAnaconda struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macAnaconda"`
		LinuxAnaconda struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"linuxAnaconda"`
		ThirdPartyLibrariesJson struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"thirdPartyLibrariesJson"`
		Windows struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windows"`
		Mac struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"mac"`
		MacM1 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macM1"`
		MacM1Anaconda struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macM1Anaconda"`
	} `json:"distributions"`
}
