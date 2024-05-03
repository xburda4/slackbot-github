package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GithubUser holds the schema definition for the GithubUser entity.
type GithubUser struct {
	ent.Schema
}

// Fields of the GithubUser.
func (GithubUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("gh_username"),
		field.String("slack_id"),
		field.String("gh_access_token"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the GithubUser.
func (GithubUser) Edges() []ent.Edge {
	return nil
}
