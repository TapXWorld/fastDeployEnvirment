package structs

type IIU struct {
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
			ThirdPartyLibrariesJson struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"thirdPartyLibrariesJson,omitempty"`
			Windows struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windows,omitempty"`
			WindowsZip struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windowsZip,omitempty"`
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
			LinuxWithoutJBR struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"linuxWithoutJBR,omitempty"`
			WindowsJBR8 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windowsJBR8,omitempty"`
			WindowsZipJBR8 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windowsZipJBR8,omitempty"`
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
			WindowsJBR11 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windowsJBR11,omitempty"`
			MacJBR11 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"macJBR11,omitempty"`
			WindowsZipJBR11 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"windowsZipJBR11,omitempty"`
			LinuxJBR11 struct {
				Link         string `json:"link"`
				Size         int    `json:"size"`
				ChecksumLink string `json:"checksumLink"`
			} `json:"linuxJBR11,omitempty"`
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
		NotesLink              *string `json:"notesLink"`
		LicenseRequired        bool    `json:"licenseRequired"`
		Version                string  `json:"version"`
		MajorVersion           string  `json:"majorVersion"`
		Build                  string  `json:"build"`
		Whatsnew               *string `json:"whatsnew"`
		UninstallFeedbackLinks struct {
			MacWithJBR              string `json:"macWithJBR"`
			SourcesArchive          string `json:"sourcesArchive"`
			WindowsJBR8             string `json:"windowsJBR8"`
			WindowsZipJBR8          string `json:"windowsZipJBR8"`
			LinuxWithoutJBR         string `json:"linuxWithoutJBR"`
			WindowsZipJBR11         string `json:"windowsZipJBR11"`
			ThirdPartyLibrariesJson string `json:"thirdPartyLibrariesJson"`
			Windows                 string `json:"windows"`
			WindowsZip              string `json:"windowsZip"`
			Mac                     string `json:"mac"`
			MacJBR8                 string `json:"macJBR8"`
			MacJBR11                string `json:"macJBR11"`
			WindowsJBR11            string `json:"windowsJBR11"`
			Linux                   string `json:"linux"`
			LinuxJBR11              string `json:"linuxJBR11"`
			LinuxJBR8               string `json:"linuxJBR8"`
			MacM1                   string `json:"macM1"`
		} `json:"uninstallFeedbackLinks"`
		PrintableReleaseType *string `json:"printableReleaseType"`
	} `json:"releases"`
	Distributions struct {
		MacWithJBR struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macWithJBR"`
		SourcesArchive struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"sourcesArchive"`
		WindowsJBR8 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windowsJBR8"`
		WindowsZipJBR8 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windowsZipJBR8"`
		LinuxWithoutJBR struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"linuxWithoutJBR"`
		WindowsZipJBR11 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windowsZipJBR11"`
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
		MacJBR8 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macJBR8"`
		MacJBR11 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macJBR11"`
		WindowsJBR11 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"windowsJBR11"`
		Linux struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"linux"`
		LinuxJBR11 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"linuxJBR11"`
		LinuxJBR8 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"linuxJBR8"`
		MacM1 struct {
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"macM1"`
	} `json:"distributions"`
}
