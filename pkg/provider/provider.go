package provdier

import (
	"github.com/itsyndicate/itsb/pkg/confparse"
	"github.com/itsyndicate/itsb/pkg/target"
)

var TargetsMap = map[string]func(confparse.Target){
	// "s3": target.S3Target
	// "local": target.LocalTarget
}
