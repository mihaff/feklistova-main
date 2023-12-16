package main

import (
	"context"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
)

// Serve - является функцией работы сервера. В ней указываются все параметры и она по сути постоянно запущена
func Serve(ctx context.Context) {
	server := http.Server{Addr: ":8080"}
	// Роутер используется для указания маршрутов на сервере
	router := mux.NewRouter()
	http.Handle("/", router)
	// Метод для подключения Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	// Роутинг по шаблонам для отображения
	router.HandleFunc("/", HomeHandlerTmpl)

	// Пути к страницам, где отрисовываются HTML шаблоны с отделами
	router.HandleFunc("/dept", DeptsHandlerTmpl)
	router.HandleFunc("/dept/newdept", NewDeptHandlerTmpl)
	router.HandleFunc("/dept/{id}", DeptHandlerTmpl)
	router.HandleFunc("/dept/{id}/edit", EditDeptHandlerTmpl)

	// Пути к страницам, где отрисовываются HTML шаблоны с сотрудниками
	router.HandleFunc("/empl", EmployeesHandlerTmpl)
	router.HandleFunc("/empl/newemployee", NewEmployeeHandlerTmpl)
	router.HandleFunc("/empl/{id}", EmployeeHandlerTmpl)
	router.HandleFunc("/empl/{id}/edit", EditEmployeeHandlerTmpl)

	//// Пути к API, то есть нет отрисовки шаблонов, данные передаются через JSON
	//// Пути к API методам сотрудника
	router.HandleFunc("/api/empl/newemployee", NewEmployeeHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/empl/{id}/edit", EditEmployeeHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/empl/{id}/delete", DeleteEmployeeHandler).Methods(http.MethodDelete)
	router.HandleFunc("/api/empl", EmployeesHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/empl/{id}", EmployeeHandler).Methods(http.MethodGet)

	//// Пути к API методам отделов
	router.HandleFunc("/api/dept/newdept", NewDeptHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/dept/{id}/edit", EditDeptHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/dept/{id}/delete", DeleteDeptHandler).Methods(http.MethodDelete)
	router.HandleFunc("/api/dept", AllDeptsHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/dept/{id}", DeptHandler).Methods(http.MethodGet)

	// В отдельном потоке запускаем сервер
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()
	// В основном потоке проверяем сигнал OS.INTERRUPT
	for {
		select {
		case <-ctx.Done():
			// Если получили сигнал об остановке, то завершаем работу сервера
			log.Println("Shutting down server")
			err := server.Shutdown(ctx)
			if err != nil {
				panic(err)
			}
			return
		}

	}
}
