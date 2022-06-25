package base

type Java interface {
	downloadIdea(path string) bool
	installIdea(path string) bool

	downloadMaven(path string) bool
	installMaven(path string) bool

	downloadGradle(path string) bool
	installGradle(path string) bool

	downloadJava(path string) bool
	installJava(path string) bool

	configEnv()
}

type Golang interface {
	downloadGoland(path string) bool
	installGoland(path string) bool

	downloadGolang(path string) bool
	installGolang(path string) bool

	configEnv()
}
