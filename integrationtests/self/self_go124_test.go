//go:build !go1.25

package self_test

import tls "github.com/nicorobinreal/utls"

func getCurveID(connState tls.ConnectionState) tls.CurveID {
	return 0
}
