package main

import (
  "context"
  "errors"
  "fmt"
  "log"
  "net"
  "os"
  "strings"

  svc "github.com/the2sang/server-streaming/service"
  "google.golang.org/grpc"
)

type userService struct {
  svc.UnimplementedUserServer
}

type repoService struct {
  svc.UnimplementedRepoServer
}

func (s *userService) GetUser(ctx context.Context, in *svc.UserGetRequest) (*svc.userGetReply, error) {
  log.Printf("Received request for user with Email: %s Id: %s\n", in.Email, in.Id)
  components := strings.Split(in.Email, "@")
  if len(components) != 2 {
    return nil, errors.New("invalid email address")
  }
  u := svc.User{
    Id: in.Id,
    FirstName: components[0],
    LastName: components[1],
    Age: 36,
  }
  return &svc.userGetReply{User: &u,}, nil
}

func (s *repoService) GetRepos(in *svc.RepoGetRequest, stream svc.Repo_GetRepoServer,) error {
  log.Printf("Recieved request for repo with CreateId: %s Id: %s\n", in.CreateId, in.Id,)
  repo := svc.Repository{
    Owner: &svc.User{Id: in.CreateId, FirstName: "Jane"},
  }
  cnt := 1
  for {
    repo.Name = fmt.Sprintf("repo-%d", cnt)
    repo.Url = fmt.Sprintf("https://git.example.com/tet/%s", repo.Name)
    r := svc.RepoGetReply{
      Repo: &repo,
    }
    if err := stream.Send(&r); err != nil {
      return err
    }
    if cnt >= 5 {
      break
    }
    cnt++
  }
  return nil
}

func registerServices(s *grpc.Server) {
  svc.RegisterUsersServer(s, &userService{})
  svc.RegisterRepoServer(s, &repoService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
  return s.Serve(l)
}

func main() {
  listenAddr := os.Getenv("LISTEN_ADDR")
  if len(listenAddr) == 0 {
    listenAddr = ":50051"
  }

  lis, err := net.Listen("tcp", listenAddr)
  if err != nil {
    log.Fatal(err)
  }

  s := grpc.NewServer()
  registerServices(s)
  log.Fatal(startServer(s, lis))
}