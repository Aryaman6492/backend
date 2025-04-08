package v1alpha1

import (
	"github.com/Aryaman6492/backend/pkg/command/types"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	OperatorCommandVersion string = "v1alpha1"
)

var SchemaGroupVersionResource = schema.GroupVersionResource{
	Group:    types.OperatorCommandGroup,
	Version:  OperatorCommandVersion,
	Resource: types.OperatorCommandPlural,
}
