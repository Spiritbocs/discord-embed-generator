package api

import (
	"github.com/merlinfuchs/embed-generator/embedg-server/actions/handler"
	"github.com/merlinfuchs/embed-generator/embedg-server/actions/parser"
	"github.com/merlinfuchs/embed-generator/embedg-server/api/access"
	"github.com/merlinfuchs/embed-generator/embedg-server/api/premium"
	"github.com/merlinfuchs/embed-generator/embedg-server/api/session"
	"github.com/merlinfuchs/embed-generator/embedg-server/bot"
	"github.com/merlinfuchs/embed-generator/embedg-server/custom_bots"
)

type managers struct {
	session    *session.SessionManager
	access     *access.AccessManager
	premium    *premium.PremiumManager
	customBots *custom_bots.CustomBotManager

	actionParser  *parser.ActionParser
	actionHandler *handler.ActionHandler
}

func createManagers(stores *stores, bot *bot.Bot) *managers {
	sessionManager := session.New(stores.pg)
	accessManager := access.New(bot.State, bot.Session)
	premiumManager := premium.New(stores.pg, bot)

	actionParser := parser.New(accessManager, stores.pg, bot.State)
	actionHandler := handler.New(stores.pg, actionParser)

	customBots := custom_bots.NewCustomBotManager(stores.pg, actionHandler)

	bot.ActionHandler = actionHandler
	bot.ActionParser = actionParser

	return &managers{
		session:       sessionManager,
		access:        accessManager,
		premium:       premiumManager,
		customBots:    customBots,
		actionParser:  actionParser,
		actionHandler: actionHandler,
	}
}
