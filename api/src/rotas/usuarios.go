package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{ usuarioID}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{ usuarioID}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
}
