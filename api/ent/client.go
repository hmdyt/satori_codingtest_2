// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/hmdyt/satori_codingtest-2/ent/migrate"

	"github.com/hmdyt/satori_codingtest-2/ent/mesuringpoint"
	"github.com/hmdyt/satori_codingtest-2/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// MesuringPoint is the client for interacting with the MesuringPoint builders.
	MesuringPoint *MesuringPointClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.MesuringPoint = NewMesuringPointClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		MesuringPoint: NewMesuringPointClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		MesuringPoint: NewMesuringPointClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		MesuringPoint.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.MesuringPoint.Use(hooks...)
	c.User.Use(hooks...)
}

// MesuringPointClient is a client for the MesuringPoint schema.
type MesuringPointClient struct {
	config
}

// NewMesuringPointClient returns a client for the MesuringPoint from the given config.
func NewMesuringPointClient(c config) *MesuringPointClient {
	return &MesuringPointClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `mesuringpoint.Hooks(f(g(h())))`.
func (c *MesuringPointClient) Use(hooks ...Hook) {
	c.hooks.MesuringPoint = append(c.hooks.MesuringPoint, hooks...)
}

// Create returns a builder for creating a MesuringPoint entity.
func (c *MesuringPointClient) Create() *MesuringPointCreate {
	mutation := newMesuringPointMutation(c.config, OpCreate)
	return &MesuringPointCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MesuringPoint entities.
func (c *MesuringPointClient) CreateBulk(builders ...*MesuringPointCreate) *MesuringPointCreateBulk {
	return &MesuringPointCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MesuringPoint.
func (c *MesuringPointClient) Update() *MesuringPointUpdate {
	mutation := newMesuringPointMutation(c.config, OpUpdate)
	return &MesuringPointUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MesuringPointClient) UpdateOne(mp *MesuringPoint) *MesuringPointUpdateOne {
	mutation := newMesuringPointMutation(c.config, OpUpdateOne, withMesuringPoint(mp))
	return &MesuringPointUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MesuringPointClient) UpdateOneID(id int) *MesuringPointUpdateOne {
	mutation := newMesuringPointMutation(c.config, OpUpdateOne, withMesuringPointID(id))
	return &MesuringPointUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MesuringPoint.
func (c *MesuringPointClient) Delete() *MesuringPointDelete {
	mutation := newMesuringPointMutation(c.config, OpDelete)
	return &MesuringPointDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MesuringPointClient) DeleteOne(mp *MesuringPoint) *MesuringPointDeleteOne {
	return c.DeleteOneID(mp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MesuringPointClient) DeleteOneID(id int) *MesuringPointDeleteOne {
	builder := c.Delete().Where(mesuringpoint.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MesuringPointDeleteOne{builder}
}

// Query returns a query builder for MesuringPoint.
func (c *MesuringPointClient) Query() *MesuringPointQuery {
	return &MesuringPointQuery{
		config: c.config,
	}
}

// Get returns a MesuringPoint entity by its id.
func (c *MesuringPointClient) Get(ctx context.Context, id int) (*MesuringPoint, error) {
	return c.Query().Where(mesuringpoint.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MesuringPointClient) GetX(ctx context.Context, id int) *MesuringPoint {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MesuringPointClient) Hooks() []Hook {
	return c.hooks.MesuringPoint
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
