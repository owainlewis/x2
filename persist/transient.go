package persist

type Transient struct {
	Store map[string]string
}

func (transient Transient) Set(k string, v string) {
	if transient.Store != nil {
		transient.Store[k] = v
	}
}

func (transient Transient) Get(k string) string {
	return transient.Store[k]
}
