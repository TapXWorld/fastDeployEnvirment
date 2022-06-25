package windows

type WindowsPlatform struct {
	path string
}

func (j *WindowsPlatform) DownloadIdea() bool {
	return false
}

func (j *WindowsPlatform) DownloadMaven() bool {
	return false
}

func (j *WindowsPlatform) DownloadGradle() bool {
	return false
}

func (j *WindowsPlatform) DownloadJava() bool {
	return false
}

func (j *WindowsPlatform) InstallIdea() bool {
	return false
}

func (j *WindowsPlatform) InstallMaven() bool {
	return false
}

func (j *WindowsPlatform) InstallGradle() bool {
	return false
}

func (j *WindowsPlatform) InstallJava() bool {
	return false
}

func (j *WindowsPlatform) ConfigEnv() bool {
	return false
}
