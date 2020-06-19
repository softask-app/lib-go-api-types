package apitask

import (
	"github.com/softask-app/lib-go-api-types/v1/pkg/apiuser"
	"time"
)

type StepDetails struct {
	Id          uint64
	Task        TaskDetails
	Description string
	Position    uint16
	Creator     apiuser.UserMeta
	Created     time.Time
}
