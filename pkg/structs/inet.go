package structs

import (
	"database/sql/driver"
	"fmt"
	"net"
)

type Inet struct {
	net.IP
}

func (i *Inet) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		if i.IP = net.ParseIP(string(v)); i.IP == nil {
			err = fmt.Errorf("invalid value for ip %q", v)
		}
	case string:
		if i.IP = net.ParseIP(v); i.IP == nil {
			err = fmt.Errorf("invalid value for ip %q", v)
		}
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

func (i Inet) Value() (driver.Value, error) {
	return i.IP.String(), nil
}
