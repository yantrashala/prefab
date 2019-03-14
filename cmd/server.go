// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>
//

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yantrashala/prefab/model"
)

type server struct {
	router *mux.Router
}

// Server struct
var Server *server

func (s *server) ListenAndServe(port string) error {
	return http.ListenAndServe(":"+port, Server.router)
}

func (s *server) routes() {
	fs := http.FileServer(http.Dir("./ui/build"))

	s.router.HandleFunc("/api/project/name", s.handleGetName()).Methods("GET")
	s.router.HandleFunc("/api/environments/build", s.handleGetBuildEnvironments()).Methods("GET")
	s.router.PathPrefix("/").Handler(http.StripPrefix("/", fs))
}

func (s *server) handleGetName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.CurrentProject.Name)
	}
}

func (s *server) handleGetBuildEnvironments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		l, err := model.GetBuildEnvironmentTypes()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(l)
	}
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starts the web UI",
	Long:  `starts a web UI in the specified port (default 9876) for a interactive prefab configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		var port = viper.GetString("port")
		fmt.Println(colors.Blue("prefab"), " confuration server running at ", colors.Bold(colors.Green("http://localhost:"+port)))
		if err := Server.ListenAndServe(port); err != nil {
			panic(err)
		}
	},
}

func init() {
	Server = &server{router: mux.NewRouter()}
	Server.routes()

	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().Uint16P("port", "p", 9876, "port for the prefab ui")

	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	viper.SetDefault("port", 9876)
}
