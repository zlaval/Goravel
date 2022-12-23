package goravel

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"goravel/render"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type Goravel struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	JetViews *jet.Set
	config   config
}

type config struct {
	port     string
	renderer string
}

func (g *Goravel) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := g.Init(pathConfig)
	if err != nil {
		return err
	}

	err = g.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	infoLog, errorLog := g.startLoggers()

	g.ErrorLog = errorLog
	g.InfoLog = infoLog
	g.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	g.Version = version
	g.RootPath = rootPath
	g.Routes = g.routes().(*chi.Mux)
	g.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}
	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)
	g.JetViews = views
	g.createRenderer()
	return nil
}

func (g *Goravel) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := g.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Goravel) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", g.config.port),
		ErrorLog:     g.ErrorLog,
		Handler:      g.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 300 * time.Second,
	}
	g.InfoLog.Printf("Listening on port %s", g.config.port)
	err := srv.ListenAndServe()
	g.ErrorLog.Fatal(err)
}

func (g *Goravel) checkDotEnv(path string) error {
	err := g.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (g *Goravel) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}

func (g *Goravel) createRenderer() {
	renderer := render.Render{
		Renderer: g.config.renderer,
		RootPath: g.RootPath,
		Port:     g.config.port,
		JetViews: g.JetViews,
	}
	g.Render = &renderer
}
