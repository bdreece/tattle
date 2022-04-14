package tattle

type Logger[F LogFormat] interface {
	Log(format F)
}