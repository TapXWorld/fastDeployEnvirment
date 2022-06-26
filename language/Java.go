package language

type Java struct {
	Path string
}

func (j *Java) DownloadIdea() bool {
	return false
}

func (j *Java) DownloadMaven() bool {
	return false
}

func (j *Java) DownloadGradle() bool {
	return false
}

func (j *Java) DownloadJava() bool {
	return false
}

func (j *Java) InstallIdea() bool {
	return false
}

func (j *Java) InstallMaven() bool {
	return false
}

func (j *Java) InstallGradle() bool {
	return false
}

func (j *Java) InstallJava() bool {
	return false
}

func (j *Java) ConfigEnv() bool {
	return false
}
