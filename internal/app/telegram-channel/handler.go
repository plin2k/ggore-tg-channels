package telegram_channel

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/plin2k/ggore-tg-channels/internal/repository"

	"github.com/plin2k/ggore-tg-channels/internal/models"
	"github.com/plin2k/ggore-tg-channels/internal/services"
)

const (
	xmlDir = "./assets/channels/"
)

type handler struct {
	config   *models.Config
	services *services.Service
	repo     *repository.Repository

	limitArticlesChannelMap map[int64]int
	wg                      sync.WaitGroup
	pubCh                   chan models.Article
}

func getChannels(dir string) (models.Channels, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln("ReadDir", err)
	}

	var channels = make(models.Channels, len(files))

	for i, f := range files {
		//if f.Name() != "invest_useful.xml" {
		//	continue
		//}

		file, err := os.Open(dir + f.Name())
		if err != nil {
			return nil, err
		}

		byteFile, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		file.Close()

		err = xml.Unmarshal(byteFile, &channels[i])
		if err != nil {
			return nil, err
		}
	}

	return channels, nil
}

func (h *handler) publish() {

	log.Println("Publisher Started")

	for a := range h.pubCh {

		if h.limitArticlesChannelMap[a.ChannelID] < 1 {
			continue
		}

		exists, err := h.repo.Article.Exists(a)
		if err != nil {
			log.Printf("ERROR article exists : %d %s \n", a.ChannelID, err)
			continue
		}

		if exists {
			continue
		}

		err = h.repo.Article.Create(a)
		if err != nil {
			log.Printf("ERROR content create : %d %s \n", a.ChannelID, err)
		}

		err = h.services.Telegram.Send(a.ChannelID, a.HTMLMessage)
		if err != nil {
			log.Printf("ERROR telegram send : %s %s \n", a.ChannelID, err)
		}

		h.limitArticlesChannelMap[a.ChannelID]--

		log.Println(a)

		rand.Seed(time.Now().UnixMicro())
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}

func Start(cfg *models.Config, services *services.Service, repo *repository.Repository) error {
	h := &handler{
		config:   cfg,
		services: services,
		repo:     repo,

		pubCh: make(chan models.Article, 10),
	}

	channels, err := getChannels(xmlDir)
	if err != nil {
		return fmt.Errorf("telegram-channel start: %w", err)
	}

	h.limitArticlesChannelMap = channels.GetLimitMap(h.config.LimitArticles)

	channels.Construct()

	go h.publish()

	h.wg.Add(len(channels))
	go func() {
		for i := range channels {
			go h.parse(&channels[i])
		}
	}()

	h.wg.Wait()

	close(h.pubCh)

	for {
		if len(h.pubCh) == 0 {
			break
		}
	}

	return nil
}

func (h *handler) parse(channel *models.Channel) error {
	defer h.wg.Done()

	log.Printf("Starting : %s", channel.Name)

	var contents models.Contents

	for _, s := range channel.Sources.Shuffle() {

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

		log.Printf("Parse : %s", s.Link)
		content, err := h.services.Parser.Parse(s)
		if err != nil {
			return err
		}

		contents.Reverse()

		contents = append(contents, content...)
	}

	for _, content := range contents.Shuffle() {
		h.pubCh <- content.ToArticle(channel)
	}

	return nil
}
