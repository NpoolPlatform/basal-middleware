// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/NpoolPlatform/basal-middleware/pkg/db/ent/schema","Package":"github.com/NpoolPlatform/basal-middleware/pkg/db/ent","Schemas":[{"name":"API","config":{"Table":""},"fields":[{"name":"created_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"deleted_at","type":{"Type":16,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":4,"Ident":"uuid.UUID","PkgPath":"github.com/google/uuid","PkgName":"uuid","Nillable":false,"RType":{"Name":"UUID","Ident":"uuid.UUID","Kind":17,"PkgPath":"github.com/google/uuid","Methods":{"ClockSequence":{"In":[],"Out":[{"Name":"int","Ident":"int","Kind":2,"PkgPath":"","Methods":null}]},"Domain":{"In":[],"Out":[{"Name":"Domain","Ident":"uuid.Domain","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"ID":{"In":[],"Out":[{"Name":"uint32","Ident":"uint32","Kind":10,"PkgPath":"","Methods":null}]},"MarshalBinary":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"MarshalText":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"NodeID":{"In":[],"Out":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"Time":{"In":[],"Out":[{"Name":"Time","Ident":"uuid.Time","Kind":6,"PkgPath":"github.com/google/uuid","Methods":null}]},"URN":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalBinary":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalText":{"In":[{"Name":"","Ident":"[]uint8","Kind":23,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Variant":{"In":[],"Out":[{"Name":"Variant","Ident":"uuid.Variant","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]},"Version":{"In":[],"Out":[{"Name":"Version","Ident":"uuid.Version","Kind":8,"PkgPath":"github.com/google/uuid","Methods":null}]}}}},"unique":true,"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"protocol","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"DefaultProtocol","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"service_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"","default_kind":24,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"method","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"DefaultMethod","default_kind":24,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"method_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"","default_kind":24,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"path","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"","default_kind":24,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"exported","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":false,"default_kind":1,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}},{"name":"path_prefix","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":"","default_kind":24,"position":{"Index":7,"MixedIn":false,"MixinIndex":0}},{"name":"domains","type":{"Type":3,"Ident":"[]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":{}}},"optional":true,"default":true,"default_value":[],"default_kind":23,"position":{"Index":8,"MixedIn":false,"MixinIndex":0}},{"name":"depracated","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":false,"default_kind":1,"position":{"Index":9,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntSQL":{"table":"apis"}}}],"Features":["entql","sql/lock","sql/execquery","sql/upsert","privacy","schema/snapshot","sql/modifier"]}`
