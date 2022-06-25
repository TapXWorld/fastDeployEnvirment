package linux

type linuxPlatform struct {
	path string
}

func (j *linuxPlatform) downloadIdea() bool {
	return false
}

func (j *linuxPlatform) downloadMaven() bool {
	return false
}

func (j *linuxPlatform) downloadGradle() bool {
	return false
}

func (j *linuxPlatform) downloadJava() bool {
	return false
}

func (j *linuxPlatform) installIdea() bool {
	return false
}

func (j *linuxPlatform) installMaven() bool {
	return false
}

func (j *linuxPlatform) installGradle() bool {
	return false
}

func (j *linuxPlatform) installJava() bool {
	return false
}

func (j *linuxPlatform) configEnv() bool {
	return false
}
