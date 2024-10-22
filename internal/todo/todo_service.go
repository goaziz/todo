package todo

import (
	"context"
	"fmt"
	"sync"

	pb "github.com/goaziz/todo/proto/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ToDoService is an implementation of the ToDoService gRPC server
type ToDoService struct {
	pb.UnimplementedToDoServiceServer
	mu          sync.Mutex
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewToDoService() *ToDoService {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	collection := client.Database("todo").Collection("todos")
	return &ToDoService{
		mongoClient: client,
		collection:  collection,
	}
}

// @Summary Create a new ToDo item
// @Description Create a new ToDo item and store in MongoDB
// @Accept  json
// @Produce  json
// @Param   todo  body     pb.CreateToDoRequest  true  "ToDo details"
// @Success 200   {object} pb.ToDo
// @Router /v1/todo [post]
func (s *ToDoService) CreateToDo(ctx context.Context, req *pb.CreateToDoRequest) (*pb.ToDo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := req.GetTodo()

	// Insert the ToDo into MongoDB
	_, err := s.collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// @Summary Get a To Do item by ID
// @Description Get a To Do item by ID from MongoDB
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "To Do ID"
// @Success 200   {object} pb.ToDo
// @Router /v1/todo/{id} [get]
func (s *ToDoService) GetToDo(ctx context.Context, req *pb.GetToDoRequest) (*pb.ToDo, error) {
	filter := bson.M{"id": req.Id} // Use bson.M for filtering
	var todo pb.ToDo
	err := s.collection.FindOne(ctx, filter).Decode(&todo)
	if err != nil {
		return nil, fmt.Errorf("ToDo not found: %v", err)
	}
	return &todo, nil
}

// @Summary List all To Do items
// @Description List all To Do items from MongoDB
// @Accept  json
// @Produce  json
// @Success 200   {object} pb.ListToDoResponse
// @Router /v1/todo [get]
func (s *ToDoService) ListToDo(ctx context.Context, req *emptypb.Empty) (*pb.ListToDoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []*pb.ToDo
	for cursor.Next(ctx) {
		var todo pb.ToDo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListToDoResponse{Todos: todos}, nil
}

// @Summary Update a To Do item
// @Description Update a To Do item in MongoDB
// @Accept  json
// @Produce  json
// @Param   todo  body     pb.UpdateToDoRequest  true  "ToDo details"
// @Success 200   {object} pb.ToDo
// @Router /v1/todo [put]
func (s *ToDoService) UpdateToDo(ctx context.Context, req *pb.UpdateToDoRequest) (*pb.ToDo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter := bson.M{"id": req.Id} // Use bson.M for filtering
	update := bson.M{
		"$set": bson.M{
			"title":       req.Title,
			"description": req.Description,
			"done":        req.Done,
		},
	}

	result, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("ToDo not found with id: %s", req.Id)
	}

	// Return the updated ToDo
	return s.GetToDo(ctx, &pb.GetToDoRequest{Id: req.Id})
}

// @Summary Delete a To Do item
// @Description Delete a To Do item from MongoDB
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "To Do ID"
// @Success 200   {object} pb.ToDo
// @Router /v1/todo/{id} [delete]
func (s *ToDoService) DeleteToDo(ctx context.Context, req *pb.GetToDoRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	filter := bson.M{"id": req.Id} // Use bson.M for filtering

	result, err := s.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("ToDo not found with id: %s", req.Id)
	}
	return &emptypb.Empty{}, nil
}
