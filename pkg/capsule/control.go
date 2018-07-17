package capsule

import "time"

type Capsule struct {
	LastRequest time.Time
	ContainerID string
	Healthy     bool
	Host        string
	Port        string
	UrlPrefix   string
}

type Ctrl struct {
	HealthScore []int8
	PortPool    []int8
	Plugins     map[string]*Capsule
}

func New(host, urlprefix string) {

}

func (c *Ctrl) Attach(name string, capsule *Capsule) {

}

func (c *Ctrl) Ping(name string) (bool, error) {
	return false, nil
}

func (c *Ctrl) Dettach(name string) {

}

func (c *Ctrl) Request(name string, urlTail string) {

}

func (c *Ctrl) HealthCheck() (bool, []error) {
	errors := make([]error, 0, len(c.Plugins))
	for k := range c.Plugins {
		_, err := c.Ping(k)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return false, errors
	}
	return true, nil
}
