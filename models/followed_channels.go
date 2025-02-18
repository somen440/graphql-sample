// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// FollowedChannel is an object representing the database table.
type FollowedChannel struct {
	FollowedChannelID int       `boil:"followed_channel_id" json:"followed_channel_id" toml:"followed_channel_id" yaml:"followed_channel_id"`
	ListenerID        string    `boil:"listener_id" json:"listener_id" toml:"listener_id" yaml:"listener_id"`
	ChannelID         string    `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	CreatedAt         time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt         time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *followedChannelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L followedChannelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FollowedChannelColumns = struct {
	FollowedChannelID string
	ListenerID        string
	ChannelID         string
	CreatedAt         string
	UpdatedAt         string
}{
	FollowedChannelID: "followed_channel_id",
	ListenerID:        "listener_id",
	ChannelID:         "channel_id",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

var FollowedChannelWhere = struct {
	FollowedChannelID whereHelperint
	ListenerID        whereHelperstring
	ChannelID         whereHelperstring
	CreatedAt         whereHelpertime_Time
	UpdatedAt         whereHelpertime_Time
}{
	FollowedChannelID: whereHelperint{field: "`followed_channels`.`followed_channel_id`"},
	ListenerID:        whereHelperstring{field: "`followed_channels`.`listener_id`"},
	ChannelID:         whereHelperstring{field: "`followed_channels`.`channel_id`"},
	CreatedAt:         whereHelpertime_Time{field: "`followed_channels`.`created_at`"},
	UpdatedAt:         whereHelpertime_Time{field: "`followed_channels`.`updated_at`"},
}

// FollowedChannelRels is where relationship names are stored.
var FollowedChannelRels = struct {
	Channel  string
	Listener string
}{
	Channel:  "Channel",
	Listener: "Listener",
}

// followedChannelR is where relationships are stored.
type followedChannelR struct {
	Channel  *Channel
	Listener *Listener
}

// NewStruct creates a new relationship struct
func (*followedChannelR) NewStruct() *followedChannelR {
	return &followedChannelR{}
}

// followedChannelL is where Load methods for each relationship are stored.
type followedChannelL struct{}

var (
	followedChannelAllColumns            = []string{"followed_channel_id", "listener_id", "channel_id", "created_at", "updated_at"}
	followedChannelColumnsWithoutDefault = []string{"listener_id", "channel_id"}
	followedChannelColumnsWithDefault    = []string{"followed_channel_id", "created_at", "updated_at"}
	followedChannelPrimaryKeyColumns     = []string{"followed_channel_id"}
)

type (
	// FollowedChannelSlice is an alias for a slice of pointers to FollowedChannel.
	// This should generally be used opposed to []FollowedChannel.
	FollowedChannelSlice []*FollowedChannel
	// FollowedChannelHook is the signature for custom FollowedChannel hook methods
	FollowedChannelHook func(context.Context, boil.ContextExecutor, *FollowedChannel) error

	followedChannelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	followedChannelType                 = reflect.TypeOf(&FollowedChannel{})
	followedChannelMapping              = queries.MakeStructMapping(followedChannelType)
	followedChannelPrimaryKeyMapping, _ = queries.BindMapping(followedChannelType, followedChannelMapping, followedChannelPrimaryKeyColumns)
	followedChannelInsertCacheMut       sync.RWMutex
	followedChannelInsertCache          = make(map[string]insertCache)
	followedChannelUpdateCacheMut       sync.RWMutex
	followedChannelUpdateCache          = make(map[string]updateCache)
	followedChannelUpsertCacheMut       sync.RWMutex
	followedChannelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var followedChannelBeforeInsertHooks []FollowedChannelHook
var followedChannelBeforeUpdateHooks []FollowedChannelHook
var followedChannelBeforeDeleteHooks []FollowedChannelHook
var followedChannelBeforeUpsertHooks []FollowedChannelHook

var followedChannelAfterInsertHooks []FollowedChannelHook
var followedChannelAfterSelectHooks []FollowedChannelHook
var followedChannelAfterUpdateHooks []FollowedChannelHook
var followedChannelAfterDeleteHooks []FollowedChannelHook
var followedChannelAfterUpsertHooks []FollowedChannelHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FollowedChannel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FollowedChannel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FollowedChannel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FollowedChannel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FollowedChannel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FollowedChannel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FollowedChannel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FollowedChannel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FollowedChannel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range followedChannelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFollowedChannelHook registers your hook function for all future operations.
func AddFollowedChannelHook(hookPoint boil.HookPoint, followedChannelHook FollowedChannelHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		followedChannelBeforeInsertHooks = append(followedChannelBeforeInsertHooks, followedChannelHook)
	case boil.BeforeUpdateHook:
		followedChannelBeforeUpdateHooks = append(followedChannelBeforeUpdateHooks, followedChannelHook)
	case boil.BeforeDeleteHook:
		followedChannelBeforeDeleteHooks = append(followedChannelBeforeDeleteHooks, followedChannelHook)
	case boil.BeforeUpsertHook:
		followedChannelBeforeUpsertHooks = append(followedChannelBeforeUpsertHooks, followedChannelHook)
	case boil.AfterInsertHook:
		followedChannelAfterInsertHooks = append(followedChannelAfterInsertHooks, followedChannelHook)
	case boil.AfterSelectHook:
		followedChannelAfterSelectHooks = append(followedChannelAfterSelectHooks, followedChannelHook)
	case boil.AfterUpdateHook:
		followedChannelAfterUpdateHooks = append(followedChannelAfterUpdateHooks, followedChannelHook)
	case boil.AfterDeleteHook:
		followedChannelAfterDeleteHooks = append(followedChannelAfterDeleteHooks, followedChannelHook)
	case boil.AfterUpsertHook:
		followedChannelAfterUpsertHooks = append(followedChannelAfterUpsertHooks, followedChannelHook)
	}
}

// One returns a single followedChannel record from the query.
func (q followedChannelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FollowedChannel, error) {
	o := &FollowedChannel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for followed_channels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all FollowedChannel records from the query.
func (q followedChannelQuery) All(ctx context.Context, exec boil.ContextExecutor) (FollowedChannelSlice, error) {
	var o []*FollowedChannel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FollowedChannel slice")
	}

	if len(followedChannelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all FollowedChannel records in the query.
func (q followedChannelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count followed_channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q followedChannelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if followed_channels exists")
	}

	return count > 0, nil
}

// Channel pointed to by the foreign key.
func (o *FollowedChannel) Channel(mods ...qm.QueryMod) channelQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`channel_id` = ?", o.ChannelID),
	}

	queryMods = append(queryMods, mods...)

	query := Channels(queryMods...)
	queries.SetFrom(query.Query, "`channels`")

	return query
}

// Listener pointed to by the foreign key.
func (o *FollowedChannel) Listener(mods ...qm.QueryMod) listenerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`listener_id` = ?", o.ListenerID),
	}

	queryMods = append(queryMods, mods...)

	query := Listeners(queryMods...)
	queries.SetFrom(query.Query, "`listeners`")

	return query
}

// LoadChannel allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (followedChannelL) LoadChannel(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFollowedChannel interface{}, mods queries.Applicator) error {
	var slice []*FollowedChannel
	var object *FollowedChannel

	if singular {
		object = maybeFollowedChannel.(*FollowedChannel)
	} else {
		slice = *maybeFollowedChannel.(*[]*FollowedChannel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &followedChannelR{}
		}
		args = append(args, object.ChannelID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &followedChannelR{}
			}

			for _, a := range args {
				if a == obj.ChannelID {
					continue Outer
				}
			}

			args = append(args, obj.ChannelID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`channels`), qm.WhereIn(`channels.channel_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Channel")
	}

	var resultSlice []*Channel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Channel")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for channels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for channels")
	}

	if len(followedChannelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Channel = foreign
		if foreign.R == nil {
			foreign.R = &channelR{}
		}
		foreign.R.FollowedChannels = append(foreign.R.FollowedChannels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ChannelID == foreign.ChannelID {
				local.R.Channel = foreign
				if foreign.R == nil {
					foreign.R = &channelR{}
				}
				foreign.R.FollowedChannels = append(foreign.R.FollowedChannels, local)
				break
			}
		}
	}

	return nil
}

// LoadListener allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (followedChannelL) LoadListener(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFollowedChannel interface{}, mods queries.Applicator) error {
	var slice []*FollowedChannel
	var object *FollowedChannel

	if singular {
		object = maybeFollowedChannel.(*FollowedChannel)
	} else {
		slice = *maybeFollowedChannel.(*[]*FollowedChannel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &followedChannelR{}
		}
		args = append(args, object.ListenerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &followedChannelR{}
			}

			for _, a := range args {
				if a == obj.ListenerID {
					continue Outer
				}
			}

			args = append(args, obj.ListenerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`listeners`), qm.WhereIn(`listeners.listener_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Listener")
	}

	var resultSlice []*Listener
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Listener")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for listeners")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for listeners")
	}

	if len(followedChannelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Listener = foreign
		if foreign.R == nil {
			foreign.R = &listenerR{}
		}
		foreign.R.FollowedChannels = append(foreign.R.FollowedChannels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ListenerID == foreign.ListenerID {
				local.R.Listener = foreign
				if foreign.R == nil {
					foreign.R = &listenerR{}
				}
				foreign.R.FollowedChannels = append(foreign.R.FollowedChannels, local)
				break
			}
		}
	}

	return nil
}

// SetChannel of the followedChannel to the related item.
// Sets o.R.Channel to related.
// Adds o to related.R.FollowedChannels.
func (o *FollowedChannel) SetChannel(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Channel) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `followed_channels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"channel_id"}),
		strmangle.WhereClause("`", "`", 0, followedChannelPrimaryKeyColumns),
	)
	values := []interface{}{related.ChannelID, o.FollowedChannelID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ChannelID = related.ChannelID
	if o.R == nil {
		o.R = &followedChannelR{
			Channel: related,
		}
	} else {
		o.R.Channel = related
	}

	if related.R == nil {
		related.R = &channelR{
			FollowedChannels: FollowedChannelSlice{o},
		}
	} else {
		related.R.FollowedChannels = append(related.R.FollowedChannels, o)
	}

	return nil
}

// SetListener of the followedChannel to the related item.
// Sets o.R.Listener to related.
// Adds o to related.R.FollowedChannels.
func (o *FollowedChannel) SetListener(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Listener) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `followed_channels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"listener_id"}),
		strmangle.WhereClause("`", "`", 0, followedChannelPrimaryKeyColumns),
	)
	values := []interface{}{related.ListenerID, o.FollowedChannelID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ListenerID = related.ListenerID
	if o.R == nil {
		o.R = &followedChannelR{
			Listener: related,
		}
	} else {
		o.R.Listener = related
	}

	if related.R == nil {
		related.R = &listenerR{
			FollowedChannels: FollowedChannelSlice{o},
		}
	} else {
		related.R.FollowedChannels = append(related.R.FollowedChannels, o)
	}

	return nil
}

// FollowedChannels retrieves all the records using an executor.
func FollowedChannels(mods ...qm.QueryMod) followedChannelQuery {
	mods = append(mods, qm.From("`followed_channels`"))
	return followedChannelQuery{NewQuery(mods...)}
}

// FindFollowedChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFollowedChannel(ctx context.Context, exec boil.ContextExecutor, followedChannelID int, selectCols ...string) (*FollowedChannel, error) {
	followedChannelObj := &FollowedChannel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `followed_channels` where `followed_channel_id`=?", sel,
	)

	q := queries.Raw(query, followedChannelID)

	err := q.Bind(ctx, exec, followedChannelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from followed_channels")
	}

	return followedChannelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *FollowedChannel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no followed_channels provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(followedChannelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	followedChannelInsertCacheMut.RLock()
	cache, cached := followedChannelInsertCache[key]
	followedChannelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			followedChannelAllColumns,
			followedChannelColumnsWithDefault,
			followedChannelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(followedChannelType, followedChannelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(followedChannelType, followedChannelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `followed_channels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `followed_channels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `followed_channels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, followedChannelPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into followed_channels")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.FollowedChannelID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == followedChannelMapping["followed_channel_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.FollowedChannelID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for followed_channels")
	}

CacheNoHooks:
	if !cached {
		followedChannelInsertCacheMut.Lock()
		followedChannelInsertCache[key] = cache
		followedChannelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the FollowedChannel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *FollowedChannel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	followedChannelUpdateCacheMut.RLock()
	cache, cached := followedChannelUpdateCache[key]
	followedChannelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			followedChannelAllColumns,
			followedChannelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update followed_channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `followed_channels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, followedChannelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(followedChannelType, followedChannelMapping, append(wl, followedChannelPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update followed_channels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for followed_channels")
	}

	if !cached {
		followedChannelUpdateCacheMut.Lock()
		followedChannelUpdateCache[key] = cache
		followedChannelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q followedChannelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for followed_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for followed_channels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FollowedChannelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), followedChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `followed_channels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, followedChannelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in followedChannel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all followedChannel")
	}
	return rowsAff, nil
}

var mySQLFollowedChannelUniqueColumns = []string{
	"followed_channel_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *FollowedChannel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no followed_channels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(followedChannelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFollowedChannelUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	followedChannelUpsertCacheMut.RLock()
	cache, cached := followedChannelUpsertCache[key]
	followedChannelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			followedChannelAllColumns,
			followedChannelColumnsWithDefault,
			followedChannelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			followedChannelAllColumns,
			followedChannelPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert followed_channels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "followed_channels", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `followed_channels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(followedChannelType, followedChannelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(followedChannelType, followedChannelMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for followed_channels")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.FollowedChannelID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == followedChannelMapping["followed_channel_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(followedChannelType, followedChannelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for followed_channels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for followed_channels")
	}

CacheNoHooks:
	if !cached {
		followedChannelUpsertCacheMut.Lock()
		followedChannelUpsertCache[key] = cache
		followedChannelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single FollowedChannel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FollowedChannel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FollowedChannel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), followedChannelPrimaryKeyMapping)
	sql := "DELETE FROM `followed_channels` WHERE `followed_channel_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from followed_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for followed_channels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q followedChannelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no followedChannelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from followed_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for followed_channels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FollowedChannelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(followedChannelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), followedChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `followed_channels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, followedChannelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from followedChannel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for followed_channels")
	}

	if len(followedChannelAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FollowedChannel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFollowedChannel(ctx, exec, o.FollowedChannelID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FollowedChannelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FollowedChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), followedChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `followed_channels`.* FROM `followed_channels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, followedChannelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FollowedChannelSlice")
	}

	*o = slice

	return nil
}

// FollowedChannelExists checks if the FollowedChannel row exists.
func FollowedChannelExists(ctx context.Context, exec boil.ContextExecutor, followedChannelID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `followed_channels` where `followed_channel_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, followedChannelID)
	}
	row := exec.QueryRowContext(ctx, sql, followedChannelID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if followed_channels exists")
	}

	return exists, nil
}
