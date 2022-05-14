# Lambda Architecture (Kaliaha Artsiom, s2110455009)

Das Ziel dieses Projekts bestand darin, eine vereinfachte Version der Lambda Architektur zu implementieren. Das ist eine Architektur, die für die Verarbeitung massiver Datenmengen gedacht wurde. Die Besonderheit der Architektur besteht in der Ausnutzung der Batch- und Streaming-Datenverarbeitung. Mittels dieser Architektur kann man sowohl aktuelle, als auch vergangene Daten/Änderungen auswerten.

Übersicht über die verwendeten Technologien:
- Programmiersprache: Golang
- Web Framework: Es wurde kein Web Framework verwendet, sondern einzelne Bibliotheken.
- Bibliotheken:
    + gorilla mux (Routing)
    + uber zap (strukturierte Logging)
    + sqlx (Wrapper für Boilerplate SQL Code)
    + viper (Konfigurationsmanagement)
    + gocql (Cassandra Treiber)
- Persistenz: MySQL (transaktionale Datenbank), Cassandra (analytische Datenbank)
- Message oriented Middleware: Kafka
- Deployment: Kubernetes im Docker Desktop, Helm, PowerShell Skripts zur Automatisierung des Deployment Prozesses

### Architektur der Lösung

Die in dieser Übung implementierte Architektur wertet Informationen über User aus, und zwar:
- die Anzahl der User, die im System gespeichert wurden (laufende Analyse oder Streaming-Processing)
- die Länge des Namens aller User (regelmäßig durchgeführte Analyse oder Batch-Processing)

Stream-Processing hilft festzustellen, wann wie viele User im System existiert haben. Stream-Processing liefert aktuelle daten. Batch-Processing seinerseits läuft jede Minute und unter allen Usern findet den User mit dem längsten Namen. In echten Systemen kann Batch-Processing unter anderem zum Trainieren KI-Modelle verwendet werden. Dafür ist wichtig, dass ein Modell auf allen vorhandenen Daten trainiert und validiert wird.

![1_architecture](/images/1_architecture.jpg)

Die Lösung besteht aus vier Services und einem CronJob Service (ein Begriff aus Kubernetes). Diese Services erfüllen folgende Funktionen:
- Ingest Service (im Code "Ingress" genannt). Über diesen Service kann ein User ins System eingespeist werden. Dieser Service leitet einen User in Form von JSON an Kafka weiter.
- User-Stream-Processor. Dieser Service kriegt Daten von Kafka und erhöht den Zähler. Anschließend wird dieser Zähler zusammen mit entsprechendem Timestamp in Cassandra Datenbank (analytische Datenbank) gespeichert.
- Transactional-Service. Der Service ebenfalls nimmt Daten von Kafka entgegen und speichert sie ohne jegliche Transformationen in MySQL Datenbank. Diese Datenbank kann in einem echten System als Storage für Kundendaten verwendet werden.
- Regulärer-Batch-Processor. Das ist ein CronJob, der jede Minute hochfährt, alle User aus der Datenbank ausliest, den User mit dem längsten Namen identifiziert und in derselben MySQL Datenbank speichert.
- Analytischer Service. Das gilt als der Einstiegspunkt für die Analyse der ins System eingespeisten Daten. Dieser Service stellt Informationen über den User mit dem längsten Name zur Verfügung und gibt den zeitlichen Verlauf der Anzahl der User im System zurück.

![2_architecture](/images/2_architecture.jpg)

Die von Services zur Verfügung gestellte REST API ist in der Tabelle abgebildet:

| Service             | Method | Routes      | Description                                                     |
| ------------------- | ------ | ----------- | --------------------------------------------------------------- |
| Ingest              | POST   | /user       | Einen User ins System einspeisen                                |
| Analytics Processor | GET    | /user/count | Den aktuellen Werte des Zählers abfragen (nur für Debug Zwecke) |
| Analytics Service   | GET    | /analytics  | Einen Freelancer erstellen                                      |

Die REST Endpoints (falls vorhanden) wurden in Handler Strukturen implementiert. CronJob und Transactional Service stellen keine HTTP Endpoints zur Verfügung.

```go
package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stakkato95/lambda-architecture/ingress/config"
	"github.com/stakkato95/lambda-architecture/ingress/domain"
	"github.com/stakkato95/lambda-architecture/ingress/logger"
	"github.com/stakkato95/lambda-architecture/ingress/service"
)

func Start() {
	router := mux.NewRouter()

	repo := domain.NewKafkaUserRepository()

	service := service.NewSimpleUserService(repo)
	handlers := UserHandlers{service}

	router.HandleFunc("/user", handlers.CreateUser).Methods(http.MethodPost)

	port := config.AppConfig.ServerPort
	logger.Info(fmt.Sprintf("started ingress at %s", port))
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Fatal("error when starting ingress " + err.Error())
	}
}
```

### Kurzer Überblick über Technologien

### Gorilla Mux
In Golang steht “mux” für “HTTP request multiplexer”. Solche Bibliotheken definieren, welche HTTP Endpoints von welchen Handlers bedient werden. Gorilla Mux bietet eine gute Leistung und Funktionalität im Vergleich mit anderen Multiplexern. Für Golang existiert eine Unmenge von Multiplexern, zumal sie die Entwicklung der Web-Services deutlich erleichtern.

```go
r := mux.NewRouter()
r.HandleFunc("/", handler)
r.HandleFunc("/products", handler).Methods("POST")
r.HandleFunc("/articles", handler).Methods("GET")
r.HandleFunc("/articles/{id}", handler).Methods("GET", "PUT")
r.HandleFunc("/authors", handler).Queries("surname", "{surname}")
```

### Uber Zap
Uber Zap bietet strukturiertes Logging. Unter strukturiertem Logging versteht man Logging in einem bestimmten Format, wie z.B. JSON. Dadurch kann das Parsen und die Verarbeitung von Logs deutlich vereinfacht werden. Zap ist eine der populärsten und die schnellste Implementierung des strukturierten Loggings für Golang.

```go
encoderConfig := zap.NewProductionEncoderConfig()
encoderConfig.TimeKey = "timestamp"
encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
encoderConfig.StacktraceKey = ""

config := zap.NewProductionConfig()
config.EncoderConfig = encoderConfig
```

```go
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
```
```json
{
    "level":"error",
    "timestamp":"2022-05-14T11:24:22.328Z",
    "caller":"domain/userProcessor.go:44",
    "msg":"error when reading a msg from kafka: write tcp 10.1.1.187:45268->10.1.1.149:9092: use of closed network connection"
}
```

### Viper
Viper ist eine Bibliothek für die Arbeit mit Konfigurationsdateien. Viper unterstützt unterschiedliche Dateiformate (JSON, TOML, YAML, HCL) und Konfigurationsquellen (Umgebungsvariablen, Command Line Flags, Remote Servers wie etcd, usw).

```go
type Config struct {
	ServerPort   string `mapstructure:"SERVER_PORT"`
	KafkaService string `mapstructure:"KAFKA_SERVICE"`
}

var AppConfig Config

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("config not found")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		logger.Panic("config can not be read")
	}

	if AppConfig == (Config{}) {
		logger.Panic("config is emtpy")
	}
}
```

```env
SERVER_PORT=:8080
KAFKA_SERVICE=kafka.default.svc.cluster.local:9092
```

### Verwendung der Message Oriented Middleware

Als MoM wurde in meinem Projekt Kafka verwendet. Kafka kommt in den Big-Data Pipelines sehr oft zum Einsatz und dient als Datenbank, Immutable Log und Queue gleichzeitig. In meinem Fall dient Kafka eher als Immutable Log und Queue. Es wurde ein einziger Topic angelegt, und zwar User Topic. In dieses Topic werden User durch Ingest Service eingespeist und danach von Transactional Service und Stream Processor konsumiert.

![3_mom](/images/3_mom.jpg)

Kafka läuft im Kubernetes Cluster als eine einzige Instanz, deswegen wurde für das User Topic keine Partitionierung und keine Replizierung konfiguriert.

![4_kafdrop](/images/4_kafdrop.jpg)

![4_kafdrop_1](/images/4_kafdrop_1.jpg)

### Packaging für Deployment

Als Packaging für den Service wurde Docker verwendet. Im Rootverzeichnis jedes Services liegt ein Dockerfile. 

```dockerfile
FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=builder /build/main /app/
COPY --from=builder /build/app.env /app/
WORKDIR /app
CMD ["./main"]
```

Build wird durch Makefile gestartet. Makefile enthielt unter anderem Befehle, um einen Service zu testen (es wurden aber keine Tests implementiert), zu starten und generierte Dateien zu löschen.

```makefile
ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif

.DEFAULT_GOAL := docker-push-image

# local dev
test:
	go test ./...
.PHONY:test

build: test
	go build main.go
.PHONY:build

run-with-env: build
	$$env:SERVER_PORT='8080'; $$env:KAFKA_SERVICE='kafka.default.svc.cluster.local:9092'; ./main
.PHONY:run-with-env

# clear local dev
clear:
	rm main.exe
.PHONY:clear

# docker
docker-build-image:
	docker build -t stakkato95/lambda-ingress:latest . -f Dockerfile
.PHONY:docker-build-image

docker-push-image: docker-build-image
	docker push stakkato95/lambda-ingress:latest
.PHONY:docker-push-image

docker-run-tmp-container:
	docker run --rm -p 8080:8080 -d stakkato95/lambda-ingress
.PHONY:docker-local-container
```

Als Container Registry wurde Docker Hub verwendet. Jeder Service hat sein eigenes Repository auf Docker Hub.

![5_dockerhub](/images/5_dockerhub.jpg)

### Deployment der Infrastruktur

Die Infrastruktur des Projekts wurde, genauso wie Services, auf Kubernetes installiert. Die Installierung von Kafka, Kafdrop (Web UI für Kafka), Cassandra und MySQL erfolgt mittels Helm. Helm ist der gängigste Package-Manager und Template-Engine für Kubernetes. Das Packaging (die Zusammenführung von Deployments, Services, Ingress und Templates) erfolgt mittels Helm-Charts.

Als Quelle der Helm-Charts wurde Artifact Hub verwendet. Artifact Hub ist eine mit Docker Hub vergleichbare Quelle der Softwarekomponenten, die mittels Helm verpackt und zur Verfügung gestellt werden. Helm Charts für MySQL, Cassandra und Kafka kommen von Bitnami, Chart für Kafdrop kommt von Bedag Informatik AG.

![6_artifacthub](/images/6_artifacthub.jpg)

![7_artifacthub_kafka](/images/7_artifacthub_kafka.jpg)

![8_artifacthub_kafka_2](/images/8_artifacthub_kafka_2.jpg)

Zur Automatisierung das Deployment der Infrastruktur wurde ein PowerShell Skript geschrieben, der dann die entsprechenden Deployment Scripts für die Teile der Infrastruktur aufruft. Nach demselben Prinzip wurde das Deployment der Microservices vereinfacht.

`Deployment von Cassandra (als Beispiel)`
```powershell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install cassandra-100500 bitnami/cassandra --set dbUser.user=cassandra --set dbUser.password=cassandra
```

`Deployment der ganzen Infrastruktur`
```powershell
./nginx-controller/helm-nginx-controller.ps1 > $null

./cassandra/helm-cassandra.ps1 > $null
./mysql/helm-mysql.ps1 > $null

./kafka/deploy/helm-1-kafka.ps1 > $null
# kafdrop should be installed via ubuntu in advance 
helm ls
./kafka/deploy/helm-4-kafdrop-port-forward.ps1
```

![9_infra](/images/9_infra.jpg)

`Deployment von Analytics Service (als Beispiel)`
```powershell
kubectl apply -f .\service-analytics-deployment.yaml
kubectl apply -f .\service-analytics-service.yaml
kubectl apply -f .\service-analytics-ingress.yaml
```

`Deployment aller Microservices`
```powershell
cd ingress/deploy
echo "deploy ingress"
./deploy.ps1

cd ../../processor-analytics/deploy
echo "`ndeploy processor-analytics"
./deploy.ps1

cd ../../service-transactional
echo "`ndeploy service-transactional"
./deploy.ps1

cd ../service-analytics
echo "`ndeploy service-analytics"
./deploy.ps1

cd ../processor-batch
echo "`ndeploy processor-batch"
./deploy.ps1

cd ..
```

![10_deployment_scripts](/images/10_deployment_scripts.jpg)

### Besonderheiten beim Deployment

Alle User-Facing Services (Ingest/Ingress und analytischer Service) wurden mittels Ingress exposed und auf localhost gemappt. Um die Überlappung der Datenpunkte zu Vermeiden wurde Target-Rewriting verwendet. Somit haben Ingress und Analytics Service bei einem HTTP Aufruf einen eigenen Prefix.

`http://localhost/service-analytics/analytics`

`http://localhost/ingress/user`

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$2
  name: service-analytics-ingress
spec:
  ingressClassName: nginx
  rules:
  - host: localhost
    http:
      paths:
      - path: /service-analytics(/|$)(.*)
        pathType: Prefix
        backend:
          service:
            name: service-analytics-service
            port:
              number: 80
```

Die zweite Besonderheit ist der CronJob Service, der jede Minute startet, um alle User zu analysieren.

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: processor-batch-cronjob
spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: processor-batch
            image: stakkato95/lambda-processor-batch
          restartPolicy: OnFailure
```

![11_cronjob](/images/11_cronjob.jpg)

### Nachteile dieser Art des Deployments und Future Work

Nicht ganz optimal wurde in diesem Projekt das Deployment mittels PowerShell Skripts und YAML implementiert:
- Kubernetes YAML Dateien sind bei fast allen Services identisch und wurden einfach kopiert. Templating durch Helm Charts könnte dieses Problem leicht lösen
- Deployment mit selbstgeschriebenen Skripts ist fehleranfällig (ein Tippfehler in einem Skript führt zu mehreren Stunden Debugging, was mir tatsächlich passiert ist)

Weitere Problem und Herausforderungen:
- wegen Zeitmangels wurde Dockerfile für jeden Service einfach kopiert, obgleich der Inhalt fast identisch ist
- aus demselben Grund wurden Module fürs Logging und Konfiguration ebenfalls kopiert. Solche Module hätte man als Bibliotheken in einem eigenen Repository implementieren sollen.

### Testdurchlauf

Nachdem die Lambda Architektur erfolgreich deployed ist, zeigt der Analytics Service, dass es noch keine User im System gibt.

![12_app_deployed](/images/12_app_deployed.jpg)

`GET http://localhost/service-analytics/analytics`
```json
{
  "longestNameUser": null,
  "userCount": []
}
```

Wenn ein User ingested wird, dann wird das auch im Response des Analytics Services abgebildet.

`Ingest Script`
```powershell
$id = [guid]::NewGuid().ToString()
$userName = "user1"
$payload = '{ \"id\": \"' + $id + '\", \"name\": \"' + $userName + '\" }'
echo $payload
curl -X POST localhost/ingress/user -H 'Content-Type: application/json' -d $payload
```

![13_user_ingested](/images/13_user_ingested.jpg)

`GET http://localhost/service-analytics/analytics`
```json
{
  "longestNameUser": {
    "id": "87d4dc5a-346d-4ae8-b6dc-4e845176de07",
    "name": "user1"
  },
  "userCount": [
    {
      "id": "594759d4-7423-4a74-8a3e-379baa9c1e67",
      "time": "2022-05-14T16:57:40Z",
      "userCount": 1
    }
  ]
}
```

Ein User mit einem noch längeren Namen ersetzt den vorherigen User.

![14_user_ingested](/images/14_user_ingested.jpg)

`GET http://localhost/service-analytics/analytics`
```json
{
  "longestNameUser": {
    "id": "fc0ef01d-bb0d-4400-b665-effb86f30bde",
    "name": "user12345"
  },
  "userCount": [
    {
      "id": "594759d4-7423-4a74-8a3e-379baa9c1e67",
      "time": "2022-05-14T16:57:40Z",
      "userCount": 1
    },
    {
      "id": "b3304405-f8fe-4bec-860b-45615d86e8cc",
      "time": "2022-05-14T17:00:00Z",
      "userCount": 2
    }
  ]
}
```

### Con­clu­sio

Im vorliegenden Projekt wurde Folgendes ausprobiert:
- ein technology Stack für die Entwicklung von Microservice auf Basis Golang (ohne große Web Frameworks)
- Werkzeuge zur Automatisierung des Infrastrukturdeployments für Microservices, und zwar Helm Charts
- Kafka als Message Oriented Middleware