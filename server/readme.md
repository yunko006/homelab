from pathlib import Path

content = """# 🗺️ Plan d'apprentissage — Construire un serveur Go

> Objectif : construire progressivement un serveur en Go en comprenant les fondamentaux, sans framework et en limitant l'utilisation de l'IA.  
> L'idée est d'écrire le code soi-même, puis d'utiliser la documentation et l'aide ponctuelle pour comprendre ce qui bloque.

---

## Vue d'ensemble

```text
TON PROJET
│
├── 1. net
│   │
│   ├── Listen
│   ├── Listener
│   ├── Accept
│   └── Conn
│
├── 2. HTTP
│   │
│   ├── requête
│   ├── réponse
│   ├── headers
│   └── routing
│
└── 3. net/http
        │
        └── comprendre ce que Go fait pour toi


PUIS CADDY
│
├── cmd/caddy/main.go
│
└── modules/caddyhttp/server.go
````

---

# 1. `net` — Construire ton serveur TCP

### Documentation principale

- [Package `net`](https://pkg.go.dev/net?utm_source=chatgpt.com)
- `net.Listener`
- `net.Conn`
- `net.Listen`

Le package `net` fournit les primitives réseau de bas niveau de Go pour TCP/IP, UDP, résolution DNS, etc.

## À étudier dans cet ordre

### 1.1 `Listen`

Comprendre :

```
listener, err := net.Listen("tcp", ":8080")
```

Questions à pouvoir expliquer :

- Qu'est-ce que `"tcp"` ?
- Que signifie `":8080"` ?
- Qu'est-ce que retourne `Listen` ?
- Pourquoi la fonction retourne-t-elle une `error` ?
- Qu'est-ce qu'un `Listener` ?

---

### 1.2 `Listener`

Comprendre l'interface :

```
type Listener interface {
    Accept() (Conn, error)
    Close() error
    Addr() Addr
}
```

Pour commencer, se concentrer sur :

```
Listener
   │
   ├── Accept()
   ├── Close()
   └── Addr()
```

Ne pas chercher à tout comprendre immédiatement.

---

### 1.3 `Accept`

`Accept()` attend qu'un client se connecte.

Conceptuellement :

```
Listener
   │
   │ Accept()
   ▼
 Conn
```

Objectif :

- lancer ton serveur ;
- attendre une connexion ;
- détecter lorsqu'un client se connecte.

---

### 1.4 `Conn`

Documentation : `net.Conn`

Une `Conn` représente la connexion entre ton serveur et un client.

Concepts importants :

```
Conn
├── Read()
├── Write()
├── Close()
└── ...
```

Objectif :

- lire des données envoyées par le client ;
- écrire des données vers le client ;
- fermer proprement la connexion.

---

# 2. HTTP — Comprendre le protocole

Une fois le serveur TCP fonctionnel, ne pas passer immédiatement à `net/http`.

Comprendre d'abord ce qu'est réellement HTTP.

Une requête peut ressembler à :

```
GET /hello HTTP/1.1
Host: localhost:8080
User-Agent: ...
Accept: */*
```

Décomposer :

```
GET
 ↓
méthode

/hello
 ↓
path

HTTP/1.1
 ↓
version

Host: ...
 ↓
header
```

Une réponse peut ressembler à :

```
HTTP/1.1 200 OK
Content-Type: text/plain

Hello World!
```

## Objectif

Comprendre le chemin :

```
TCP
 ↓
bytes
 ↓
HTTP request
 ↓
routing
 ↓
HTTP response
```

À ce stade, l'objectif est de comprendre ce qu'un serveur HTTP fait réellement avant de laisser `net/http` le faire pour nous.

---

# 3. `net/http` — Comprendre ce que Go fait pour toi

### Documentation

- [Package `net/http`](https://pkg.go.dev/net/http?utm_source=chatgpt.com)
- `http.Request`
- `http.ResponseWriter`
- `http.Handler`
- `http.ServeMux`
- `http.Server`

Le package `net/http` fournit les implémentations HTTP client et serveur de Go.

## Concepts à découvrir

```
net/http
│
├── Request
├── Response
├── Handler
├── ServeMux
├── Server
└── ...
```

Par exemple :

```
func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello")
}
```

Ne pas simplement copier ce genre de code.

Chercher à comprendre :

- Pourquoi `w` est un `ResponseWriter` ?
- Pourquoi `r` est un `Request` ?
- Qui appelle la fonction `hello` ?
- Comment le serveur sait qu'elle correspond à `/hello` ?
- Qu'est-ce qu'un `Handler` ?
- Qu'est-ce que `ServeHTTP` ?
- Quel rôle joue `ServeMux` ?

---

# 4. Puis Caddy

Une fois les bases suffisamment comprises, commencer à lire un vrai projet Go.

## 4.1 Point d'entrée

[Caddy — `cmd/caddy/main.go`](https://github.com/caddyserver/caddy/blob/master/cmd/caddy/main.go)

Commencer par comprendre :

- où démarre le programme ;
- comment Caddy initialise son application ;
- comment les différents composants sont assemblés.

---

## 4.2 Serveur HTTP

[Caddy — `modules/caddyhttp/server.go`](https://github.com/caddyserver/caddy/blob/master/modules/caddyhttp/server.go?utm_source=chatgpt.com)

C'est un fichier particulièrement intéressant à étudier après avoir compris `net/http`.

On y retrouvera notamment des concepts comme :

```
Server
ServeHTTP
Request
ResponseWriter
net/http
```

L'objectif n'est pas de comprendre tout Caddy.

L'objectif est de pouvoir progressivement répondre à :

> « Comment un projet Go réel construit-il une architecture de serveur autour de la bibliothèque standard ? »

---

# 🧭 Progression globale

```
                    TON APPRENTISSAGE

                         Go
                          │
                          ▼
                    ┌─────────┐
                    │  net    │
                    └────┬────┘
                         │
                 TCP / Connexions
                         │
                         ▼
                    ┌─────────┐
                    │  HTTP   │
                    └────┬────┘
                         │
             Requests / Responses
                         │
                         ▼
                   ┌──────────┐
                   │ net/http │
                   └────┬─────┘
                        │
                Handlers / Routing
                        │
                        ▼
                    ┌───────┐
                    │ Caddy │
                    └───────┘
```

---

# 🎯 Premier objectif

Pour commencer, ne chercher à apprendre que :

```
net
 │
 ├── Listen
 ├── Listener
 ├── Accept
 └── Conn
```

### Premier serveur à construire

Sans framework et sans copier une solution complète :

1. écouter sur TCP `:8080` ;
2. afficher `Server started` ;
3. attendre une connexion ;
4. détecter lorsqu'un client se connecte ;
5. envoyer `Hello from Go!` ;
6. fermer la connexion.

Pour tester, utiliser éventuellement `nc` / netcat depuis un deuxième terminal :

```
nc localhost 8080
```

---

# 🧠 Méthode d'apprentissage

Pour ce projet :

1. **Essayer seul.**
2. Lire la documentation officielle.
3. Chercher à comprendre les concepts plutôt que copier du code.
4. Écrire le code soi-même.
5. En cas de blocage, montrer le code et l'erreur.
6. Chercher d'abord une explication ou un indice.
7. Ne regarder une solution complète qu'en dernier recours.

Le but n'est pas de construire le serveur le plus rapidement possible.

Le but est de pouvoir expliquer **pourquoi chaque morceau du serveur existe et comment il fonctionne**.

---

# 📚 Ressources principales

- [Documentation Go](https://go.dev/doc/)
- [Package `net`](https://pkg.go.dev/net?utm_source=chatgpt.com)
- `net.Listener`
- `net.Conn`
- `net.Listen`
- [Package `net/http`](https://pkg.go.dev/net/http?utm_source=chatgpt.com)
- [Caddy — GitHub](https://github.com/caddyserver/caddy?utm_source=chatgpt.com)
- [Caddy — `cmd/caddy/main.go`](https://github.com/caddyserver/caddy/blob/master/cmd/caddy/main.go)
- [Caddy — `modules/caddyhttp/server.go`](https://github.com/caddyserver/caddy/blob/master/modules/caddyhttp/server.go?utm_source=chatgpt.com)  
    """

path = Path("/mnt/data/plan-apprentissage-serveur-go.md")  
path.write_text(content, encoding="utf-8")  
print(f"Fichier créé : {path}")