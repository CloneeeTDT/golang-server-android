package server

func Init() {
	r := NewRouter()
	err := r.Run() // Run in 8080
	if err != nil {
		return
	}
}
