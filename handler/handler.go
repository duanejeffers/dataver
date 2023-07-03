package handler

import "github.com/duanejeffers/dataver/wrapper"

type VersionHandler interface {
	Upgrade() (wrapper.Wrapper, error)
	Downgrade() (wrapper.Wrapper, error)
}
