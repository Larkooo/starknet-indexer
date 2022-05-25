// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/cartridge-gg/starknet-indexer/ent"
	"github.com/dontpanicdao/caigo/types"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Query() QueryResolver
	Subscription() SubscriptionResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Block struct {
		BlockHash           func(childComplexity int) int
		BlockNumber         func(childComplexity int) int
		ID                  func(childComplexity int) int
		ParentBlockHash     func(childComplexity int) int
		StateRoot           func(childComplexity int) int
		Status              func(childComplexity int) int
		Timestamp           func(childComplexity int) int
		TransactionReceipts func(childComplexity int) int
		Transactions        func(childComplexity int) int
	}

	BlockConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	BlockEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Contract struct {
		CreatedAt    func(childComplexity int) int
		ID           func(childComplexity int) int
		Transactions func(childComplexity int) int
		Type         func(childComplexity int) int
		UpdatedAt    func(childComplexity int) int
	}

	ContractConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	ContractEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	Event struct {
		Data        func(childComplexity int) int
		From        func(childComplexity int) int
		ID          func(childComplexity int) int
		Keys        func(childComplexity int) int
		Transaction func(childComplexity int) int
	}

	EventConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	EventEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	L1Message struct {
		Payload   func(childComplexity int) int
		ToAddress func(childComplexity int) int
	}

	L2Message struct {
		FromAddress func(childComplexity int) int
		Payload     func(childComplexity int) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Blocks       func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.BlockOrder, where *BlockWhereInput) int
		Events       func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *EventWhereInput) int
		Node         func(childComplexity int, id string) int
		Nodes        func(childComplexity int, ids []string) int
		Transactions func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *TransactionWhereInput) int
	}

	Subscription struct {
		WatchEvent func(childComplexity int, address string, keys []*types.Felt) int
	}

	Transaction struct {
		Block              func(childComplexity int) int
		Calldata           func(childComplexity int) int
		ContractAddress    func(childComplexity int) int
		EntryPointSelector func(childComplexity int) int
		Events             func(childComplexity int) int
		ID                 func(childComplexity int) int
		Nonce              func(childComplexity int) int
		Receipt            func(childComplexity int) int
		Signature          func(childComplexity int) int
		TransactionHash    func(childComplexity int) int
	}

	TransactionConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	TransactionEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	TransactionReceipt struct {
		Block           func(childComplexity int) int
		ID              func(childComplexity int) int
		L1OriginMessage func(childComplexity int) int
		MessagesSent    func(childComplexity int) int
		Status          func(childComplexity int) int
		StatusData      func(childComplexity int) int
		Transaction     func(childComplexity int) int
		TransactionHash func(childComplexity int) int
	}

	TransactionReceiptConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	TransactionReceiptEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Block.blockHash":
		if e.complexity.Block.BlockHash == nil {
			break
		}

		return e.complexity.Block.BlockHash(childComplexity), true

	case "Block.blockNumber":
		if e.complexity.Block.BlockNumber == nil {
			break
		}

		return e.complexity.Block.BlockNumber(childComplexity), true

	case "Block.id":
		if e.complexity.Block.ID == nil {
			break
		}

		return e.complexity.Block.ID(childComplexity), true

	case "Block.parentBlockHash":
		if e.complexity.Block.ParentBlockHash == nil {
			break
		}

		return e.complexity.Block.ParentBlockHash(childComplexity), true

	case "Block.stateRoot":
		if e.complexity.Block.StateRoot == nil {
			break
		}

		return e.complexity.Block.StateRoot(childComplexity), true

	case "Block.status":
		if e.complexity.Block.Status == nil {
			break
		}

		return e.complexity.Block.Status(childComplexity), true

	case "Block.timestamp":
		if e.complexity.Block.Timestamp == nil {
			break
		}

		return e.complexity.Block.Timestamp(childComplexity), true

	case "Block.transactionReceipts":
		if e.complexity.Block.TransactionReceipts == nil {
			break
		}

		return e.complexity.Block.TransactionReceipts(childComplexity), true

	case "Block.transactions":
		if e.complexity.Block.Transactions == nil {
			break
		}

		return e.complexity.Block.Transactions(childComplexity), true

	case "BlockConnection.edges":
		if e.complexity.BlockConnection.Edges == nil {
			break
		}

		return e.complexity.BlockConnection.Edges(childComplexity), true

	case "BlockConnection.pageInfo":
		if e.complexity.BlockConnection.PageInfo == nil {
			break
		}

		return e.complexity.BlockConnection.PageInfo(childComplexity), true

	case "BlockConnection.totalCount":
		if e.complexity.BlockConnection.TotalCount == nil {
			break
		}

		return e.complexity.BlockConnection.TotalCount(childComplexity), true

	case "BlockEdge.cursor":
		if e.complexity.BlockEdge.Cursor == nil {
			break
		}

		return e.complexity.BlockEdge.Cursor(childComplexity), true

	case "BlockEdge.node":
		if e.complexity.BlockEdge.Node == nil {
			break
		}

		return e.complexity.BlockEdge.Node(childComplexity), true

	case "Contract.createdAt":
		if e.complexity.Contract.CreatedAt == nil {
			break
		}

		return e.complexity.Contract.CreatedAt(childComplexity), true

	case "Contract.id":
		if e.complexity.Contract.ID == nil {
			break
		}

		return e.complexity.Contract.ID(childComplexity), true

	case "Contract.transactions":
		if e.complexity.Contract.Transactions == nil {
			break
		}

		return e.complexity.Contract.Transactions(childComplexity), true

	case "Contract.type":
		if e.complexity.Contract.Type == nil {
			break
		}

		return e.complexity.Contract.Type(childComplexity), true

	case "Contract.updatedAt":
		if e.complexity.Contract.UpdatedAt == nil {
			break
		}

		return e.complexity.Contract.UpdatedAt(childComplexity), true

	case "ContractConnection.edges":
		if e.complexity.ContractConnection.Edges == nil {
			break
		}

		return e.complexity.ContractConnection.Edges(childComplexity), true

	case "ContractConnection.pageInfo":
		if e.complexity.ContractConnection.PageInfo == nil {
			break
		}

		return e.complexity.ContractConnection.PageInfo(childComplexity), true

	case "ContractConnection.totalCount":
		if e.complexity.ContractConnection.TotalCount == nil {
			break
		}

		return e.complexity.ContractConnection.TotalCount(childComplexity), true

	case "ContractEdge.cursor":
		if e.complexity.ContractEdge.Cursor == nil {
			break
		}

		return e.complexity.ContractEdge.Cursor(childComplexity), true

	case "ContractEdge.node":
		if e.complexity.ContractEdge.Node == nil {
			break
		}

		return e.complexity.ContractEdge.Node(childComplexity), true

	case "Event.data":
		if e.complexity.Event.Data == nil {
			break
		}

		return e.complexity.Event.Data(childComplexity), true

	case "Event.from":
		if e.complexity.Event.From == nil {
			break
		}

		return e.complexity.Event.From(childComplexity), true

	case "Event.id":
		if e.complexity.Event.ID == nil {
			break
		}

		return e.complexity.Event.ID(childComplexity), true

	case "Event.keys":
		if e.complexity.Event.Keys == nil {
			break
		}

		return e.complexity.Event.Keys(childComplexity), true

	case "Event.transaction":
		if e.complexity.Event.Transaction == nil {
			break
		}

		return e.complexity.Event.Transaction(childComplexity), true

	case "EventConnection.edges":
		if e.complexity.EventConnection.Edges == nil {
			break
		}

		return e.complexity.EventConnection.Edges(childComplexity), true

	case "EventConnection.pageInfo":
		if e.complexity.EventConnection.PageInfo == nil {
			break
		}

		return e.complexity.EventConnection.PageInfo(childComplexity), true

	case "EventConnection.totalCount":
		if e.complexity.EventConnection.TotalCount == nil {
			break
		}

		return e.complexity.EventConnection.TotalCount(childComplexity), true

	case "EventEdge.cursor":
		if e.complexity.EventEdge.Cursor == nil {
			break
		}

		return e.complexity.EventEdge.Cursor(childComplexity), true

	case "EventEdge.node":
		if e.complexity.EventEdge.Node == nil {
			break
		}

		return e.complexity.EventEdge.Node(childComplexity), true

	case "L1Message.payload":
		if e.complexity.L1Message.Payload == nil {
			break
		}

		return e.complexity.L1Message.Payload(childComplexity), true

	case "L1Message.toAddress":
		if e.complexity.L1Message.ToAddress == nil {
			break
		}

		return e.complexity.L1Message.ToAddress(childComplexity), true

	case "L2Message.fromAddress":
		if e.complexity.L2Message.FromAddress == nil {
			break
		}

		return e.complexity.L2Message.FromAddress(childComplexity), true

	case "L2Message.payload":
		if e.complexity.L2Message.Payload == nil {
			break
		}

		return e.complexity.L2Message.Payload(childComplexity), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.blocks":
		if e.complexity.Query.Blocks == nil {
			break
		}

		args, err := ec.field_Query_blocks_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Blocks(childComplexity, args["after"].(*ent.Cursor), args["first"].(*int), args["before"].(*ent.Cursor), args["last"].(*int), args["orderBy"].(*ent.BlockOrder), args["where"].(*BlockWhereInput)), true

	case "Query.events":
		if e.complexity.Query.Events == nil {
			break
		}

		args, err := ec.field_Query_events_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Events(childComplexity, args["after"].(*ent.Cursor), args["first"].(*int), args["before"].(*ent.Cursor), args["last"].(*int), args["where"].(*EventWhereInput)), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(string)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]string)), true

	case "Query.transactions":
		if e.complexity.Query.Transactions == nil {
			break
		}

		args, err := ec.field_Query_transactions_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Transactions(childComplexity, args["after"].(*ent.Cursor), args["first"].(*int), args["before"].(*ent.Cursor), args["last"].(*int), args["where"].(*TransactionWhereInput)), true

	case "Subscription.watchEvent":
		if e.complexity.Subscription.WatchEvent == nil {
			break
		}

		args, err := ec.field_Subscription_watchEvent_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Subscription.WatchEvent(childComplexity, args["address"].(string), args["keys"].([]*types.Felt)), true

	case "Transaction.block":
		if e.complexity.Transaction.Block == nil {
			break
		}

		return e.complexity.Transaction.Block(childComplexity), true

	case "Transaction.calldata":
		if e.complexity.Transaction.Calldata == nil {
			break
		}

		return e.complexity.Transaction.Calldata(childComplexity), true

	case "Transaction.contractAddress":
		if e.complexity.Transaction.ContractAddress == nil {
			break
		}

		return e.complexity.Transaction.ContractAddress(childComplexity), true

	case "Transaction.entryPointSelector":
		if e.complexity.Transaction.EntryPointSelector == nil {
			break
		}

		return e.complexity.Transaction.EntryPointSelector(childComplexity), true

	case "Transaction.events":
		if e.complexity.Transaction.Events == nil {
			break
		}

		return e.complexity.Transaction.Events(childComplexity), true

	case "Transaction.id":
		if e.complexity.Transaction.ID == nil {
			break
		}

		return e.complexity.Transaction.ID(childComplexity), true

	case "Transaction.nonce":
		if e.complexity.Transaction.Nonce == nil {
			break
		}

		return e.complexity.Transaction.Nonce(childComplexity), true

	case "Transaction.receipt":
		if e.complexity.Transaction.Receipt == nil {
			break
		}

		return e.complexity.Transaction.Receipt(childComplexity), true

	case "Transaction.signature":
		if e.complexity.Transaction.Signature == nil {
			break
		}

		return e.complexity.Transaction.Signature(childComplexity), true

	case "Transaction.transactionHash":
		if e.complexity.Transaction.TransactionHash == nil {
			break
		}

		return e.complexity.Transaction.TransactionHash(childComplexity), true

	case "TransactionConnection.edges":
		if e.complexity.TransactionConnection.Edges == nil {
			break
		}

		return e.complexity.TransactionConnection.Edges(childComplexity), true

	case "TransactionConnection.pageInfo":
		if e.complexity.TransactionConnection.PageInfo == nil {
			break
		}

		return e.complexity.TransactionConnection.PageInfo(childComplexity), true

	case "TransactionConnection.totalCount":
		if e.complexity.TransactionConnection.TotalCount == nil {
			break
		}

		return e.complexity.TransactionConnection.TotalCount(childComplexity), true

	case "TransactionEdge.cursor":
		if e.complexity.TransactionEdge.Cursor == nil {
			break
		}

		return e.complexity.TransactionEdge.Cursor(childComplexity), true

	case "TransactionEdge.node":
		if e.complexity.TransactionEdge.Node == nil {
			break
		}

		return e.complexity.TransactionEdge.Node(childComplexity), true

	case "TransactionReceipt.block":
		if e.complexity.TransactionReceipt.Block == nil {
			break
		}

		return e.complexity.TransactionReceipt.Block(childComplexity), true

	case "TransactionReceipt.id":
		if e.complexity.TransactionReceipt.ID == nil {
			break
		}

		return e.complexity.TransactionReceipt.ID(childComplexity), true

	case "TransactionReceipt.l1OriginMessage":
		if e.complexity.TransactionReceipt.L1OriginMessage == nil {
			break
		}

		return e.complexity.TransactionReceipt.L1OriginMessage(childComplexity), true

	case "TransactionReceipt.messagesSent":
		if e.complexity.TransactionReceipt.MessagesSent == nil {
			break
		}

		return e.complexity.TransactionReceipt.MessagesSent(childComplexity), true

	case "TransactionReceipt.status":
		if e.complexity.TransactionReceipt.Status == nil {
			break
		}

		return e.complexity.TransactionReceipt.Status(childComplexity), true

	case "TransactionReceipt.statusData":
		if e.complexity.TransactionReceipt.StatusData == nil {
			break
		}

		return e.complexity.TransactionReceipt.StatusData(childComplexity), true

	case "TransactionReceipt.transaction":
		if e.complexity.TransactionReceipt.Transaction == nil {
			break
		}

		return e.complexity.TransactionReceipt.Transaction(childComplexity), true

	case "TransactionReceipt.transactionHash":
		if e.complexity.TransactionReceipt.TransactionHash == nil {
			break
		}

		return e.complexity.TransactionReceipt.TransactionHash(childComplexity), true

	case "TransactionReceiptConnection.edges":
		if e.complexity.TransactionReceiptConnection.Edges == nil {
			break
		}

		return e.complexity.TransactionReceiptConnection.Edges(childComplexity), true

	case "TransactionReceiptConnection.pageInfo":
		if e.complexity.TransactionReceiptConnection.PageInfo == nil {
			break
		}

		return e.complexity.TransactionReceiptConnection.PageInfo(childComplexity), true

	case "TransactionReceiptConnection.totalCount":
		if e.complexity.TransactionReceiptConnection.TotalCount == nil {
			break
		}

		return e.complexity.TransactionReceiptConnection.TotalCount(childComplexity), true

	case "TransactionReceiptEdge.cursor":
		if e.complexity.TransactionReceiptEdge.Cursor == nil {
			break
		}

		return e.complexity.TransactionReceiptEdge.Cursor(childComplexity), true

	case "TransactionReceiptEdge.node":
		if e.complexity.TransactionReceiptEdge.Node == nil {
			break
		}

		return e.complexity.TransactionReceiptEdge.Node(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Subscription:
		next := ec._Subscription(ctx, rc.Operation.SelectionSet)

		var buf bytes.Buffer
		return func(ctx context.Context) *graphql.Response {
			buf.Reset()
			data := next()

			if data == nil {
				return nil
			}
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "entgql.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
type Block implements Node {
  id: ID!
  blockHash: String!
  parentBlockHash: String!
  blockNumber: Long!
  stateRoot: String!
  status: BlockStatus!
  timestamp: Time!
  transactions: [Transaction!]
  transactionReceipts: [TransactionReceipt!]
}
"""A connection to a list of items."""
type BlockConnection {
  """A list of edges."""
  edges: [BlockEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  totalCount: Int!
}
"""An edge in a connection."""
type BlockEdge {
  """The item at the end of the edge."""
  node: Block
  """A cursor for use in pagination."""
  cursor: Cursor!
}
input BlockOrder {
  direction: OrderDirection! = ASC
  field: BlockOrderField!
}
enum BlockOrderField {
  BLOCK_NUMBER
  TIMESTAMP
}
"""BlockStatus is enum for the field status"""
enum BlockStatus @goModel(model: "github.com/cartridge-gg/starknet-indexer/ent/block.Status") {
  ACCEPTED_ON_L1
  ACCEPTED_ON_L2
}
"""
BlockWhereInput is used for filtering Block objects.
Input was generated by ent.
"""
input BlockWhereInput {
  not: BlockWhereInput
  and: [BlockWhereInput!]
  or: [BlockWhereInput!]
  """block_hash field predicates"""
  blockHash: String
  blockHashNEQ: String
  blockHashIn: [String!]
  blockHashNotIn: [String!]
  blockHashGT: String
  blockHashGTE: String
  blockHashLT: String
  blockHashLTE: String
  blockHashContains: String
  blockHashHasPrefix: String
  blockHashHasSuffix: String
  blockHashEqualFold: String
  blockHashContainsFold: String
  """parent_block_hash field predicates"""
  parentBlockHash: String
  parentBlockHashNEQ: String
  parentBlockHashIn: [String!]
  parentBlockHashNotIn: [String!]
  parentBlockHashGT: String
  parentBlockHashGTE: String
  parentBlockHashLT: String
  parentBlockHashLTE: String
  parentBlockHashContains: String
  parentBlockHashHasPrefix: String
  parentBlockHashHasSuffix: String
  parentBlockHashEqualFold: String
  parentBlockHashContainsFold: String
  """block_number field predicates"""
  blockNumber: Long
  blockNumberNEQ: Long
  blockNumberIn: [Long!]
  blockNumberNotIn: [Long!]
  blockNumberGT: Long
  blockNumberGTE: Long
  blockNumberLT: Long
  blockNumberLTE: Long
  """state_root field predicates"""
  stateRoot: String
  stateRootNEQ: String
  stateRootIn: [String!]
  stateRootNotIn: [String!]
  stateRootGT: String
  stateRootGTE: String
  stateRootLT: String
  stateRootLTE: String
  stateRootContains: String
  stateRootHasPrefix: String
  stateRootHasSuffix: String
  stateRootEqualFold: String
  stateRootContainsFold: String
  """status field predicates"""
  status: BlockStatus
  statusNEQ: BlockStatus
  statusIn: [BlockStatus!]
  statusNotIn: [BlockStatus!]
  """timestamp field predicates"""
  timestamp: Time
  timestampNEQ: Time
  timestampIn: [Time!]
  timestampNotIn: [Time!]
  timestampGT: Time
  timestampGTE: Time
  timestampLT: Time
  timestampLTE: Time
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """transactions edge predicates"""
  hasTransactions: Boolean
  hasTransactionsWith: [TransactionWhereInput!]
  """transaction_receipts edge predicates"""
  hasTransactionReceipts: Boolean
  hasTransactionReceiptsWith: [TransactionReceiptWhereInput!]
}
type Contract implements Node {
  id: ID!
  type: Type!
  createdAt: Time!
  updatedAt: Time!
  transactions: [Transaction!]
}
"""A connection to a list of items."""
type ContractConnection {
  """A list of edges."""
  edges: [ContractEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  totalCount: Int!
}
"""An edge in a connection."""
type ContractEdge {
  """The item at the end of the edge."""
  node: Contract
  """A cursor for use in pagination."""
  cursor: Cursor!
}
input ContractOrder {
  direction: OrderDirection! = ASC
  field: ContractOrderField!
}
enum ContractOrderField {
  CREATED_AT
}
"""
ContractWhereInput is used for filtering Contract objects.
Input was generated by ent.
"""
input ContractWhereInput {
  not: ContractWhereInput
  and: [ContractWhereInput!]
  or: [ContractWhereInput!]
  """type field predicates"""
  type: Type
  typeNEQ: Type
  typeIn: [Type!]
  typeNotIn: [Type!]
  """created_at field predicates"""
  createdAt: Time
  createdAtNEQ: Time
  createdAtIn: [Time!]
  createdAtNotIn: [Time!]
  createdAtGT: Time
  createdAtGTE: Time
  createdAtLT: Time
  createdAtLTE: Time
  """updated_at field predicates"""
  updatedAt: Time
  updatedAtNEQ: Time
  updatedAtIn: [Time!]
  updatedAtNotIn: [Time!]
  updatedAtGT: Time
  updatedAtGTE: Time
  updatedAtLT: Time
  updatedAtLTE: Time
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """transactions edge predicates"""
  hasTransactions: Boolean
  hasTransactionsWith: [TransactionWhereInput!]
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
type Event implements Node {
  id: ID!
  from: String!
  keys: [Felt]!
  data: [Felt]!
  transaction: Transaction
}
"""A connection to a list of items."""
type EventConnection {
  """A list of edges."""
  edges: [EventEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  totalCount: Int!
}
"""An edge in a connection."""
type EventEdge {
  """The item at the end of the edge."""
  node: Event
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
EventWhereInput is used for filtering Event objects.
Input was generated by ent.
"""
input EventWhereInput {
  not: EventWhereInput
  and: [EventWhereInput!]
  or: [EventWhereInput!]
  """from field predicates"""
  from: String
  fromNEQ: String
  fromIn: [String!]
  fromNotIn: [String!]
  fromGT: String
  fromGTE: String
  fromLT: String
  fromLTE: String
  fromContains: String
  fromHasPrefix: String
  fromHasSuffix: String
  fromEqualFold: String
  fromContainsFold: String
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """transaction edge predicates"""
  hasTransaction: Boolean
  hasTransactionWith: [TransactionWhereInput!]
}
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node {
  """The id of the object."""
  id: ID!
}
enum OrderDirection {
  ASC
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """When paginating forwards, are there more items?"""
  hasNextPage: Boolean!
  """When paginating backwards, are there more items?"""
  hasPreviousPage: Boolean!
  """When paginating backwards, the cursor to continue."""
  startCursor: Cursor
  """When paginating forwards, the cursor to continue."""
  endCursor: Cursor
}
"""Status is enum for the field status"""
enum Status @goModel(model: "github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt.Status") {
  UNKNOWN
  RECEIVED
  PENDING
  ACCEPTED_ON_L2
  ACCEPTED_ON_L1
  REJECTED
}
type Transaction implements Node {
  id: ID!
  contractAddress: String!
  entryPointSelector: String!
  transactionHash: String!
  calldata: [String!]!
  signature: [String!]
  nonce: String!
  block: Block
  receipt: TransactionReceipt
  events: [Event!]
}
"""A connection to a list of items."""
type TransactionConnection {
  """A list of edges."""
  edges: [TransactionEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  totalCount: Int!
}
"""An edge in a connection."""
type TransactionEdge {
  """The item at the end of the edge."""
  node: Transaction
  """A cursor for use in pagination."""
  cursor: Cursor!
}
input TransactionOrder {
  direction: OrderDirection! = ASC
  field: TransactionOrderField!
}
enum TransactionOrderField {
  NONCE
}
type TransactionReceipt implements Node {
  id: ID!
  transactionHash: String!
  status: Status!
  statusData: String!
  l1OriginMessage: L2Message!
  block: Block
  transaction: Transaction!
}
"""A connection to a list of items."""
type TransactionReceiptConnection {
  """A list of edges."""
  edges: [TransactionReceiptEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  totalCount: Int!
}
"""An edge in a connection."""
type TransactionReceiptEdge {
  """The item at the end of the edge."""
  node: TransactionReceipt
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
TransactionReceiptWhereInput is used for filtering TransactionReceipt objects.
Input was generated by ent.
"""
input TransactionReceiptWhereInput {
  not: TransactionReceiptWhereInput
  and: [TransactionReceiptWhereInput!]
  or: [TransactionReceiptWhereInput!]
  """transaction_hash field predicates"""
  transactionHash: String
  transactionHashNEQ: String
  transactionHashIn: [String!]
  transactionHashNotIn: [String!]
  transactionHashGT: String
  transactionHashGTE: String
  transactionHashLT: String
  transactionHashLTE: String
  transactionHashContains: String
  transactionHashHasPrefix: String
  transactionHashHasSuffix: String
  transactionHashEqualFold: String
  transactionHashContainsFold: String
  """status field predicates"""
  status: Status
  statusNEQ: Status
  statusIn: [Status!]
  statusNotIn: [Status!]
  """status_data field predicates"""
  statusData: String
  statusDataNEQ: String
  statusDataIn: [String!]
  statusDataNotIn: [String!]
  statusDataGT: String
  statusDataGTE: String
  statusDataLT: String
  statusDataLTE: String
  statusDataContains: String
  statusDataHasPrefix: String
  statusDataHasSuffix: String
  statusDataEqualFold: String
  statusDataContainsFold: String
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """block edge predicates"""
  hasBlock: Boolean
  hasBlockWith: [BlockWhereInput!]
  """transaction edge predicates"""
  hasTransaction: Boolean
  hasTransactionWith: [TransactionWhereInput!]
}
"""
TransactionWhereInput is used for filtering Transaction objects.
Input was generated by ent.
"""
input TransactionWhereInput {
  not: TransactionWhereInput
  and: [TransactionWhereInput!]
  or: [TransactionWhereInput!]
  """contract_address field predicates"""
  contractAddress: String
  contractAddressNEQ: String
  contractAddressIn: [String!]
  contractAddressNotIn: [String!]
  contractAddressGT: String
  contractAddressGTE: String
  contractAddressLT: String
  contractAddressLTE: String
  contractAddressContains: String
  contractAddressHasPrefix: String
  contractAddressHasSuffix: String
  contractAddressEqualFold: String
  contractAddressContainsFold: String
  """entry_point_selector field predicates"""
  entryPointSelector: String
  entryPointSelectorNEQ: String
  entryPointSelectorIn: [String!]
  entryPointSelectorNotIn: [String!]
  entryPointSelectorGT: String
  entryPointSelectorGTE: String
  entryPointSelectorLT: String
  entryPointSelectorLTE: String
  entryPointSelectorContains: String
  entryPointSelectorHasPrefix: String
  entryPointSelectorHasSuffix: String
  entryPointSelectorIsNil: Boolean
  entryPointSelectorNotNil: Boolean
  entryPointSelectorEqualFold: String
  entryPointSelectorContainsFold: String
  """transaction_hash field predicates"""
  transactionHash: String
  transactionHashNEQ: String
  transactionHashIn: [String!]
  transactionHashNotIn: [String!]
  transactionHashGT: String
  transactionHashGTE: String
  transactionHashLT: String
  transactionHashLTE: String
  transactionHashContains: String
  transactionHashHasPrefix: String
  transactionHashHasSuffix: String
  transactionHashEqualFold: String
  transactionHashContainsFold: String
  """nonce field predicates"""
  nonce: String
  nonceNEQ: String
  nonceIn: [String!]
  nonceNotIn: [String!]
  nonceGT: String
  nonceGTE: String
  nonceLT: String
  nonceLTE: String
  nonceContains: String
  nonceHasPrefix: String
  nonceHasSuffix: String
  nonceIsNil: Boolean
  nonceNotNil: Boolean
  nonceEqualFold: String
  nonceContainsFold: String
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """block edge predicates"""
  hasBlock: Boolean
  hasBlockWith: [BlockWhereInput!]
  """receipt edge predicates"""
  hasReceipt: Boolean
  hasReceiptWith: [TransactionReceiptWhereInput!]
  """events edge predicates"""
  hasEvents: Boolean
  hasEventsWith: [EventWhereInput!]
}
"""Type is enum for the field type"""
enum Type @goModel(model: "github.com/cartridge-gg/starknet-indexer/ent/contract.Type") {
  UNKNOWN
  ERC20
  ERC721
}
`, BuiltIn: false},
	{Name: "schema.graphql", Input: `scalar Time
scalar Long
scalar JSON
scalar Felt

type L1Message {
  toAddress: String!
  payload: [Felt!]
}

type L2Message {
  fromAddress: String!
  payload: [Felt!]
}

extend type TransactionReceipt {
  messagesSent: [L1Message]!
}

type Query {
  node(id: ID!): Node
  nodes(ids: [ID!]!): [Node]!
  blocks(
    after: Cursor
    first: Int
    before: Cursor
    last: Int
    orderBy: BlockOrder
    where: BlockWhereInput
  ): BlockConnection
  events(
    after: Cursor
    first: Int
    before: Cursor
    last: Int
    where: EventWhereInput
  ): EventConnection
  transactions(
    after: Cursor
    first: Int
    before: Cursor
    last: Int
    where: TransactionWhereInput
  ): TransactionConnection
}

type Subscription {
  watchEvent(address: String!, keys: [Felt!]): Event
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
