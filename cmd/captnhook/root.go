package captnhook

import (
	"github.com/bareish/captnHook/pkg/broker/alpaca"
	"github.com/bareish/captnHook/pkg/config"
	"github.com/bareish/captnHook/pkg/http/rest"
	"github.com/bareish/captnHook/pkg/market/stock"
)

// Run loads configuration variables starts our services and initializes our server
func Run() {
	// create config service
	configService := &config.Service{}
	configService.Load()

	// create market data service
	marketDataService := &stock.MarketDataService{ConfigService: configService}
	marketDataService.Setup()

	// create the new broker service
	brokerService := &alpaca.BrokerService{
		ConfigService: configService,
		MarketDataService: marketDataService,
	}
	brokerService.Setup()
	
	// HTTP/2 REST server
	server := rest.NewRESTServer(configService, brokerService)
	server.Start()

}
