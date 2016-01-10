package agent

type Memory interface {
	Set(k string, v string)
	Get(k string) string
}
