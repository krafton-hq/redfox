package domain_helper

import "github.com/krafton-hq/red-fox/apis/idl_common"

type Metadatable interface {
	GetApiVersion() string
	GetKind() string
	GetMetadata() *idl_common.ObjectMeta
}

type MetadatableFactory[T Metadatable] interface {
	Create() T
}
