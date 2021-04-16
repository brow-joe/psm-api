module br.com.jonathan/psm

go 1.16

replace br.com.jonathan/psm/api => ./src/api

replace br.com.jonathan/psm/api/domain => ./src/api/domain

replace br.com.jonathan/psm/api/usecases => ./src/api/usecases

replace br.com.jonathan/psm/api/userinterfaces => ./src/api/userinterfaces

require (
	br.com.jonathan/psm/api v0.0.0-00010101000000-000000000000
	br.com.jonathan/psm/api/domain v0.0.0-00010101000000-000000000000 // indirect
	br.com.jonathan/psm/api/userinterfaces v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
)
