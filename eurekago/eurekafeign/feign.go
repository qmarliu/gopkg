package feign

import (
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/qmarliu/gopkg/eurekago"
	"github.com/qmarliu/gopkg/log"
)

const (
	DEFAULT_REFRESH_APP_URLS_INTERVALS = 30
)

type Feign struct {
	// Discovery client to get Apps and instances
	discoveryClient eurekago.EurekaHttpClient

	// assign app => urls
	appUrls map[string][]string

	// Counter for calculate next url'index
	appNextUrlIndex map[string]*uint32

	// seconds of updating app's urls periodically
	refreshAppUrlsIntervals int

	// ensure some daemon task only run one time
	once sync.Once

	mu sync.RWMutex
}

// return resty.Client
func (t *Feign) Init(serviceUrlList []string, username, password string) {
	// daemon to update app urls periodically
	// only execute once globally
	t.once.Do(func() {

		t.appUrls = make(map[string][]string)
		t.appNextUrlIndex = make(map[string]*uint32)
		t.discoveryClient = eurekago.NewEurekaHttpClient(serviceUrlList, username, password, true)

		t.updateAppUrlsIntervals()
	})

	// try update app's urls.
	// if app's urls is exist, do nothing
	t.tryRefreshAppUrls()
}

func (t *Feign) SetRefreshAppUrlsIntervals(intervals int) {
	t.refreshAppUrlsIntervals = intervals
}

func (t *Feign) GetAppUrls(app string) []string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if _, ok := t.appUrls[app]; !ok {
		return nil
	}

	return t.appUrls[app]
}

// assign static app => urls
func (t *Feign) UseUrls(appUrls map[string][]string) *Feign {
	t.mu.Lock()
	defer t.mu.Unlock()

	//v := uint32(time.Now().UnixNano())
	//appNextUrlIndex[t.app] = &v
	for app, urls := range appUrls {

		// reset app'urls
		tmpAppUrls := make([]string, 0)
		for _, u := range urls {
			_, err := url.Parse(u)
			if err != nil {
				log.Errorf("Invalid url=%s, parse err=%s", u, err.Error())
				continue
			}

			tmpAppUrls = append(tmpAppUrls, u)
		}

		if len(tmpAppUrls) == 0 {
			log.Errorf("Empty valid urls for app=%s, skip to set app's urls", app)
			continue
		}

		t.appUrls[app] = tmpAppUrls
		log.Infof("%s app urls update %v", app, tmpAppUrls)
		if t.appNextUrlIndex[app] == nil {
			v := uint32(time.Now().UnixNano())
			t.appNextUrlIndex[app] = &v
		}
	}

	return t
}

func (t *Feign) PicUrl(app string) string {
	urls := t.GetAppUrls(app)
	if len(urls) == 0 {
		log.Errorf("Failed to pick server, reason: no available urls for app=%s", app)

		return ""
	}

	idx := atomic.AddUint32(t.appNextUrlIndex[app], 1)
	idx %= uint32(len(urls))
	atomic.CompareAndSwapUint32(t.appNextUrlIndex[app], uint32(len(urls)), 0)

	return urls[idx]
}

func (t *Feign) updateAppUrlsIntervals() {
	if t.refreshAppUrlsIntervals <= 0 {
		t.refreshAppUrlsIntervals = DEFAULT_REFRESH_APP_URLS_INTERVALS
	}

	go func() {
		for {
			t.updateAppUrls()

			time.Sleep(time.Second * time.Duration(t.refreshAppUrlsIntervals))
			log.Debugf("Update app urls interval...ok")
			for app, urls := range t.appUrls {
				log.Debugf("app=> %s, urls => %v", app, urls)
			}
		}
	}()
}

func (t *Feign) updateAppUrls() {
	appList, err := t.discoveryClient.GetApplications()
	if err != nil {
		log.Warnf("get applications failed %v", err)
		return
	}

	tmpAppUrls := make(map[string][]string)
	for _, apps := range appList.Applications.Applications {
		var curAppUrls []string
		var isUpdate bool

		if curAppUrls = t.GetAppUrls(apps.Name); len(curAppUrls) > 0 {
			if len(curAppUrls) != len(apps.Instance) {
				isUpdate = true
			} else {
				for _, instances := range apps.Instance {
					isExist := false
					for _, v := range curAppUrls {
						insHomePageUrl := strings.TrimRight(instances.HomePageUrl, "/")
						if v == insHomePageUrl {
							isExist = true
							break
						}
					}

					if !isExist {
						isUpdate = true
						break
					}
				}
			}
		}

		// app are not exist in t.appUrls or app's urls has been update
		if len(curAppUrls) == 0 || isUpdate {
			tmpAppUrls[apps.Name] = make([]string, 0)

			for _, ins := range apps.Instance {
				tmpAppUrls[apps.Name] = append(tmpAppUrls[apps.Name], strings.TrimRight(ins.HomePageUrl, "/"))
			}
		}
	}

	// update app's urls to feign
	t.UseUrls(tmpAppUrls)
}

func (t *Feign) tryRefreshAppUrls() {
	if len(t.appUrls) > 0 {
		return
	}
	if t.discoveryClient == nil {
		log.Debugf("no discovery client, no need to update app'urls.")
		return
	}

	t.updateAppUrls()
}
