package web

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "runtime"

    "github.com/gin-contrib/gzip"
    "github.com/gin-gonic/gin"
)

var (
    rootPath = getRootPath()
    tempPath = getTempPath()

    staticPath = filepath.Join(rootPath, `static`)
    htmlPath   = filepath.Join(rootPath, `templates`, `*.html`)
)

func Service() *gin.Engine {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.Use(gzip.Gzip(gzip.DefaultCompression))
    // router.Use(gin.Logger())
    router.Static("/tmp", tempPath)
    router.Static("/static", staticPath)

    router.LoadHTMLGlob(htmlPath)

    // 全部導到 /
    redirectToRoot := func(c *gin.Context) {
        if c.Request.URL.Path != "/" {
            c.Redirect(http.StatusMovedPermanently, "/")
        }
    }
    router.Use(redirectToRoot)

    return router
}

func getRootPath() string {
    path, err := os.Executable()
    printError(err)
    if string(path[:11]) == `z:\go-build` {
        path, err = os.Getwd()
        printError(err)
        path += `\`
    }
    return filepath.Dir(path)
}

func getTempPath() string {
    if runtime.GOOS == "windows" {
        return `z:\jdc\`
    }
    return `/tmp/`
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}

var _ = fmt.Println
