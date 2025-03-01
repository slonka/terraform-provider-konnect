// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Item struct {
	// The custom plugin schema; `jq -Rs '.' schema.lua`.
	LuaSchema *string `json:"lua_schema,omitempty"`
	// The custom plugin name determined by the custom plugin schema.
	Name *string `json:"name,omitempty"`
	// An ISO-8604 timestamp representation of custom plugin schema creation date.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// An ISO-8604 timestamp representation of custom plugin schema update date.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o *Item) GetLuaSchema() *string {
	if o == nil {
		return nil
	}
	return o.LuaSchema
}

func (o *Item) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Item) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Item) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// PluginSchemas - A response for a single custom plugin schema.
type PluginSchemas struct {
	Item *Item `json:"item,omitempty"`
}

func (o *PluginSchemas) GetItem() *Item {
	if o == nil {
		return nil
	}
	return o.Item
}
