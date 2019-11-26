package middleware

import (
	"bufio"
	"context"
	useCase "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/prometheus/client_golang/prometheus"
	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"strconv"
	"time"
)

var (
	hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hits",
		},
		[]string{"status", "path"},
	)

	timings = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "timings",
			Help:       "Timings",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"path"},
	)
)

type HandlersMiddlwares struct {
	Users  useCase.UsersUseCase
	Logger *logrus.Logger
}

type LogResponse struct {
	w      http.ResponseWriter
	status int
}

func (l *LogResponse) Header() http.Header {
	return l.w.Header()
}

func (l *LogResponse) Write(body []byte) (int, error) {
	return l.w.Write(body)
}

func (l *LogResponse) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return l.w.(http.Hijacker).Hijack()
}

func (l *LogResponse) WriteHeader(code int) {
	l.w.WriteHeader(code)
	l.status = code
}

func (m *HandlersMiddlwares) AuthMiddleware(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err != nil {
			logrus.SetFormatter(&logrus.TextFormatter{})
			logrus.WithFields(logrus.Fields{
				"method":      r.Method,
				"remote_addr": r.RemoteAddr,
			}).Error("not authorised: request not contains cookie")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		logrus.Info("cookie = " + session.Value)
		id, err := m.Users.GetUserBySession(session.Value)
		if err != nil {
			logrus.SetFormatter(&logrus.TextFormatter{})
			logrus.WithFields(logrus.Fields{
				"method":      r.Method,
				"remote_addr": r.RemoteAddr,
			}).Error("Cookie not valid" + err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "user_id", uint64(id)))

		//token := r.Header.Get("X-CSRF-Token")
		//_, err = utils.Tokens.Check(id, session.Value, token)
		//if err != nil {
		//	logrus.SetFormatter(&logrus.TextFormatter{})
		//	logrus.WithFields(logrus.Fields{
		//		"method":      r.Method,
		//		"remote_addr": r.RemoteAddr,
		//	}).Error("not authorised: request not contains csrf token")
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		////create csrf token
		//tokenExpiration := time.Now().Add(24 * time.Hour)
		//csrfToken, err := utils.Tokens.Create(id, session.Value, tokenExpiration.Unix())
		//w.Header().Set("X-CSRF-Token", csrfToken)
		next(w, r)

	})
}

func (m *HandlersMiddlwares) LogMiddleware(next http.Handler, logrusLogger *logrus.Logger) http.Handler {
	prometheus.MustRegister(hits)
	prometheus.MustRegister(timings)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logWriter := LogResponse{
			w:      w,
			status: 200,
		}
		next.ServeHTTP(&logWriter, r)

		m.Logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			"work_time":   time.Since(start),
			"status_code": logWriter.status,
		}).Info(r.URL.Path)

		hits.
			WithLabelValues(strconv.Itoa(logWriter.status), r.URL.Path).
			Inc()

		timings.
			WithLabelValues(r.URL.Path).
			Observe(float64(time.Since(start).Milliseconds()))
	})
}

func (m *HandlersMiddlwares) PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				m.Logger.WithFields(logrus.Fields{
					"method":      r.Method,
					"remote_addr": r.RemoteAddr,
					"panic":       err,
				}).Error(r.URL.Path)
				http.Error(w, "Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
