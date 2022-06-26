package language

type Golang struct {
	Path string
}

func (g *Golang) DownloadIdea() bool {
	return false
}

func (g *Golang) DownloadMaven() bool {
	return false
}

func (g *Golang) DownloadGradle() bool {
	return false
}

func (g *Golang) DownloadJava() bool {
	return false
}

func (g *Golang) InstallIdea() bool {
	return false
}

func (g *Golang) InstallMaven() bool {
	return false
}

func (g *Golang) InstallGradle() bool {
	return false
}

func (g *Golang) InstallJava() bool {
	return false
}

func (g *Golang) ConfigEnv() bool {
	return false
}
