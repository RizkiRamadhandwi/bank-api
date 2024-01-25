package delivery

import (
	"bank-api/config"
	"bank-api/delivery/controller"
	"bank-api/delivery/middleware"
	"bank-api/repository"
	"bank-api/shared/service"
	"bank-api/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	transUc    usecase.TransactionUseCase
	userUc     usecase.UserUseCase
	authUsc    usecase.AuthUseCase
	jwtService service.JwtService
	engine     *gin.Engine
	host       string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	authMid := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewAuthController(s.authUsc, rg, s.jwtService).Route()
	controller.NewTransactionController(s.transUc, rg, authMid).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	jwtService := service.NewJwtService(cfg.TokenConfig)

	userRepo := repository.NewUserRepository("repository/data/customers.json")
	merchRepo := repository.NewMerchantRepository("repository/data/merchant.json")
	transRepo := repository.NewTransactionRepository("repository/data/transactions.json", userRepo, merchRepo)

	transUC := usecase.NewTransactionUseCase(transRepo)
	userUc := usecase.NewUserUseCase(userRepo)
	authUc := usecase.NewAuthUseCase(userUc, jwtService)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		transUc:    transUC,
		userUc:     userUc,
		authUsc:    authUc,
		jwtService: jwtService,
		engine:     engine,
		host:       host,
	}
}
