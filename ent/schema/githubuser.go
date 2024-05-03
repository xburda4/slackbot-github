package schema

import (
	"time"

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
		field.UUID("id", uuid.New()).Default(uuid.New),
		field.String("gh_username"),
		field.String("slack_id").Unique(),
		field.String("gh_access_token").Sensitive(),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Edges of the GithubUser.
func (GithubUser) Edges() []ent.Edge {
	return nil
}
