package windows

type windowsPlatform struct {
	path string
}

func (j *windowsPlatform) downloadIdea() bool {
	return false
}

func (j *windowsPlatform) downloadMaven() bool {
	return false
}

func (j *windowsPlatform) downloadGradle() bool {
	return false
}

func (j *windowsPlatform) downloadJava() bool {
	return false
}

func (j *windowsPlatform) installIdea() bool {
	return false
}

func (j *windowsPlatform) installMaven() bool {
	return false
}

func (j *windowsPlatform) installGradle() bool {
	return false
}

func (j *windowsPlatform) installJava() bool {
	return false
}

func (j *windowsPlatform) configEnv() bool {
	return false
}
