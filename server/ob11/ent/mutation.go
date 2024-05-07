// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lianhong2758/RosmBot-MUL/server/ob11/ent/chat"
	"github.com/lianhong2758/RosmBot-MUL/server/ob11/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeChat = "Chat"
)

// ChatMutation represents an operation that mutates the Chat nodes in the graph.
type ChatMutation struct {
	config
	op              Op
	typ             string
	id              *int64
	created_at      *time.Time
	updated_at      *time.Time
	group_id        *int64
	addgroup_id     *int64
	user_id         *int64
	adduser_id      *int64
	time            *int64
	addtime         *int64
	message_type    *string
	message_content *string
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*Chat, error)
	predicates      []predicate.Chat
}

var _ ent.Mutation = (*ChatMutation)(nil)

// chatOption allows management of the mutation configuration using functional options.
type chatOption func(*ChatMutation)

// newChatMutation creates new mutation for the Chat entity.
func newChatMutation(c config, op Op, opts ...chatOption) *ChatMutation {
	m := &ChatMutation{
		config:        c,
		op:            op,
		typ:           TypeChat,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withChatID sets the ID field of the mutation.
func withChatID(id int64) chatOption {
	return func(m *ChatMutation) {
		var (
			err   error
			once  sync.Once
			value *Chat
		)
		m.oldValue = func(ctx context.Context) (*Chat, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Chat.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withChat sets the old Chat of the mutation.
func withChat(node *Chat) chatOption {
	return func(m *ChatMutation) {
		m.oldValue = func(context.Context) (*Chat, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ChatMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ChatMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Chat entities.
func (m *ChatMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ChatMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ChatMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Chat.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *ChatMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ChatMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ChatMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ChatMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ChatMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ChatMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetGroupID sets the "group_id" field.
func (m *ChatMutation) SetGroupID(i int64) {
	m.group_id = &i
	m.addgroup_id = nil
}

// GroupID returns the value of the "group_id" field in the mutation.
func (m *ChatMutation) GroupID() (r int64, exists bool) {
	v := m.group_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupID returns the old "group_id" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldGroupID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGroupID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGroupID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGroupID: %w", err)
	}
	return oldValue.GroupID, nil
}

// AddGroupID adds i to the "group_id" field.
func (m *ChatMutation) AddGroupID(i int64) {
	if m.addgroup_id != nil {
		*m.addgroup_id += i
	} else {
		m.addgroup_id = &i
	}
}

// AddedGroupID returns the value that was added to the "group_id" field in this mutation.
func (m *ChatMutation) AddedGroupID() (r int64, exists bool) {
	v := m.addgroup_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetGroupID resets all changes to the "group_id" field.
func (m *ChatMutation) ResetGroupID() {
	m.group_id = nil
	m.addgroup_id = nil
}

// SetUserID sets the "user_id" field.
func (m *ChatMutation) SetUserID(i int64) {
	m.user_id = &i
	m.adduser_id = nil
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *ChatMutation) UserID() (r int64, exists bool) {
	v := m.user_id
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldUserID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// AddUserID adds i to the "user_id" field.
func (m *ChatMutation) AddUserID(i int64) {
	if m.adduser_id != nil {
		*m.adduser_id += i
	} else {
		m.adduser_id = &i
	}
}

// AddedUserID returns the value that was added to the "user_id" field in this mutation.
func (m *ChatMutation) AddedUserID() (r int64, exists bool) {
	v := m.adduser_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetUserID resets all changes to the "user_id" field.
func (m *ChatMutation) ResetUserID() {
	m.user_id = nil
	m.adduser_id = nil
}

// SetTime sets the "time" field.
func (m *ChatMutation) SetTime(i int64) {
	m.time = &i
	m.addtime = nil
}

// Time returns the value of the "time" field in the mutation.
func (m *ChatMutation) Time() (r int64, exists bool) {
	v := m.time
	if v == nil {
		return
	}
	return *v, true
}

// OldTime returns the old "time" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldTime(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTime: %w", err)
	}
	return oldValue.Time, nil
}

// AddTime adds i to the "time" field.
func (m *ChatMutation) AddTime(i int64) {
	if m.addtime != nil {
		*m.addtime += i
	} else {
		m.addtime = &i
	}
}

// AddedTime returns the value that was added to the "time" field in this mutation.
func (m *ChatMutation) AddedTime() (r int64, exists bool) {
	v := m.addtime
	if v == nil {
		return
	}
	return *v, true
}

// ResetTime resets all changes to the "time" field.
func (m *ChatMutation) ResetTime() {
	m.time = nil
	m.addtime = nil
}

// SetMessageType sets the "message_type" field.
func (m *ChatMutation) SetMessageType(s string) {
	m.message_type = &s
}

// MessageType returns the value of the "message_type" field in the mutation.
func (m *ChatMutation) MessageType() (r string, exists bool) {
	v := m.message_type
	if v == nil {
		return
	}
	return *v, true
}

// OldMessageType returns the old "message_type" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldMessageType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessageType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessageType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessageType: %w", err)
	}
	return oldValue.MessageType, nil
}

// ResetMessageType resets all changes to the "message_type" field.
func (m *ChatMutation) ResetMessageType() {
	m.message_type = nil
}

// SetMessageContent sets the "message_content" field.
func (m *ChatMutation) SetMessageContent(s string) {
	m.message_content = &s
}

// MessageContent returns the value of the "message_content" field in the mutation.
func (m *ChatMutation) MessageContent() (r string, exists bool) {
	v := m.message_content
	if v == nil {
		return
	}
	return *v, true
}

// OldMessageContent returns the old "message_content" field's value of the Chat entity.
// If the Chat object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ChatMutation) OldMessageContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessageContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessageContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessageContent: %w", err)
	}
	return oldValue.MessageContent, nil
}

// ResetMessageContent resets all changes to the "message_content" field.
func (m *ChatMutation) ResetMessageContent() {
	m.message_content = nil
}

// Where appends a list predicates to the ChatMutation builder.
func (m *ChatMutation) Where(ps ...predicate.Chat) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ChatMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ChatMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Chat, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ChatMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ChatMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Chat).
func (m *ChatMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ChatMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.created_at != nil {
		fields = append(fields, chat.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, chat.FieldUpdatedAt)
	}
	if m.group_id != nil {
		fields = append(fields, chat.FieldGroupID)
	}
	if m.user_id != nil {
		fields = append(fields, chat.FieldUserID)
	}
	if m.time != nil {
		fields = append(fields, chat.FieldTime)
	}
	if m.message_type != nil {
		fields = append(fields, chat.FieldMessageType)
	}
	if m.message_content != nil {
		fields = append(fields, chat.FieldMessageContent)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ChatMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case chat.FieldCreatedAt:
		return m.CreatedAt()
	case chat.FieldUpdatedAt:
		return m.UpdatedAt()
	case chat.FieldGroupID:
		return m.GroupID()
	case chat.FieldUserID:
		return m.UserID()
	case chat.FieldTime:
		return m.Time()
	case chat.FieldMessageType:
		return m.MessageType()
	case chat.FieldMessageContent:
		return m.MessageContent()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ChatMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case chat.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case chat.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case chat.FieldGroupID:
		return m.OldGroupID(ctx)
	case chat.FieldUserID:
		return m.OldUserID(ctx)
	case chat.FieldTime:
		return m.OldTime(ctx)
	case chat.FieldMessageType:
		return m.OldMessageType(ctx)
	case chat.FieldMessageContent:
		return m.OldMessageContent(ctx)
	}
	return nil, fmt.Errorf("unknown Chat field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMutation) SetField(name string, value ent.Value) error {
	switch name {
	case chat.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case chat.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case chat.FieldGroupID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupID(v)
		return nil
	case chat.FieldUserID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case chat.FieldTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTime(v)
		return nil
	case chat.FieldMessageType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessageType(v)
		return nil
	case chat.FieldMessageContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessageContent(v)
		return nil
	}
	return fmt.Errorf("unknown Chat field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ChatMutation) AddedFields() []string {
	var fields []string
	if m.addgroup_id != nil {
		fields = append(fields, chat.FieldGroupID)
	}
	if m.adduser_id != nil {
		fields = append(fields, chat.FieldUserID)
	}
	if m.addtime != nil {
		fields = append(fields, chat.FieldTime)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ChatMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case chat.FieldGroupID:
		return m.AddedGroupID()
	case chat.FieldUserID:
		return m.AddedUserID()
	case chat.FieldTime:
		return m.AddedTime()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ChatMutation) AddField(name string, value ent.Value) error {
	switch name {
	case chat.FieldGroupID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGroupID(v)
		return nil
	case chat.FieldUserID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUserID(v)
		return nil
	case chat.FieldTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTime(v)
		return nil
	}
	return fmt.Errorf("unknown Chat numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ChatMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ChatMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ChatMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Chat nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ChatMutation) ResetField(name string) error {
	switch name {
	case chat.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case chat.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case chat.FieldGroupID:
		m.ResetGroupID()
		return nil
	case chat.FieldUserID:
		m.ResetUserID()
		return nil
	case chat.FieldTime:
		m.ResetTime()
		return nil
	case chat.FieldMessageType:
		m.ResetMessageType()
		return nil
	case chat.FieldMessageContent:
		m.ResetMessageContent()
		return nil
	}
	return fmt.Errorf("unknown Chat field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ChatMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ChatMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ChatMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ChatMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ChatMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ChatMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ChatMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Chat unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ChatMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Chat edge %s", name)
}
