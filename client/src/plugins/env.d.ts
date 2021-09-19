declare var process: {
	env: {
		NODE_ENV: string
		VUE_APP_API_HOST: string // for dev
		VUE_APP_SOCKET_HOST: string // for dev
		VUE_APP_API_PATH: string // for production
		VUE_APP_HOST: string // for production
		VUE_APP_SOCKET_PATH: string // for production
	}
}