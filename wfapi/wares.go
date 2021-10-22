package wfapi

import (
	"fmt"

	"github.com/ipld/go-ipld-prime/schema"
)

func init() {
	TypeSystem.Accumulate(schema.SpawnStruct("WareID",
		[]schema.StructField{
			schema.SpawnStructField("packtype", "Packtype", false, false),
			schema.SpawnStructField("hash", "String", false, false),
		},
		schema.SpawnStructRepresentationStringjoin(":")))
	TypeSystem.Accumulate(schema.SpawnString("Packtype"))
	TypeSystem.Accumulate(schema.SpawnMap("Map__String__WareID",
		"String", "WareID", false))
}

type WareID struct {
	Packtype Packtype // f.eks. "tar", "git"
	Hash     string   // what it says on the tin.
}

func (w WareID) String() string {
	return fmt.Sprintf("%s:%s", w.Packtype, w.Hash)
}

type Packtype string

func init() {
	TypeSystem.Accumulate(schema.SpawnString("WarehouseAddr"))
	TypeSystem.Accumulate(schema.SpawnMap("Map__WareID__WarehouseAddr",
		"WareID", "WarehouseAddr", false))
}

// WarehouseAddr is typically parsed as roughly a URL, but we don't deal with that at the API type level.
type WarehouseAddr string

func init() {
	TypeSystem.Accumulate(schema.SpawnMap("FilterMap",
		"String", "String", false)) // FIXME: want support for map representation stringpairs!
}

// Placeholder type.  May need better definition.
type FilterMap struct {
	Keys   []string
	Values map[string]string
}

func init() {
	TypeSystem.Accumulate(schema.SpawnStruct("Mount",
		[]schema.StructField{
			schema.SpawnStructField("mode", "String", false, false), // Ideally an enum, punting on that for now.
			schema.SpawnStructField("hostPath", "String", false, false),
		},
		schema.SpawnStructRepresentationStringjoin(":")))
}

type Mount struct {
	Mode     MountMode
	HostPath string
}

type MountMode string

const (
	MountMode_Readonly  MountMode = "ro"
	MountMode_Readwrite MountMode = "rw"
	MountMode_Overlay   MountMode = "overlay"
)

func init() {
	TypeSystem.Accumulate(schema.SpawnUnion("Ingest",
		[]schema.TypeName{
			"GitIngest",
		},
		schema.SpawnUnionRepresentationStringprefix("", map[string]schema.TypeName{
			"git:": "GitIngest",
		})))
	TypeSystem.Accumulate(schema.SpawnStruct("GitIngest",
		[]schema.StructField{
			schema.SpawnStructField("hostPath", "String", false, false), // Ideally an enum, punting on that for now.
			schema.SpawnStructField("ref", "String", false, false),
		},
		schema.SpawnStructRepresentationStringjoin(":")))
}

type Ingest struct {
	GitIngest *GitIngest
}

type GitIngest struct {
	HostPath string
	Ref      string
}
