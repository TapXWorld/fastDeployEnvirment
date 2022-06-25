package linux

type LinuxPlatform struct {
	path string
}

func (j *LinuxPlatform) DownloadIdea() bool {
	return false
}

func (j *LinuxPlatform) DownloadMaven() bool {
	return false
}

func (j *LinuxPlatform) DownloadGradle() bool {
	return false
}

func (j *LinuxPlatform) DownloadJava() bool {
	return false
}

func (j *LinuxPlatform) InstallIdea() bool {
	return false
}

func (j *LinuxPlatform) InstallMaven() bool {
	return false
}

func (j *LinuxPlatform) InstallGradle() bool {
	return false
}

func (j *LinuxPlatform) InstallJava() bool {
	return false
}

func (j *LinuxPlatform) ConfigEnv() bool {
	return false
}
