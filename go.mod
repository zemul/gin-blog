module gin-blog

go 1.13

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/jinzhu/gorm v1.9.16
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace (
	gin-blog/conf => /Users/zzm/test/gin-blog/pkg/conf
	gin-blog/middleware => /Users/zzm/test/gin-blog/middleware
	gin-blog/models => /Users/zzm/test/gin-blog/models
	gin-blog/pkg/setting => /Users/zzm/test/gin-blog/pkg/setting
	gin-blog/routers => /Users/zzm/test/gin-blog/routers
)
