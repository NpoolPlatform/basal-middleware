// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApisColumns holds the columns for the "apis" table.
	ApisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "protocol", Type: field.TypeString, Nullable: true, Default: "DefaultProtocol"},
		{Name: "service_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "method", Type: field.TypeString, Nullable: true, Default: "DefaultMethod"},
		{Name: "method_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "path", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "exported", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "path_prefix", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "domains", Type: field.TypeJSON, Nullable: true},
		{Name: "depracated", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// ApisTable holds the schema information for the "apis" table.
	ApisTable = &schema.Table{
		Name:       "apis",
		Columns:    ApisColumns,
		PrimaryKey: []*schema.Column{ApisColumns[0]},
	}
	// PubsubMessagesColumns holds the columns for the "pubsub_messages" table.
	PubsubMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "message_id", Type: field.TypeString, Nullable: true, Default: "DefaultMsgID"},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultMsgState"},
		{Name: "resp_to_id", Type: field.TypeUUID, Nullable: true},
		{Name: "undo_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arguments", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// PubsubMessagesTable holds the schema information for the "pubsub_messages" table.
	PubsubMessagesTable = &schema.Table{
		Name:       "pubsub_messages",
		Columns:    PubsubMessagesColumns,
		PrimaryKey: []*schema.Column{PubsubMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pubsubmessage_state_resp_to_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[5], PubsubMessagesColumns[6]},
			},
			{
				Name:    "pubsubmessage_state_undo_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[5], PubsubMessagesColumns[7]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApisTable,
		PubsubMessagesTable,
	}
)

func init() {
	ApisTable.Annotation = &entsql.Annotation{
		Table: "apis",
	}
}
