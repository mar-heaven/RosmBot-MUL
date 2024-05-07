// Code generated by ent, DO NOT EDIT.

package chat

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lianhong2758/RosmBot-MUL/server/ob11/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldUpdatedAt, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldGroupID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldUserID, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldTime, v))
}

// MessageType applies equality check predicate on the "message_type" field. It's identical to MessageTypeEQ.
func MessageType(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldMessageType, v))
}

// MessageContent applies equality check predicate on the "message_content" field. It's identical to MessageContentEQ.
func MessageContent(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldMessageContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldUpdatedAt, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int64) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int64) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldGroupID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldUserID, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldTime, v))
}

// MessageTypeEQ applies the EQ predicate on the "message_type" field.
func MessageTypeEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldMessageType, v))
}

// MessageTypeNEQ applies the NEQ predicate on the "message_type" field.
func MessageTypeNEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldMessageType, v))
}

// MessageTypeIn applies the In predicate on the "message_type" field.
func MessageTypeIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldMessageType, vs...))
}

// MessageTypeNotIn applies the NotIn predicate on the "message_type" field.
func MessageTypeNotIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldMessageType, vs...))
}

// MessageTypeGT applies the GT predicate on the "message_type" field.
func MessageTypeGT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldMessageType, v))
}

// MessageTypeGTE applies the GTE predicate on the "message_type" field.
func MessageTypeGTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldMessageType, v))
}

// MessageTypeLT applies the LT predicate on the "message_type" field.
func MessageTypeLT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldMessageType, v))
}

// MessageTypeLTE applies the LTE predicate on the "message_type" field.
func MessageTypeLTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldMessageType, v))
}

// MessageTypeContains applies the Contains predicate on the "message_type" field.
func MessageTypeContains(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContains(FieldMessageType, v))
}

// MessageTypeHasPrefix applies the HasPrefix predicate on the "message_type" field.
func MessageTypeHasPrefix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasPrefix(FieldMessageType, v))
}

// MessageTypeHasSuffix applies the HasSuffix predicate on the "message_type" field.
func MessageTypeHasSuffix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasSuffix(FieldMessageType, v))
}

// MessageTypeEqualFold applies the EqualFold predicate on the "message_type" field.
func MessageTypeEqualFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEqualFold(FieldMessageType, v))
}

// MessageTypeContainsFold applies the ContainsFold predicate on the "message_type" field.
func MessageTypeContainsFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContainsFold(FieldMessageType, v))
}

// MessageContentEQ applies the EQ predicate on the "message_content" field.
func MessageContentEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEQ(FieldMessageContent, v))
}

// MessageContentNEQ applies the NEQ predicate on the "message_content" field.
func MessageContentNEQ(v string) predicate.Chat {
	return predicate.Chat(sql.FieldNEQ(FieldMessageContent, v))
}

// MessageContentIn applies the In predicate on the "message_content" field.
func MessageContentIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldIn(FieldMessageContent, vs...))
}

// MessageContentNotIn applies the NotIn predicate on the "message_content" field.
func MessageContentNotIn(vs ...string) predicate.Chat {
	return predicate.Chat(sql.FieldNotIn(FieldMessageContent, vs...))
}

// MessageContentGT applies the GT predicate on the "message_content" field.
func MessageContentGT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGT(FieldMessageContent, v))
}

// MessageContentGTE applies the GTE predicate on the "message_content" field.
func MessageContentGTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldGTE(FieldMessageContent, v))
}

// MessageContentLT applies the LT predicate on the "message_content" field.
func MessageContentLT(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLT(FieldMessageContent, v))
}

// MessageContentLTE applies the LTE predicate on the "message_content" field.
func MessageContentLTE(v string) predicate.Chat {
	return predicate.Chat(sql.FieldLTE(FieldMessageContent, v))
}

// MessageContentContains applies the Contains predicate on the "message_content" field.
func MessageContentContains(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContains(FieldMessageContent, v))
}

// MessageContentHasPrefix applies the HasPrefix predicate on the "message_content" field.
func MessageContentHasPrefix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasPrefix(FieldMessageContent, v))
}

// MessageContentHasSuffix applies the HasSuffix predicate on the "message_content" field.
func MessageContentHasSuffix(v string) predicate.Chat {
	return predicate.Chat(sql.FieldHasSuffix(FieldMessageContent, v))
}

// MessageContentEqualFold applies the EqualFold predicate on the "message_content" field.
func MessageContentEqualFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldEqualFold(FieldMessageContent, v))
}

// MessageContentContainsFold applies the ContainsFold predicate on the "message_content" field.
func MessageContentContainsFold(v string) predicate.Chat {
	return predicate.Chat(sql.FieldContainsFold(FieldMessageContent, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Chat) predicate.Chat {
	return predicate.Chat(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Chat) predicate.Chat {
	return predicate.Chat(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Chat) predicate.Chat {
	return predicate.Chat(sql.NotPredicates(p))
}
