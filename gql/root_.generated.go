// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/tarrencev/starknet-indexer/ent"
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

	BuiltinInstanceCounter struct {
		BitwiseBuiltin    func(childComplexity int) int
		EcOpBuiltin       func(childComplexity int) int
		EcdsaBuiltin      func(childComplexity int) int
		OutputBuiltin     func(childComplexity int) int
		PedersenBuiltin   func(childComplexity int) int
		RangeCheckBuiltin func(childComplexity int) int
	}

	ExecutionResources struct {
		BuiltinInstanceCounter func(childComplexity int) int
		NMemoryHoles           func(childComplexity int) int
		NSteps                 func(childComplexity int) int
	}

	L1ToL2ConsumedMessage struct {
		FromAddress func(childComplexity int) int
		Payload     func(childComplexity int) int
		Selector    func(childComplexity int) int
		ToAddress   func(childComplexity int) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Blocks       func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.BlockOrder, where *BlockWhereInput) int
		Node         func(childComplexity int, id string) int
		Nodes        func(childComplexity int, ids []string) int
		Transactions func(childComplexity int, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *TransactionWhereInput) int
	}

	Transaction struct {
		Block              func(childComplexity int) int
		Calldata           func(childComplexity int) int
		ContractAddress    func(childComplexity int) int
		EntryPointSelector func(childComplexity int) int
		EntryPointType     func(childComplexity int) int
		ID                 func(childComplexity int) int
		Nonce              func(childComplexity int) int
		Receipts           func(childComplexity int) int
		Signature          func(childComplexity int) int
		TransactionHash    func(childComplexity int) int
		Type               func(childComplexity int) int
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
		Block                 func(childComplexity int) int
		Events                func(childComplexity int) int
		ExecutionResources    func(childComplexity int) int
		ID                    func(childComplexity int) int
		L1ToL2ConsumedMessage func(childComplexity int) int
		L2ToL1Messages        func(childComplexity int) int
		Transaction           func(childComplexity int) int
		TransactionHash       func(childComplexity int) int
		TransactionIndex      func(childComplexity int) int
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

	case "BuiltinInstanceCounter.bitwiseBuiltin":
		if e.complexity.BuiltinInstanceCounter.BitwiseBuiltin == nil {
			break
		}

		return e.complexity.BuiltinInstanceCounter.BitwiseBuiltin(childComplexity), true

	case "BuiltinInstanceCounter.ecOpBuiltin":
		if e.complexity.BuiltinInstanceCounter.EcOpBuiltin == nil {
			break
		}

		return e.complexity.BuiltinInstanceCounter.EcOpBuiltin(childComplexity), true

	case "BuiltinInstanceCounter.ecdsaBuiltin":
		if e.complexity.BuiltinInstanceCounter.EcdsaBuiltin == nil {
			break
		}

		return e.complexity.BuiltinInstanceCounter.EcdsaBuiltin(childComplexity), true

	case "BuiltinInstanceCounter.outputBuiltin":
		if e.complexity.BuiltinInstanceCounter.OutputBuiltin == nil {
			break
		}

		return e.complexity.BuiltinInstanceCounter.OutputBuiltin(childComplexity), true

	case "BuiltinInstanceCounter.pedersenBuiltin":
		if e.complexity.BuiltinInstanceCounter.PedersenBuiltin == nil {
			break
		}

		return e.complexity.BuiltinInstanceCounter.PedersenBuiltin(childComplexity), true

	case "BuiltinInstanceCounter.rangeCheckBuiltin":
		if e.complexity.BuiltinInstanceCounter.RangeCheckBuiltin == nil {
			break
		}

		return e.complexity.BuiltinInstanceCounter.RangeCheckBuiltin(childComplexity), true

	case "ExecutionResources.builtinInstanceCounter":
		if e.complexity.ExecutionResources.BuiltinInstanceCounter == nil {
			break
		}

		return e.complexity.ExecutionResources.BuiltinInstanceCounter(childComplexity), true

	case "ExecutionResources.nMemoryHoles":
		if e.complexity.ExecutionResources.NMemoryHoles == nil {
			break
		}

		return e.complexity.ExecutionResources.NMemoryHoles(childComplexity), true

	case "ExecutionResources.nSteps":
		if e.complexity.ExecutionResources.NSteps == nil {
			break
		}

		return e.complexity.ExecutionResources.NSteps(childComplexity), true

	case "L1ToL2ConsumedMessage.fromAddress":
		if e.complexity.L1ToL2ConsumedMessage.FromAddress == nil {
			break
		}

		return e.complexity.L1ToL2ConsumedMessage.FromAddress(childComplexity), true

	case "L1ToL2ConsumedMessage.payload":
		if e.complexity.L1ToL2ConsumedMessage.Payload == nil {
			break
		}

		return e.complexity.L1ToL2ConsumedMessage.Payload(childComplexity), true

	case "L1ToL2ConsumedMessage.selector":
		if e.complexity.L1ToL2ConsumedMessage.Selector == nil {
			break
		}

		return e.complexity.L1ToL2ConsumedMessage.Selector(childComplexity), true

	case "L1ToL2ConsumedMessage.toAddress":
		if e.complexity.L1ToL2ConsumedMessage.ToAddress == nil {
			break
		}

		return e.complexity.L1ToL2ConsumedMessage.ToAddress(childComplexity), true

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

	case "Transaction.entryPointType":
		if e.complexity.Transaction.EntryPointType == nil {
			break
		}

		return e.complexity.Transaction.EntryPointType(childComplexity), true

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

	case "Transaction.receipts":
		if e.complexity.Transaction.Receipts == nil {
			break
		}

		return e.complexity.Transaction.Receipts(childComplexity), true

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

	case "Transaction.type":
		if e.complexity.Transaction.Type == nil {
			break
		}

		return e.complexity.Transaction.Type(childComplexity), true

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

	case "TransactionReceipt.events":
		if e.complexity.TransactionReceipt.Events == nil {
			break
		}

		return e.complexity.TransactionReceipt.Events(childComplexity), true

	case "TransactionReceipt.executionResources":
		if e.complexity.TransactionReceipt.ExecutionResources == nil {
			break
		}

		return e.complexity.TransactionReceipt.ExecutionResources(childComplexity), true

	case "TransactionReceipt.id":
		if e.complexity.TransactionReceipt.ID == nil {
			break
		}

		return e.complexity.TransactionReceipt.ID(childComplexity), true

	case "TransactionReceipt.l1ToL2ConsumedMessage":
		if e.complexity.TransactionReceipt.L1ToL2ConsumedMessage == nil {
			break
		}

		return e.complexity.TransactionReceipt.L1ToL2ConsumedMessage(childComplexity), true

	case "TransactionReceipt.l2ToL1Messages":
		if e.complexity.TransactionReceipt.L2ToL1Messages == nil {
			break
		}

		return e.complexity.TransactionReceipt.L2ToL1Messages(childComplexity), true

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

	case "TransactionReceipt.transactionIndex":
		if e.complexity.TransactionReceipt.TransactionIndex == nil {
			break
		}

		return e.complexity.TransactionReceipt.TransactionIndex(childComplexity), true

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
enum BlockStatus @goModel(model: "github.com/tarrencev/starknet-indexer/ent/block.Status") {
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
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
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
type Transaction implements Node {
  id: ID!
  contractAddress: String!
  entryPointSelector: String!
  entryPointType: String!
  transactionHash: String!
  calldata: [String!]!
  signature: [String!]!
  type: Type!
  nonce: String!
  block: Block
  receipts: TransactionReceipt
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
  transactionIndex: Int!
  transactionHash: String!
  l1ToL2ConsumedMessage: L1ToL2ConsumedMessage!
  executionResources: ExecutionResources!
  events: JSON!
  l2ToL1Messages: JSON!
  block: Block
  transaction: Transaction
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
input TransactionReceiptOrder {
  direction: OrderDirection! = ASC
  field: TransactionReceiptOrderField!
}
enum TransactionReceiptOrderField {
  INDEX
}
"""
TransactionReceiptWhereInput is used for filtering TransactionReceipt objects.
Input was generated by ent.
"""
input TransactionReceiptWhereInput {
  not: TransactionReceiptWhereInput
  and: [TransactionReceiptWhereInput!]
  or: [TransactionReceiptWhereInput!]
  """transaction_index field predicates"""
  transactionIndex: Int
  transactionIndexNEQ: Int
  transactionIndexIn: [Int!]
  transactionIndexNotIn: [Int!]
  transactionIndexGT: Int
  transactionIndexGTE: Int
  transactionIndexLT: Int
  transactionIndexLTE: Int
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
  entryPointSelectorEqualFold: String
  entryPointSelectorContainsFold: String
  """entry_point_type field predicates"""
  entryPointType: String
  entryPointTypeNEQ: String
  entryPointTypeIn: [String!]
  entryPointTypeNotIn: [String!]
  entryPointTypeGT: String
  entryPointTypeGTE: String
  entryPointTypeLT: String
  entryPointTypeLTE: String
  entryPointTypeContains: String
  entryPointTypeHasPrefix: String
  entryPointTypeHasSuffix: String
  entryPointTypeEqualFold: String
  entryPointTypeContainsFold: String
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
  """type field predicates"""
  type: Type
  typeNEQ: Type
  typeIn: [Type!]
  typeNotIn: [Type!]
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
  """receipts edge predicates"""
  hasReceipts: Boolean
  hasReceiptsWith: [TransactionReceiptWhereInput!]
}
"""Type is enum for the field type"""
enum Type @goModel(model: "github.com/tarrencev/starknet-indexer/ent/transaction.Type") {
  INVOKE_FUNCTION
  DEPLOY
}
`, BuiltIn: false},
	{Name: "schema.graphql", Input: `scalar Time
scalar Long
scalar JSON

type L1ToL2ConsumedMessage {
  fromAddress: String!
  toAddress: String!
  selector: String!
  payload: [String!]
}

type BuiltinInstanceCounter {
  pedersenBuiltin: Long
  rangeCheckBuiltin: Long
  bitwiseBuiltin: Long
  outputBuiltin: Long
  ecdsaBuiltin: Long
  ecOpBuiltin: Long
}

type ExecutionResources {
  nSteps: Long
  builtinInstanceCounter: BuiltinInstanceCounter
  nMemoryHoles: Long
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
  transactions(
    after: Cursor
    first: Int
    before: Cursor
    last: Int
    where: TransactionWhereInput
  ): TransactionConnection
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
