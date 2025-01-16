package processors

import (
	"github.com/cx-rotems/StremSimple/types"
)

type ETLProcess interface {
    Process(job types.Job) (types.Job, error) 
}